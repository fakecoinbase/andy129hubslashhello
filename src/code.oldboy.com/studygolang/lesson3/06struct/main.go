package main

import "fmt"

type test8 struct {
	a int8
	b int8
	c int8
}

type test16 struct {
	a int16
	b int16
	c int16
}

type student struct {
	name string
	age  int
}

// 结构体的内存布局
// 内存地址 是以字节为单位的 十六进制
// 1字节= 8位 = 8bit, 根据结构体中属性的类型，十六进制地址数值 自动增加 N
func main() {
	//test1()
	//test2()
	//test3()
	test4()
}

// 语法糖，对比 test1():    fmt.Println(&t.a)  也可以写成 fmt.Println(&(t.a))
func test4() {
	var t = test8{
		a: 1,
		b: 2,
		c: 3,
	}
	fmt.Println(t)              // "{1 2 3}"
	fmt.Println(&t)             // "&{1 2 3}"
	fmt.Printf("t的地址：%p\n", &t) // "0xc00006c068"     // 结构体的地址与 结构体中第一个属性的地址相同，类似于数组，slice

	fmt.Println("结构体中个属性的值的地址为：")
	fmt.Println(&t.a) // "0xc00006c068"     // 由于 a 为int8 ,所以占8位，一个字节，所以后面的 b, c 的地址连续 加 1
	fmt.Println(&t.b) // "0xc00006c069"
	fmt.Println(&t.c) // "0xc00006c06a"
}

// 通过指针达到 结构体的引用传递的作用
func test3() {
	var wang = student{
		name: "王",
		age:  28,
	}

	yang := &wang            // yang 是一个指针, 是一个 main 包下面的 student 的指针
	fmt.Printf("%T\n", yang) // "*main.student"
	(*yang).name = "阳"
	fmt.Println(wang.name)    // "阳"    // 达到了修改 wang.name 的操作
	wang.name = "蛮王"          // 修改 wang.name
	fmt.Println((*yang).name) // "蛮王"    // 也达到了 修改 yang 的操作
}

// 验证结构体是 值类型
func test2() {
	var wang = student{
		name: "王",
		age:  28,
	}

	yang := wang // 通过下面的修改结果，得出此条代码 只是拷贝，并没有进行引用传递， 说明结构体是 值类型
	yang.name = "阳"
	fmt.Println(yang.name) // "阳"
	fmt.Println(wang.name) // "王"
}

// 验证结构体内部的内存是连续的
func test1() {

	var t = test8{
		a: 1,
		b: 2,
		c: 3,
	}
	fmt.Println(t)              // "{1 2 3}"
	fmt.Println(&t)             // "&{1 2 3}"
	fmt.Printf("t的地址：%p\n", &t) // "0xc0000120c0"     // 结构体的地址与 结构体中第一个属性的地址相同，类似于数组，slice

	fmt.Println("结构体中个属性的值的地址为：")
	fmt.Println(&(t.a)) // "0xc0000120c0"     // 由于 a 为int8 ,所以占8位，一个字节，所以后面的 b, c 的地址连续 加 1
	fmt.Println(&(t.b)) // "0xc0000120c1"
	fmt.Println(&(t.c)) // "0xc0000120c2"

	fmt.Println("-----------------------------------------------------------")

	var t2 = test16{
		a: 1,
		b: 2,
		c: 3,
	}

	fmt.Println(t2)               // "{1 2 3}"
	fmt.Println(&t2)              // "&{1 2 3}"
	fmt.Printf("t2的地址：%p\n", &t2) // "0xc0000120f4"   // 结构体的地址与 结构体中第一个属性的地址相同，类似于数组，slice

	fmt.Println("结构体中个属性的值的地址为：")
	fmt.Println(&(t2.a)) // "0xc0000120f4"     // 由于 a 为int16 ,所以占16位，两个字节，所以后面的 b, c 的地址连续 加 2
	fmt.Println(&(t2.b)) // "0xc0000120f6"
	fmt.Println(&(t2.c)) // "0xc0000120f8"

}
