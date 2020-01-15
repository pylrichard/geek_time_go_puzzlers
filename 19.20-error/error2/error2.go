package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

//underlyingError 会返回已知的操作系统相关错误的潜在错误值
func underlyingError(err error) error {
	/*
		类型不同可以如此分辨，但是在错误值类型相同的情况下就不适用
		Go语言标准库中有不少以相同方式创建的同类型的错误值
	*/
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}

func main() {
	//示例1
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Printf("unexpected error: %s\n", err)
		return
	}
	//人为制造*os.PathError类型的错误
	r.Close()
	_, err = w.Write([]byte("hi"))
	uError := underlyingError(err)
	fmt.Printf("underlying error: %s (type: %T)\n",
		uError, uError)
	fmt.Println()

	//示例2
	paths1 := []string{
		os.Args[0],           // 当前的源码文件或可执行文件
		"/it/must/not/exist", // 肯定不存在的目录
		os.DevNull,           // 肯定存在的目录
	}
	//err代表某个文件操作相关的错误
	printError1 := func(i int, err error) {
		if err == nil {
			fmt.Println("nil error")
			return
		}
		uError = underlyingError(err)
		//可以使用if语句和判等操作符
		switch uError {
		case os.ErrClosed:
			fmt.Printf("error(closed)[%d]: %s\n", i, uError)
		case os.ErrInvalid:
			fmt.Printf("error(invalid)[%d]: %s\n", i, uError)
		case os.ErrPermission:
			fmt.Printf("error(permission)[%d]: %s\n", i, uError)
		}
	}
	var f *os.File
	var index int
	{
		index = 0
		f, err = os.Open(paths1[index])
		if err != nil {
			fmt.Printf("unexpected error: %s\n", err)
			return
		}
		//人为制造潜在错误为os.ErrClosed的错误
		f.Close()
		_, err = f.Read([]byte{})
		printError1(index, err)
	}
	{
		index = 1
		//人为制造os.ErrInvalid错误
		f, _ = os.Open(paths1[index])
		_, err = f.Stat()
		printError1(index, err)
	}
	{
		index = 2
		//人为制造潜在错误为os.ErrPermission的错误
		_, err = exec.LookPath(paths1[index])
		printError1(index, err)
	}
	if f != nil {
		f.Close()
	}
	fmt.Println()

	//示例3
	paths2 := []string{
		runtime.GOROOT(),     // 当前环境下的Go语言根目录
		"/it/must/not/exist", // 肯定不存在的目录
		os.DevNull,           // 肯定存在的目录
	}
	printError2 := func(i int, err error) {
		if err == nil {
			fmt.Println("nil error")
			return
		}
		uError = underlyingError(err)
		if os.IsExist(uError) {
			fmt.Printf("error(exist)[%d]: %s\n", i, uError)
		} else if os.IsNotExist(uError) {
			fmt.Printf("error(not exist)[%d]: %s\n", i, uError)
		} else if os.IsPermission(uError) {
			fmt.Printf("error(permission)[%d]: %s\n", i, uError)
		} else {
			fmt.Printf("error(other)[%d]: %s\n", i, uError)
		}
	}
	{
		index = 0
		err = os.Mkdir(paths2[index], 0700)
		printError2(index, err)
	}
	{
		index = 1
		f, err = os.Open(paths2[index])
		printError2(index, err)
	}
	{
		index = 2
		_, err = exec.LookPath(paths2[index])
		printError2(index, err)
	}
	if f != nil {
		f.Close()
	}
}