package main

import (
	"fmt"
)

/*
	去掉SetName()指针方法，Dog类型变成Pet接口的实现类型
*/
type Pet interface {
	Name() string
	Category() string
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

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	//示例1
	dog := Dog{"little pig"}
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	var pet Pet = dog
	//SetName()的接收者是指向dog的指针值的副本
	dog.SetName("monster")
	//dog的name字段的值是"monster"
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	/*
		pet的name字段的值依然是"little pig"
		使用一个变量给另外一个变量赋值，赋给后者的并不是前者持有的值，而是该值的一个副本
	*/
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
	fmt.Println()

	//示例2
	dog1 := Dog{"little pig"}
	fmt.Printf("The name of first dog is %q.\n", dog1.Name())
	dog2 := dog1
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	dog1.name = "monster"
	fmt.Printf("The name of first dog is %q.\n", dog1.Name())
	//dog2的name字段的值依然会是"little pig"
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	fmt.Println()

	//示例3
	dog = Dog{"little pig"}
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	pet = &dog
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
}