package main

import "fmt"

func main() {
	array1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("array1: %v (len: %d, cap: %d)\n",
		array1, len(array1), cap(array1))
	slice9 := array1[1:4]
	slice9[0] = 1
	fmt.Printf("slice9: %v (len: %d, cap: %d)\n",
		slice9, len(slice9), cap(slice9))
	for i := 1; i <= 5; i++ {
		slice9 = append(slice9, i)
		fmt.Printf("slice9(%d): %v (len: %d, cap: %d)\n",
			i, slice9, len(slice9), cap(slice9))
	}
	fmt.Printf("array1: %v (len: %d, cap: %d)\n",
		array1, len(array1), cap(array1))
	fmt.Println()
}