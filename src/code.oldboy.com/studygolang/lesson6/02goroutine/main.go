package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// goroutine 和线程的区别

/*  可增长的栈

OS 线程(操作系统线程)一般都有固定的栈内存（通常为2MB）， 一个goroutine 的栈在其生命周期开始只有很小的栈 (典型情况下2KB)
goroutine 的栈不是固定的，他可以按需增大和缩小，goroutine 的栈大小限制可以达到 1GB， 虽然极少会用到这么大。
所以在 Go语言中一次创建 十万左右的 goroutine 也是可以的。

*/

/* goroutine 调度

OS 线程是由 OS内核来调度的， goroutine 则是由 Go运行时 (runtime) 自己的调度器调度的，
	这个调度器使用一个称为 m：n 调度的技术 (复用/调度 m个 goroutine 到 n 个OS线程)。
	goroutine 的调度不需要切换内核语境，所以调用一个 goroutine 比调度一个线程成本低很多。

*/

/* GOMAXPROCS

Go 运行时的调度器使用 GOMAXPROCS 参数来确定需要使用多少个 OS 线程来同时执行Go 代码。
默认值是机器上的 CPU 核心数。例如在一个 8 核心的机器上，调度器会把 Go代码同时调度到 8 个OS线程上 (GOMAXPROCS 是 m:n 调度中的 n)

Go语言中可以通过 runtime.GOMAXPROCS() 函数设置当前程序并发时 占用的CPU 逻辑核心数。
Go 1.5 版本之前，默认使用的是 单核心执行。  Go1.5 版本之后， 默认使用全部的CPU 逻辑核心数。


	go 语言源码分析： func GOMAXPROCS(n int) int {  (设置可以被执行的CPU最大数量, 可以设置1，2，3等。。，如果不设置的话，默认为机器的最大CPU 数)

		// GOMAXPROCS sets the maximum number of CPUs that can be executing
		// simultaneously and returns the previous setting. If n < 1, it does not
		// change the current setting.
		// The number of logical CPUs on the local machine can be queried with NumCPU.
		// This call will go away when the scheduler improves.
		func GOMAXPROCS(n int) int {

*/

/*  Go语言中的操作系统线程和 goroutine 的关系：
1, 一个操作系统线程对应 用户态多个 goroutine
2, go 程序可以同时使用多个操作系统线程
3, goroutine 和 OS 线程是多对多的关系，即 m:n .
		m 是go 程序中创建的 goroutine 数量
		n 是真正干活的线程数量 (虽然不等于 CPU 数量，但是值一般设置为 CPU的数量)
*/

func main() {

	test1()
}

func test1() {

	// runtime.GOMAXPROCS()
	fmt.Println("本机CPU 数量：", runtime.NumCPU()) // 本机CPU 数量： 8

	// 以下测试的目的是： 设置最大可被执行的CPU 逻辑核心数为 1， 然后创建两个 goroutine, 所以会先把一个函数执行完毕 再执行另一个函数 (先 a 后 b 或者 先 b 后 a)
	// 当我们设置 GOMAXPROCS(2) 为 2 时， 推理应该是 a() 和 b() 的打印交错在一起。（目前测试条件不允许，暂时看不到效果）
	runtime.GOMAXPROCS(1) // 设置CPU逻辑核心数
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}

func a() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("b() err ：v%\n", err)
		}
		wg.Done()
	}()

	for i := 0; i < 100; i++ {
		fmt.Printf("a() : %d\n", i)
	}
}

func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("b() err ：v%\n", err)
		}
		wg.Done()
	}()

	for i := 0; i < 100; i++ {
		fmt.Printf("b() : %d\n", i)
	}
}
