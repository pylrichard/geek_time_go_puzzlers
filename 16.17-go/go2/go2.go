package main

import (
	"fmt"
	//"time"
)

/*
	等其他的goroutine运行完毕后，再让主goroutine结束运行
*/
func main() {
	num := 10
	/*
		struct{}类似空接口类型interface{}
		它代表了既不包含任何字段也不拥有任何方法的空结构体类型
		占用的内存空间是0字节
		这个值在Go程序中永远只存在一份。虽然可以多次使用这个值字面量，但是用到的都是同一个值
		把通道当作传递信号的介质时，用struct{}作为其元素类型再好不过
	*/
	sign := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func() {
			fmt.Println(i)
			//发送表达式放在go函数体的最后
			sign <- struct{}{}
		}()
	}

	/*
		办法1
		time.Sleep()让当前goroutine暂停一段时间，直到到达指定的恢复运行时间
	*/
	//time.Sleep(time.Millisecond * 500)

	/*
		办法2
		但不容易预估睡眠时间，可以让其他goroutine在运行完毕时发出通知
		先创建一个通道，长度与手动启用的goroutine的数量一致
		在每个手动启用的goroutine即将运行完毕时，向该通道发送一个值
		比使用通道更好的方法是标准库中的sync.WaitGroup类型
	*/
	for j := 0; j < num; j++ {
		//从通道接收元素值，接收次数与手动启用的goroutine的数量保持一致
		<-sign
	}
}