package main
 
import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	/*
		flag.PanicOnError和flag.ExitOnError都是预定义在flag包中的常量
		flag.PanicOnError的区别是在最后抛出“运行时恐慌(panic)“
		两种情况都会在调用flag.Parse()时被触发。“运行时恐慌”是Go程序错误处理方面的概念
	 */
	//flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
    /*
		输出与cmd3.go一致，但定制方法更加灵活
		flag.ExitOnError的含义是告诉命令参数容器，当命令后跟--help或者参数设置的不正确时，打印命令参数使用说明后以状态码2结束当前程序
     */
    flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}