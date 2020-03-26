package main

import "fmt"

func main() {
	fmt.Println("数组学习")

	// 数组中， 长度 和 类型 ，作为数组的组成，所以 a 与 b 是不同的数组。
	var a [5]int  // 定义一个长度为 5 存放 int 类型的数组
	var b [10]int // 定义一个长度为 10 存放 int 类型的数组

	// 初始化
	a = [5]int{1, 2, 3, 4, 5}
	b = [10]int{1, 2, 3, 4} // 其他元素，默认补 0   （针对 int 类型的数组）
	fmt.Println(a)          // "[1 2 3 4 5]"
	fmt.Println(b)          // "[1 2 3 4 0 0 0 0 0 0]"

	var c = [3]string{"北京", "上海", "深圳"}
	fmt.Println(c) //  "[北京 上海 深圳]"

	// [...]  表示让编译器根据后面的初始值去 判断数组的长度，然后给变量赋值
	var d = [...]int{2, 3, 45564, 3, 43, 5}
	fmt.Printf("%T\n", d) // "[6]int"

	// 根据索引值初始化
	var e [20]int
	// 注意： 下标不能超过 数组长度
	e = [20]int{19: 1} // 下标索引19 的元素为 1，其他的默认为 0
	fmt.Println(e)     // "[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]"

	// 数组的基本使用
	fmt.Println(e[19]) // "1"
	// 遍历数组的方式1
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// for range 循环
	for i, v := range a {
		fmt.Println(i, v)
	}
}
