package main

import "fmt"

// 希尔排序 (是一种分组插入排序算法)  时间复杂度: 视gap值 的情况而定 （本文的 gap 就是 N/2, N/4 ... 直到 等于1），最坏时间复杂度为：O(n平方)
/*  原理：

	第一步：首先取一个整数 d = n/2 （n 为数组长度）, 将元素分为 d 个组，每组相邻元素之间距离为 d, 在各组内进行直接插入排序;
	第二步：取第二个整数 d = d/2, 重复上述分组排序过程，直到 d =1, 即所有元素在同一组内进行直接插入排序。
	第三步：希尔排序每趟并不使某些元素有序，而是使整体数据越来越接近有序；最后一趟排序使得所有数据有序。

 */
func main() {
	data := []int{5,7,4,6,3,1,2,9,8}

	shellSort(data)
	fmt.Println(data)
}


// 希尔排序是一种 分组插入排序算法，基于 插入排序
func insertSortGap(data []int, gap int){
	for i:=gap;i<len(data);i++ {
		tmp := data[i]   // 摸到的牌 (想象把摸到的牌取出，则空出一个位置)
		j:= i -gap
		for j>=0 && data[j] > tmp{   // 这里的操作就是 将手里的牌 与 摸到的牌比较，如果大于，则大的牌往右边移动一个位置
			data[j+gap] = data[j]   // 移动位置
			j = j-gap
		}
		data[j+gap] = tmp   // 插入到合适的位置
	}
}
// 希尔分组原理:  d = n/2的k次方 , 直到 d = 1  ( n 除2 分一次组，然后再除2 再分一次组，然后再除2 再分一次组，直到 d == 1)
func shellSort(data []int){
	d := len(data)/2
	for d>=1 {
		insertSortGap(data, d)
		d = d/2
	}
}




// 插入排序 回顾
func insertSort(data []int){
	for i:=1;i<len(data);i++ {
		tmp := data[i]   // 摸到的牌 (想象把摸到的牌取出，则空出一个位置)
		j:= i -1
		for j>=0 && data[j] > tmp{   // 这里的操作就是 将手里的牌 与 摸到的牌比较，如果大于，则大的牌往右边移动一个位置
			data[j+1] = data[j]   // 移动位置
			j--
		}
		data[j+1] = tmp   // 插入到合适的位置
	}
}
