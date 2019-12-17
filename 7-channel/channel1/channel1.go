package main

import "fmt"

func main() {
	channel1 := make(chan int, 3)
	/*
		向通道先后发送三个元素值，通道变量名称、接送操作符、发送元素值三者之间最好用空格进行分割
	*/
	channel1 <- 2
	channel1 <- 1
	channel1 <- 3
	//将最先进入channel1的元素2接收并赋值给变量element1
	element1 := <-channel1
	fmt.Printf("The first element received from channel1: %v\n",
		element1)
}