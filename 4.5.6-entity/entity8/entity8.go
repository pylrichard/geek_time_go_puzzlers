package main

import "fmt"

/*
	type MyString2 string //注意这里没有等号
	MyString2和string是两个不同的类型，MyString2是一个新的类型，不同于其他任何类型
	这种方式叫对类型的再定义，需要尽量避免
	string可以被称为MyString2的潜在类型。潜在类型的含义是某个类型在本质上是哪个类型或者是哪个类型的集合
	潜在类型相同的不同类型的值之间是可以进行类型转换的。因此MyString2类型的值与string类型的值可以使用类型转换表达式进行互转
	集合类的类型[]MyString2与[]string这样做是不合法的，因为[]MyString2与[]string的潜在类型不同，分别是MyString2和string
	另外即使两个类型的潜在类型相同，它们的值之间也不能进行判等或比较，它们的变量之间也不能赋值
*/
func main() {
	/*
		示例1
		可以用关键字type声明自定义的各种类型，这些类型必须在Go语言基本类型和高级类型的范畴之内
		type MyString = string
		MyString是string类型的别名类型，别名类型与其源类型的区别只是名称
		源类型与别名类型是一对概念，是两个对立的称呼。别名类型主要是为了代码重构而存在的
		Go语言内建的基本类型中存在两个别名类型。byte是uint8的别名类型，rune是int32的别名类型
	*/
	{
		type MyString = string
		str := "BCD"
		myStr1 := MyString(str)
		myStr2 := MyString("A" + str)
		fmt.Printf("%T(%q) == %T(%q): %v\n",
			str, str, myStr1, myStr1, str == myStr1)
		fmt.Printf("%T(%q) > %T(%q): %v\n",
			str, str, myStr2, myStr2, str > myStr2)
		fmt.Printf("Type %T is the same as type %T.\n", myStr1, str)

		strs := []string{"E", "F", "G"}
		myStrs := []MyString(strs)
		fmt.Printf("A value of type []MyString: %T(%q)\n",
			myStrs, myStrs)
		fmt.Printf("Type %T is the same as type %T.\n", myStrs, strs)
		fmt.Println()
	}
	/*
		示例2
	*/
	{
		type MyString string
		str := "BCD"
		myStr1 := MyString(str)
		myStr2 := MyString("A" + str)
		_ = myStr2
		//这里的判等不合法，会引发编译错误
		//fmt.Printf("%T(%q) == %T(%q): %v\n",
		//	str, str, myStr1, myStr1, str == myStr1)
		//这里的比较不合法，会引发编译错误
		//fmt.Printf("%T(%q) > %T(%q): %v\n",
		//	str, str, myStr2, myStr2, str > myStr2)
		fmt.Printf("Type %T is different from type %T.\n", myStr1, str)

		strs := []string{"E", "F", "G"}
		var myStrs []MyString
		//这里的类型转换不合法，会引发编译错误
		//myStrs := []MyString(strs)
		//fmt.Printf("A value of type []MyString: %T(%q)\n",
		//	myStrs, myStrs)
		fmt.Printf("Type %T is different from type %T.\n", myStrs, strs)
		fmt.Println()
	}
	/*
		示例3
	*/
	{
		type MyString1 = string
		type MyString2 string
		str := "BCD"
		myStr1 := MyString1(str)
		myStr2 := MyString2(str)
		myStr1 = MyString1(myStr2)
		myStr2 = MyString2(myStr1)

		myStr1 = str
		//这里的赋值不合法，会引发编译错误
		//myStr2 = str 
		//myStr1 = myStr2
		//myStr2 = myStr1
	}
}