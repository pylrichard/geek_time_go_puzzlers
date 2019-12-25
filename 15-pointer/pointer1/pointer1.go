package main

type Named interface {
	Name() string
}

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	const num = 123
	//常量不可寻址
	//_ = &num
	//基本类型值的字面量不可寻址
	//_ = &(123)

	var str = "abc"
	_ = str
	//对字符串变量的索引结果值不可寻址
	//_ = &(str[0])
	//对字符串变量的切片结果值不可寻址
	//_ = &(str[0:2])
	str2 := str[0]
	//但这样的寻址就是合法的
	_ = &str2

	//算术操作的结果值不可寻址
	//_ = &(123 + 456)
	num2 := 456
	_ = num2
	//算术操作的结果值不可寻址
	//_ = &(num + num2)

	//对数组字面量的索引结果值不可寻址
	//_ = &([3]int{1, 2, 3}[0])
	//对数组字面量的切片结果值不可寻址
	//_ = &([3]int{1, 2, 3}[0:2])
	//对切片字面量的索引结果值却是可寻址的
	_ = &([]int{1, 2, 3}[0])
	//对切片字面量的切片结果值不可寻址
	//_ = &([]int{1, 2, 3}[0:2])
	//对字典字面量的索引结果值不可寻址
	//_ = &(map[int]string{1: "a"}[0])

	var map1 = map[int]string{1: "a", 2: "b", 3: "c"}
	_ = map1
	//对字典变量的索引结果值不可寻址
	//_ = &(map1[2])

	//字面量代表的函数不可寻址
	//_ = &(func(x, y int) int {
	//	return x + y
	//})
	//标识符代表的函数不可寻址
	//_ = &(fmt.Sprintf)
	//对函数的调用结果值不可寻址
	//_ = &(fmt.Sprintln("abc"))

	dog := Dog{"little pig"}
	_ = dog
	//标识符代表的函数不可寻址
	//_ = &(dog.Name)
	//对方法的调用结果值不可寻址
	//_ = &(dog.Name())

	//结构体字面量的字段不可寻址
	//_ = &(Dog{"little pig"}.name)

	//类型转换表达式的结果值不可寻址
	//_ = &(interface{}(dog))
	dogI := interface{}(dog)
	_ = dogI
	//类型断言表达式的结果值不可寻址
	//_ = &(dogI.(Named))
	named := dogI.(Named)
	_ = named
	//类型断言表达式的结果值不可寻址
	//_ = &(named.(Dog))

	var chan1 = make(chan int, 1)
	chan1 <- 1
	//接收表达式的结果值不可寻址
	//_ = &(<-chan1)
}