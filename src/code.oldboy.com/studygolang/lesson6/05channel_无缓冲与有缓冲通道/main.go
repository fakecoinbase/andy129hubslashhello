package main

import (
	"fmt"
	"time"
)

// 无缓冲通道
// 光有发送行不行？ 不行，  无缓冲区 channel 又称为 同步 channel, 必须有接收才能发送，否则会一直阻塞。
// 光有接收行不行？ 可以，但是 无缓冲区 channel 直接在 main() 函数里取值，它又取不到值则会阻塞 （例如：test1() 中的错误1 ）
// 有一个方法就是 给他开启一个 goroutine ,让它 异步取值，这样就算取不到值，也不会阻塞main()，并且当 main()执行完毕后，go f1() 也就释放了。

// 有缓冲通道
// 有缓冲区的 channel， 超过容量就阻塞
func main() {

	// test1() // 无缓冲通道 测试

	// test2()   // 有缓冲通道 测试
}

func test1() {
	var ch = make(chan int) // 无缓冲通道

	// 错误1 ：
	/*
		v := <-ch  //  它会一直阻塞在这里，会造成 main() 函数的阻塞。结果就报错
		fmt.Println(v)
	*/

	// 正确操作
	go f1(ch) // 必须先创建一个 取值的 goroutine

	ch <- 10
	// 向无缓冲通道发送值时，要确保前面必须要有从 通道里取值的操作，
	// 因为执行到这一步的时候，它会阻塞在这里，直到有人做好取值的准备，它就可以发送值到通道里了。

	fmt.Println("hello test1()")

	time.Sleep(time.Second * 5)
}

func test2() {
	var ch = make(chan int, 10) // 无缓冲通道

	ch <- 10
	ch <- 10
	ch <- 10
	ch <- 10
	ch <- 10
	ch <- 10
	ch <- 10
	ch <- 10
	ch <- 10
	ch <- 10

	go f1(ch) // 有缓冲的通道，无需在前面先执行 接收的操作 ， 对比 test1() 中 的顺序

	fmt.Println("hello test2()")

	time.Sleep(time.Second)
}

func f1(ch chan int) {

	// time.Sleep(time.Second)
	v := <-ch // 取不到值也会阻塞在这里，但是还好 f1 是一个 goroutine 调用，不会阻塞 main()
	fmt.Println("f1 : ", v)
}
