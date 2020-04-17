package main

import (
	"fmt"
	"sync"
	//"time"
)

var wg sync.WaitGroup

// context 学习
func main() {
	 
	// testSimpleChan()
	testSimpleChan2()

	// testExitGoroutine()

}

func testSimpleChan(){

	// 无缓冲通道 , 两点注意：
	// 1, 如果先发送再接收，则会报错　
	// 2, 必须先接收再发送，并且接收方必须是 开启一个 goroutine 先去准备取值
	c := make(chan bool)
	 <- c    // 这里不是一个 goroutine, 并且又取不到值，所以会一直阻塞在这里，不会执行到 c <- true , 所以最终会报错
	 c <- true
}

func testSimpleChan2(){

	// 有缓冲通道， 三点注意：
	/*
	// 1，可以先发送，再接收, 
		  发送量超过容量，则会报错，
		  发送完毕，将通道关闭，则还可以从通道中取值，前提是通道里面还有值，没有值则会接收到 通道类型的零值
	// 2，通道没有值之前，就去接收值则会阻塞 报错
	// 3, 通道已关闭，则不能从里面发送值
	*/ 
	c := make(chan bool, 1)
	// <- c    // 这里不是一个 goroutine, 并且又取不到值，所以会一直阻塞在这里，不会执行到 c <- true , 所以最终会报错
	c <- true
	// c <- false  // 超过容量，则报错
	close(c)
	// c <- false  // "panic: send on closed channel"
}

func testExitGoroutine(){
	fmt.Println("main start")
	exitChan := make(chan bool)

	wg.Add(1)

	go worker(exitChan)

	// time.Sleep(time.Second*1)

	exitChan <- false    // 对于无缓冲区的通道来说， 发送值到通道之前，必须要有其它goroutine 从通道中取值的操作，所以他不能写在  go worker() 之前

	wg.Wait()

	fmt.Println("main over")

}

func worker(exitChan <-chan bool){

	defer wg.Done()
	LABEL:
	for {

		fmt.Println("worker ....")

		// 不用 select 多路复用，会在这里造成阻塞，直到从通道中拿到值
		// 如果长时间没有获取到值，则会报错
		/*
		exit := <- exitChan
		if exit {
			break
		}
		*/
		select {
		// case <- exitChan  只要能取到值，不管是0 还是1， true 还是 false, 只要能取到值就会继续往下执行
		case <- exitChan : 
			break LABEL
		default :
			fmt.Println("default ....")
		}
	
	}
}