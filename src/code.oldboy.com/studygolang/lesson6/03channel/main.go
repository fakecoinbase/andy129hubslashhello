package main

import "fmt"

// channel
func main() {

	// test1()
	// test2()
	test3()
}

// 初识 channel
func test1() {
	// 定义一个 ch1 变量，它是一个 channel 类型(通道)，这个 channel 内部传递的数据是 int 类型
	var ch1 chan int
	// var ch2 chan string

	// channel 是引用类型
	fmt.Println("ch1 : ", ch1) // "ch1 :  <nil>"

	// 引用类型可以用 make() 进行初始化 (分配内存)
	ch1 = make(chan int, 1) // 给通道创建一个容量，可以接收多少个数值 (如果不创建容量，给通道传入值，则会 deadlock)
	// 通道的操作： 发送、接收、关闭
	// 发送和接收都用同一个符号： <-
	ch1 <- 10 // 把10 发送到 ch1 中

	<-ch1 // 从ch1 中接收值，直接丢弃 (只要进行取值，则通道里面的容量就释放 一个，所以后面再往通道里传值，可正常运行，不会 deadlock)

	ch1 <- 20 // 容量只有一个，在这里又接收了一个数值，则会导致： deadlock

	var value = <-ch1
	fmt.Println("从通道中接收的值 value : ", value)

	// 通道的关闭
	// 1， 关闭通道之后再从通道里取值，如果里面的值都被取完了，你再取，则会取到对应类型的 零值
	close(ch1)
	value2 := <-ch1
	fmt.Println("关闭通道之后再从通道中接收的值 value2 : ", value2) // "0"

	// 2，往关闭的通道中发送值，会引发 panic
	// ch1 <- 30 // "panic: send on closed channel"

	// 3, 关闭一个已经关闭的通道会引发 panic
	// close(ch1)  // "panic: close of closed channel"

}

// channel 测试  (向通道里传入多个值，然后关闭通道，看是否还能拿到值)
func test2() {
	ch1 := make(chan int, 10)
	ch1 <- 3
	v1 := <-ch1
	fmt.Println(v1) // 3

	ch1 <- 7
	ch1 <- 8
	ch1 <- 9
	close(ch1)

	// 从关闭的通道里，可以拿值
	v2 := <-ch1
	fmt.Println(v2) // 7     // 先入先出的原则  (First in First out)

	// 不能往关闭的通道里 发送值，会引发 panic  (panic: send on closed channel)
	// ch1 <- 1

	v3 := <-ch1
	fmt.Println(v3) // 8    // 先入先出的原则  (First in First out)

	v4 := <-ch1
	fmt.Println(v4) // 9    // 先入先出的原则  (First in First out)

	v5 := <-ch1
	fmt.Println(v5) // 0    // 通道已关闭的前提下：当通道里面的值取完之后， 再向通道里面取值，则会得到 对应类型的 零值

}

// chanel 测试2 (向通道里面传入多个值，然后把值取完之后，再去取值，是否能取对应类型的 零值)
func test3() {
	ch1 := make(chan int, 10)

	ch1 <- 7
	ch1 <- 8
	ch1 <- 9

	fmt.Println(<-ch1) // 7
	fmt.Println(<-ch1) // 8
	fmt.Println(<-ch1) // 9

	// 通道没关闭的时候，当里面的值全都被读取之后，再去取 通道里的值，则会 报错
	fmt.Println(<-ch1) // "fatal error: all goroutines are asleep - deadlock!"
}
