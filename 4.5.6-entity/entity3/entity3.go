package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	//用短变量声明对新变量num和旧变量err进行了声明并赋值
	num, err := io.WriteString(os.Stdout, "Hello Everyone\n")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("%d bytes write\n", num)
}
