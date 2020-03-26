package main

import "fmt"

// panic , recover
func main() {
	// test1()
	// fmt.Println("test1() 崩溃之后，不能执行这条语句")

	test2() // 在 test2() 中进行了 recover() 捕捉，不会影响 函数外语句的执行。
	fmt.Println("test2() 崩溃之后，被recover() 捕捉，还能继续执行这条语句")

	test3()
	fmt.Println("test3() 没有运行错误，依旧被recover() 捕捉，正常继续执行这条语句")
}

// panic 错误
func test1() {
	var a []int
	a[0] = 100 // 运行错误：panic: runtime error: index out of range [0] with length 0

	fmt.Println("这是test1()最后一条语句，看会不会执行到这里") // 由于之前 panic 了，所以这行代码不会执行
}

// 采用 defer 与 recover 一起使用，处理 panic
func test2() {
	defer func() {
		// recover
		err := recover()
		fmt.Println("捕获错误：", err) // "捕获错误： runtime error: index out of range [0] with length 0"
		/*
			这里可以处理一些，内存释放，关闭数据库，关闭I/O 等程序崩溃之后的 收尾工作
		*/
	}() // 注意 这里要加(), 代表一个 逻辑语句块
	var a []int
	a[0] = 100 // 运行错误：panic: runtime error: index out of range [0] with length 0

	// 注意，尽管运行错误被 recover() 捕捉，但是下面这行代码依旧不会执行
	fmt.Println("这是test2()最后一条语句，看会不会执行到这里")
}

func test3() {
	defer func() {
		// recover
		err := recover()
		fmt.Println("捕获错误：", err) // "捕获错误： <nil>,   // 代表没有补货到 任何错误
	}() // 注意 这里要加(), 代表一个 逻辑语句块
	var a [3]int
	a[0] = 100
	fmt.Println(a) // "[100 0 0]"

	// 运行没出错，正常执行下面这条语句
	fmt.Println("这是test3()最后一条语句，看会不会执行到这里")
}
