package main

import "fmt"

func main() {
	//示例1
	//value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	/*
		Go语言中只有类型相同的值之间才有可能进行判等操作
		如果switch表达式的结果值是无类型的常量，1+3的结果值就是无类型的常量4，这个常量会被自动转换为此种常量的默认类型值
		比如整数4的默认类型是int，浮点数3.14的默认类型是float64
		switch表达式的结果类型是int，case表达式中子表达式的结果类型是int8，它们的类型并不相同
		所以这条语句无法编译通过
	*/
	//switch 1 + 3 {
	//case value1[0], value1[1]:
	//	fmt.Println("0 or 1")
	//case value1[2], value1[3]:
	//	fmt.Println("2 or 3")
	//case value1[4], value1[5], value1[6]:
	//	fmt.Println("4 or 5 or 6")
	//}

	/*
		示例2
		switch表达式的结果值是int8，case表达式中子表达式的结果值是无类型常量
		它的类型会被自动转换为switch表达式的结果类型
		又由于上述几个整数都可以被转换为int8，所以进行判等操作没有问题
		如果自动转换没成功，switch语句通不过编译
	*/
	value2 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value2[4] {
	case 0, 1:
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
}