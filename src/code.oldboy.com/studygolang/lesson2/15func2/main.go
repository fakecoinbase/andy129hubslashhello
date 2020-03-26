package main

import (
	"fmt"
	"strings"
)

// 匿名函数和闭包
func main() {

	// 将匿名函数直接赋值给一个变量
	sayHello := func() {
		fmt.Println("匿名函数")
	}

	// 变量执行匿名函数
	sayHello()

	func() {
		fmt.Println("立即执行函数")
	}() // 后面加(), 代表立即执行

	r := a()
	r()

	// 闭包 = 函数+外层变量的引用
	// r2 此时就是一个闭包
	r2 := b()
	r2() // "hello :  沙河小王子"

	r2 = b2("沙河大魔王")
	r2() // "hello : 沙河大魔王"

	// fmt.Println("直接返回：", c()) // 由于 c() 返回的是一个 函数，所以打印出来的值就是 函数的地址
	r3 := c()
	ret := r3()
	fmt.Println("匿名函数返回值：", ret)

	// test1()

	// testCalc()

	testCalc2()

}

// 定义一个函数它的返回值时一个函数
// 把函数作为返回值
func a() func() {
	return func() {
		fmt.Println("返回匿名函数")
	}
}

// 闭包 = 函数+外层变量的引用
func b() func() {
	name := "沙河小王子"
	return func() {
		fmt.Println("hello : ", name)
	}
}

func b2(name string) func() {
	return func() {
		fmt.Println("hello : " + name)
	}
}

func c() func() int {
	return func() int {
		return 0
	}
}
func test1() {
	// 闭包 = 函数+外层变量的引用
	// r 是一个闭包
	r := makeSuffixFunc(".txt")
	ret := r("沙河小王子")
	fmt.Println(ret) // "沙河小王子.txt"

	// r2 是一个闭包
	r2 := makeSuffixFunc(".avi")
	ret2 := r2("沙河蛮王")
	fmt.Println(ret2) // "沙河蛮王.avi"

	// r3 是一个闭包
	r3 := makeSuffixFunc(".doc")
	ret3 := r3("我的简历.doc")
	fmt.Println(ret3) // "我的简历.doc"
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) { // 判断 name 是不是 不以 suffix 为后缀
			return name + suffix
		}
		return name
	}
}

func testCalc() {

	add, sub := calc(100)
	ret := add(100)
	fmt.Println(ret) // "200"

	// 注意，calc(100)时，base =100，  调用add(100)之后，base 等于 200，所以再调用 sub(200)时，内部是这样的：
	// base = base - i   >>>  base = 200 - 200 >>>  base = 0,  return base
	// 由于内部操作会 base = base - i,  会更新到  base 的值，并且 add 与 sub 又共用 这个外层引用，所以会互相影响
	// 如果不想互相影响，就只需 不更新 base 的值，采用一个临时变量储存计算后的值并返回
	ret2 := sub(200)
	fmt.Println(ret2) // "0"
}

// 与 testCalc 做对比，修改了内部 base 的值 （不更新赋值）
func testCalc2() {
	add, sub := calc2(100)
	ret := add(100)
	fmt.Println(ret) // "200"

	ret2 := sub(200)
	fmt.Println(ret2) // "-100"
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base = base + i
		return base
	}

	sub := func(i int) int {
		base = base - i
		return base
	}
	return add, sub
}

func calc2(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		return base + i
	}

	sub := func(i int) int {
		return base - i
	}
	return add, sub
}
