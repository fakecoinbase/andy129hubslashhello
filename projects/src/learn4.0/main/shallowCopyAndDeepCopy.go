package main

import (
	"fmt"
	"reflect"
)

// 赋值，浅拷贝与深拷贝
/*   参考文档：
	http://www.imooc.com/article/272014
	https://blog.csdn.net/enjoy_sun_moon/article/details/84778551
	https://blog.csdn.net/weixin_40165163/article/details/90680466
	https://studygolang.com/topics/2896
 */
func main() {
	fmt.Println("赋值，浅拷贝与深拷贝")
	copyFunc1()
}

// 赋值
func copyFunc1(){

	// 基本数据类型 （赋值时，n 将 m的值 拷贝一份在栈里，两个变量的值互相不影响）
	var m = 4
	var n = m
	fmt.Println(n)   // 4
	m = 5
	fmt.Println(n)   // 4

	// 浅拷贝，只拷贝引用地址，但指向同一块内存
	// 引用类型时 (赋值时，b 将 a 的引用地址 拷贝一份在栈里，并且同时指向 堆里同一块数据)
	var a = []int{1,2,3}   //
	var b = a     // b 与 a 地址相同
	fmt.Println(reflect.TypeOf(b))  // []int
	fmt.Printf("%p, %p\n", a, b)   // 0xc00000c3c0, 0xc00000c3c0  ,   相同
	fmt.Println(b)   // [1 2 3]
	a[2] = 4
	fmt.Println(b)   // [1 2 4]
	b[0] = 0
	fmt.Println(a)   // [0 2 4]

	fmt.Println("---------------------------------------")

	// 那么如何操作 才能让 b 的修改不会影响到 a , a 的修改也不会影响到 b 呢？
	var a1 = [3]int{1,2,3}    // 回顾之前的学习，[]int{} 作为引用传递，如果[3]int{} 设定长度，则为 值传递，那么 a 与 b 就互不影响了
	var b1 = a1               //  回顾之前的学习，如果[3]int{} 设定长度，则 &a != &b, 也就互不影响了
	fmt.Println(reflect.TypeOf(b1))  // [3]int
	fmt.Printf("%p, %p\n", &a1, &b1)   // 0xc00000c3e0, 0xc00000c400，  不相同
	fmt.Println(b1)   // [1 2 3]
	a1[2] = 4
	fmt.Println(b1)   // [1 2 3]
	b1[0] = 0
	fmt.Println(a1)   // [1 2 4]
	// 以上 各自修改值，互不影响

	fmt.Println("---------------------------------------")

	// 那么还有没有其他方法：
	// 深拷贝，直接创建一个新的内存，新的引用地址
	var a2 = []int{1, 2, 3}
	var b2 = make([]int, len(a2))   // 回顾前面学的，make()用法
	i := copy(b2, a2[:])   // copy(dst, src []Type) int，  注意 a2[:] 的这种写法。
	/*  关于 copy()函数的返回值 说明：
		returns the number of elements copied, which will be the minimum of
		len(src) and len(dst).
	 */
	fmt.Println("copy()函数返回值：", i)   // 3
	fmt.Println(reflect.TypeOf(b2))   // []int
	fmt.Println(b2)   // [1 2 3]
	a2[2] = 4
	fmt.Println(b2)   // [1 2 3]
	b2[0] = 0
	fmt.Println(a2)   // [1 2 4]
	// 以上 各自修改值，互不影响

	/*	Go 语言中 copy()函数说明

	// The copy built-in function copies elements from a source slice into a
		// destination slice. (As a special case, it also will copy bytes from a
		// string to a slice of bytes.) The source and destination may overlap. Copy
		// returns the number of elements copied, which will be the minimum of
		// len(src) and len(dst).
		func copy(dst, src []Type) int

	 */
}
