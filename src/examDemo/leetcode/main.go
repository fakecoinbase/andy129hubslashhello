package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	// testLeetCode()

	// testLeetCode2()
	testLeetCode3()

}

// 如果a+b+c=1000，且a*a+b*b=c*c（a,b,c为自然数），如何求出所有a,b,c可能的组合？不使用数学公式和math包，代码实现，并且尽可能保证算法效率高效
// 考核点：for 循环内套 (时间复杂度)
func testLeetCode3(){

	start := time.Now()

	/*  普通版本
	// 用时： 912.1473ms
	for a:=0;a<1001;a++ {
		for b:=0;b<1001;b++ {
			for c:=0;c<1001;c++ {
				if a+b+c == 1000 && a*a+b*b == c*c {
					fmt.Println(a,b,c)
				}
			}
		}
	}
	*/

	// 优化版本1：
	// 用时： 1.4973ms
	/*
	for a:=0;a<1001;a++ {
		for b:=0;b<1001;b++ {
			c := 1000 -a -b           // 去除 c 的 for 循环，由 a,b 可得知 c 的值
			if a*a+b*b == c*c {
				fmt.Println(a,b,c)
			}
		}
	}
	*/

	// 优化版本2：
	// 用时： 1.0007ms
	for a:=0;a<1001;a++ {
		for b:=0;b<1001-a;b++ {    // a 与 b 的关系是， 当a == 0 时， b 的范围在[0,1000], 当a == 200时，b 的范围在 [0,800], 不会超过 1000
			c := 1000 -a -b
			if a*a+b*b == c*c {
				fmt.Println(a,b,c)
			}
		}
	}


	end := time.Since(start)
	fmt.Println("用时：", end)
}



// 两个协程，交互打印 1-20个自然数，一个打印奇数，一个打印偶数
// 考核点：协程的使用， 无缓冲通道注意事项
func testLeetCode2(){

	ch1 :=  make(chan int)   // 无缓冲通道

	ch2 :=  make(chan int)   // 无缓冲通道

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i:=1;i<=10;i++ {
			ch1 <- 2*i-1    // 奇数
			fmt.Println("打印偶数：", <-ch2)
		}

	}()

	go func() {
		defer wg.Done()
		for i:=1;i<=10;i++ {
			fmt.Println("打印奇数：", <-ch1)    // 由于定义的是无缓冲通道，所以必须先从 ch1 通道中取值,否则会造成死锁
			ch2 <- 2*i    // 奇数
		}
	}()

	wg.Wait()

}



// 从一个数组中找出 只出现一次的数字
// 考核点：异或操作
func testLeetCode(){

	nums := []int{1,2,2,4,1}
	single := singleNumber(nums)

	fmt.Println()
	fmt.Println("result : ", single)
}

func singleNumber(nums []int) int{
	ret := nums[0]
	for i:=1;i< len(nums);i++ {
		ret ^= nums[i]   // 从第一个数字开始与后面的数字 异或操作，然后返回结果给 ret , 然后继续与下一个元素进行 异或操作
		fmt.Println(ret)

		/*  打印结果：
		    3
			1
			5
			4
		*/
	}
	return ret
}

