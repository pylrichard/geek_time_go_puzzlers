package main

import "fmt"

type Cat struct {
	name           string
	scientificName string
	category       string
}

func New(name, scientificName, category string) Cat {
	return Cat {
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func (cat Cat) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat Cat) Category() string {
	return cat.category
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.category, cat.name)
}

func main() {
	cat := New("little pig", "American Shorthair", "cat")
	//自动转译为(&cat).SetName("monster")，设置原值的名字
	cat.SetName("monster")
	fmt.Printf("The cat: %s\n", cat)

	//设置副本的名字
	cat.SetNameOfCopy("little pig")
	//原值不会改变
	fmt.Printf("The cat: %s\n", cat)

	type Pet interface {
		SetName(name string)
		Name() string
		Category() string
		ScientificName() string
	}

	//基本类型和它的指针类型的方法集合是不同的
	_, ok := interface{}(cat).(Pet)
	fmt.Printf("Cat implements interface Pet: %v\n", ok)
	//指针类型包含SetName()指针方法
	_, ok = interface{}(&cat).(Pet)
	fmt.Printf("*Cat implements interface Pet: %v\n", ok)
}