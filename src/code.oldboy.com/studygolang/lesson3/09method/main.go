package main

import "fmt"

type MyInt int

// 结构体中的 匿名字段
type student struct {
	name string
	string
	// string           // 省略字段，不能定义多个 string
	int
}

// 任意类型 添加方法
// 不能给别的包 (第三方包 或者 go 语言内置的包)定义的类型 添加方法
func main() {

	// test1()
	test2()
}

// 匿名字段
func test2() {
	var stu1 = student{
		name:   "阳",
		string: "男",
	}
	fmt.Println(stu1.name)
	fmt.Println(stu1.string) // "男"
	//省略字段的名字，可以用 .类型访问，但是要保证 结构体中省略的字段类型 只有一个 string , 定义多个省略的字段 string, 则会报错
}

// 给自定义类型添加 方法
func test1() {
	var m MyInt
	(&m).sayHi()
}

func (m *MyInt) sayHi() {
	// fmt.Println("sayHi : ", m) // sayHi :  0xc00006e068    // m 为指针类型，所以直接打印，打印的是地址
	fmt.Println("sayHi : ", *m) // sayHi :  0
}

/*  不能给别的包 (第三方包 或者 go 语言内置的包)定义的类型 添加方法
func (i *int) sayHello() {
	fmt.Println("sayHello")
}
*/
