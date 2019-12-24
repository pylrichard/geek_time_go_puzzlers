package main

import (
	"fmt"
	"reflect"
)

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
	/*
		示例1
		dog1未初始化，值为nil
	*/
	var dog1 *Dog
	fmt.Println("The first dog is nil.")
	//dog2值为nil
	dog2 := dog1
	fmt.Println("The second dog is nil.")
	/*
		pet的值不会是nil，因为这个动态值只是pet值的一部分
		字面量nil表示的值叫无类型的nil。这是真正的nil，因为它的类型也是nil
	*/
	var pet Pet = dog2
	/*
		Go语言会识别出赋予pet的值是一个*Dog类型的nil。然后Go语言会用一个iface的实例对它进行包装
		把一个有类型的nil赋给接口变量，那么这个变量的值不会是真正的nil
		判断pet是否与字面量nil相等，结果为false
		让一个接口变量的值真正为nil，要么只声明它但不初始化，要么把字面量nil赋给它
	*/
	if pet == nil {
		fmt.Println("The pet is nil.")
	} else {
		fmt.Println("The pet is not nil.")
	}
	/*
		验证pet的动态类型是*Dog
	*/
	fmt.Printf("The type of pet is %T.\n", pet)
	fmt.Printf("The type of pet is %s.\n", reflect.TypeOf(pet).String())
	fmt.Printf("The type of second dog is %T.\n", dog2)
	fmt.Println()

	//示例2
	wrap := func(dog *Dog) Pet {
		if dog == nil {
			return nil
		}
		return dog
	}
	pet = wrap(dog2)
	if pet == nil {
		fmt.Println("The pet is nil.")
	} else {
		fmt.Println("The pet is not nil.")
	}
}