package main

import (
	"flag"
	"geek_time_go_puzzlers/3-lib/lib2/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	/*
		lib为限定符，指明右边的程序实体所在的代码包

		此处使用lib2.SayHello会报错
		导入了geek_time_go_puzzlers/3-lib/lib2/lib包
		但没有实际使用其中的任何程序实体。这在Go语言中是不被允许的，在编译时就会导致失败
		“as lib2”说明虽然导入了代码包geek_time_go_puzzlers/3-lib/lib2/lib，但是使用其中的程序实体的时候应该以lib2.为限定符
		./cmd2.go:5:2: imported and not used: "geek_time_go_puzzlers/3-lib/lib2/lib" as lib2
		./cmd2.go:16:2: undefined: lib
		源码文件所在的目录相对于src目录的相对路径就是它的代码包导入路径(geek_time_go_puzzlers/3-lib/lib2/lib)
		使用其程序实体时给定的限定符要与它声明所属的代码包名称(lib2)对应
	*/
	lib.SayHello(name)
}
