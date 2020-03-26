package main

import "fmt"

// 指针
// 指针和地址有什么区别？
/*
	地址：就是内存地址(用字节来描述的内存地址)
	指针：指针是带类型。
	& 和 *
	& : 表示 取地址
	* : 根据地址取值
*/
func main() {

	fmt.Println("指针")
	// test2()
	test3()
}

// 尝试修改指针
func test3() {

	// go 语言中指针只读，不能进行修改， 这个修改到底是指什么样的修改？
	a := "沙河"
	s := &a
	fmt.Println(a) // "沙河"
	*s = "天河"
	fmt.Println(a) // "天河"

	b := "黄河"
	s = &b
	fmt.Println(*s) // "黄河"
}

func test2() {
	// 指针的应用
	a := [3]int{1, 2, 3}
	modifyArray(a)
	fmt.Println(a) // "[1 2 3]"

	modifyArrayByPtr(&a)
	fmt.Println(a) // "[100 2 3]"
}

// 带长度的数组作为参数，是值传递 (拷贝一份)，所以在函数内部对拷贝的数组进行修改，不会影响到 原来的数组
func modifyArray(a1 [3]int) {
	a1[0] = 100
}

// 传入指针类型的数组，则可以修改到原数组
func modifyArrayByPtr(a1 *[3]int) {
	(*a1)[0] = 100 // 个人推崇这种写法，因为逻辑清楚明白
	// 一定要加 (), 不然编译会报错：invalid indirect of a1[0] (type int)

	// 还可以这样写，如下：
	// 语法糖(简化写法)：因为 Go 语言中指针不支持修改，所以它默认就知道了你是想 取a1的值
	// a1[0] = 100
}

func test1() {
	var a int
	fmt.Println(a) // "0"

	// & 取地址
	// b 是指针
	b := &a                    // 取a 的地址赋值给 b
	fmt.Println(b)             // "0xc00006e088"
	fmt.Printf("b的类型：%T\n", b) // "*int"

	// c := "沙河"
	// 由于 b 在上面初始化为 *int 类型，只能接收 int 类型的地址
	// b = &c // 编译报错：cannot use &c (type *string) as type *int in assignment
	fmt.Println(b) // "0xc0000120e0"

	d := 100
	// b 是指针
	b = &d               // d 是 int 类型，所以可以取地址 赋值给 *int
	fmt.Println(b)       // "0xc000012100"
	fmt.Println(b == &d) // "true"

	// * 取地址对应的值
	fmt.Println(*b) // "100"

}
