package main

import "fmt"

// 检查 几个工人是否都干完活了
var doneChan chan struct{}

// work pool (工作池)
/*
	在工作中我们通常会使用可以指定启动的goroutine数量–worker pool模式，控制goroutine的数量，防止goroutine泄漏和暴涨。
*/
func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)
	doneChan = make(chan struct{}, 100)

	// 产线生产
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs) // 往通道里发布任务完了之后，关闭通道

	// 开启goroutine
	// 开启三个goroutine 从jobs通道里取值，然后将结果放入到 results中
	// 形象说明：一条生产线上产出5部手机， 分三个工人去将手机包装完成放到 另一条销售线上。
	// 假设生产线上已经将5部手机生产完毕， 5部手机都摆放在生产线上，并且产线关闭。
	// 此时三个工人就回去产线上拿手机进行包装，每次只能拿一个，包装完毕放到 销售线上才代表一次  完整的工作，
	// 那么 有的工人经验比较丰富，比较熟练，有的是新手，那么经验丰富的做的快，那么它做完一个之后，继续到产线上取手机进行包装。
	// 做的慢的再去拿的时候，要么没货了，要么还能拿一个。 再回到我们这个模型来看：

	// 产线为一个通道，负责生产手机，假设一次性生产完毕，通道关闭
	// 每个工人类似于 一个 goroutine， 每个 goroutine 的工作就是不停的去  产线上(通道里)取手机， 包装手机，然后放到销售线上 （另一个通道）
	// 有的做的快，有的做的慢，直到把产线上的手机全部取完
	// main() 函数 就类似于 车间主任或者 销售经理， 负责让 产线生产，让 工人包装，最后从 销售线上把 成品 拿到。

	// 工人干活
	for j := 0; j < 3; j++ {
		go worker(j, jobs, results)
	}

	// 车间主任监督
	go closeResult(doneChan, results)

	// 获取结果方式1： (都需要先把 results 通道关闭，否则会阻塞在这里 )
	/*
		for {
			v, ok := <-results   // ok 这个 bool值，只有当 results 通道关闭，并且里面的值被取完了之后，才会返回 false, 如果没有关闭，里面的值虽然也取完了，也会一直阻塞在这里
			if !ok {
				break
			}
			fmt.Println("从通道中获取最终结果：", v)
		}
	*/

	// 获取结果方式2： (都需要先把 results 通道关闭，否则会阻塞在这里 )
	for v := range results { // 和 方式1 一样，如果通道没有关闭，值也取完了，那么会一直阻塞在这里，除非 results 通道里面一直有值。
		fmt.Println("从通道中获取最终结果：", v)
	}

}

func worker(i int, jobs <-chan int, results chan<- int) {
	for v := range jobs {
		fmt.Printf("worker : %d start %d job!\n", i, v)
		results <- v * v // 从 jobs 通道中获取值，然后计算值的平凡，将结果发送到 通道 results 中
		fmt.Printf("worker : %d end %d job!\n", i, v)
	}
	doneChan <- struct{}{} // 从 生产线上已经拿不到手机了，代表这个工人的 工作就结束
}

func closeResult(doneChan chan struct{}, results chan int) {
	// 循环3次，代表有3个工人 (三个 goroutine) ，我们不关心产线上到底有多少东西，
	// 我们只关心这3个工人还 能不能从 产线上拿到手机，当每个人不能从产线上拿到手机，就代表这个工人工作结束，所以给他分配一个标识
	// 当三个工人的标志都能拿到，那就代表 三个工人的工作结束了，车间主任的监督工作(检查通道)可以关闭了， 销售线 上的通道可以关闭了。
	for i := 0; i < 3; i++ {
		<-doneChan
	}
	close(doneChan) // 关闭检查通道
	close(results)  // 关闭销售线的通道
}
