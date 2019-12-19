package main

import "fmt"

func main() {
	channel1 := make(chan int, 2)
	/*
		发送方
	*/
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Sender: sending elementent %v...\n", i)
			channel1 <- i
		}
		fmt.Println("Sender: close the channel...")
		//由发送方关闭通道，因为向已关闭的通道发送数据，会引发panic
		close(channel1)
	}()

	/*
		接收方
	*/
	for {
		element, ok := <-channel1
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		fmt.Printf("Receiver: received an elementent: %v\n", element)
	}

	fmt.Println("End.")
}