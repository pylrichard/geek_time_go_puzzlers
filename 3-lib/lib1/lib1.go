/*
	同一目录下的源码文件都需要被声明为属于同一个代码包
	go run cmd1.go lib1.go
	go build ./
	源码文件声明的包名可以与其所在目录的名称不同，只要这些文件声明的包名一致就可以
 */
package main

import "fmt"
 
func SayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}