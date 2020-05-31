package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

// 堆排序： 堆 是一个特殊的完全二叉树结构 (// 时间复杂度：O(nlogn))
/*
	大根堆：一棵完全二叉树，满足任一节点都比其孩子节点大
	小根堆：一棵完全二叉树，满足任一节点都比其孩子节点小
 */

/*	二叉树的存储方式 (表达方式)
	-- 链式存储方式
	-- 顺序存储方式  (数组，列表等, 本文以此为例)

		父节点下标和左孩子节点 的下标的关系为： i   -->   2i+1
        父节点下标和右孩子节点 的下标的关系为： i   -->   2i+2

        通过子节点找 父节点:	a  --->  (a -1)/2   (无论左右节点，都能套用这个公式)

 */
func main() {

	// heapSortDemo()   // 堆排序原理

	topKDemo()   // 堆排序应用

}

// 堆排序的应用： topK 问题 (热搜榜)
// 现有 n 个数，设计算法得到前 k 大的数. (k<n)
/*  实现步骤：
	1，采用堆排序
	2, 修改 sift() 使其进行 小根堆调整  (为什么使用小根堆? 因为最后出数 可以达到 降序的目的)
 */
func topKDemo(){
	// 解决思路：
	/*
		方案一： 排序后切片,  时间复杂度： O(nlogn)
		方案二： 冒泡，选择，插入排序，  时间复杂度：O(kn)
		方案三： 堆排序 ,  时间复杂度：O(nlogk)
	 */

	length := 100000

	// 初始化一个 字节切片，指定长度为  100000
	byteArr := make([]byte, length)  // 字节切片里的元素，默认都为 0
	rand.Read(byteArr)    // "crypto/rand"  包下的 Read([]byte) 可生成字节数组长度 个数字  (随机数, 满足一个字节的取值范围 0-255)
	// fmt.Println(byteArr)

	data := make([]int, length)
	for i:=0;i<length;i++ {
		data[i] = int(byteArr[i])    // 将字节强制转换为 int
	}

	arr := topK(data, 500)
	fmt.Println(arr)

}
// 小根堆 调整
func siftLittleHeap(data []int, low,high int){
	i := low  // 根节点
	j := 2*i + 1 // 子节点 (左边)
	tmp := data[i]   // 保存根节点的值

	// 保证数组不越界
	for j <= high {
		if j+1 <= high && data[j+1] < data[j]  {
			j = j+1    // 子节点指向 右边的节点
		}
		if data[j] < tmp{
			data[i] = data[j]
			i = j
			j = 2*i+1
		}else {
			data[i] = tmp
			break
		}
	}

	if j > high {
		data[i] = tmp
	}
}
func topK(data []int, k int) []int{
	heap := data[0:k]
	// 建堆 (针对 heap列表 进行建堆)
	for i:= (k-2)/2; i>=0;i -- {
		siftLittleHeap(heap, i, k-1)
	}
	// 遍历
	// 从下表为 k 的元素开始，遍历整个列表，依次与 heap 中根节点的值对比
	// heap 建堆之后，由于是小根堆, heap[0]中一定是最小的值
	for i:=k;i< len(data)-1;i++ {
		if data[i] > heap[0] {   // 小于 heap[0] 不考虑，只考虑大于  heap[0]的值
			heap[0] = data[i]  // 将一个新成员的值 赋值到 heap[0]
			siftLittleHeap(heap, 0, k-1)  // 然后继续小根堆调整，调整直到 heap[0] 为小根堆中最小的元素
		}
	}
	// 出数 (效果是 最小的数放在了 末尾，达到 降序目的)
	for i:= k-1;i>=0;i-- {
		heap[0],heap[i] = heap[i],heap[0]   // 将根节点的值赋值到 i 的位置上 (实际就是 从列表最后一个元素依次向前追加,例如：xxxxxxxxx9, xxxxxxxx89,xxxxxxx789)
		siftLittleHeap(heap, 0, i-1)   // i-1 新的 high
	}

	return heap
}

// 堆排序
/*	两个重要的步骤:
	a, 建堆    (建堆完成的标志： 根节点是最大元素或 最小元素，每一个父节点元素值 全大于子节点，或者 全小于 子节点)
		在堆排序中，我们使用的是 大根堆，排出来的顺序是 从小到大
	b, 调整堆
 */

// 时间复杂度：O(nlogn)
/*
	sift() 时间复杂度为： O(logn)
	heapSort() 时间复杂度为：O(nlogn)
 */
func heapSortDemo(){
	length := 100000

	// 初始化一个 字节切片，指定长度为  100000
	byteArr := make([]byte, length)  // 字节切片里的元素，默认都为 0
	rand.Read(byteArr)    // "crypto/rand"  包下的 Read([]byte) 可生成字节数组长度 个数字  (随机数, 满足一个字节的取值范围 0-255)
	// fmt.Println(byteArr)

	data := make([]int, length)
	for i:=0;i<length;i++ {
		data[i] = int(byteArr[i])    // 将字节强制转换为 int
	}

	start1 := time.Now()
	heapSort(data)
	end1 := time.Since(start1)
	fmt.Println("用时：", end1.Seconds())   // 用时： 0.0094915
}

// 调整树，建堆过程
// 参数1：需要排序的列表；参数2：根节点；参数3：最后一个叶节点  (可以把一个完整的二叉树理解为 多个堆，每一个堆又分 根节点和叶节点)
// 大根堆 调整
func sift(data []int, low,high int){

	i := low  // 根节点
	j := 2*i + 1 // 子节点 (左边)
	tmp := data[i]   // 保存根节点的值

	// 保证数组不越界
	for j <= high {
		// 当右边的节点存在(不越界)， 并且 右节点的值 大于 左节点的值, 则将 j 指向右边
		if j+1 <= high && data[j+1] > data[j]  {
			j = j+1    // 子节点指向 右边的节点
		}
		// 当 子节点的值 大于 父节点的值
		if data[j] > tmp{
			data[i] = data[j]   // 将子节点的值 提升到 父节点上  （注意这里的提升，只是值的提升，原本子节点的位置上依旧保存着 旧的值）
			// 下面两步操作，将父节点与子节点 向下调整一层
			i = j   // 更新父节点的位置为 子节点的位置
			j = 2*i+1  // 通过 子节点的位置再算出 下一个子节点的位置
		}else {
			data[i] = tmp   // 将父节点的值重新调整到 合适的位置 (合适的位置： 小于它的父节点的值)
			break   // 调整完成，跳出循环
		}
	}

	// 如果 i 这个位置下面没有任何子节点了，则将 i 位置保存的值(还是以前的子节点的旧值) 替换为 根节点的值
	if j > high {
		data[i] = tmp   // 将保存父节点值的变量赋值给 最后的 父节点 (其实是最底层的某个叶节点)， 从而达到 将父节点的值移到最下面(如果它是最小值的话)
	}

}

// 堆排序
func heapSort(data []int){
	n := len(data)
	// n-1 最后一个叶节点的位置,  通过子节点计算父节点的位置 : ((n-1)-1)/2 == (n-2)/2
	for i:= (n-2)/2; i>=0;i -- {    // 计算出每个父节点的位置，然后依次进行 调整  (其中 n-1 为 high, 这里取巧均指向 列表最后一个 叶节点)
		sift(data, i, n-1)
	}
	// 以上建堆完成 (最大的数一定是在堆顶，也就是 根节点)
	// 开始挨个出数 (原理：让列表最后的未排序的元素一个一个与根节点互换值，将根节点的值一个一个的放置到 列表结尾， 然后 对未排序的列表再进行堆调整， )
	//  (效果是 最大的数放在了 末尾，达到 升序目的)
	for i:= n-1;i>=0;i-- {
		data[0],data[i] = data[i],data[0]   // 将根节点的值赋值到 i 的位置上 (实际就是 从列表最后一个元素依次向前追加,例如：xxxxxxxxx9, xxxxxxxx89,xxxxxxx789)
		sift(data, 0, i-1)   // i-1 新的 high
	}

}

