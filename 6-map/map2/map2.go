package main

import "fmt"

func main() {
	var container map[string]int

	key := "two"
	element, ok := container["two"]
	fmt.Printf("The element paired with key %q in nil map: %d (%v)\n",
		key, element, ok)

	fmt.Printf("The length of nil map: %d\n",
		len(container))

	fmt.Printf("Delete the key-element pair by key %q...\n",
		key)
	delete(container, key)

	element = 2
	fmt.Println("Add a key-element pair to a nil map...")
	//这里会引发panic
	container["two"] = element
}