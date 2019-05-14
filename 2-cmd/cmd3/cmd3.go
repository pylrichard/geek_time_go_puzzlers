package main
 
import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	/*
		自定义命令源码文件的参数使用说明就是对变量flag.Usage重新赋值，flag.Usage的类型是func()，即无参数声明且无结果声明的函数类型
		flag.Usage变量在声明时就已经被赋值
		对flag.Usage的赋值必须在调用flag.Parse()之前
		调用flag包函数(比如StringVar()、Parse()等)时，实际是在调用flag.CommandLine变量的对应方法
		flag.CommandLine相当于默认情况下的命令参数容器
	 */
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}