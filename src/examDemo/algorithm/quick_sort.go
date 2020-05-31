package main

import (
	"crypto/rand"
	"fmt"
	"sort"

	 r "math/rand"
	"time"
)

// 快速排序 (时间复杂度为：平均情况：o(nlogn), 最坏情况：O(n方).,  空间复杂度：平均 O(logn), 最坏：O(n) )
func main() {

	// quickSortDemo()   // 快速排序的原理

	// quickSortDemo2()     // 快速排序的用时

	 quickSortDemo3()     // 快速排序的局限
}

// 对比测试
/*
	// 局限：
		// 1, 因为使用了 递归调用，所以递归深度达到一定值时，某些语言会执行不了，比较占用空间 (例如：python语言，但可以修改深度值)
		// 2, 当针对一个 已经排好序的列表时， 每一层的时间复杂度为 O(n)，总共又有 n层 （而不再是 logn 层） 所以时间复杂度为 O(n的平方)

		// 测试步骤：
		/*
			1, 初始化一个 长度为 100000 的 []int ，初始化数据并打乱
				a, 创建一个长度为 100000 的 []byte
				b, 通过rand.Read() 随机生成 []byte 里元素的值并 打乱数据
				c, 将 []byte 转换为 []int
			2, 测试1：通过 快速排序算法进行排序，计算用时
			3, 测试2：将 []int 通过 sort.Sort() 先排好序， 再使用 快速排序算法 排序，计算用时
			4, 对比 两个测试的用时

			测试结果：
				测试1:  // 针对未排序的列表，排序后用时：0.0119941
			    测试2:  // 针对升序列表，排序后用时：2.1785672
	            测试3:  // 针对降序列表，排序后用时：2.1695454

			结论：
				快速排序针对 已经排好序的列表，其真实的 时间复杂度要 高于 未排序好的列表
	*/

	// 优化： (局限2 的原因就是 data[left] 是一个最大值或者最小值, 与 left 和 right 比较时，左右分区动作不明显)
	/*   在 quickSort()  tmp := data[left] 之前进行 两步操作

	// 针对 快排对 已经排序好的列表 进行排序有一定的局限，那么我们可以如下操作 ()
	index := r.Intn(right-left) + left    // 从 left, right 区间，随机找到一个元素 (这里取下标)
	data[left], data[index] = data[index], data[left]   // 然后将 这个值与 left 做交换，某种程度上达到了 破坏已经排序好的 data

	tmp := data[left]
	*/

	// 再测试：(重复以上测试步骤)
	/*
		测试结果：
			测试1:  // 针对未排序的列表，排序后用时：0.0150053
			测试2:  // 针对升序列表，排序后用时：0.0100078
	        测试3:  // 针对降序列表，排序后用时：0.0130141

		结论：
			优化后，用时已经明显 减少
	*/

	// 优化后：问题
	/*  index 依旧有可能随机到 最大的值，所以也有可能是 O(nlogn), 但这种情况会很少出现，不用再纠结了

	index := r.Intn(right-left) + left    // 从 left, right 区间，随机找到一个元素 (这里取下标)
	data[left], data[index] = data[index], data[left]   // 然后将 这个值与 left 做交换，某种程度上达到了 破坏已经排序好的 data
*/

func quickSortDemo3() {

	length := 100000

	// 初始化一个 字节切片，指定长度为  100000
	byteArr := make([]byte, length)  // 字节切片里的元素，默认都为 0
	rand.Read(byteArr)    // "crypto/rand"  包下的 Read([]byte) 可生成字节数组长度 个数字  (随机数, 满足一个字节的取值范围 0-255)
	// fmt.Println(byteArr)

	data := make([]int, length)
	for i:=0;i<length;i++ {
		data[i] = int(byteArr[i])    // 将字节强制转换为 int
	}

	// 通过快速排序将 打乱的 data 进行排序，查看用时
	start1 := time.Now()
	quickSort(data, 0, len(data) - 1)
	end1 := time.Since(start1)
	fmt.Printf("针对未排序的列表，排序后用时：%v\n", end1.Seconds())    // 未排序用时：0.0110063

	fmt.Println("-------------------------------------------------------------------")

	// fmt.Println(data)
	// 给 内置包 sort 计算用时，不是本次测试的目的，只是研究一下 sort 功能的排序大概要花多长时间
	//start := time.Now()
	// 升序
	sort.Sort(sort.IntSlice(data))    // 使用 sort 包对 data 进行排序，然后再使用 快速排序，查看用时
	//end := time.Since(start)
	//fmt.Println("sort 用时: ", end.Seconds())   //  sort 用时:  0.0040119

	start2 := time.Now()
	quickSort(data, 0, len(data) - 1)
	end2 := time.Since(start2)
	fmt.Printf("针对升序列表，排序后用时：%v\n", end2.Seconds())    // 排序后用时：2.1655513

	fmt.Println("-------------------------------------------------------------------")

	// 降序
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	start3 := time.Now()
	quickSort(data, 0, len(data)-1)
	end3 := time.Since(start3)
	fmt.Printf("针对降序列表，排序后用时：%v\n", end3.Seconds())    // 排序后用时：2.1655513
}

// 随机一个 100000 长度的 数组，并打乱数据进行排序，查看 快速排序的 用时
func quickSortDemo2(){

	length := 100000

	// 初始化一个 字节切片，指定长度为  100000
	byteArr := make([]byte, length)  // 字节切片里的元素，默认都为 0
	rand.Read(byteArr)    // "crypto/rand"  包下的 Read([]byte) 可生成字节数组长度 个数字  (随机数, 满足一个字节的取值范围 0-255)
	// fmt.Println(byteArr)

	data := make([]int, length)
	for i:=0;i<length;i++ {
		data[i] = int(byteArr[i])    // 将字节强制转换为 int
	}
	// fmt.Println(data)

	// 开始时间
	start := time.Now()
	quickSort(data,0,len(data) -1)
	// 从开始时间到 结束时间， 时间间隔
	end := time.Since(start)
	fmt.Println(data)
	fmt.Println("用时：", end.Seconds())   // 打印 用时:  0.0120079


	// 洗牌算法(打乱数据)
	// "math/rand" 包下的 Shuffle 方法可以根据自定义的 打乱算法，将 数组里的指定长度的元素 打乱, 由于上面已经被 Read() 打乱了，所以这里暂时用不上
	/*
		rand.Seed(time.Now().UnixNano()) //设置种子

	    sixah := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	    rand.Shuffle(len(sixah), func(i, j int) { //调用算法
	        sixah[i], sixah[j] = sixah[j], sixah[i]
	    })

	    fmt.Println(sixah)
	*/
	// fmt.Println(byteArr)
}

// 快速排序：
// 时间复杂度： 每个 partition 的时间复杂度为 O(n), 总共有 logn 层， 所以时间复杂度为：o(nlogn)
func quickSortDemo(){

	data := []int{5,7,4,6,3,1,2,9,8}
	// mid := partition(data, 0, len(data)-1)
	// fmt.Println(mid)
	 quickSort(data, 0, len(data)-1)
	 fmt.Println(data)
}

// 参数1: 要排序的列表;  参数2：列表左边起始位；  参数3： 列表右边末尾；  返回值：列表中位的 下标
func partition(data []int, left, right int) int{

	// 优化
	// 针对 快排对 已经排序好的列表 进行排序有一定的局限，那么我们可以如下操作:
	// 备注： r "math/rand"
	index := r.Intn(right-left) + left    // 从 left, right 区间，随机找到一个元素 (这里取下标)
	data[left], data[index] = data[index], data[left]   // 然后将 这个值与 left 做交换，某种程度上达到了 破坏已经排序好的 data

	tmp := data[left]
	for left < right{   // left < right , 继续走下面
		for left < right && data[right] >= tmp{   // 从右边找, 如果大于等于 tmp, 则 right --, 继续查找， 如果 小于等于 tmp, 则退出循环
			right --   // 往左走一位
			//fmt.Println("right : ", data)
		}
		data[left] = data[right]   // 找到小于tmp 的数字之后，把这个值赋值到 左边去， 完成一次 小号的左边分区

		for left < right && data[left] <= tmp{    // 从左边找，如果小于等于 tmp, 则 left ++, 继续查找； 如果 大于等于 tmp, 则退出循环
			left++     // 往右走一位
			//fmt.Println("left : ", data)
		}
		// 最后的一种情况是：left ++ 后， left == right , data[right] = data[left]  把最后一个数字自己对自己赋值，不影响整体代码, 最后跳出外层的 for 循环
		data[right] = data[left]   // 找到大于 tmp 的数字之后，把这个值赋值到 右边去，完成了一次 大号的右边分区
	}

	// 当 left == right ，说明已经找到 中位， 则退出上面那个最外层的 for 循环
	data[left] = tmp    // 最后将 tmp 值放置到中位上，则完成 tmp 左边是小于 tmp 的值，右边是大于 tmp 的值， 此时 left == right ,  data[right] = tmp 也可以

	//fmt.Println("all : ", data)
	return left   // 返回 left 或 right 都可以，此时 left == right
}
func quickSort(data []int, left, right int){
	// 递归调用，直到 left == right ,代表左右分区完成，直至 列表全部完成排序
	if left < right {
		mid := partition(data, left,right)  // 找出列表 中位，让列表分出 左右两块
		// fmt.Println(mid)
		quickSort(data, left, mid-1)    // 对列表左边再进行分区
		quickSort(data, mid+1, right)   // 对列表右边再进行分区
	}
}
