package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var onlyOnce sync.Once

// goroutine 与 闭包
func main() {
	test1()
}

// goroutine 与 闭包一起使用会带来的 坑
func test1() {

	/*  问题坑：
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("hello : ", i)
			wg.Done()
		}()
	}
	wg.Wait()
	*/

	// 会出现多个相同的数据，与我们期望效果不一致, 为什么呢?
	// 因为 func() 匿名函数在这里是个 闭包， 匿名函数里面使用 i 变量 是需要到外部去找的，
	// 循环和这个闭包， 你把它当成两个分开的部分， 循环一直跑， 多个 goroutine 一起执行的时候，有的 goroutine 运行的快，有的慢，
	// 快的 goroutine 就能早一些拿到 i 的值， i 的值就 小一些， 慢的 goroutine 拿到的 i 值就大一些，当 for 循环足够快，其他的 goroutine 还在准备阶段，
	// for 循环已经循环完毕，那么其他 goroutine 拿到的 i 值，就是这个 循环到最后的 最大值。

	/*  打印结果：
	hello :  9
	hello :  10
	hello :  10
	hello :  10
	hello :  10
	hello :  10
	hello :  10
	hello :  10
	hello :  10
	hello :  10

	*/

	// 解决方法：
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			fmt.Println("hello : ", index)
			wg.Done()
		}(i) // 将外部变量的值传入到 func() 里， 这样就能保证每个 goroutine 拿到的 index 都是不一样的值
	}
	wg.Wait()

	// 打印结果：
	/*
		hello :  9
		hello :  2
		hello :  5
		hello :  7
		hello :  6
		hello :  0
		hello :  1
		hello :  8
		hello :  4
		hello :  3

	*/

}

func say(i int) {
	fmt.Println("say : ", i)
}

func speak(i int) func() {
	return func() {
		say(i)
	}
}

func hello() {

}

// sync.Once 执行带参数的方法， 闭包的巧妙使用
/*
	sync.Once其实内部包含一个互斥锁和一个布尔值，
	互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成。
	这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。
*/
func test2() {

	onlyOnce.Do(hello) // sync.Once Do() 里面参数为 函数名，  那么我要执行一个带参数的 函数怎么办呢?

	s := speak(10) // 调用一个 speak(10) 返回一个 闭包 s
	onlyOnce.Do(s)
}
