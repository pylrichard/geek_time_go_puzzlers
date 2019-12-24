package main

import (
	"fmt"
)

/*
	Go语言规范推荐声明体量较小的接口，建议通过接口间组合来扩展程序，增加程序的灵活性
	相比于包含很多方法的大接口，小接口更加专注地表达某一种能力，更容易被组合在一起
*/
type Animal interface {
	ScientificName() string
	Category() string
}

type Named interface {
	Name() string
}

type Pet interface {
	Animal
	Named
}

type PetTag struct {
	name  string
	owner string
}

func (pt PetTag) Name() string {
	return pt.name
}

func (pt PetTag) Owner() string {
	return pt.owner
}

type Dog struct {
	PetTag
	scientificName string
}

func (dog Dog) ScientificName() string {
	return dog.scientificName
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	petTag := PetTag{name: "little pig"}
	_, ok := interface{}(petTag).(Named)
	fmt.Printf("PetTag implements interface Named: %v\n", ok)
	dog := Dog{
		PetTag:         petTag,
		scientificName: "Labrador Retriever",
	}
	_, ok = interface{}(dog).(Animal)
	fmt.Printf("Dog implements interface Animal: %v\n", ok)
	_, ok = interface{}(dog).(Named)
	fmt.Printf("Dog implements interface Named: %v\n", ok)
	_, ok = interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
}