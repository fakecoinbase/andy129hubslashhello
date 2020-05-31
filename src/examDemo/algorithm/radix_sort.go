package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"time"
)

// 基数排序
func main() {

	radixSortDemo()

}

func radixSortDemo(){
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
	radixSort(data, 255)
	// 从开始时间到 结束时间， 时间间隔
	end := time.Since(start)
	fmt.Println("用时：", end.Seconds())   //  0.0105076

	fmt.Println(data)

}


// 基数排序
// 原理：
/*
	比较 列表中每个数字的个位，十位，百位，千位 等等
	比较个位，来一次 入桶过程，然后出桶， 然后再比较 十位，再来一次入桶过程，再出桶，然后再比较百位，再来一次入桶过程，然后再出桶。。。。。
	入桶过程就是 排序的过程，直到最后出桶时，所有的数字已经排序好了。

    为什么入桶的过程是 排序的过程，因为 分了10个桶，对应 (0,1,2,3,4,5,6,7,8,9) 个数字，
	列表中每个数字的每一位不会超过这几个数字，所以数字分在了哪个桶，就能直到 该元素的某一位 大小顺序了
 */
// 时间复杂度：O(nk),  k 取决于 maxNum   k = log(10, maxNum) 最大数值决定了 要迭代多少次
func radixSort(data []int, maxNum int) {

	it :=0    // 迭代多少次 (取决于 下面这个判断)
	for int(math.Pow(10, float64(it))) <= maxNum {   // 例如最大值为 255, 那么 it = 2 (迭代3次， 0,1,2, 几位数就迭代几次)

		buckets := make([][]int, 10)   // 创建10个桶 对应 (0,1,2,3,4,5,6,7,8,9)
		for i:=0;i<len(data);i++ {
			digit := (data[i]/ (int(math.Pow(10, float64(it)))) ) %10
			// 例如：180， 第1次，it==0, 180/1%10 取余为 0， 第二次, it==1, 180/10%10 取余为8，第三次，it==2, 180/100%10 取余为 1
			buckets[digit] = append(buckets[digit], data[i])   // 根据 如上计算，算出 digit (桶的标识),  存放到对应的桶中
		}

		// 清空 data, 为下面做存储准备 (下面更新了 元素的顺序)
		data = data[0:0]
		// 将每一次 排序后的 buckets 里面的元素 迭代出来，存放到 data 列表中
		for k:= range buckets {
			data = append(data, buckets[k]...)
		}

		it ++   // 进行下一位的迭代  (个位，中位，十位，百位等，依次进行下去)
	}

}