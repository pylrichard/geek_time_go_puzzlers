package main

func main() {
	//示例1
	//这里会引发编译错误
	var badMap1 = map[[]int]int{}
	_ = badMap1

	//示例2
	var badMap2 = map[interface{}]int{
		"1":      1,
		//这里会引发panic
		[]int{2}: 2,
		3:        3,
	}
	_ = badMap2

	//示例3
	//这里会引发编译错误
	var badMap3 map[[1][]string]int
	_ = badMap3

	//示例4
	type BadKey1 struct {
		slice []string
	}
	//这里会引发编译错误
	var badMap4 map[BadKey1]int
	_ = badMap4

	//示例5
	//这里会引发编译错误
	var badMap5 map[[1][2][3][]string]int
	_ = badMap5

	//示例6
	type BadKey2Field1 struct {
		slice []string
	}
	type BadKey2 struct {
		field BadKey2Field1
	}
	//这里会引发编译错误
	var badMap6 map[BadKey2]int
	_ = badMap6
}