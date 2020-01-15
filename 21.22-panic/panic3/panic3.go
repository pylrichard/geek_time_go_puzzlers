package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")
	caller()
	fmt.Println("Exit function main.")
}

func caller() {
	fmt.Println("Enter function caller.")
	//正例
	panic(errors.New("something wrong"))
	//反例
	panic(fmt.Println)
	fmt.Println("Exit function caller.")
}
