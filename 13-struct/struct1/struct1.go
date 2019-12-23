package main

import "fmt"

/*
	示例1
	AnimalCategory代表动物分类学中的基本分类法
*/
type AnimalCategory struct {
	//界
	kingdom string
	//门
	phylum  string
	//纲
	class   string
	//目
	order   string
	//科
	family  string
	//属
	genus   string
	//种
	species string
}

/*
	String为方法名称，ac为接收者
	Go语言中可以为一个类型编写String()，自定义该类型的字符串表示形式
	String()不需要任何参数声明，但需要有一个string类型的结果声明
*/
func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

//示例2
type Animal struct {
	//学名
	scientificName string
	/*
		动物基本分类
		AnimalCategory代表了Animal类型的一个嵌入字段
		Go语言规范规定，如果一个字段的声明中只有字段的类型名而没有字段的名称，那么它就是一个嵌入字段(匿名字段)
		通过此类型变量的名称后跟.，再跟嵌入字段类型的方式引用到该字段。也就是说嵌入字段的类型既是类型也是名称
	*/
	AnimalCategory
}

func (a Animal) String() string {
	/*
		对嵌入字段的String()的调用结果融入到了Animal同名方法的结果中
		这种将同名方法的结果逐层包装的手法很常见
	*/
	return fmt.Sprintf("%s (category: %s)",
		a.scientificName, a.AnimalCategory)
}

//接收者类型是Animal
func (a Animal) Category() string {
	//通过链式的选择表达式，选择调用嵌入字段的方法或字段
	return a.AnimalCategory.String()
}

/*
	示例3
	多层嵌入，屏蔽会以嵌入的层级为依据，嵌入层级越深的字段或方法越可能被屏蔽
*/
type Cat struct {
	name string
	Animal
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.Animal.AnimalCategory, cat.name)
}

func main() {
	//示例1
	category := AnimalCategory {species: "cat"}
	/*
		调用fmt.Printf()时，使用占位符%s和category本身就可以打印出后者的字符串表示形式
		而无需显式地调用String()，fmt.Printf()会自己去寻找它
	*/
	fmt.Printf("The animal category: %s\n", category)

	//示例2
	animal := Animal {
		scientificName: "American Shorthair",
		AnimalCategory: category,
	}
	/*
		如果没有为Animal实现String()，会调用嵌入字段AnimalCategory的String()
		实现了String()，只要名称相同，无论两个方法的签名是否一致，都会屏蔽嵌入字段中的同名方法
		存在同名的字段，那么嵌入字段的字段也会被屏蔽
		即使在两个同名的成员中一个是字段，另一个是方法的情况下，屏蔽依然会存在
	*/
	fmt.Printf("The animal: %s\n", animal)

	//示例3
	cat := Cat {
		name:   "little pig",
		Animal: animal,
	}
	/*
		当Cat和Animal都没有实现String()时，AnimalCategory的String()才会被调用

		如果处于同一个层级的多个嵌入字段拥有同名的字段或方法
		那么从被嵌入类型的值选择此名称时会引发一个编译错误
		因为编译器无法确定被选择的成员到底是哪一个
	*/
	fmt.Printf("The cat: %s\n", cat)
}