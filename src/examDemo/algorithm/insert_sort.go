package main

import "fmt"

// 插入排序
func main() {
	data := []int{12,31,14,51,16,33,0,9,17,25}

	insertSortA(data)
	fmt.Println(data)   // [0 9 12 14 16 17 25 31 33 51]
}

func insertSortA(data []int){

	// 从 data 的第二张牌开始 摸牌
	for i:=1;i<len(data);i++ {
		tmp := data[i]
		j := i-1   // 手里的牌的下标 （初始为 data 中的第一张牌）

		// 开始比较, 手里的牌与 摸到的牌 一一比较 (手里的牌是作为排序，手机的牌越来越多)
		for j>=0 && data[j] > tmp {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = tmp
	}
}
