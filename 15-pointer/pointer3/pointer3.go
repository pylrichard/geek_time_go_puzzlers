package main

import (
	"fmt"
	"unsafe"
)

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	/*
		示例1
		转换规则如下：
			一个指针值可以被转换为一个unsafe.Pointer类型的值，反之亦然
			一个uintptr类型的值也可以被转换为一个unsafe.Pointer类型的值，反之亦然
			一个指针值无法被直接转换成一个uintptr类型的值，反之亦然
			对于指针值和uintptr类型值之间的转换，必须使用unsafe.Pointer类型的值作为中转
	*/
	dog := Dog{"little pig"}
	dogP := &dog
	dogPtr := uintptr(unsafe.Pointer(dogP))

	/*
		把指针值转换成uintptr类型的值的意义
		unsafe.Offsetof()用于获取两个值在内存中的起始存储地址之间的偏移量，以字节为单位
		这两个值一个是某个字段的值，另一个是该字段值所属的结构体值
		在调用这个函数时，需要把针对字段的选择表达式传给它，比如dogP.name
		有了这个偏移量，又有了结构体值在内存中的起始存储地址(dogPtr)，把它们相加就可以得到dogP的name字段值的起始存储地址(namePtr)
	*/
	namePtr := dogPtr + unsafe.Offsetof(dogP.name)
	/*
		通过两次类型转换把namePtr的值转换成一个*string类型的值，这样就得到了指向dogP的name字段值的指针值
		用取址表达式&(dogP.name)也能拿到这个指针值
		如果根本不知道结构体类型是什么，也拿不到dogP变量，如何访问它的name字段
		有namePtr就可以，它就是一个无符号整数，同时也是一个指向程序内部数据的内存地址
	*/
	nameP := (*string)(unsafe.Pointer(namePtr))
	fmt.Printf("nameP == &(dogP.name)? %v\n",
		nameP == &(dogP.name))
	fmt.Printf("The name of dog is %q.\n", *nameP)

	*nameP = "monster"
	fmt.Printf("The name of dog is %q.\n", dogP.name)
	fmt.Println()

	/*
		示例2
		这种不匹配的转换虽然不会引发panic，但是其结果往往不符合预期
	*/
	numP := (*int)(unsafe.Pointer(namePtr))
	num := *numP
	fmt.Printf("This is an unexpected number: %d\n", num)
}