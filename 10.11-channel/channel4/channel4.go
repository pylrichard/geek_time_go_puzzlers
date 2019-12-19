package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//示例1
	//只能发不能收的通道
	var sendChannel = make(chan<- int, 1)
	//只能收不能发的通道
	var recvChannel = make(<-chan int, 1)
	//这里打印的是可以分别代表两个通道的指针的16进制表示
	fmt.Printf("The unidirectional channels: %v, %v\n",
		sendChannel, recvChannel)

	//示例2
	intChannel1 := make(chan int, 3)
	SendInt(intChannel1)

	//示例4
	intChannel2 := getIntChannel()
	for element := range intChannel2 {
		fmt.Printf("The element in intChannel2: %v\n", element)
	}

	//示例5
	_ = GetIntChannel(getIntChannel)
}

//示例2
func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000)
}

//示例3
type Notifier interface {
	SendInt(ch chan<- int)
}

//示例4
func getIntChannel() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}
		
//示例5
type GetIntChannel func() <-chan int