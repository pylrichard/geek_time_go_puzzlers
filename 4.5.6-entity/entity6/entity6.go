package main

import "fmt"

var container = []string{"zero", "one", "two"}

func getElement(containerI interface{}) (element string, err error) {
	switch container := containerI.(type) {
	case []string:
		element = container[1]
	case map[int]string:
		element = container[1]
	default:
		err = fmt.Errorf("unsupported container type: %T", containerI)
		return
	}
	return
}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	/*
		value, ok := interface{}(container).([]string)
		类型断言表达式包括用来把container变量的值转换为空接口值的interface{}(container)
		以及一个用于判断前者的类型是否为切片类型[]string的.([]string)
		这个表达式的结果可以被赋给两个变量，在这里由value和ok代表
		变量ok是布尔类型，代表类型判断的结果，true或false
		如果是true，那么被判断的值将会被自动转换为[]string类型的值，并赋给变量value，否则value将被赋予nil
		这里的ok也可以没有，就是说类型断言表达式的结果可以只被赋给一个变量，在这里是value
		但是这样当判断为否时就会引发异常
		这种异常在Go语言中被叫做panic，它是一种在Go程序运行期间才会被抛出的异常
		除非显式地恢复panic，否则它会使Go程序崩溃并停止

		类型断言表达式的语法形式是x.(T)。其中的x代表要被判断类型的那个值。这个值当下的类型必须是接口类型
		interface{}代表空接口，任何类型都是它的实现类型
		一对不包裹任何东西的花括号，除了可以代表空的代码块之外，还可以用于表示不包含任何内容的数据结构或者数据类型
		struct{}代表不包含任何字段和方法、空的结构体类型
		interface{}代表不包含任何方法定义、空的接口类型
		对于集合类的数据类型，{}可以表示其值不包含任何元素，比如空的切片值[]string{}，空的字典值map[int]string{}

		[]string是一个类型字面量。类型字面量就是用来表示数据类型本身的若干个字符
		string表示字符串类型的字面量，uint8表示8位无符号整数类型的字面量
		[]string表示元素类型为string的切片类型
		map[int]string表示键类型为int、值类型为string的字典类型

		类型转换表达式的语法形式是T(x)，x可以是一个变量，可以是一个代表值的字面量(比如1.23和struct{})，可以是一个表达式
		表达式的结果只能是一个值，而不能是多个值
		x叫做源值，它的类型就是源类型，而T代表的类型就是目标类型
		在语言规范或其他官方文档中已经说明在编程语言层面很难检测的东西才是应该关注的
	*/
	_, ok1 := interface{}(container).([]string)
	_, ok2 := interface{}(container).(map[int]string)
	if !(ok1 || ok2) {
		fmt.Printf("Error: unsupported container type: %T\n", container)
		return
	}
	fmt.Printf("The element is %q. (container type: %T)\n", container[1], container)

	element, err := getElement(container)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("The element is %q, (container type: %T)\n", element, container)
}
