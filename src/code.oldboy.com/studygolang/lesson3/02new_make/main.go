package main

import "fmt"

// new , make
/*
	new:  用来初始化 值类型 指针
	make: 用来初始化 slice, map, chan
*/
func main() {

	// test1()
	test2()
}

func test2() {
	var a = new(int)
	fmt.Println(a) // "0xc0000120c0"

	*a = 10
	fmt.Println(a)  // "0xc0000120c0"
	fmt.Println(*a) // "10"

	// 创建一个数组的指针
	var c = new([3]int)
	fmt.Println(c) // "&[0 0 0]"

	(*c)[0] = 1
	(*c)[1] = 2
	fmt.Println(*c) // "[1 2 0]"
}

// 错误写法
func test1() {

	// 错误写法
	var a *int
	fmt.Printf("a的类型：%T\n", a) // "*int"
	// fmt.Println(*a)            // a 指针没有初始化，所以不能进行取值操作， 运行报错:panic: runtime error: invalid memory address or nil pointer dereference

	// var b *int = 0 //  编译报错：cannot use 0 (type int) as type *int in assignment
	// fmt.Println(*b)

}
