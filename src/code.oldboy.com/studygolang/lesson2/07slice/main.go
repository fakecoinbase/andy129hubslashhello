package main

import "fmt"

// 切片是引用类型
func main() {
	a := []int{1, 2, 3}
	b := a
	b[0] = 100
	fmt.Println(a) // "[100 2 3]"
	fmt.Println(b) // "[100 2 3]"

	var c []int
	num := copy(c, a)
	// 因为 c 的长度和容量就是 0
	fmt.Println("元素拷贝成功的个数：", num) // "元素拷贝成功的个数： 0"
	fmt.Println(c)                 // "[]",   为什么没有拷贝成功呢？  原因是  var c []int ， 这只是变量声明，并没有初始化所以没有申请内存空间

	// d 虽然与 c 效果一致， 但是 d 被初始化了，只是 d 的长度和容量都为0 ， 当将 a 拷贝到 d 时，发现空间不足，无法拷贝
	var d = []int{}
	copy(d, a)
	fmt.Println(d)

	// 正确初始化
	var e = []int{0, 0}
	copy(e, a)     // 深拷贝，重新申请一块内存，把内存中的值拷贝一份过来， 这样 e, a 就互不影响了
	fmt.Println(e) // "[100 2]"     // 拷贝成功
	// 两块不同的内存地址，互不影响
	fmt.Printf("类型:%T, e:%v, ptr:%p\n", e, e, e) // "类型:[]int, e:[100 2], ptr:0xc000012150"
	fmt.Printf("类型:%T, a:%v, ptr:%p\n", a, a, a) // "类型:[]int, a:[100 2 3], ptr:0xc00000a400"

	// make 申请内存
	var f []int
	f = make([]int, 3, 3) // type, len , cap
	copy(f, a)
	fmt.Println(f) // "[100 2 3]"

}
