package main

import "fmt"

// 自定义类型
// NewInt 这是一个新类型
type NewInt int

// 类型别名:  只存在代码编写过程中，代码编译之后根本不存在  shahe
// 类型别名 可提高代码的可读性
type shahe = int

// go语言中常见的 别名
/*
	byte : uint8
	rune : int32
*/

func main() {
	test1()
}

func test1() {
	var a NewInt
	fmt.Println(a)
	fmt.Printf("%T\n", a) // "main.newInt"    // main 包里的 newInt 类型

	var b shahe
	fmt.Println(b)
	fmt.Printf("%T\n", b) // "int"

	// go 语言中常见的内置别名
	var c byte
	fmt.Println(c)
	fmt.Printf("%T\n", c) // "uint8"

	var d rune
	fmt.Println(d)
	fmt.Printf("%T\n", d) // "int32"
}
