package main

import "fmt"

// panic 和 recover
// 1, recover() 必须搭配 defer.
// 2, defer 一定要在可能引发 panic 的语句之前定义。
func main() {
	/*
		a() // "a()"           // 正常打印
		b() // 报错：panic: panic in b      // 程序崩溃退出
		c() // 不再执行
	*/

	// 采用 recover() 捕获panic 之后，不会影响到后续代码的执行，所以 c() 能正常调用
	a() // "a()"
	d() // "func d error :  panic in d"
	c() // "c()"

}

func a() {
	fmt.Println("a()")
}

func b() {
	panic("panic in b")
}

func c() {
	fmt.Println("c()")
}

// 回顾一下 defer 延迟执行的特性
// 执行 d(), 按照代码顺序，先把 defer 语句块放到一遍，继续向下执行
// 执行 panic()，注意这里虽然程序崩溃了，但依旧不影响 defer 语句块的执行
// 执行完 panic() 之后，d() 快执行完了，所以会找到 defer 语句块，执行它
// 然后 recover() 捕获程序异常信息， 然后做一些操作
// recover 用在程序会 panic 之前，采用 defer 延迟执行
func d() {
	defer func() {
		err := recover() // recover() 捕获程序异常信息，并返回 err
		if err != nil {  // 如果 err != nil , 代表有错误信息
			fmt.Println("func d error : ", err) // err 信息，即是： panic("panic in d"),  panic里面指定的错误信息
		}
	}() // 注意这里要加上(),  因为这是语句块，保证正常执行
	panic("panic in d")
}
