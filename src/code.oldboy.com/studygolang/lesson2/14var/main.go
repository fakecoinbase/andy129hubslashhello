package main

import "fmt"

// 函数进阶之作用域

// 定义全局变量 num
var num int = 10

// 定义函数
func testGlobal() {
	// 可以在函数中访问全局变量
	fmt.Println("全局变量", num) // "全局变量 10"
}

// num 会先在 函数内部找，如果有则打印内部的 num ,如果没有则会在全局找
func test2() {
	num := 100
	name := "沙河"
	fmt.Println("变量num : ", num) // "变量num :  100"
	fmt.Println("变量name : ", name)
}

func test3() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// i 只能在 for 语句块中使用
	// fmt.Println(i) // 编译报错："undefined: i"
}
func main() {

	// testGlobal()
	test2()
	// 外部无法访问 test2() 内部定义的变量 (局部变量)
	// fmt.Println(name) //  编译报错："undefined: name"

	// 如果想访问 test2() 里面内部变量的值，则需要添加返回值，将内部函数的值返回给其他地方调用

	// 函数可以作为变量
	abc := testGlobal
	fmt.Printf("%T\n", abc) // "func()"

	// 然后就直接可以调用 abc 函数, 效果与调用 testGlobal() 一致
	abc() // "全局变量 10"

	// 将 add 函数作为参数传入到 calc 函数里，
	r1 := calc(100, 200, add)
	fmt.Println(r1) // "300"
	r2 := calc(100, 200, sub)
	fmt.Println(r2) // "-100"
}

func add(x, y int) int {
	return x + y
}

// 函数可以作为参数
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func sub(x, y int) int {
	return x - y
}
