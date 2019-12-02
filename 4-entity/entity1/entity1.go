package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "everyone", "The greeting object")
	/*
		返回字符串指针，类型推断是一种编程语言在编译期自动解释表达式类型的能力，它只能用于对变量或常量的初始化
		表达式类型就是对表达式进行求值后得到结果的类型
		对flag.String()的调用就是一个调用表达式，这个表达式的类型是*string，即字符串的指针类型
		类型推断的好处在于代码重构，见entity2.go
	*/
	var namePtr = flag.String("name", "everyone", "The greeting object")
	/*
		短变量是在类型推断基础上加了语法糖，只能在函数体内部使用短变量声明
		在编写if、for、switch语句时，经常把它安插在初始化子句中，用来声明临时的变量
	*/
	shortVar := *flag.String("shortVar", "everyone", "The greeting object")
	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)
	fmt.Printf("Hello, %v!\n", *namePtr)
	fmt.Printf("Hello, %v!\n", shortVar)
}
