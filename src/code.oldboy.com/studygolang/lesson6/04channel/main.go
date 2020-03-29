package main

import (
	"fmt"
	"time"
)

// 无缓冲通道和 缓冲通道
func main() { // main() 函数其实也是一个 goroutine

	// test1()

	// test2()

	// test3()

	test4()
}

func test1() {
	// 无缓冲通道
	ch := make(chan bool) // 不设置容量(缓冲区)
	go receiveValue(ch)

	ch <- true

	time.Sleep(time.Second)
}

func receiveValue(ch chan bool) {

	ret := <-ch // 会在这里阻塞， 从通道里取值， 会在这里一直等待, 直到 ch <- true (通道里面有值)
	fmt.Println("接收的值：ret , ", ret)
}

// len(ch), cap(ch)  查看通道里 值的数量，和 通道的容量 （传入一个，数量就加1， 取出一个值，数量就减 1）
func test2() {
	// 有缓冲通道
	ch := make(chan bool, 8)                    // 设置容量(缓冲区)
	fmt.Println("----------", len(ch), cap(ch)) // "0 8"

	go receiveValue2(ch)

	ch <- true

	fmt.Println("----------2", len(ch), cap(ch)) // "1 8"

	time.Sleep(time.Second)
}

func receiveValue2(ch chan bool) {
	ret := <-ch // 会在这里阻塞， 从通道里取值， 会在这里一直等待, 直到 ch <- true (通道里面有值)
	fmt.Println("接收的值：ret , ", ret)
	fmt.Println("----------3", len(ch), cap(ch)) // "0 8"
}

// 测试 channel 里的容量是否可以自动扩容
func test3() {
	// 有缓冲通道
	ch := make(chan bool, 3)
	fmt.Println(len(ch), cap(ch)) // "0 3"
	ch <- true
	ch <- false
	ch <- true

	fmt.Println(len(ch), cap(ch)) // "3 3"
	// 已经超过 通道设定的容量值，则报错：fatal error: all goroutines are asleep - deadlock!
	// ch <- false  // fatal error: all goroutines are asleep - deadlock!
}

// 关闭后的通道，查看len 和 cap
func test4() {

	ch := make(chan int, 10)
	fmt.Println(len(ch), cap(ch)) // "0 10"

	ch <- 1234
	ch <- 345
	fmt.Println(len(ch), cap(ch)) // "2 10"

	close(ch)
	fmt.Println(len(ch), cap(ch)) // "2 10"    // 虽然通道关闭了，但是依旧可以获取通道里值的数量 和 容量
}
