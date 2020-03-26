package main

import "fmt"

// slice 补充
func main() {

	/*
		var a []string
		var b []int
		var c = []bool{true, false}

		fmt.Println(a) // "[]"
		fmt.Println(b) // "[]"
		fmt.Println(c) // "[true false]"
	*/

	fmt.Println("--------------------切片的获取方式---------------------")

	// 基于数组得到切片
	a := [5]int{55, 56, 57, 58, 59}
	b := a[1:4]
	fmt.Println(b)        // "[56 57 58]"
	fmt.Printf("%T\n", b) // "[]int"

	// 切片再切片
	c := b[0:len(b)]
	fmt.Println(c)        // "[56 57 58]"
	fmt.Printf("%T\n", c) // "[]int"

	// make 函数构造切片
	d := make([]int, 5, 10)
	fmt.Println(d)        // "[0 0 0 0 0]"
	fmt.Printf("%T\n", d) // "[]int"
	fmt.Println(len(d))   // "5"
	fmt.Println(cap(d))   // "10"

	fmt.Println("--------------------切片的 nil 比较---------------------")

	// 切片不能直接比较, 例如: == ,  切片唯一合法的比较运算符是 与  nil 进行比较
	// 声明 int 类型的切片
	var k []int
	fmt.Println(k, len(k), cap(k)) // "[] 0 0"
	fmt.Println(k == nil)          // "true"

	// 一旦带 {} 则必须是要带 = ， 并且 在 = 右边
	// 声明并初始化
	var j = []int{}
	fmt.Println(j, len(j), cap(j)) // "[] 0 0"
	fmt.Println(j == nil)          // "false"

	// make 创建
	m := make([]int, 0)
	fmt.Println(m, len(m), cap(m)) // "[] 0 0"
	fmt.Println(m == nil)          // "false"

	// 所以，我们判断数组是不是为空，一般不用 nil 来判断，而是判断长度 len

	fmt.Println("--------------------切片的赋值拷贝---------------------")
	index1 := make([]int, 3)
	index2 := index1 // 共用同一块内存空间

	index2[0] = 100
	fmt.Println(index2) //"[100 0 0]"
	fmt.Println(index1) //"[100 0 0]"

	fmt.Println("--------------------切片的遍历---------------------")
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}

	fmt.Println()
	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Println("--------------------切片的扩容---------------------")
	// 注意:
	var p []int
	// p[0] = 100        // 运行报错：panic: runtime error: index out of range [0] with length 0
	fmt.Println(p)
	// var p []int   这行代码只是对变量 p 进行了声明并没有初始化，所以没有申请内存，里面的元素为0，所以当进行 p[0] 时，发现找不到元素，数组越界了。

	// 可以使用 append()进行初始化
	var q []int
	q = append(q, 10)
	// 可追加多个元素
	q = append(q, 12, 34, 4, 34, 343)
	fmt.Println(q)

	// q = append(q, a...) // 编译报错： cannot use a (type [5]int) as type []int in append
	fmt.Println(q)
	q = append(q, a[:]...)
	fmt.Println(q) // "[10 12 34 4 34 343 55 56 57 58 59]"

	/*  查看 append() 源码, 第二种情况,针对第二个参数的介绍：  anotherSlice...
	//	slice = append(slice, elem1, elem2)
	//	slice = append(slice, anotherSlice...)

	// 由于 a 是指定长度的数组，非 slice 类型，所以编译报错
		q = append(q, a...)   // 错误
		但可通过这样来写：
		q = append(d, a[:]...)   // 转换为 slice 类型
	*/

	// for 循环打印出 append()扩容前后的  len 与 cap 的关系
	for i := 0; i < 10; i++ {
		// 省略
	}

	fmt.Println("--------------------切片的copy---------------------")
	// 切片的 copy
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 5, 5)
	s3 := s2
	copy(s2, s1)
	s2[0] = 100
	fmt.Println(s2) // "[100 2 3 4 5]"
	fmt.Println(s1) // "[1 2 3 4 5]"
	fmt.Println(s3) // "[100 2 3 4 5]"
	// fmt.Println(s2 == s3)  // 编译错误：invalid operation: s2 == s3 (slice can only be compared to nil)
	// 再次说明，slice 之间不能直接进行比较，唯一能使用 比较运算符的是和 nil  比较

	fmt.Println("--------------------切片的删除---------------------")
	x := []string{"北京", "上海", "广州", "深圳"}
	x = append(x[0:2], x[3:]...)
	fmt.Println(x) // "[北京 上海 深圳]"

	// 删除下标为index的元素的公式： append(a[:index], a[index+1:]...)

}
