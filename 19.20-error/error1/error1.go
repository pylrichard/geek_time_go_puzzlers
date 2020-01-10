package main

import (
	"errors"
	"fmt"
)

/*
	echo 判断request为空字符串返回，response会是一个空字符串
	request不是空字符串时err会是nil
*/
func echo(request string) (response string, err error) {
	//卫述语句
	if request == "" {
		err = errors.New("empty request")
		return
	}
	response = fmt.Sprintf("echo: %s", request)
	return
}

func main() {
	//示例1
	for _, req := range []string{"", "hello!"} {
		fmt.Printf("request: %s\n", req)
		resp, err := echo(req)
		if err != nil {
			/*
				调用fmt.Printf()并给定占位符%s可以打印某个值的字符串表示形式
				对于其他类型值只要为这个类型编写String()
				对于error类型值的字符串表示形式取决于Error()
				fmt.Printf()发现被打印的值是error类型，会调用它的Error()
			*/
			fmt.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("response: %s\n", resp)
	}
	fmt.Println()

	/*
		示例2
		当想通过模板化方式生成错误信息，并得到错误值时，使用fmt.Errorf()
		该函数先调用fmt.Sprintf()，得到错误信息
		再调用errors.New()，得到包含该错误信息的error类型值并返回
	*/
	err1 := fmt.Errorf("invalid contents: %s", "#$%")
	//err2的静态类型是error，动态类型是errors包中包级私有的类型*errorString
	err2 := errors.New(fmt.Sprintf("invalid contents: %s", "#$%"))
	/*
		errorString类型的一个指针方法实现了error接口的Error()，这个方法会返回传入的错误信息
		error类型值的Error()相当于其他类型值的String()
	*/
	if err1.Error() == err2.Error() {
		fmt.Println("The error messages in err1 and err2 are the same.")
	}
}