package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生产者消费者模型
// 使用 goroutine 和 channel 实现一个简易的生产者消费者模型

// 生产者：产生随机数   math/rand

// 消费者：计算每个随机数的每个位的数字的和

// 1个生产者 20个消费者

type item struct {
	id  int64
	num int64
}

type result struct {
	item *item
	sum  int64
}

var itemChan chan *item
var resultChan chan *result

// 生产者做的工作
func producer(itemChan chan *item) {

	var id int64

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("-------err : ", err)
		}
	}()

	// 生产者一直进行着 产生随机数的工作
	for {
		id++
		number := rand.Int63() // int64 的正整数
		itemObj := &item{
			id:  id,
			num: number,
		}
		fmt.Println("-++-------写入：id : ", id)
		itemChan <- itemObj // 将生产的随机数传入通道
	}
}

// 计算随机数 每个位数的之和
func calc(num int64) int64 {

	// 例如： num = 123
	// 123%10 == 3
	// 12%10 = 2
	// 1%10 = 1
	var sum int64
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}
	return sum
}

// 消费者从通道里拿到 生产者产生的随机数，然后计算总和 把结果放入到 另一个通道里 (resultChan)
func consumer(itemChan chan *item, resultChan chan *result) {

	// 生产者一直在往一个缓冲通道里产生随机数 (注意该缓冲通道有固定容量)
	// 所以为了不让缓冲通道阻塞，爆满，所以 消费者也要一直从 通道里取值

	for itemObj := range itemChan { // 从生产者产生随机数的通道里 取值
		sum := calc(itemObj.num)
		resultObj := &result{
			item: itemObj,
			sum:  sum,
		}
		resultChan <- resultObj // 消费者将结果往 resultChan 通道里传值
	}
}

// 把 消费者传入resultChan 通道里的结果 取出来并打印
func printResult(resultChan chan *result) {
	for ret := range resultChan {
		fmt.Printf("id : %v, num : %v, sum : %v\n", ret.item.id, ret.item.num, ret.sum)

		time.Sleep(time.Microsecond * 500)
	}
}

func main() {

	// test1()

	testInt63n()
}

// 练习 生产者消费者模型
func test1() {
	// 设置生产者产生随机数 的通道的最大容量
	// len(itemChan)  (产生一个，长度加1， 被取走一个， 则长度减 1 )
	itemChan = make(chan *item, 100)
	resultChan = make(chan *result, 100)

	go producer(itemChan) // 创建一个 goroutine

	// 创建多个 消费者 goroutine
	for i := 0; i < 20; i++ {
		go consumer(itemChan, resultChan) // 参数1： 需要取值的通道， 参数2：需要将结果传入的另一个通道
	}

	printResult(resultChan)
}

// 测试 rand.Int63() 与 rand.Int63n() 的区别

// rand.Int63n() 源码说明：
/*
	// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n)
	// from the default Source.
	// It panics if n <= 0.
	func Int63n(n int64) int64 { return globalRand.Int63n(n) }
*/
func testInt63n() {
	for i := 0; i < 20; i++ {

		number := rand.Int63n(100) // 与 Int63() 不同之处， Int63n() 可以指定区间 [0,n) (左包含右不包含), 返回值都是 int64类型
		fmt.Println(number)
	}

}
