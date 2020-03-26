package main

import "fmt"

// defer与 return  面试题

// 分析 return , defer 底层实现流程
/*

			函数中 return 语句底层实现                           defer 语句执行的时机

	                  （汇编层面）返回值 = x                                （汇编层面）返回值 = x
	return x  ====>                                   return x  ====>           运行 defer
                      （汇编层面）RET 指令                                  （汇编层面）RET 指令
*/

// 总结：
/*
	1, 看上图.
	2, 函数指定 返回值变量，则代表 在 RET 指令之前， defer 语句是可以针对该值进行修改的。
	3, 其他情况下的 return value,  都将看成是一次性赋值拷贝操作，后续操作不再受影响
*/
func main() {

	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
	// 第一步：x = 5, (汇编)返回值 = 5   (返回值已被赋值，不会再被修改)
	// 第二步：defer语句, x++
	// 第三步： (汇编)RET 指令 ===> 5
}

func f2() (x int) { // 指定返回值变量
	defer func() {
		x++ // 内部没有 x 的变量，所以就向外面找， 找到一个返回值 x 变量，就执行 x++
	}()
	return 5
	// 第一步：(汇编)返回值 = x （虽然此时 x 的值为5， 但是 x 在最终 RET指令之前 还能被 defer 语句里修改）
	// 第二步：defer语句, x++
	// 第三步： (汇编)RET 指令 ===> 6
}

func f3() (y int) { // 指定返回值变量
	x := 5
	defer func() {
		x++
	}()
	return x
	// 第一步：(汇编)返回值 = y (虽然此时 y 的值通过 x 的值赋值为5， defer 语句里面没有再针对 y 变量的处理，所以y 最终值依然是 5)
	// 第二步：defer语句, x++
	// 第三步： (汇编)RET 指令 ===> 5
}

func f4() (x int) { // 指定返回值变量
	defer func(x int) {
		x++
	}(x)
	return 5
	// 第一步：(汇编)返回值 = x （虽然此时 x 的值是5，x 的变量在 RET 指令之前还能被修改，但是 defer 语句里面对匿名函数进行了传参， 此时匿名函数里面的 x 与外部的x 不是同一个变量）
	// 第二步：defer语句, x++， 此时的x 是匿名函数内部的x , 与外层 x 不是同一个变量，所以 x++ 不影响返回值 x
	// 第三步： (汇编)RET 指令 ===> 5
}
