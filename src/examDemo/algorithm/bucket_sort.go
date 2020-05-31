package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

// 桶排序
func main() {
	bucketSortDemo()
}

func bucketSortDemo(){
	length := 100000

	// 初始化一个 字节切片，指定长度为  100000
	byteArr := make([]byte, length)  // 字节切片里的元素，默认都为 0
	rand.Read(byteArr)    // "crypto/rand"  包下的 Read([]byte) 可生成字节数组长度 个数字  (随机数, 满足一个字节的取值范围 0-255)
	// fmt.Println(byteArr)

	data := make([]int, length)
	for i:=0;i<length;i++ {
		data[i] = int(byteArr[i])    // 将字节强制转换为 int
	}

	start := time.Now()
	bucketSort(data, 100, 255)
	// 从开始时间到 结束时间， 时间间隔
	end := time.Since(start)
	fmt.Println("用时：", end.Seconds())   // 0.2791831

	fmt.Println(data)

}

// 桶排序
// 参数1：需要排序的列表；参数2：分多少个桶；参数3：列表中已知的最大值
// 应用
// 应用于：已知最大值
// 时间复杂度： 平均情况： O(n+k),  最坏情况：O(n方k)   （k 是根据 n 与 m 计算出来的 , m 为桶的个数,  k 大概表示 一个桶里平均能装多少个数，桶的容量）
// 空间复杂度：O(nk)
// 桶排序的表现取决于数据的分布。也就是需要对不同数据排序时采取不同的分桶策略。
func bucketSort(data []int, bucketNum, maxNum int){

	buckets := make([][]int, bucketNum)   // 初始化 bucketNum 个桶
	for i:=0;i<len(data);i++ {
		j:= data[i] / (maxNum/bucketNum)
		// 假如 bucketNum == 10, maxNum == 10000,  遍历 data[i]  == 10000， 此时 j = 100 ，会导致数组越界，所以我们在这里进行 -1 的操作，把它放到 99
		if j > bucketNum-1 {
			j = bucketNum-1
		}
		// 把 data 列表中的值分为 若干个桶，每个桶保存 一定范围的值
		buckets[j] = append(buckets[j], data[i])

		// 保持桶内的数据是 有序的
		for n := len(buckets[j])-1;n>0;n-- {
			if buckets[j][n] < buckets[j][n-1] {
				buckets[j][n], buckets[j][n-1] = buckets[j][n-1], buckets[j][n]
			} else {
				break
			}
		}
	}

	// 先清空  data
	data = data[0:0]
	// 遍历 桶的集合，把每个桶的数据拿出来 添加到 data  中
	for k:= range buckets {   // 遍历数组，如果只用一个参数接收，则是 index
		data = append(data,buckets[k]...)
		// append() 追加值，可以是单个数字，也可以是多个 数字，多个数字也就是 slice 时，需要添加 ... (系统是当做可变参数来处理)
	}

}
