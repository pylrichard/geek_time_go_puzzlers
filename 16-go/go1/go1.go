package main

import "fmt"

/*
	运行结果可能如下：
	1 不会有任何内容被打印(绝大多数情况)
	for语句会很快执行完毕。当它执行完毕时，10个包装了go函数的goroutine往往还没有获得运行的机会
	2 打印10个10
	go函数中的那个对fmt.Println函数的调用是以for语句中的变量i作为参数的。你可以想象一下，如果当for语句执行完毕的时候，这些go函数都还没有执行，那么它们引用的变量i的值都会是10
	3 打印乱序的0~9

*/
func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}