package main

import "fmt"

// 全局变量声明
var globalVar string = "全局变量"

// string = "赋值操作"    // 编译报错，不能在全局进行代码语句赋值，只能在声明时就赋值

func foo() (string, int) {
	return "alex", 9000
}

func main() {

	// 变量声明， var 名称  类型
	var name string
	var age int

	fmt.Println(name)
	fmt.Println(age)

	// 批量定义：
	var (
		a string = "哈哈哈"
		b int    = 1
		c bool
		d float64 = 1.00
	)

	fmt.Println(a, b, c, d) // "哈哈哈 1 false 1"

	var x string = "老男孩"
	fmt.Println(x)
	// 占位符使用
	fmt.Printf("%s 嘿嘿嘿\n", x)

	// 类型推导（编译器根据变量初始值的类型，指定给变量）
	var y = 200
	var z = true
	fmt.Println(y, z)

	// 全局变量
	fmt.Println(globalVar)

	// 短变量声明（只能在函数内部使用，无法在全局声明）
	nazha := "短变量"
	fmt.Println(nazha)

	// 调用 foo 函数
	// _ (匿名变量，也叫 空标识符) 用于接收不需要的值
	aa, _ := foo()
	fmt.Println(aa) // "alex"
	_, bb := foo()
	fmt.Println(bb) // "9000"

	// 不能重复声明同名变量 （在同一个作用域中）
	var cc = "alex"
	// var cc = "hh"
	fmt.Println(cc) // "alex"
}
