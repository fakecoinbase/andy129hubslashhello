package main

import "fmt"

// 学习第四章-复合数据类型--4.2.2-slice 就地修改
func main() {
	fmt.Println("learn4.2.2")

	// nonemptyTest()
	// sliceImplStackTest()
	sliceRemoveTest()
}

/*
	我们多看一些就地使用 slice 的例子，比如 rotate 和 reverse 这种可以就地修改 slice 元素的函数。
	下面的函数 nonempty 可以从给定的一个字符串列表中去除 空字符串并返回一个新的 slice .
	nonempty 演示了 slice 的就地修改算法
 */
// nonempty 返回一个新的 slice , slice 中的元素都是非空字符串
// 在函数的调用过程中，底层数组的元素发生了改变
func nonempty(strings []string) []string{
	i:=0
	for _, s:= range strings {
		if s != "" {
			strings[i] = s     // 就地修改原 slice
			i++
		}
	}
	return strings[:i]        // 返回除去 "" 之后 长度的 slice
}
/*
	这里有一点是 输入的 slice 和输出的 slice 拥有相同的底层数组，这样就避免在函数内部重新分配一个数组。
	当然，这种情况下，底层数组的元素只是部分被修改，示例如下：
 */
func nonemptyTest(){
	data := []string{"one", "", "three"}
	nonemptySlice := nonempty(data)
	fmt.Println(len(nonemptySlice), cap(nonemptySlice))   // 2 3      // 容量与 原底层数组 data 保持不变
	fmt.Println(nonemptySlice)     // [one three]
	fmt.Println(data)              // [one three three] , 发现底层数组只是进行了部分修改，并没有达到效果。

	fmt.Println("-------------经常使用的是下面的写法----------------")

	// 因此，通常我们会这样来写
	data2 := []string{"one", "", "three"}
	data2 = nonempty(data2)
	fmt.Println(len(data2), cap(data2))   // 2 3      // 容量依旧与 底层数组 data2 保持不变
	fmt.Println(data2)                    // [one three]

	fmt.Println("-----------appen()来实现nonempty()----------------")

	data3 := []string{"one", "", "three","four","", ""}
	data3 = nonempty2(data3)
	fmt.Println(len(data3), cap(data3))   // 3 6      // 容量依旧与 底层数组 data3 保持不变
	fmt.Println(data3)                    // [one three four]

}

// 函数  nonempty 还可以利用 append 函数来写：
func nonempty2(strings []string) []string {
	out := strings[:0]   // 引用原始 slice 的新的零长度的 slice
	for _,s:=range strings {
		if s != "" {
			out = append(out , s)
		}
	}
	return out
}
/*
	无论使用哪种方式，重用底层数组的结果是 每一个输入值的 slice 最多只有一个输出的结果 slice,
	很多从序列中过滤元素再组合结果的算法都是这样做的。这种精细的 slice 使用方式只是一个特例，
	并不是规则，但是偶尔这样做可以让实现更清晰、高效、有用。
 */
// 使用 slice 来进行栈的 push ,pop 工作
func sliceImplStackTest() {
	var stack []string = []string{"one","two","three","four"}
	var v string = "five"
	stack = append(stack, v)     // push v
	fmt.Println(stack)           // [one two three four five]
	fmt.Println(len(stack), cap(stack))  // 5 8  ,  append()函数将扩容后的 slice 赋值给了 stack, stack 便同步更新，进行了扩容
	// 栈的顶部是最后一个元素：
	top := stack[len(stack)-1]   // 栈顶元素
	fmt.Println(top)             // five
	// 通过弹出最后一个元素来缩减栈：
	stack = stack[:len(stack)-1]  // 扩容后的 stack 并不会因为 弹出了一个元素，导致容量也回到原来的值，弹出元素改变的只是 len
	fmt.Println(stack)            // [one two three four]
	fmt.Println(len(stack), cap(stack))  // 4 8
}

func sliceRemoveTest(){
	s := []int{5,6,7,8,9}
	fmt.Println(remove(s,2))     // "[5 6 8 9]"

	s2 := []int{5,6,7,8,9}
	fmt.Println(remove2(s2, 2))  // "[5 6 9 8]"
}
/*
	为了从 slice 的中间移除一个元素，并保留剩余元素的顺序，可以使用函数 copy 来将高位索引的元素
	向前移动来覆盖被移除元素所在位置：
 */
func remove(slice []int , i int) []int {
	count := copy(slice[i:], slice[i+1:])
	fmt.Println("成功拷贝元素的个数：",count)     //  2
	/*  代码分析，假设 i = 2
		copy(slice[2:], slice[2+1]:)
		copy([7 8 9], [8 9])     // 将 src [8 9] 拷贝到 dst [7 8 9] 里面
		== [8,9,9]
	 */
	fmt.Println(slice[i:])         // [8 9 9]
	fmt.Println(slice)             //  [5 6 8 9 9]
	return slice[:len(slice)-1]    // [5 6 8 9]
	// 由于元素向前移动，覆盖了要删除的元素，所以原数组最后一个元素也要舍弃。最终达到删除一个元素的目的
}
/*
	如果不需要维持 slice 中剩余元素的顺序，可以简单地 将slice 的最后一个元素赋值给被移除元素所在的索引位置
 */
func remove2(slice []int, i int) []int{
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}