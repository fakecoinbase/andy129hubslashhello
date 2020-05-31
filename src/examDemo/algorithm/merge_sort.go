package main

import "fmt"

// 归并排序  ( 时间复杂度: O(nlogn),  空间复杂度：O(n) 递归函数也占用空间 O(logn)，但是由于 O(n)大，取最大值为 空间复杂度)
func main() {

	// testCopyAndAppend()  // // 回顾 append 和 copy 的用法

	mergeSortDemo()    // 归并排序原理

}

// 回顾 append 和 copy 的用法
func testCopyAndAppend(){
	data := []int{1,2,3,4,5}
	data2 := []int{6,7,8,9}

	// copy 用法
	// copy(data, data2)
	// fmt.Println(data)   // [6 7 8 9 5]

	// append 用法
	for i:=0;i<len(data2);i++ {
		data = append(data, data2[i])
	}
	fmt.Println(data)   // [1 2 3 4 5 6 7 8 9]

}

func mergeSortDemo(){
	data := []int{2,5,3,6,1,4,7,8,0,9}

	mergeSort(data,0, len(data)-1, "start")

	fmt.Println(data)
}

// 参数1：需要排序的列表(必须是 有两段排序好的列表的合集， 例如：[3,7,8,0,2,4,5,6,9] 中的 3,7,8, 和 0,2,4,5,6,9 )
// 参数2：起始位置;  参数3：中间位置; 参数4： 末尾位置
func merge(data []int, low, mid, high int){

	fmt.Printf("----------merge, low : %d, mid : %d, high : %d\n", low,mid,high)
	i := low
	j := mid+1

	tmpArr := make([]int,0,10)   // 创建了临时空间，增加了空间复杂度

	// 两段元素 分别对比大小
	for i<= mid && j<= high {
		if data[i] < data[j] {
			tmpArr = append(tmpArr, data[i])   // 将 锁定的元素存入临时列表中
			i++
		}
		if data[i] > data[j] {
			tmpArr = append(tmpArr, data[j])
			j++
		}
	}

	// 当跳出 以上 for 循环之后，说明 其中肯定是有一段列表 已经没有值了，所以我们就把另外一段 的元素依次添加到 临时列表中，此时 排序已完成
	for i<= mid {
		tmpArr = append(tmpArr, data[i])
		i++
	}
	for j<=high {
		tmpArr = append(tmpArr, data[j])
		j++
	}

	copy(data[low:high+1], tmpArr)   // 将排序好的列表 赋值到 data 指定长度中，此时 data 中的这段长度的列表就是有序的了
	// fmt.Println(tmpArr)
}

// 时间复杂度: O(nlogn),  空间复杂度：O(n)
func mergeSort(data []int, low,high int, flag string){

	fmt.Println("----mergeSort : ", flag)
	if low < high {
		fmt.Println("----enter : ", flag)
		mid := (low + high)/2
		point := fmt.Sprintf("left (%d,%d)", low,mid)
		mergeSort(data, low, mid, point)

		point2 := fmt.Sprintf("right (%d,%d)", mid+1,high)
		mergeSort(data, mid+1, high, point2)
		merge(data, low, mid, high)
		// fmt.Println(data[low:high+1])
	}
}

/*
----mergeSort :  start
----enter :  start
----mergeSort :  left (0,4)
----enter :  left (0,4)
----mergeSort :  left (0,2)
----enter :  left (0,2)
----mergeSort :  left (0,1)
----enter :  left (0,1)
----mergeSort :  left (0,0)
----mergeSort :  right (1,1)
----------merge, low : 0, mid : 0, high : 1
----mergeSort :  right (2,2)
----------merge, low : 0, mid : 1, high : 2
----mergeSort :  right (3,4)
----enter :  right (3,4)
----mergeSort :  left (3,3)
----mergeSort :  right (4,4)
----------merge, low : 3, mid : 3, high : 4
----------merge, low : 0, mid : 2, high : 4
----mergeSort :  right (5,9)
----enter :  right (5,9)
----mergeSort :  left (5,7)
----enter :  left (5,7)
----mergeSort :  left (5,6)
----enter :  left (5,6)
----mergeSort :  left (5,5)
----mergeSort :  right (6,6)
----------merge, low : 5, mid : 5, high : 6
----mergeSort :  right (7,7)
----------merge, low : 5, mid : 6, high : 7
----mergeSort :  right (8,9)
----enter :  right (8,9)
----mergeSort :  left (8,8)
----mergeSort :  right (9,9)
----------merge, low : 8, mid : 8, high : 9
----------merge, low : 5, mid : 7, high : 9
----------merge, low : 0, mid : 4, high : 9
[0 1 2 3 4 5 6 7 8 9]

 */