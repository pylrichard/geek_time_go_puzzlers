package main

import "fmt"

//Printer 函数签名
type Printer func(contents string) (n int, err error)

func printToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}

func main() {
	var printer Printer
	printer = printToStd
	printer("something")
}