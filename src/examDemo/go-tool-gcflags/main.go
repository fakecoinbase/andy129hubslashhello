package main

import "fmt"

var p *int

var number [3]func()

var number2 [3]func()int

func main() {
	// test()

	// test2()

	// test3()

	// test4()

	// test5()

	test6()
}

// 闭包参考文档： https://www.jianshu.com/p/faf7ef7fbcf8

// 对比 test4()， 为什么呢?
// 如果想让打印 正常输出 (按照我们原本的意思打印)，那么闭包中使用的外部变量，最好传入进去
/*
	func(i int) {
         // 针对 i 进行操作
    }(i)
*/
func test6(){
	number2 = [3]func()int{}
	for i:=0;i<3;i ++ {
		j:= i   // move to heap : j ，那为什么 j 不是最后的 2 值呢，依旧还是 0,1,2 三个不同的值呢
		number2[i] = func()int{
			j++
			return j
		}
	}

	for index,v := range number2 {
		fmt.Println(index)  // number2[i],  给数组下标的赋值也就还是 0,1,2
		fmt.Println(v())
		fmt.Println()
	}

	/*  打印结果：
	    0
		1

		1
		2

		2
		3

	*/

	fmt.Println("+++++数组长度： ", len(number))
}

func test5(){
	number2 = [3]func()int{}
	for i:=0;i<3;i ++ {   // move to heap : i
		// j:= i
		number2[i] = func()int{
			i++
			return i
		}
	}

	for index,v := range number2 {
		fmt.Println(index)  // number[i],  给数组下标的赋值也就还是 0,1,2
		fmt.Println(v())
		fmt.Println()
	}

	/*  打印结果：
	    0
		4

		1
		5

		2
		6

	*/
	fmt.Println("=====数组长度： ", len(number))
}

func test4(){
	number = [3]func(){}
	for i:=0;i<3;i ++ {
		j:= i
		number[i] = func(){
			fmt.Println(j)   // 初步分析, fmt.Println() 语句会出现 : escapes to heap 的情况
			// 使用 go build -gcflags "-m -l" main.go  查看变量逃逸情况,
			// 发现 .\main.go:24:15: j escapes to heap
			// 但是并没有出现 test3() 中 i 的 move to heap 的情况, 所以预测 j 并没有被移入 heap 中，所以 j 的值还只是 值拷贝与 test3() 中 i 不同
		}
	}

	for index,v := range number {
		fmt.Println(index)  // number[i],  给数组下标的赋值也就还是 0,1,2
		v()
		fmt.Println()
	}

	/*  打印结果：
		0
		0

		1
		1

		2
		2

	*/

	fmt.Println("---数组长度： ", len(number))
}

func test3(){

	number = [3]func(){}
	for i:=0;i<3;i ++ {       // 使用 go build -gcflags "-m -l" main.go  查看变量逃逸情况 ：  .\main.go:55:6: moved to heap: i
		number[i] = func(){
			fmt.Println(i)
		}
	}

	for index,v := range number {
		fmt.Println(index)  // number[i],  给数组下标的赋值也就还是 0,1,2
		v()  // 都打印 3 (代表 i 值为 3)
		fmt.Println()
	}
	/*  打印结果：
		0
		3

		1
		3

		2
		3

	*/
	fmt.Println("数组长度： ", len(number))

}



/*
	// 描述 golang 中的 stack 和 heap 的区别，分别在什么情况下会分配 stack , 又在何时会分配到 heap 中
	// 资料参考：https://blog.csdn.net/u010853261/article/details/102846449

	// 通过命令查看 变量逃逸情况
	// 第一： 编译器命令
	/*
		可以看到详细的逃逸分析过程。而指令集 -gcflags 用于将标识参数传递给 Go 编译器，涉及如下：

		-m 会打印出逃逸分析的优化策略，实际上最多总共可以用 4 个 -m，但是信息量较大，一般用 1 个就可以了
		-l 会禁用函数内联，在这里禁用掉 inline 能更好的观察逃逸情况，减少干扰

		go build -gcflags "-m -l" main.go

*/

// 第二：反编译命令查看
// go tool compile -S main.go

func test2(){
	var f = add()
	fmt.Println(f(1))   // 1
	fmt.Println(f(2))   // 3
	fmt.Println(f(3))   // 6

	// 局部变量x 在外部还会被引用，所以就被分配到了 heap(堆空间) 中，所以才会出现 值累加的情况
}

func add() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
	// 使用 go build -gcflags "-m -l" main.go  查看变量逃逸情况, 发现 局部变量 x 在外部还会被引用，所以就被移动到了 heap (堆)中, 变量 x 值不会被释放
	// .\main.go:39:6: moved to heap: x
}

func test(){
	var a [1]int
	_ = a[:]

	var b = 10

	p = &b
}

/*  使用命令查看 变量是否逃逸： go build -gcflags "-m -l" main.go

	go build -gcflags "-m -l" main.go
	# command-line-arguments
	.\main.go:12:6: moved to heap: b

 */