package main

import "fmt"

func main() {

	test1()
	// test2()
	fmt.Println("-------------------------------------------------------------")
	testYouhua()
}

// 打印 200-1000 之间的水仙花数 (素数)
func test1() {
	b := true
	for i := 200; i <= 1000; i++ {
		for j := 2; j < i; j++ { // j 除去1 和 自身
			if i%j == 0 { // 当发现有被整除的情况，则一定不是素数
				b = false // b 置为 false 并跳出
				break
			} else { // 这里一定要记住，当不能被整除的时候要将值 赋值为 true, 只有
				b = true
			}
		}
		// 一个n 在与 2 ~ n-1 之间的数进行整除之后，发现其中没有数能够整除，则视为 素数。
		if b {
			fmt.Printf("%d ", i)
		}
	}
}

// 与 test1() 做比较，注意我们将 标志位 b  放到了 第一层循环中，就省略了 else 里面 b = true 的代码
func testYouhua() {

	for i := 200; i <= 1000; i++ {
		b := true
		for j := 2; j < i; j++ { // j 除去1 和 自身
			if i%j == 0 { // 当发现有被整除的情况，则一定不是素数
				b = false // b 置为 false 并跳出
				break
			}
		}
		// 一个n 在与 2 ~ n-1 之间的数进行整除之后，发现其中没有数能够整除，则视为 素数。
		if b {
			fmt.Printf("%d ", i)
		}
	}
}

// 9*9 乘法表
func test2() {

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d  ", j, i, j*i)
		}
		fmt.Println()
	}

	/*  打印信息：

	1x1=1
	1x2=2  2x2=4
	1x3=3  2x3=6  3x3=9
	1x4=4  2x4=8  3x4=12  4x4=16
	1x5=5  2x5=10  3x5=15  4x5=20  5x5=25
	1x6=6  2x6=12  3x6=18  4x6=24  5x6=30  6x6=36
	1x7=7  2x7=14  3x7=21  4x7=28  5x7=35  6x7=42  7x7=49
	1x8=8  2x8=16  3x8=24  4x8=32  5x8=40  6x8=48  7x8=56  8x8=64
	1x9=9  2x9=18  3x9=27  4x9=36  5x9=45  6x9=54  7x9=63  8x9=72  9x9=81

	*/

}
