package main

import (
	"fmt"
)

// slice 扩容
func main() {

	test()

	var a = []int{}
	fmt.Println(a, len(a), cap(a)) // "[] 0 0"

	for i := 0; i < 5; i++ {
		// append() 函数是往切片中追加元素。
		a = append(a, 1)
		fmt.Printf("类型:%T, a:%v\t, len:%d, cap:%d, ptr:%p\n", a, a, len(a), cap(a), a) // %p 打印切片地址
	}

	/*  // 只要发生了扩容，地址都发生改变
	类型:[]int, a:[1]      	 	, len:1, cap:1, ptr:0xc0000120e8
	类型:[]int, a:[1 1]     	, len:2, cap:2, ptr:0xc000012140
	类型:[]int, a:[1 1 1]   	, len:3, cap:4, ptr:0xc00000a420
	类型:[]int, a:[1 1 1 1]		, len:4, cap:4, ptr:0xc00000a420
	类型:[]int, a:[1 1 1 1 1]   , len:5, cap:8, ptr:0xc000010140
	*/

	// 切片三要素:
	/*
		-- 地址(切片中第一个元素指向的内存空间)
		-- 大小(切片中目前元素的个数)              len()
		-- 容量(底层数组最大能存放的元素的个数)     cap()
	*/

	// 切片支持自动扩容
	/*  扩容策略：
	每一次都是上一次的 2 倍。
	*/
}

func test() {

	var a = make([]int, 4, 5)
	fmt.Println(len(a), cap(a)) // "4 5"
}
