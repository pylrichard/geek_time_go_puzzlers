package main

import (
	"flag"
	"fmt"
)

func main() {
	// 不改变某个程序与外界的任何交互方式和规则
	// 而只改变其内部实现的代码修改方式叫重构
	// 不显式指定变量name的类型，使得它可以被赋予任何类型的值
	var name = getTheFlag()
	flag.Parse()
	fmt.Printf("Hello, %v!\n", *name)
}

// 随意改变getTheFlag()内部实现及返回结果的类型，而不用修改main()中的任何代码
// 程序灵活性的提升在一些编程语言中是用程序的可维护性和运行效率换来的
// Go是静态类型语言，在初始化变量时确定了变量类型，之后就不可能再改变
// 这就避免了后面程序维护问题，这种类型确定是在编译期完成的，不会对程序运行效率产生任何影响
func getTheFlag() *string {
	return flag.String("name", "everyone", "The greeting object")
}

// 改变getTheFlag()的结果类型之后，Go语言的编译器会在再次构建该程序时，自动更新变量name的类型
//func getTheFlag() *int {
//	return flag.Int("num", 1, "The number of greeting object.")
//}
