package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// goroutine 使用
// 1, 当 main() 函数执行完毕之后， main() 函数中创建的 goroutine，也就会消亡
//
// 2,  time.Sleep(time.Second)  使用等待 验证一下 main() 和  goroutine 的关系
//  一般情况下，不使用 time.Sleep()
// 3, 我们使用 sync.WaitGroup 来处理 goroutine，  实现优雅的等待。
//      wg.Add(int), wg.Done(),  wg.Wait() 三个方法成套使用，缺一不可。
func main() {

	// test1()
	// test2()
	// test3()
	test4()

}

func test4() {
	fmt.Println("test4() start")
	wg.Add(10) // 一次性指定goroutine 个数 (或者可以在循环中 执行wg.Add(1)， 每次循环都 Add(1)一次)
	// 循环创建多个 goroutine
	for i := 0; i < 10; i++ {
		go hello(i) // hello3() 中，在函数结尾要加上  wg.Done()  声明函数已经执行完毕
	}
	wg.Wait() // wg.Add(),  wg.Done(),  wg.Wait()   这三个方法成套使用，缺一不可。
	fmt.Println("test4() end")

	/*	打印信息：

		test4() start
		hello :  0
		hello :  9
		hello :  3
		hello :  4
		hello :  1
		hello :  7
		hello :  6
		hello :  8
		5 出错了
		hello :  2
		test4() end

	*/
}

func hello(i int) {

	// 1, 注意异常捕获，
	// 2, wg.Done()  WaitGroup 里的 goroutine 的数量 要减 1
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		wg.Done() // 如果报错的话，确保 wg.Done() 这里能被执行到。
	}()

	if i == 5 {
		panic(fmt.Errorf(" %d 出错了", i))
	}

	fmt.Println("hello : ", i)
}

func test3() {

	fmt.Println("test3() start")
	wg.Add(10) // 一次性指定goroutine 个数 (或者可以在循环中 执行wg.Add(1)， 每次循环都 Add(1)一次)
	// 循环创建多个 goroutine
	for i := 0; i < 10; i++ {
		go hello3() // hello3() 中，在函数结尾要加上  wg.Done()  声明函数已经执行完毕
	}
	wg.Wait() // wg.Add(),  wg.Done(),  wg.Wait()   这三个方法成套使用，缺一不可。
	fmt.Println("test3() end")

	/*  打印信息：
	test3() start
	Hello3
	Hello3
	Hello3
	Hello3
	Hello3
	Hello3
	Hello3
	Hello3
	Hello3
	Hello3
	test3() end

	*/

}

// goroutine,  sync.WaitGroup  (wg.Add(int),  wg.Done(),  wg.Wait())
func test2() {
	/*   注意：以下这种创建多个 goroutine 的写法是错误的
	wg.Add(1)
	go hello2()

	wg.Add(2)
	go hello3()

	wg.Add(3)
	go hello4()

	wg.Add(4)
	go hello5()
	*/

	wg.Add(2) // Add(delta int) , 里面这个int 值，表示  wg.Add() 开始 到 wg.Wait()之间，中间创建了多少个 goroutine
	go hello2()
	// hello4() // 注意，一个goroutine 对应一个函数，如果这个函数无需 goroutine, 请把函数里面的 wg.Done() 去掉，以免影响其他 goroutine
	go hello3()

	fmt.Println("test2()-----")
	wg.Wait() // 等待，直到 wg 添加的所有任务(多个goroutine )都执行完毕
}

func hello2() {
	fmt.Println("Hello2")
	wg.Done()
}

func hello3() {
	fmt.Println("Hello3")
	wg.Done()
}

func hello4() {
	fmt.Println("Hello4")
	wg.Done()
}

func hello5() {
	fmt.Println("Hello5")
	wg.Done()
}

func test1() {
	// defer fmt.Println("hahahah")   ///注意 defer 和 goroutine 的区别
	go hello1() // 1, 创建一个 goroutine  2. 在新的 goroutine 中执行 hello 函数
	fmt.Println("Hello main test1.")

	time.Sleep(time.Second)
}

func hello1() {
	fmt.Println("Hello 沙河!")
}
