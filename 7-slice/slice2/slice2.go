package main

import "fmt"

func main() {
	/*
		示例1
	*/
	slice6 := make([]int, 0)
	fmt.Printf("The capacity of slice6: %d\n", cap(slice6))
	for i := 1; i <= 5; i++ {
		slice6 = append(slice6, i)
		fmt.Printf("slice6(%d): len: %d, cap: %d\n", i, len(slice6), cap(slice6))
	}
	fmt.Println()

	/*
		示例2
	*/
	slice7 := make([]int, 1024)
	fmt.Printf("The capacity of slice7: %d\n", cap(slice7))
	slice7e1 := append(slice7, make([]int, 200)...)
	fmt.Printf("slice7e1: len: %d, cap: %d\n", len(slice7e1), cap(slice7e1))
	slice7e2 := append(slice7, make([]int, 400)...)
	fmt.Printf("slice7e2: len: %d, cap: %d\n", len(slice7e2), cap(slice7e2))
	slice7e3 := append(slice7, make([]int, 600)...)
	fmt.Printf("slice7e3: len: %d, cap: %d\n", len(slice7e3), cap(slice7e3))
	fmt.Println()

	/*
		示例3
	*/
	slice8 := make([]int, 10)
	fmt.Printf("The capacity of slice8: %d\n", cap(slice8))
	slice8e1 := append(slice8, make([]int, 11)...)
	fmt.Printf("slice8e1: len: %d, cap: %d\n", len(slice8e1), cap(slice8e1))
	slice8e2 := append(slice8e1, make([]int, 23)...)
	fmt.Printf("slice8e2: len: %d, cap: %d\n", len(slice8e2), cap(slice8e2))
	slice8e3 := append(slice8e2, make([]int, 45)...)
	fmt.Printf("slice8e3: len: %d, cap: %d\n", len(slice8e3), cap(slice8e3))
}