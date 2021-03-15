package main

import (
	"flag"
	"geek_time_go_puzzlers/3-lib/lib3/lib"
	// 此行无法通过编译，因为internal代码包中声明的公开程序实体仅能被该代码包的直接父包及其子包中的代码引用
	// in "geek_time_go_puzzlers/3-lib/lib3/lib/internal"
	// "os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.SayHello(name)
	// in.SayHello(os.Stdout, name)
}
