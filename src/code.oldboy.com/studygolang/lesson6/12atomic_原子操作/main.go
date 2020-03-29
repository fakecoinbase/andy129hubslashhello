package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// atomic 原子操作

/*
	代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。
	针对基本数据类型我们还可以使用原子操作来保证并发安全，因为原子操作是Go语言提供的方法它在用户态就可以完成，
	因此性能比加锁操作更好。Go语言中原子操作由内置的标准库sync/atomic提供。
*/

/*  建议： atomic 使用场景较少， 我们一般使用  通道或者 sync 包的函数/类型 实现同步更好。
atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用。
这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，使用通道或者sync包的函数/类型实现同步更好。

*/

// Counter 是一个接口， 定义了 Inc() 增加， Load() 加载 两个方法
type Counter interface {
	Inc()
	Load() int64
}

// CommonCounter 是一个结构体，定义了一个 counter 变量
type CommonCounter struct {
	counter int64
}

// Inc 计数功能
func (c *CommonCounter) Inc() {
	c.counter = c.counter + 1
}

// Load 获取数值
func (c *CommonCounter) Load() int64 {

	return c.counter
}

// MutexCounter 是一个结构体， 定义了 counter , lock 两个变量
type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

// Inc 计数功能
func (m *MutexCounter) Inc() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

// Load 获取数值
func (m *MutexCounter) Load() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

// AtomicCounter 是一个结构体，定义了 counter 变量
type AtomicCounter struct {
	counter int64
}

// Inc 计数功能
func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

// Load 获取数值
func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func test(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}

	wg.Wait()
	end := time.Now()
	fmt.Printf("%d\t%v\n", c.Load(), end.Sub(start))
}

// sync/atomic 包
func main() {

	test1()

	// test2()
}

func test1() {
	// 为了达到修改结构体里值的目的，必须要传入指针类型
	c1 := CommonCounter{}
	test(&c1)

	c2 := MutexCounter{}
	test(&c2)

	c3 := AtomicCounter{}
	test(&c3)

	/*  上面的代码，我们运行了四次，四次的结果对比如下：
	// c1 得出的结果， 不是并发安全的 ,  c2,c3 安全
	// c2 加锁的操作 对比 c1,c3 最耗时
	// c3 既安全，又相对比较 耗时少。     （原子操作 胜出！！）

	9798963 		3.2572973s
	10000000        3.4454452s
	10000000        3.2763256s

	9801179 		3.2563114s
	10000000        3.3713933s
	10000000        3.262315s

	9798118 		3.2613145s
	10000000        3.3964104s
	10000000        3.2683196s

	9791337 		3.3003417s
	10000000        3.3964111s
	10000000        3.264317s

	*/
}

// 再次回顾以前的内容， 结构体是 值类型， 当调用一个函数修改结构体内部的值的时候，需要传入指针，否则 函数修改的只是一个副本
func test2() {
	c1 := CommonCounter{}
	testValueStruct(c1)

	// fmt.Println(c1.Load())
	fmt.Println(c1.counter)
}

func testValueStruct(c CommonCounter) {

	c.counter = 5
	fmt.Println(c.counter)
}
