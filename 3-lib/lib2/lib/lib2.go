/*
	在构建或者安装这个代码包时，提供给go命令的路径应该是目录的相对路径
	go install geek_time_go_puzzlers/3-lib/lib2/lib
	pkg子目录下会产生相应的归档文件
	为了不让该代码包的使用者产生困惑，总是应该让声明的包名与其父目录的名称一致
*/
package lib

import "fmt"

/*
	名称首字母为大写的程序实体才可以被当前包外的代码引用，否则它就只能被当前包内的其他代码引用
	通过名称，Go语言把程序实体的访问权限划分为包级私有和包级公开
	对于包级私有的程序实体，即使导入了它所在的代码包也无法引用到它
*/
func SayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
