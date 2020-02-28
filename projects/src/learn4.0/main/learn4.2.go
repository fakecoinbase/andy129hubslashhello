package main

import (
	"fmt"
)

// 学习第四章-复合数据类型--4.2-slice
/*
	slice 表示一个拥有相同类型元素的可变长度的序列。 slice 通常写成 []T, 其中元素的类型都是 T；
		它看上去像没有长度的数组类型。
	数组和 slice 是紧密关联的。 slice 是一种轻量级的数据结构，可以用来访问数组的部分或者全部的元素，
		而这个数组称为 slice 的底层数组。slice 有三个属性：指针、长度和容量。
		指针指向数组的第一个可以从 slice 中访问的元素，这个元素并不一定是数组的第一个元素。
		长度是指 slice 中的元素个数，它不能超过slice 的容量。容量的大小通常是从 slice 的起始元素
		到底层数组的最后一个元素间元素的个数。Go 的内置函数 len 和 cap 用来返回 slice 的长度和容量。
 */
func main() {
	fmt.Println("learn4.2")

	// sliceFunc()
	// sliceReverseFunc()
	sliceCompareFunc()
}

/*
	一个底层数组可以对应多个 slice， 这些slice 可以引用数组的任何位置，彼此之间的元素还可以重叠。
 */
func sliceFunc(){
	months:= [...]string{1:"January", 2:"February",3:"March",4:"April",5:"May",6:"June",7:"July",
							8:"August",9:"September",10:"October",11:"November",12:"December"}
	/*	以上省略了 数组的第一个元素，months[0] 默认为""
		所以 January 就是 months[1], December 是 months[12]。一般来讲，数组中索引 0 的位置存放数组的第一个元素，
		但是由于月份总是从 1 开始，因此我们可以不设置索引为 0 的元素，这样它的值就是空字符串。
	 */

	/*
		slice 操作符 s[i:j] (其中 0<=i<=j<=cap(s) ) 创建了一个新的 slice, 这个新的 slice 引用了序列 s 中
		从 i 到 j-1 索引位置的所有元素，这里的 s 既可以是数组或者 指向数组的指针，也可以是 slice。
		新的 slice 的元素个数是 j-i 个。 如果上面的表达式中省略了 i , 那么新 slice 的起始索引位置就是 0，即 i=0;
		如果省略了 j , 那么新  slice 的结束索引位置是 len(s)-1，即 j=len(s)。
		因此 slice months[1:13] 引用了所有的有效月份，同样的写法可以是 months[1:]。
		slice months[:] 引用了整个数组。接下来，我们定义元素重叠的 slice，分别用来表示第二季度的月份和北半球的夏季月份：
	 */

	Q2:= months[4:7]
	summer:= months[6:9]
	fmt.Println(Q2)        // "[April May June]"
	fmt.Println(summer)    // "[June July August]"

	for _, s:= range summer {
		for _, q:= range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)   // "June appears in both"
			}
		}
	}

	/*
		如果 slice 的引用超过了被引用对象的容量，即 cap(s)，那么会导致程序宕机；
		但是如果 slice 的引用超出了被引用对象的长度，即 len(s)，那么最终 slice 会比原 slice 长：
	 */
	//fmt.Println(summer[:20])     // 运行出错： panic: runtime error: slice bounds out of range [:20] with capacity 7
	endlessSummer:= summer[:5]   // 在 slice 容量范围内扩展了 slice
	fmt.Println(endlessSummer)   // "[June July August September October]"

	str := endlessSummer[0][1:]
	fmt.Printf("slice endlessSummer[0][1:] ：%s\n", str)   // "une"
	// 取 endlessSummer 里面下标为0的元素的值，也就是 "June",
	//	然后截取这个字符串：(下标为1 开始到最后的 字符) 得到 "une"
}

/*
	另外，注意求字符串(string) 子串操作和对字节 slice([]byte) 做 slice 操作这两者的相似性。
	它们都写作 x[m:n]， 并且都返回原始字节的一个字序列，同时它们的底层引用方式也是相同的，
	所以两个操作都消耗常量时间。区别在于： 如果 x 是字符串，那么x[m:n] 返回的是一个字符串；
	如果 x 是字节 slice, 那么返回的结果是 字节 slice 。

	因为 slice 包含了指向数组元素的指针，所以将 一个 slice 传递给函数的时候，可以在函数内部修改底层数组的元素。
	换言之，创建一个数组的 slice 等于为数组创建了一个别名（见 2.3.2节）。下面的函数 reverse 就地反转了整型 slice
	中的元素，它适用于任意长度的整型 slice.
 */
func sliceReverseFunc(){
	a := [...]int{0,1,2,3,4,5}
	reverse(a[:])
	/*
		注意 数组作为参数传入时的写法，a[:],  直接传入 a, 会导致类型不匹配，
		具体说明请参考 learn4.1 里arrayFunc7() 里面的示例。
	 */
	fmt.Println(a)    // "[5,4,3,2,1,0]"

	s := []int{0,1,2,3,4,5}
	// 向左移动两个元素
	reverse(s[:2])
	fmt.Println(s)     // "[1 0 2 3 4 5]"
	reverse(s[2:])
	fmt.Println(s)     // "[1 0 5 4 3 2]"
	reverse(s)
	fmt.Println(s)     // "[2 3 4 5 0 1]"
}

func reverse(s []int){   // 引用修改，会直接修改到数组原数据，详细说明 请参考 learn4.1 里arrayFunc7() 里面的示例。
	for i,j := 0, len(s)-1; i<j; i,j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*  强调重点： slice 与 数组的 区别。
	Go 语言中的切片 (slice)
		type slice struct {
			 array unsafe.Pointer
			 len int
			 cap int
			}
		slice是一个特殊的引用类型,但是它自身也是个结构体
       属性len表示可用元素数量,读写操作不能超过这个限制,不然就会panic
       属性cap表示最大扩张容量,当然这个扩张容量也不是无限的扩张,它是受到了底层数组array的长度限制,超出了底层array的长度就会panic

	注意初始化 slice s 的表达式和初始化数组 a 的表达式的区别。 slice 字面量看上去和数组字面量很像，
	都是用逗号分隔并用花括号括起来的一个元素序列，但是 slice 没有指定长度。（看下面的例子）
				a := [...]int{0,1,2,3,4,5}
				b := [6]int{0,1,2,3,4,5}
				c := a[2:5]
				d := b[2:5]
			a,b 是数组，有固定的长度，类型是：[n]int{} ；
			c,d 是 slice 没有指定长度 ,类型是： []int{}；
			a 与 b 可以用 == 进行比较，
			c 与 d 不行， a 与 c 也不行，
			== 运算只能用在 固定长度的数组之间进行比较。

	这种隐式区别的结果分别是： 创建有固定长度的数组和 创建指向数组的 slice。 和数组一样， slice 也按照顺序指定元素，
	也可以通过索引来指定元素，或者 两者结合。

 */
func sliceCompareFunc(){
	a := [...]int{0,1,2,3,4,5}
	b := [6]int{0,1,2,3,4,5}
	// c := []int{0,1,2,3,4,5}
	fmt.Println(a == b)     // true,  [...]int{} 与 [6]int{} 可以互相转换，比较。(数组比较，比较的是里面的元素值)
	// fmt.Println(a == c)  // 编译错误， [...]int{} 与 []int{} 是不同类型，所以不能比较
	// fmt.Println(b == c)  // 编译错误， [6]int{} 与 []int{} 是不同类型，所以不能比较

	a1 := a[2:5]
	b1 := b[2:5]
	fmt.Println(a1)
	fmt.Println(b1)

	 // fmt.Println(a == a1)   // 编译错误: Invalid operation: a == a1 (mismatched types [6]int and []int)
	// fmt.Println(a1 == b1)   // 编译错误: Invalid operation: a1 == b1 (operator == not defined on []int)

	/*
		和数组不同的是， slice 无法做比较，因此不能用 == 来测试两个 slice 是否拥有相同的元素。
		标准库里面提供了高度优化的函数 bytes.Equal 来比较两个字节 slice([]byte) 。
		但是对于其他类型的 slice ，我们必须自己来写函数比较。
	 */
}

func equal(x,y []string) bool {

	return false
}






