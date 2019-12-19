/**
	本文件包含全域代码块、main包代表的代码块、main()代表的代码块，main()中用花括号包起来的代码块
 	如果在当前源码文件中导入了其他代码包，那么引用其中的程序实体时是需要以限定符为前缀的
	所以程序在找代表变量的那个未加限定符的名字(即标识符)时，是不会去被导入的代码包中查找的
*/
package main

/*
	把代码包导入语句写成import . xxx的形式(注意中间的.)，那么就会让xxx包中公开的程序实体被当前源码文件中的代码，视为当前代码包中的程序实体
	比如代码包导入语句import . fmt，那么在当前源码文件中引用fmt.Printf()时直接用Printf()就可以了
	在这个特殊情况下，程序在查找当前源码文件后会先去查用这种方式导入的代码包
*/
import "fmt"

var block = "package"

/*
	声明重名的变量是无法通过编译的，用短变量声明对已有变量进行重声明除外，但这只是对于同一个代码块而言
*/
func main() {
	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s\n", block)
	}
	fmt.Printf("The block is %s\n", block)
}
