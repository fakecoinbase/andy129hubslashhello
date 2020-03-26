package main

import (
	"fmt"
)

type Address struct {
	province string
	city     string
}

type Email struct {
	province string
}

type Student struct {
	name string
	age  int
	addr Address
}

// 嵌套匿名结构体
type Teacher struct {
	name    string
	age     int
	Address // 匿名结构体
	Email
}

// 结构体的嵌套
func main() {
	fmt.Println("结构体的嵌套")
	// test1()
	test2()

}

// 结构体中嵌套匿名结构体   (可简写访问字段)
func test2() {
	var teacher1 = Teacher{
		name: "李永乐",
		age:  37,
		Address: Address{ // 匿名结构体初始化方式
			"吉林",
			"长春",
		},
		Email: Email{
			"https://",
		},
	}

	fmt.Println(teacher1) // "{李永乐 37 {吉林 长春}}"

	fmt.Println(teacher1.name) // "李永乐"

	fmt.Println(teacher1.Address.province) //"吉林"

	// 结构体的匿名嵌套，可以有如下省略写法, 它会先去 Teacher 结构体中找 city 字段，
	// 如果没有发现，则会到 Address 里面去找
	// fmt.Println(teacher1.province) //"吉林"   // 针对 匿名结构体，可以这样访问

	// 如果又加入了一个 匿名结构体 Email, 里面也有province 字段，
	// 则此时 teacher1.province 会编译报错： ambiguous selector teacher1.province
	fmt.Println(teacher1.Address.province)
	fmt.Println(teacher1.Email.province)
}

// 结构体的嵌套
func test1() {
	var stu1 = Student{
		name: "杨",
		age:  18,
		addr: Address{
			"湖南",
			"长沙",
		},
	}

	fmt.Println(stu1) // "{杨 18 {湖南 长沙}}"

	fmt.Println(stu1.name)          // "杨"
	fmt.Println(stu1.addr.province) // "湖南"
	fmt.Println(stu1.addr.city)     // "长沙"
}
