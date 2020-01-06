package main

import "fmt"

func main() {
	/*
		示例1
		由于在三个case表达式中存在结果值相等的子表达式，这条语句无法编译通过
		这个约束本身还有个约束，就是只针对结果值为常量的子表达式
		比如子表达式1+1和2不能同时出现
	*/
	//value3 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch value3[4] {
	//case 0, 1, 2:
	//	fmt.Println("0 or 1 or 2")
	//case 2, 3, 4:
	//	fmt.Println("2 or 3 or 4")
	//case 4, 5, 6:
	//	fmt.Println("4 or 5 or 6")
	//}

	/*
		示例2
		绕过子表达式限制
		第一个case表达式和第二个case表达式都包含了value4[2]
	*/
	value4 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value4[4] {
	case value4[0], value4[1], value4[2]:
		fmt.Println("0 or 1 or 2")
	case value4[2], value4[3], value4[4]:
		fmt.Println("2 or 3 or 4")
	case value4[4], value4[5], value4[6]:
		fmt.Println("4 or 5 or26")
	}

	/*
		示例3
		类型判断switch语句的case表达式的子表达式，必须由类型字面量表示，而无法通过间接的方式表示
		value5是空接口类型，byte类型是uint8类型的别名类型，本质是同一个类型，只是类型名称不同
		这条语句无法编译通过
	*/
	//value5 := interface{}(byte(127))
	//switch t := value5.(type) {
	//case uint8, uint16:
	//	fmt.Println("uint8 or uint16")
	//case byte:
	//	fmt.Printf("byte")
	//default:
	//	fmt.Printf("unsupported type: %T", t)
	//}
}