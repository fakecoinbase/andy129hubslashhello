package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 接收值时判断通道是否关闭
func main() {

	// test1()
	// test2()
	test3()
}

// 从通道中取值，返回的 bool 类型变量 来判断通道是否被关闭。
func test1() {
	var ch = make(chan int, 10)
	go send(ch) // 创建一个 goroutine 循环给 通道传值
	// goroutine 对应的函数 不能有返回值。

	for {
		v, ok := <-ch // 当通道关闭时， 返回的 bool 类型 ok == false
		if !ok {
			break
		}
		fmt.Println("v : ", v)
	}
}

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i // 给通道传值
	}
	close(ch) // 关闭通道  （不像 文件操作，无需特意关闭通道，它可以被 垃圾回收机制回收）
}

// (推荐这种)使用 for range 来判断通道是否关闭
func test2() {

	var ch = make(chan int, 10)

	go send(ch)

	for v := range ch { // 能取到值，则为 true,  通道关闭，则为false, 取不到正确的值
		fmt.Println("for range v : ", v)
	}
}

// 测试 go 后面接函数的返回值 (写法错误)
func test3() {
	// var ch = make(chan int, 10)
	fmt.Println("test3() start")

	go returnValue()
	// go v := returnValue()   // 虽然 returnValue() 函数带有返回值，但是不能这样写，  go 后面必须接  函数名

	// fmt.Println(v)

	fmt.Println("test3() end ")

	time.Sleep(time.Second)

}

func returnValue() int {
	// 生成随机数，先 初始化随机种子函数， 不然可能随机拿出的数字都是一样的。
	rand.Seed(time.Now().Unix())
	v := rand.Intn(100)

	fmt.Println("returnValue : ", v)
	return v
}
