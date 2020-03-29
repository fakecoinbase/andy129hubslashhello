package main

import (
	"fmt"
	"sync"
)

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex // 互斥锁
)

func main() {

	// test1()   // 演示为什么要加锁

	test2() // 特殊情况下， 加锁解锁的正确方式
}

func test2() {

	getValue()
}

// 下面这种情况，语句很短，只有一句 return x,  如何加锁解锁呢
func getValue() int64 {
	lock.Lock()
	defer lock.Unlock() // 正确，使用 defer 语句，保证 加锁和解锁 能 成对执行。
	return x
	// lock.Unlock()    // 编译错误，放在了 return 后面，错误
}

// 为什么要加锁？
func test1() {
	wg.Add(2)
	// 开启了两个 goroutine 都去争着修改 全局变量 x
	// go add()
	// go add()

	// 给全局变量加 互斥锁
	go addByLock()
	go addByLock()
	wg.Wait()

	fmt.Println(x)
}

// 未加锁
func add() {
	for i := 0; i < 50000; i++ {
		x = x + 1
	}
	wg.Done()
}

// 添加互斥锁
func addByLock() {
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
