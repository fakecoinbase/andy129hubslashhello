package main

import (
	"fmt"
	"runtime"
	"time"
)

// goroutine 中可能会遇到的  runtime 方法
// runtime.Gosched() ， runtime.Goexit()，runtime.GOMAXPROCS()
func main() {

	// test1()   

	// test2()

	test3()
}

// 测试 runtime.Gosched() 
func test1(){

	go func(s string) {
		for i:=0;i<20;i++ {
			fmt.Println(s)
		}
	}("world")

	// 
	for i:=0;i<20;i++ {
		runtime.Gosched()   // 出让CPU 时间片,再次分配任务 (分配给其它 任务)
		fmt.Println("hello")
	}


	/*   打印信息：  
		world
		world
		world
		world
		world
		world
		world
		world
		world
		world
		hello
		world
		hello
	*/

	// 总结：
	/*
		打印 "hello" 的主协程，出让CPU 时间片给其它任务， 打印 "world"的goroutine 获得时间片开始打印, 
		但是 主协程 并不是挂起，并不是等待其它 goroutien 执行完毕。 而是动态再切换回来，让主协程 与 其它goroutine 交互打印。
	*/
	
}

// 测试 runtime.Goexit()
func test2(){
	go func() {
		defer fmt.Println("A defer")
		func (){
			defer fmt.Println("B defer")
			// 结束协程
		    runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	time.Sleep(time.Second)
}

// 打印信息：(加 runtime.Goexit())
/*
	B defer
	A defer
*/

// 打印信息： (不加 runtime.Goexit())
/*
	B
	B defer
	A
	A defer
*/

// 测试 runtime.GOMAXPROCS()
func test3(){

	runtime.GOMAXPROCS(3)   // 设置CPU 核心数，可以让程序 交互打印处理的更频繁
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}