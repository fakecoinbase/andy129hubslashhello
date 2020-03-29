package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)

var wg sync.WaitGroup

var rwLock sync.RWMutex

var syncMap sync.Map

// go 语言中内置的 map 不是并发安全的。不能进行 并发写操作
// 为了解决这个问题，我们有两种方法：
// 1, test2() 针对map 的读写操作，添加 读写锁
// 2, test4() 使用内置包  sync.Map
/*
	Go语言的sync包中提供了一个开箱即用的并发安全版map–sync.Map。
	开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。
	同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。
*/

func main() {
	// test1()

	// test2()

	// test3()

	test4()
}

// 并发 对 map 进行读写操作
// 会报错误： fatal error: concurrent map writes   （并发针对map 进行写操作）
// go 语言中内置的 map 不是并发安全的。不能进行 并发写操作
func test1() {
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func get(key string) int {

	v := m[key]
	return v
}

func set(key string, value int) {
	m[key] = value
}

// 添加 读写锁(或者 互斥锁也行)，下面示例 读与写的操作 相差不多。
// 添加锁之后，对 map 进行读写操作 正常。
func test2() {

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			setbyLock(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, getbyLock(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func getbyLock(key string) int {
	rwLock.RLock()
	v := m[key]
	rwLock.RUnlock()
	return v
}

func setbyLock(key string, value int) {
	rwLock.Lock()
	m[key] = value
	rwLock.Unlock()
}

// 上面的例子 验证了对 map 进行并发写操作有问题，那么 并发读 会不会也有问题呢
// 目前来看，并发读的时候没有报错
func test3() {
	m = make(map[string]int, 1000)

	for i := 0; i < 1000; i++ {
		key := strconv.Itoa(i)
		m[key] = i
	}

	wg.Add(10)
	go readMap()
	go readMap()
	go readMap()
	go readMap()
	go readMap()
	go readMap()
	go readMap()
	go readMap()
	go readMap()
	go readMap()
	wg.Wait()

}

func readMap() {

	for i := 0; i < len(m); i++ {
		key := strconv.Itoa(i)
		v := m[key]
		fmt.Println(v)
	}
	wg.Done()
}

// 使用内置包 sync.Map
/*
	Go语言的sync包中提供了一个开箱即用的并发安全版map–sync.Map。
	开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。
	同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。
*/
func test4() {

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			syncMap.Store(key, n) // 使用 sync.Map 进行并发写操作，不会出现问题
			v, _ := syncMap.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, v)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
