package main

import "fmt"

// 选择排序
// 时间复杂度： O(n的平方)
func main() {
	data := []int{12,31,14,51,16,33,0,9,17,25}

	selectSort(data)
	fmt.Println(data)   // [0 9 12 14 16 17 25 31 33 51]
}

func selectSort(data []int) {

	// 与冒泡排序一样，同样需要 n-1 趟比较
	for i:=0;i<len(data)-1;i++ {
		minIndex := i    // 记录最小数的下标
		// 无序区比较
		for j:=i+1;j<len(data);j++ {
			if data[j] < data[minIndex] {
				minIndex = j   // 更新最小值的下标
			}
		}
		// 每比较一趟data, 则将最小值更新到 有序区
		data[i], data[minIndex] = data[minIndex], data[i]
	}
}
