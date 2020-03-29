package main

import (
	"fmt"
	"time"
)

// select 多路复用
/*
	使用 select 语句能提高代码的可读性。

		. 可处理一个或多个 channel 的发送/接收操作。
		. 如果多个 case 同时满足, select 会随机选择一个.
		. 对于没有 case 的 select{} 会一直等待，可用于阻塞 main 函数。
		注意： select{} 会阻塞 main 函数，并不代表你可以直接这样写 空的 select{}
		 还是需要配上 case , default 语句，否则会报错

*/
func main() {

	// test1()

	// test2()

	// test3()

	test4()

	// select {} // "报错： goroutine 1 [select (no cases)]:"
}

// selct{} 使用定律1 ： 哪个条件满足，就执行哪一个
func test1() {
	// 设置一个容量为1 的带缓冲区通道
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select { // 以下条件，除去default, 每次只会满足一个条件，哪个条件满足就执行哪一个
		case x := <-ch:
			fmt.Println("取值：", x)
		case ch <- i:
			fmt.Println("存放值：", i)
		default:
			fmt.Println("啥都不干")
		}
	}

	/*  打印结果：
	存放值： 0
	取值： 0
	存放值： 2
	取值： 2
	存放值： 4
	取值： 4
	存放值： 6
	取值： 6
	存放值： 8
	取值： 8

	*/
}

//  selct{} 使用定律2 : 多个条件都满足时，随机执行一个
func test2() {
	// 定义一个容量为 10 的带缓冲的通道
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		select { // 以下条件，会出现 同时满足的情况， select 原则就是 随机执行某个满足的条件。
		case x := <-ch:
			fmt.Println("取值：", x)
		case ch <- i:
			fmt.Println("存放值：", i)
		default:
			fmt.Println("啥都不干")
		}
	}

	fmt.Println("------------------------------")
	fmt.Println(len(ch), cap(ch))

	/*  打印结果： 从打印结果可以看出， 即可以存放值，也可以取值的时候，它会随机执行 存值或者 取值的操作。
	存放值： 0
	取值： 0
	存放值： 2
	存放值： 3
	取值： 2
	存放值： 5
	取值： 3
	存放值： 7
	存放值： 8
	取值： 5
	*/
}

// select{} 中 default 的妙用：
func test3() {

	var ch chan string
	fmt.Println(ch) // "<nil>"

	var ch2 = make(chan string, 5)

	// 案例1：
	// 针对一个 <nil> 的通道取值时，会引发：fatal error :fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [select (no cases)]:
	/*
		fmt.Println("案例1：")
		select {
		case v := <-ch:
			fmt.Println("received value", v)
		}
	*/

	// 案例2：
	// 从一个既没有值，也没有关闭的通道里取值，会报错。 fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan receive]:
	/*
		fmt.Println("案例2：")
		select {
		case v2 := <-ch2:
			fmt.Println("received ch2 value", v2)
		}
	*/

	// 为什么要用 default :  (可以避免由于上面两种情况导致的 程序崩溃)
	// 运行结果： "默认输出"
	select {
	case v := <-ch:
		fmt.Println("received value", v)
	default:
		fmt.Println("默认输出")
	}

	// 运行结果： "默认输出"
	select {
	case v2 := <-ch2:
		fmt.Println("received ch2 value", v2)
	default:
		fmt.Println("默认输出")
	}

}

func test4() {

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go server1(ch1)
	go server2(ch2)

	// 模拟情况一：
	/*
		time.Sleep(time.Second) // 主程序等待一秒之后， server1() 和 server2() 都响应完成，所以进入 select{} 之后，两个条件都满足，则会随机打印

		select {
		case v := <-ch1:
			fmt.Println(v)
		case v2 := <-ch2:
			fmt.Println(v2)
		default:
			fmt.Println("暂时无法获取数据！")
		}
	*/

	// 模拟情况二：ch1, ch2 模拟两个服务器的响应通道，结果： 我们可以先拿到 响应快的服务器的通道里的值 (其他通道的就丢弃掉)
	// (不加 default 时，会阻塞在 case 那里，直到通道有值，取到值就退出 select
	// 注意 不加 default 会不会有风险
	select {
	case v := <-ch1:
		fmt.Println(v)
	case v2 := <-ch2:
		fmt.Println(v2)
		//default:
		//	fmt.Println("暂时无法获取数据！")
	}

}

func server1(ch1 chan int) {
	time.Sleep(time.Millisecond * 500) // 模拟服务器响应速度   (等待时间必须放在 将结果放入到通道之前，这样才能正确模拟请求服务的环境)
	ch1 <- 1000
}

func server2(ch2 chan int) {
	time.Sleep(time.Second) // 模拟服务器响应速度
	ch2 <- 2000
}
