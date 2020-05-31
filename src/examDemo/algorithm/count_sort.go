package main

import (
	"crypto/rand"
	"fmt"
	"sort"
	"time"
)

// 计数排序
func main() {
	// countSortDemo()   // 计数排序原理

	 countSortDemo2()  // 计数排序 尝试排序 无重复数组 (经过测试，正常)

	// countSortDemo3()     // 测试网上 计数排序实现方式 , 有局限性
}

// 网络参考： (局限：只能针对 无重复元素的数组，并且开辟了与原数组一样长度的 slice, 占用资源)
/*
func countingSort(theArray[] int)[]int{
	lastArray := make([]int, len(theArray))
	for i := 0; i<len(theArray); i++ {
		count := 0
		for j :=0; j<len(theArray); j++ {
			if theArray[i] > theArray[j]{
				count ++
			}
		}
		lastArray[count] = theArray[i]
	}
	return lastArray
}

func countSortDemo3(){
	data := []int{6, 4, 5, 1, 6, 3, 2, 3}
	data = countingSort(data)
	fmt.Println(data)   // [1 2 3 0 4 5 6 0]
}
 */

func countSortDemo2(){
	data := []int{6, 4, 5, 1, 8, 7, 2, 3}
	countSort(data, 0)
	/*
		data 长度： 8
		map 长度：  8
	*/
	fmt.Println(data)  // [1 2 3 4 5 6 7 8]

}

func countSortDemo(){
	length := 100000

	// 初始化一个 字节切片，指定长度为  100000
	byteArr := make([]byte, length)  // 字节切片里的元素，默认都为 0
	rand.Read(byteArr)    // "crypto/rand"  包下的 Read([]byte) 可生成字节数组长度 个数字  (随机数, 满足一个字节的取值范围 0-255)
	// fmt.Println(byteArr)

	data := make([]int, length)
	for i:=0;i<length;i++ {
		data[i] = int(byteArr[i])    // 将字节强制转换为 int
	}

	data[99999] = 100000
	// fmt.Println("未排序：", data)
	/*   100 个长度的元素 示例：
	未排序： [244 156 77 82 170 106 8 189 25 22 104 44 147 21 122 254 233 97 137 131 252 19 151 86 7 75 5 246 118 227 25 14 160 215 69 3 92 62 91 83 240 134 140 133 19 184 161 164 99 31 168 158 95 13 9 130 141 81 155 182 155 2
	3 199 244 59 206 117 48 168 53 33 146 99 152 184 226 208 166 204 164 221 32 69 200 174 69 7 114 26 48 20 192 74 32 153 179 4 76 39 173]
	*/

	fmt.Println("------------------------------------------------------------------")

	start := time.Now()
	countSort(data, 255)  // 255代表我们已知的 data 列表中的最大值，(go语言统计排序的实现版本： 255 其实没用到）
	// 从开始时间到 结束时间， 时间间隔
	end := time.Since(start)
	fmt.Println("用时：", end.Seconds())   // 0.0049947

	// fmt.Println("拷贝到：", data)
	/*    100 个长度的元素 示例：
	拷贝到： [3 4 5 7 7 8 9 13 14 19 19 20 21 22 23 25 25 26 31 32 32 33 39 44 48 48 53 59 62 69 69 69 74 75 76 77 81 82 83 86 91 92 95 97 99 99 104 106 114 117 118 122 130 131 133 134 137 140 141 146 147 151 152 153 155 155 1
	56 158 160 161 164 164 166 168 168 170 173 174 179 182 184 184 189 192 199 200 204 206 208 215 221 226 227 233 240 244 244 246 252 254]
	*/
}
// 计数排序

// 应用于： 已知最大最小值范围，并且可能会有相同值的 列表

// 原理：将列表中的每一个值作为一个 下标保存到一个集合里(map)作为下标或者key, 遍历整个列表，如果发现相同的则 计数加1
// 上面的步骤完成之后，map 中就已经是排好序的 元素集合了，然后我们需要遍历 map， 把元素一个一个追加到 data 列表中
/*	两个问题：
		1， 需要清空 data, 用于存放排序号的元素
		2,  由于map 的遍历是无序的，如果想 有序遍历，则需要使用 先遍历 map 拿到所有的 key, 然后使用 sort.Sort 对key 进行排序
		3,  根据排序好的 key 遍历 map,  拿到每一个元素的数量，然后 循环追加到 data 列表中，从而完成 排序
 */

// 时间复杂度：O(n), 空间复杂度:O(n)

// 局限 :
/*
    其他语言实现版本：
    计数排序：已知列表的范围  （虽然我们这里没用到 maxCount）
    假设 data 中最大的数值为 maxCount = 100000, 则需要创建 一个长度为 100000+1 的 数组, 因为要将值 当做下标存放到新的 数组中

	局限：需要创建一个 map , 空间复杂度为 O(n)

	go 语言实现版本（虽然需要对key 进行排序，操作复杂一些，但是因祸得福，可以避免开辟 maxCount+1 个长度的 slice ）：
	maxCount 为什么没用到呢？ 由于我们下面要使用 判断 val是否存在 data中，go 语言中 slice 不支持这个方法，所以我们用map 来保存判断
	因为 map[100000] = 2, 与 slice[100000] 的区别:  map[100000] ,不需要创建100000个长度的map, 但是 slice[100000] 就真的要创建 100000个长度的 slice 了

 */
func countSort(data []int, maxCount int){   // go 语言版本中， maxCount 无需知道，也用不上
	countMap := make(map[int]int)
	for i:=0;i<len(data);i++ {
		val := data[i]    // 遍历 data 中每一个元素
		_,ok := countMap[val]  // 将 data中每个元素的值作为 key , 进入 map 中进行判断
		if ok {  // 如果 map 中存在这个 key 则将其值加1  (计数过程)
			countMap[val] = countMap[val] +1
		}else {
			countMap[val] = 1  // 如果不存在，则代表是新key, 则赋值为1 (计数过程)
		}
	}

	fmt.Println("data 长度：", len(data))

	fmt.Println("map 长度： ", len(countMap))

	// fmt.Println(countMap)

	var keyArr = make([]int, 0, len(countMap))   // 创建一个 slice 专门存放 key 值
	for key := range countMap {
		keyArr = append(keyArr, key)
		// fmt.Println("key : ", key)
	}

	sort.Sort(sort.IntSlice(keyArr))    // 将 map 中的key 进行排序

	// 清空原始的 data, 然后存放排序好的元素
	data = data[0:0]    // 清空 data中的元素, 然后将排序好的元素 重新追加到 data 中
	for _,v:= range keyArr {
		num := countMap[v]   // 查看 v 到底有多少个
		for i:=1;i<=num;i++ {  // 有多少个 v 就往新的 slice 里面 把 v 追加多少次
			data = append(data, v)
		}
	}


	/*  创建一个新的 slice 保存排序好的元素
	newData := make([]int,0,len(data))  // 创建一个新的 slice 存放 map 中统计好的 元素值

	// 遍历slice, 拿到 value (这里的 value 就是 data列表里的元素值)
	for _,v:= range keyArr {
		num := countMap[v]   // 查看 v 到底有多少个
		for i:=1;i<=num;i++ {  // 有多少个 v 就往新的 slice 里面 把 v 追加多少次
			newData = append(newData, v)
		}
	}
	fmt.Println("排序后：", newData)

	copy(data, newData)
	*/

}
