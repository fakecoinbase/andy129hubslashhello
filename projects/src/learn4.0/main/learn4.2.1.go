package main

import (
	"fmt"
)

// 学习第四章-复合数据类型--4.2.1-append 函数
func main() {

	fmt.Println("learn4.2.1")

	// appendFunc1()
	// appendIntTest()
	appendTest()
	// appendFunc2()
}

// 内置函数 append 用来将元素追加到 slice 的后面
func appendFunc1() {
	var runes []rune    // 回顾 rune 与 range 的妙用。
	fmt.Println(len(runes), cap(runes))
	for _,r := range "Hello, 世界" {
		runes = append(runes, r)     //
		fmt.Println("for",len(runes), cap(runes))
	}
	fmt.Println(len(runes), cap(runes))
	fmt.Printf("%q\n", runes)  // ['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']

	// 回顾一下 []rune 与 string 的互相转换方法
	rune2 := []rune("Hello, 世界")
	fmt.Printf("%q\n", rune2)  // ['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']
	str := string(rune2)
	fmt.Println(str)   //  Hello, 世界
}

func appendInt(x []int, y int) []int{
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// slice 仍有增长空间，扩展 slice 内容
		z = x[:zlen]   // 没扩容，依旧使用同一个 底层数组。
	} else {
		// slice 已无空间，为它分配一个新的底层数组
		// 为了达到分摊线性复杂性，容量扩展一倍
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2*len(x)
		}
		z = make([]int, zlen, zcap) // 新创建一个slice
		copy(z, x)   // 内置 copy 函数,  回顾一下 copy():  将 x 的内容 重新拷贝一份给 z,   x 与 z 互相独立，互不影响。
	}
	z[len(x)] = y
	return z
}

/*
	每一次 appendInt 调用都必须检查 slice 是否仍有足够容量来存储数组中的新元素。
	如果 slice 容量足够，那么它就会定义一个新的 slice (仍然引用原始底层数组)，
	然后将新元素 y 复制到新的位置，并返回这个新的 slice。 输入参数 slice x 和 函数返回值 slice z
	拥有相同的底层数组。

	如果slice 的容量不够容纳增长的元素， appendInt 函数必须创建一个拥有足够容量的新的底层数组来存储新元素，
	然后将元素从 slice x 复制到这个数组，再将 新元素 y 追加到数组后面。
	返回值 slice z 将和 输入参数 slice x 引用不同的底层数组。

	使用循环语句来复制元素看上去直观一点，但是使用内置函数 copy 将更简单，
	copy 函数用来为两个拥有相同类型元素的 slice 复制元素。copy 函数的第一个参数是目标 slice,
	第二个参数是 源 slice, copy 函数将源 slice 中的元素复制到目标 slice 中，
	这个和一般的元素赋值有点儿像，比如 dest = src。 不同的slice 可能对应相同的底层数组，
	甚至可能存在元素重叠。copy 函数有返回值，它返回实际上复制的元素个数，这个值是两个 slice 长度的较小值。
	所以这里不存在由于元素复制而导致的索引越界的问题。

	出于效率的考虑，新创建的数组容量会比实际容纳 slice x 和 slice y 所需要的最小长度更大一点。
	在每次数组容量扩展时，通过扩展一倍的容量来减少内存分配的次数，这样也可以保证追加一个元素所消耗的是固定时间。
	下面的程序演示了这个效果：
 */

func appendIntTest(){
	var x,y []int
	for i := 0;i<10;i++ {
		y = appendInt(x, i)
		fmt.Printf("%d   cap=%d\t%v\n", i,cap(y), y)
		x = y
	}

	// 每次 slice 容量的改变都意味着一次底层数组重新分配和元素复制：
	/*  打印结果：
			0   cap=1       [0]
			1   cap=2       [0 1]
			2   cap=4       [0 1 2]
			3   cap=4       [0 1 2 3]
			4   cap=8       [0 1 2 3 4]
			5   cap=8       [0 1 2 3 4 5]
			6   cap=8       [0 1 2 3 4 5 6]
			7   cap=8       [0 1 2 3 4 5 6 7]
			8   cap=16      [0 1 2 3 4 5 6 7 8]
			9   cap=16      [0 1 2 3 4 5 6 7 8 9]
	 */

	/*
		我们来仔细看一下当 i = 3 时的情况。这个时候 slice x 拥有三个元素 [0 1 2]，
		但是容量是 4，这个时候 slice 最后还有一个空位置，所以调用 appendInt 追加元素3 的时候，
		没有发生底层数组重新分配。调用的结果是 slice 的长度和容量都是 4 ，并且这个结果 slice 和 x
		一样拥有相同的底层数组。

		在下一次循环中 i = 4， 这个时候原来的 slice 已经没有空间了，所以 appendInt 函数分配了一个长度为 8 的新数组。
		然后将 x 中的4 个元素 [0 1 2 3] 都复制到新的数组中，最后再追加新元素 i 。这样结果 slice 的长度就是 5，
		而容量是 8。多分配的三个位置就留给接下来的循环添加值使用，在接下来的三次循环中，就不需要再重新分配空间。
		所以 y 和 x 是不同数组的 slice.
	 */
}

/*
	内置的 append 函数使用了比这里的 appendInt 更复杂的增长策略。
	通常情况下，我们并不清楚一次 append 调用会不会导致一次新的内存分配，
	所以我们不能假设原始的 slice 和调用 append 后的结果 slice 指向同一个底层数组，
	也无法证明它们就指向不同的底层数组。同样，我们也无法假设旧 slice 上对元素的操作会 或 不会影响新的 slice 元素。
	所以，通常我们将 append 的调用结果再次赋值给传入 append 函数的 slice:

		runes = append(runes, r)

	不仅仅是在调用 append 函数的情况下需要更新 slice 变量。
	另外，对于任何函数，只要有可能改变 slice 的长度或者容量，
	抑或 是使得 slice 指向不同的底层数组，都需要更新 slice 变量。
	为了正确地使用 slice，必须记住，虽然底层数组的元素是间接引用的，
	但是 slice 的指针、长度和容量不是。要更新一个 slice 的指针，长度或容量必须使用如上所示的显式赋值。
	从这个角度看，slice 并不是纯 引用类型，而是像下面这种聚合类型：

	type IntSlice struct {
		ptr      *int
		len, cap int
	}
 */

/*
	appendInt 函数只能给 slice 添加一个元素，但是内置的 append 函数可以同时给 slice 添加多个元素，
	甚至添加另一个 slice 里的所有元素。
 */
func appendTest(){
	var x []int
	x = append(x, 1)      // 追加数字 1
	x = append(x, 2,3)    // 追加数字 2 和 3
	x = append(x, 4,5,6)  // 追加数字 4,5,6
	x = append(x, x...)           // 追加 x 中的所有元素
	fmt.Println(x)                // [1 2 3 4 5 6 1 2 3 4 5 6]

}

/*
	可以简单修改一下 appendInt 函数来匹配 append 的功能。函数 appendInt 参数声明中的省略号 "..."
	表示该函数可以接受  可变长度参数列表。 上面例子中 append 函数的参数后面的省略号表示如何将一个
	slice 转换为参数列表。 5.7 节会详细解释这种机制。
*/
func appendIntD(x []int, y ...int) []int {
	var z []int
	// zlen := len(x) + len(y)       // 扩展 slice z 底层数组的逻辑和上面一样，所以就不重复了。
	// ... 扩展 slice z 的长度至少到 zlen ...
	copy(z[len(x):], y)
	return z
}


// 参考 appendInt() 函数 深刻理解 slice 追加的工作原理。
func appendFunc2(){

	fmt.Println("------------------扩容示例-----------------")

	a := []int{1,2,3,4}
	fmt.Println(len(a), cap(a))   // 4 4
	c := append(a, 2)
	fmt.Println(len(c), cap(c))   // 5 8  , c 的len 超过 a 的 cap, 所以要扩容，并且扩容 a cap 的一倍( cap(a)*2 == 4*2)
	fmt.Println(c)                // [1 2 3 4 2]
	c[0] = 0                      // 扩容后，a, c 是不同数组的 slice, 修改互不影响
	fmt.Println(c)                // [0 2 3 4 2]
	fmt.Println(a)                // [1 2 3 4]

	fmt.Printf("%p,%p\n", &a, &c)   // 0xc000064460,0xc000064480,  扩容后，a, c 有着不同的地址

	a = append(a, 6)
	fmt.Println(a)
	fmt.Println(len(a), cap(a))
	// 扩容后，a 的容量增大，但是 a = append(a, 6), 将扩容后slice 返回给 a , 所以 a 地址保持原有的不变
	fmt.Printf("%p\n", &a)   //  0xc000064460

	fmt.Println("------------------扩容示例-----------------")

	fmt.Println("------------------未扩容示例-----------------")
	b := [5]int{1,2,3,4}
	b1 := b[:4]
	fmt.Println(b, b1)   //  [1 2 3 4 0] [1 2 3 4]
	fmt.Println(len(b1), cap(b1))   // 4 5
	b2 := append(b1, 3)    // 若没扩容，则 追加的数字会更新到 底层数组 b 里。
	fmt.Println(b, b1)   //  [1 2 3 4 3] [1 2 3 4]
	fmt.Println(len(b2), cap(b2))   // 5 5 ,  b2的 len 没有超过 b1 的cap, 所以不扩容，依然遵循 b1的容量
	fmt.Println(b2)                 // [1 2 3 4 3]
	b[0] = 0          // b2 并没有在 b1 的基础上扩容，所以， b1,b2 都有共同的底层数组 b, 一但某个被修改，则都会影响，看下面三个值
	fmt.Println("------------------b,b1,b2 互为影响-----------------")
	fmt.Println(b)    // [0 2 3 4 3]
	fmt.Println(b1)   // [0 2 3 4]
	fmt.Println(b2)   // [0 2 3 4 3]

	fmt.Println("------------------未扩容示例-----------------")


	fmt.Printf("%p\n", &b1)   //  0xc000004580
	b1 = append(b1, 0)
	fmt.Printf("%p\n", &b1)   //  0xc000004580
	fmt.Println(len(b1), cap(b1))
	fmt.Println(b1)


	/*   总结:
		1，func append(slice []Type, elems ...Type) []Type  会根据传入参数 slice 的容量来决定是否扩容
		2，如果扩容，则扩容量为： cap(slice)*2,  例如：原 slice 容量为 4，则新slice 则会扩容到 4*2 == 8
		3，如果不扩容，则保持 cap(slice) 容量不变。
		4，如果扩容，则原 slice 和 新 slice 会是两个不同的 slice, 分别指向不同的底层数组，各自修改互不影响
		5，如果不扩容，则原 slice 和 新 slice 共同指向同一个底层数组，
			一旦其中一个slice 被修改，则会同步到 底层数组，还有可能会更新到有重叠元素的另一个 slice.
		注意： 既然调用append() 我们不知道会不会扩容，并且扩容后还有各种各样不同的情况，例如：第5 条，
			所以，我们一般情况下都是这样操作： （将返回值 赋值给 传入参数 a ）

			a := []int{1,2,3,4,5}
			a = append(a, 追加的数值)
	 */

}
