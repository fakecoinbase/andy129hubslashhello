package main

import "fmt"

// slice 学习
func main() {
	var a = [3]int{1, 2, 3}
	var b = []int{1, 2, 3}

	fmt.Println(a, b)
	fmt.Printf("a:%T   b:%T\n", a, b) // "a:[3]int   b:[]int"

	// 切片底层是 数组，所以可以直接用下标 取值
	fmt.Println(b[1]) // "2"

	// 声明切片的方式：  从数组得到切片
	var c []int
	// 冒号切[下标值i:下标值j]，  左包含右不包含
	// 长度： j - i
	c = a[:]       // a[0:3]  ， 从头取到尾
	fmt.Println(c) // "[1,2,3]"
	d := a[1:2]
	fmt.Println(d) // "[2]"

	// 切片的大小（目前元素的数量）
	fmt.Println(len(b))
	// 容量(底层数组最大能放多少元素)
	x := [...]string{"北京", "上海", "深圳", "广州", "成都", "杭州", "重庆"}
	y := x[1:4]         // 切片之后，返回的就是 引用类型
	fmt.Println(y)      // "[上海 深圳 广州]"
	fmt.Println(len(y)) // "3"
	// 从切片后的第一个元素开始计算 容量
	fmt.Println(cap(y)) // "6"

	// 修改切片里元素的值，其实就修改了 底层数组 x,  修改了 x 则也会更新到 y ,   x,y 共用一个底层数组, (只要 x,y 之间涉及到了 [:] 切片操作)
	y[0] = "sz"
	fmt.Println(y)
	fmt.Println(x)
	fmt.Println("--------------------------------------------------")
	x[2] = "武汉"
	fmt.Println(y)
	fmt.Println(x)

}
