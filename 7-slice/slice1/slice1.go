package main

import "fmt"

func main() {
	/*
		示例1
		内建函数make()声明切片，第2个参数指明切片长度，第3个参数指明切片容量
		切片slice1和slice2的容量分别是5和8
		make()初始化切片时如果不指明容量，那么就和长度一致
		slice2同时指定切片长度和容量，容量实际代表底层数组的长度，注意底层数组长度不可变
		切片长度和容量的关系：
		假设通过一个窗口看到一个数组，但是不一定能看到该数组中的所有元素。这个数组就是slice2的底层数组，这个窗口就是slice2
		slice2的长度指明窗口宽度，决定透过slice2可以看到其底层数组中的哪几个连续元素
		由于slice2长度是5，对应的底层数组的索引范围是[0, 4]
		切片代表的窗口也会被划分成一个个小格子，就像窗户，每个小格子对应底层数组的某一个元素
		用make()或切片值字面量([]int{1, 2, 3})初始化切片时，该窗口最左边的小格子对应底层数组的第1个元素
	*/
	slice1 := make([]int, 5)
	fmt.Printf("The length of slice1: %d\n", len(slice1))
	fmt.Printf("The capacity of slice1: %d\n", cap(slice1))
	fmt.Printf("The value of slice1: %d\n", slice1)
	slice2 := make([]int, 5, 8)
	fmt.Printf("The length of slice2: %d\n", len(slice2))
	fmt.Printf("The capacity of slice2: %d\n", cap(slice2))
	fmt.Printf("The value of slice2: %d\n", slice2)
	fmt.Println()

	/*
		示例2
		通过切片表达式基于数组或切片生成新切片
		slice3长度和容量都是8
	*/
	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	/*
		用切片表达式初始化，[3:6]表示透过新窗口能看到slice3中索引范围[3, 5](即[3, 6)不包括6)的元素
		3称为起始索引，6称为结束索引，slice4的长度就是6-3=3
		因此可以说slice4的索引[0, 2]指向的元素对应slice3底层数组中索引[3, 5]这3个元素
		切片容量代表底层数组的长度，但仅限于使用make()或者切片值字面量初始化切片的情况
		更通用的规则是切片容量可以被看作是透过窗口最多可以看到的底层数组元素的个数
		由于slice4是在slice3上进行切片操作得来的，所以slice3底层数组就是slice4底层数组
		又因为在底层数组不变的情况下，切片代表的窗口可以向右扩展，直至底层数组末尾
		所以slice4容量就是其底层数组的长度8-切片表达式的起始索引3=5
	*/
	slice4 := slice3[3:6]
	fmt.Printf("The length of slice4: %d\n", len(slice4))
	fmt.Printf("The capacity of slice4: %d\n", cap(slice4))
	fmt.Printf("The value of slice4: %d\n", slice4)
	fmt.Println()

	/*
		示例3
		注意，切片代表的窗口是无法向左扩展的
		把切片的窗口向右扩展到最大的方法：使用切片表达式slice4[0:cap(slice4)]得到一个新的切片[]int{4, 5, 6, 7, 8}，长度/容量都是5
	*/
	slice5 := slice4[:cap(slice4)]
	fmt.Printf("The length of slice5: %d\n", len(slice5))
	fmt.Printf("The capacity of slice5: %d\n", cap(slice5))
	fmt.Printf("The value of slice5: %d\n", slice5)
}