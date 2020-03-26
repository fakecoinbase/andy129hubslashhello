package main

import "fmt"

// defer: 延迟执行
/*
	由于 defer 语句延迟调用的特性，所以 defer 语句能非常方便的处理 资源释放问题。
	比如：资源清理、文件关闭、解锁及记录时间等。
*/
func main() {
	testDefer()
	/* 打印结果：
	start...
	end...
	3
	2
	1
	*/
}

// 代码按顺序执行，先执行fmt.Println("start...")
//  当发现 defer 的时候，会把 defer 语句放到一边，可以想象成一个栈,
// fmt.Println(1) 被压入了栈，然后 fmt.Println(2) 也被压入了栈，然后 fmt.Println(3) 也被压入了栈
// 然后继续往下面走，执行了 fmt.Println("end..."), 当 函数快调用结束的时候，再把 defer 语句拿出来执行
// 此时会从 栈里拿出第一条，知道栈的原理，就知道 现在栈顶的一条语句是 最后被压入到 栈的，也就是 fmt.Println(3)
// 然后依次取出栈里的其他语句， fmt.Println(2),  fmt.Println(1)
func testDefer() {
	fmt.Println("start...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end...")
}
