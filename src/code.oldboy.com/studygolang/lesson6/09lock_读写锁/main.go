package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex   // 互斥锁
	rwLock sync.RWMutex // 读写锁
)

// 读写锁
// ## 需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。

// 本程序测试时间： 添加互斥锁，大概耗时 12秒,   添加 读写锁之后， 大概耗时 2 秒，所以可见： 在读多写少的场景下， 读写锁的优势还是很大的。
func main() {

	now := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()

	duration := time.Since(now)
	fmt.Println(duration)
}

func read() {
	//lock.Lock()// 互斥锁
	rwLock.RLock() // 读锁 (读锁和 之前的互斥锁 方法名不一样了，具体是什么机制，以后再研究)
	time.Sleep(time.Millisecond * 10)
	rwLock.RUnlock()
	//lock.Unlock()
	wg.Done()
}

func write() {
	//lock.Lock() // 互斥锁
	rwLock.Lock() // 写锁 （目前来看，写锁和互斥锁的写法一样，大概就是 我写的时候 谁都不能进来）
	x = x + 1
	time.Sleep(time.Millisecond * 200)
	rwLock.Unlock()
	//lock.Unlock()
	wg.Done()
}
