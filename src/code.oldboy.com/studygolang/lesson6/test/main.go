package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 通道练习题
func main() {

	test1()
	// testRand()
	// testRand2()
}

func test1() {
	var max = 20

	ch := make(chan int, max)

	/*
		defer func() {
			close(ch)
		}()
	*/

	wg.Add(1)
	go randomNumber(ch, max) // 创建一个goroutine(类似于生产者) 往通道里产生随机数

	for i := 0; i < max; i++ { // 创建多个goroutine(类似于多个消费者) 从通道里取值
		wg.Add(1)
		go getNumber(ch)
	}

	wg.Wait()
}

// 往通道里产生随机数
func randomNumber(ch chan int, max int) {

	rand.Seed(time.Now().UnixNano()) // 创建随机数因子，只需要定义一次，无需再产生随机数之前频繁调用。

	for i := 0; i < max; i++ {
		// rand.Seed(time.Now().UnixNano())     // 不能再这里调用，否则将会产生大量相同的随机数
		number := rand.Intn(100)
		fmt.Println("产生数字：", number)
		ch <- number // 产生一个随机数 就往通道里传入
	}
	close(ch)
	wg.Done()

}

func getNumber(ch chan int) {
	v, ok := <-ch
	if !ok { // 如果取不到值，则代表通道已关闭，不能再进行取值行为,否则只会取到 零值
		return
	}

	fmt.Println("获取数字：", v)
	wg.Done()
}

// 另一种产生随机数的 写法
func testRand() {
	// rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000)

	source := rand.NewSource(time.Now().UnixNano())
	randObj := rand.New(source) // 不能放到循环里面，只需要在外部 调用一次就行了，否则将会生成大量重复的 随机数

	for i := 0; i < 10; i++ {

		number := randObj.Intn(100)
		fmt.Println(number)
	}

}

func testRand2() {
	for i := 0; i < 10; i++ {

		number := rand.Intn(100)
		fmt.Println("无需随机因子：", number)
	}

	/*   每次执行程序调用这个 testRand2() 方法都会生成下面这几个 随机数 (每次都相同)

	无需随机因子： 81
	无需随机因子： 87
	无需随机因子： 47
	无需随机因子： 59
	无需随机因子： 81
	无需随机因子： 18
	无需随机因子： 25
	无需随机因子： 40
	无需随机因子： 56
	无需随机因子： 0

	*/
}

/*
	何时需要 设置随机因子：
	1, 如果不设置 随机因子， 在 for 循环中 ，每次执行 randObj.Intn() 都会产生不一样的随机数
		但是 你这个程序，每次执行 这个 for 循环，则都会得到 相同的 那几个随机数
	2, 如果你想让 程序每次执行 for 循环，都会产生不一样的打印结果，则需要设置 随机因子
*/
