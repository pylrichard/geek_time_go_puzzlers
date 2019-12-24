package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

/*
	基本类型Dog的方法集合包含2个方法(值方法)
	指针类型*Dog的方法集合包含3个方法(值方法和指针方法)
	这3个方法分别是Pet接口中方法的实现，所以*Dog类型就成为了Pet接口的实现类型
*/
type Dog struct {
	name string
}

//指针方法
func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	//示例1
	dog := Dog{"little pig"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)
	fmt.Println()

	/*
		示例2
		对于接口类型的变量pet，赋给它的值叫做它的实际值(动态值)，而该值的类型叫做它的实际类型(动态类型)，此处为*Dog
		动态类型是相对于静态类型而言的。pet的静态类型就是Pet，但是它的动态类型却会随着赋给它的动态值而变化
		把Pet接口的实现类型*Fish的值赋给pet，它的动态类型就变为*Fish
		给一个接口类型的变量赋予实际值之前，它的动态类型是不存在的
	*/
	var pet Pet = &dog
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
}