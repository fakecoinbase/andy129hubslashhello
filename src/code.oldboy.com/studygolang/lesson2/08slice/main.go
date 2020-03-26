package main

import "fmt"

// slice 中删除某个元素
func main() {

	a := []string{"北京", "上海", "深圳", "广州"}
	// 把a[1] 元素删掉
	// a = append(a[:1], a[2:]) // 编译报错：cannot use a[2:] (type []string) as type string in append
	// append([]string, string)  参数形式
	a = append(a[:1], a[2:]...) // a[2:]...,  这种写法就是把 每一个元素 追加到 a[:1] 里
	fmt.Println(a)              // "[北京 深圳 广州]"

	// 得到切片的四种方式
	/*
		-- 直接声明 a := []int{1,2,3}
		-- 基于数组得到的切片 m := [3]int{1,2,3},  b := m[:]
		-- 基于切片得到切片 m := [3]int{1,2,3},  b := m[:] , c := b[:2]
		-- 通过make() 创建一个指定 len ,cap 的 slice，  f := make([]int, 3, 3)
	*/

}
