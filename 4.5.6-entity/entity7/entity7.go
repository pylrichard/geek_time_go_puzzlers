package main

import "fmt"

func main() {
	/*
		重点1的示例
		对于整数类型值、整数常量之间的类型转换，原则上只要源值在目标类型的可表示范围内就是合法的
		uint8(255)可以把无类型的常量255转换为uint8类型的值，是因为255在[0, 255]的范围内
		注意源整数类型的可表示范围较大，而目标类型的可表示范围较小的情况
		整数在Go语言以及计算机中都是以补码形式存储，补码其实是原码个位求反再加1
		int16类型值-255的补码是1111111100000001。如果把该值转换为int8类型的值，Go语言会把在高位上的8位二进制数截掉
		从而得到00000001。又由于其最左边一位是0，表示它是正整数，正整数的补码等于其原码，所以dstInt的值就是1
		当整数值的类型的有效范围由宽变窄时，只需在补码形式下截掉一定数量的高位二进制数即可
		类似规则还有：把一个浮点数类型的值转换为整数类型值时，前者的小数部分会被全部截掉
	*/
	var srcInt = int16(-255)
	/*
		执行uint16(srcInt)是因为只有这样才能得到全二进制的表示
		例如fmt.Printf("%b", srcInt)将打印出"-11111111"，后者是负数符号再加上srcInt的绝对值的补码
		而fmt.Printf("%b", uint16(srcInt))才会打印出srcInt原值的补码"1111111100000001"
	*/
	fmt.Printf("The complement of srcInt: %b (%b)\n",
		uint16(srcInt), srcInt)
	dstInt := int8(srcInt)
	fmt.Printf("The complement of dstInt: %b (%b)\n",
		uint8(dstInt), dstInt)
	fmt.Println()

	/*
		重点2的示例
		把一个整数值转换为一个string类型的值是可行的
		被转换的整数值要可以代表一个有效的Unicode代码点，否则转换的结果将会是"�"(仅由高亮的问号组成的字符串值)
		字符"�"的Unicode代码点是U+FFFD。它是Unicode标准中定义的Replacement Character
		专用于替换那些未知的、不被认可的以及无法展示的字符
		比如string(-1)，-1无法代表一个有效的Unicode代码点，所以得到的总会是"�"
	*/
	fmt.Printf("The Replacement Character: %s\n", string(-1))
	fmt.Printf("The Unicode codepoint of Replacement Character: %U\n", '�')
	fmt.Println()

	/*
		重点3的示例
		一个值在从string类型向[]byte类型转换时代表着以UTF-8编码的字符串会被拆分成零散、独立的字节
		除了与ASCII编码兼容的那部分字符集，以UTF-8编码的某个单一字节是无法代表一个字符的
		string([]byte{'\xe4', '\xbd', '\xa0', '\xe5', '\xa5', '\xbd'})//你好
		UTF-8编码的三个字节\xe4、\xbd、\xa0合在一起才能代表字符"你"，而\xe5、\xa5、\xbd合在一起才能代表字符"好"

		一个值在从string类型向[]rune类型转换时代表着字符串会被拆分成一个个Unicode字符
		string([]rune{'\u4F60', '\u597D'})//你好
	*/
	srcStr := "你好"
	fmt.Printf("The string: %q\n", srcStr)
	fmt.Printf("The hex of %q: %x\n", srcStr, srcStr)
	fmt.Printf("The byte slice of %q: % x\n", srcStr, []byte(srcStr))
	fmt.Printf("The string: %q\n", string([]byte{'\xe4', '\xbd', '\xa0', '\xe5', '\xa5', '\xbd'}))
	fmt.Printf("The rune slice of %q: %U\n", srcStr, []rune(srcStr))
	fmt.Printf("The string: %q\n", string([]rune{'\u4F60', '\u597D'}))
}