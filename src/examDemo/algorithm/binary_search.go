package main

import (
	"fmt"
	"sort"
)

// 二分查找  (二分查找需要 先将data 排序好)
func main() {
	data := []int{12,31,14,51,16,33,0,9,17,25}
	sort.Ints(data)

	index := binarySearchA(data, 25)   // 返回的是 排序后的 数值对应的下标
	fmt.Println(index)
}

func binarySearchA(data []int, val int) int{
	left := 0
	right := len(data)-1
	for left <= right {
		mid := (left+right) /2
		if data[mid] == val {
			return mid
		}else if data[mid] > val {
			right = mid -1
		}else {
			left = mid +1
		}
	}
	return -1
}
