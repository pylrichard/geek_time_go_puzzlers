package main
 
import (
	"flag"
	"fmt"
)

var name string

func init() {
	//	第1个参数是用于存储该命令参数的值的地址
	//	第2个参数是为了指定该命令参数的名称
	//	第3个参数是为了指定在未追加该命令参数时的默认值
	//	第4个参数是该命令参数的简短说明，打印命令说明时显示
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	// 返回一个已经分配好的用于存储命令参数值的地址
	// var ret = flag.String("name", "everyone", "The greeting object.")
}

func main() {
	//	flag.Parse()解析命令参数，并赋值给相应的变量
	//	对该函数的调用必须在所有命令参数存储载体的声明(对变量name的声明)和设置(对flag.StringVar()的调用)之后
	//	并且在读取任何命令参数值之前进行
	//	所以flag.Parse()放在main()第一行
	//	go run cmd.go -name="Richard"
	//	go run cmd.go --help 输出临时生成的可执行文件的完整路径
	//	go build cmd.go 在同级目录下生成可执行文件
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}