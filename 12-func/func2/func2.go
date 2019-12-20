package main

import (
	"errors"
	"fmt"
)

//示例1
type operate func(x, y int) int

//示例2
func calculate(x int, y int, op operate) (int, error) {
	//用卫述语句检查参数，如果op为nil，返回0和代表具体错误的error类型值
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

//示例3
type calculateFunc func(x int, y int) (int, error)

//示例4
func genCalculator(op operate) calculateFunc {
	//返回匿名函数
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	x, y := 12, 23
	//实现operate类型的匿名函数
	op := func(x, y int) int {
		return x + y
	}
	result, err := calculate(x, y, op)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)
	result, err = calculate(x, y, nil)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)

	x, y = 56, 78
	/*
		闭包可以动态地生成部分程序逻辑
		在程序运行的过程中，根据需要生成功能不同的函数，继而影响后续的程序行为
		这与GoF设计模式中的模板方法模式有异曲同工之妙
	*/
	add := genCalculator(op)
	result, err = add(x, y)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)
}