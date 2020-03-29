package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// select 多路复用，实现 监听终端输入
func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)
	scanChan := make(chan int, 100)

	// 产线开足马力生产
	go produce(jobs)

	// 分配20个工人，并指定 产线和 销售线
	// jobs : 产线， 工人从这里取手机
	// results : 销售线， 工人把从产线取到的手机进行包装，最后的成品放入到 销售线上。
	initWorkers(20, jobs, results)

	go checkResult(results)

	// 检测外部不可抗力的因素，例如：地震，停电，等情况 需要立即停止销售线
	go scanInput(scanChan)
	// go scanInput2(scanChan)

	// 上面 产线，工人，车间主任查看销售线，监听外部因素 全都开启了 goroutine 模式
	// select {} 语句块会堵塞在这里，一直监听 scanChan 通道。
	select {
	case <-scanChan:
		closeAllChan(jobs, results, scanChan)
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@close()")
		break
	}

	/*  // 注意上面 break 和 下面 return 的使用场景
	for {
		select {
		case <-scanChan:
			closeAllChan(jobs, results, scanChan)
			fmt.Println("####################################close()")
			return // 这里想要跳出 for 循环，不能使用 break , break 可以跳出 select, 外面还有一层 for 循环，所以程序不会终止
		case v := <-results:
			fmt.Println("成品：", v)
		default:
			time.Sleep(time.Millisecond * 100)
			fmt.Println("无所事事！")
		}
	}
	*/
}

// 产线生产
func produce(jobs chan int) {
	for {
		i := rand.Intn(100)
		time.Sleep(time.Millisecond * 500)
		jobs <- i
	}
}

// 创建工人数量
func initWorkers(num int, jobs <-chan int, results chan<- int) {
	// 工人干活
	for j := 0; j < num; j++ {
		go work(j, jobs, results)
	}
}

// 工人工作
func work(id int, jobs <-chan int, results chan<- int) {
	for v := range jobs {
		fmt.Printf("worker : %d start %d job!\n", id, v)
		results <- v * v // 从 jobs 通道中获取值，然后计算值的平凡，将结果发送到 通道 results 中
		fmt.Printf("worker : %d end %d job!\n", id, v)
	}
}

// 车间主任或者 销售经理，一直在 销售线上查看 成品结果
func checkResult(results chan int) {

	for v := range results {
		fmt.Println("成品：", v)
	}
}

// 监听终端输入事件 (方法1)
func scanInput(scanChan chan int) {
	var input int
	// fmt.Scanln(&input)
	fmt.Scanln(&input)
	fmt.Println("---------------------------------------------------ok")
	fmt.Println("---------------------------------------------------ok")
	fmt.Println("---------------------------------------------------ok")
	fmt.Println("---------------------------------------------------input : ", input)
	fmt.Println("---------------------------------------------------ok")
	fmt.Println("---------------------------------------------------ok")
	fmt.Println("---------------------------------------------------ok")
	fmt.Println("---------------------------------------------------ok")
	scanChan <- input
}

//  监听终端输入事件 (方法2)
func scanInput2(scanChan chan int) {

	var b [1]byte
	os.Stdin.Read(b[:])
	scanChan <- 100
}

func closeAllChan(jobs, results, scanChan chan int) {
	close(jobs)
	close(results)
	close(scanChan)
}
