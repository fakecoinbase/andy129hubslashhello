package main

import "fmt"

// 多维数组
func main() {

	var a [3]int
	a = [3]int{1, 2, 3}
	fmt.Println(a) // "[1 2 3]"

	var b [3][2]int
	b = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4}, // 单独起一行，必须后面加 逗号
	}
	fmt.Println(b) // "[[1 2] [3 4] [0 0]]"

	var c = [3][2]int{{1, 2}, {3, 4}, {2}} // 最后一维数组中，只初始化了一个元素，所以另外一个默认补 0 为默认值
	fmt.Println(c)                         // "[[1 2] [3 4] [2 0]]"

	// 注意事项： 多维数组除了第一层可以用 [...],  其他层不能用 ...
	var d = [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(b)

	// 多维数组的使用
	fmt.Println(d[2][1])

	// 多维数组的遍历
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d[i]); j++ {
			fmt.Printf("第%d层：%d\n", i, d[i][j])
		}
	}

	// 多维数组的 for range 遍历
	for _, v := range d {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	// 数组是值类型(在长度指定的情况下)，值类型赋值是 拷贝一份内容， 各自互不影响了
	m := [2]int{1, 2}
	n := m
	n[0] = 100
	fmt.Println(m) // "[1 2]"
	fmt.Println(n) // "[100 2]"

	// 以下是 引用类型，共用同一个数组，一旦j 被修改，则会影响到 k
	k := []int{1, 2}
	j := k
	j[0] = 100
	fmt.Println(k) // "[100 2]"
	fmt.Println(j) // "[100 2]"

}
