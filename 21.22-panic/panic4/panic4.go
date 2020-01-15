package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")

	/*
		defer语句写在函数体开始处，引发panic的语句之后的所有语句不会执行
		recover()会拦截并恢复defer语句所属的函数，及其调用的代码中发生的所有panic
	*/
	defer func() {
		fmt.Println("Enter defer function.")

		//recover()的正确用法
		if p := recover(); p != nil {
			//panic已发生
			fmt.Printf("panic: %s\n", p)
		}

		fmt.Println("Exit defer function.")
	}()

	/*
		recover()的错误用法
		先调用recover()，再调用panic()也是不正确的
		因为如果在调用recover()时未发生panic，那么该函数就不会做任何事情，并且只会返回一个nil
	*/
	fmt.Printf("no panic: %v\n", recover())

	//引发panic
	panic(errors.New("something wrong"))

	/*
		recover()的错误用法
		recover()调用不会起到任何作用，没有机会执行
		panic一旦发生，控制权会沿着调用栈的反方向传播。所以在panic()调用之后的代码没有执行的机会
	*/
	p := recover()
	fmt.Printf("panic: %s\n", p)

	fmt.Println("Exit function main.")
}
