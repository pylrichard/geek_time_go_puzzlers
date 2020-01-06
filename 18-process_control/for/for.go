package main

import "fmt"

func main() {
	//示例1
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
		}
	}
	//输出[1 2 3 7 5 6]
	fmt.Println(numbers1)
	fmt.Println()

	/*
		示例2
		被迭代对象是range表达式结果值的副本而不是原值
		在第一次迭代时，改变numbers2第二个元素的值，新值为3
		但被迭代对象的第二个元素没有改变，它是numbers2的副本
	*/
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	//输出[7 3 5 7 9 11]
	fmt.Println(numbers2)
	fmt.Println()

	/*
		示例3
		注意：切片是引用类型，数组是值类型
	*/
	numbers3 := []int{1, 2, 3, 4, 5, 6}
	maxIndex3 := len(numbers2) - 1
	for i, e := range numbers3 {
		if i == maxIndex3 {
			numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	fmt.Println(numbers3)
}