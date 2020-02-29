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
	// sliceCompareFunc()
	// shallowCopyAndDeepCopy()
	sliceCompareFunc2()
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

	m := [...]string{"apple","banana","orange","grape"}
	n := [4]string{"banana","orange","apple"}
	fmt.Println(m == n)    // "false",  数组可以用 == 进行比较，但前提是 数组的长度已经指定。 [...]string,  [n]string 之间才能比较

	m1 := m[1:3]
	n1 := n[:2]
	// fmt.Println(m1 == n1)    // 编译报错： Invalid operation: m1 == n1 (operator == not defined on []string)
	fmt.Println(equal(m1,n1))   // "true" ,  自定义 equal() 函数 来比较 两个 slice

	// 那么上面的 equal(m1,n1) 函数还有没有改善的地方呢？
	// 我们做一个例子，比较极端的例子，数组里面都是空的。
	x := []string{}
	y := []string{}
	y = nil
	z := []string{""}
	fmt.Println(len(x), len(y), len(z))   // "0 0 1"
	fmt.Println(x,y,z)   // "[] [] []"

	fmt.Println(equal(x,y))    // true,  // equal() 函数功能不健全，有漏洞，要舍弃。
	fmt.Println(equal2(x,y))   // false, x 与 y 虽然数组长度相同都是0 ，但是 []string{}, []string(nil) 是不同的性质。

	// 小插曲：
	//j := []string{nil}     // 虽然编译没有报错，但是运行就会报错: cannot use nil as type string in array or slice literal

	// 字符串数组，还有另外比较偏门的写法，  new, 详见 2.3.3 new 函数
	k := new([]string)
	fmt.Println(*k, len(*k))   //  "[] 0"
}

/*	扩展学习： https://studygolang.com/articles/9699  （多种方法实现两个字符串string slices的比较）
			https://blog.csdn.net/luopotaotao/article/details/79410581   （两个任意类型的 slice进行比较）
 */

// equal() 函数功能不健全，有漏洞，要舍弃。详见参考文档： https://studygolang.com/articles/9699
func equal(x,y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i:= range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func equal2(x,y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// 与 equal() 不同的是，加入了下面这几行代码, 避免出现 []string{}, []string(nil) 相等的情况出现
	if (x == nil) != (y == nil) {
		return false
	}
	for i:= range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

/*  这种方法暂时保留，待学习 go test 之后再回过头来 验证效率。
	BCE优化 ： https://studygolang.com/articles/9699
	Golang提供BCE特性，即Bounds-checking elimination，关于Golang中的BCE，
	推荐一篇大牛博客Bounds Check Elimination (BCE) In Golang 1.7
	https://go101.org/article/bounds-check-elimination.html
 */
func equal3(x,y []string) bool {
	if len(x) != len(y) {
		return false
	}

	if (x == nil) != (y == nil) {
		return false
	}
	/*
		equal2() 的基础上，加入了下面这行代码, 上述代码通过y = y[:len(x)]处的bounds check
		能够明确保证 y[i]不会出现越界错误，从而避免了y[i]中的越界检查从而提高效率.
	 */
	y = y[:len(x)]

	for i:= range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
/*
	这种深度比较看上去很简单，并且运行的时候并不比字符串数组使用 == 做比较多耗费时间。
	你或许奇怪为什么 slice 比较不可以直接使用 == 操作符比较。这里有两个原因。
	首先，和数组元素不同，slice 的元素是非直接的，有可能 slice 可以包含它自身。
	虽然有办法处理这种特殊的情况，但是没有一种方法是简单、高效、直观的。

	其次，因为slice 的元素不是直接的，所以如果底层数组元素改变，同一个 slice 在不同的时间会拥有不同的元素。
	由于 散列表 (例如 Go的 map 类型)仅对元素的键做浅拷贝，这就要求散列表里面键在散列表的整个生命周期内
	必须保持不变。因为 slice 需要深度比较，所以就不能用 slice 作为 map 的键。对于引用类型，例如指针和通道，
	操作符 == 检查的是引用相等性，即它们是否指向相同的元素。如果有一个相似的 slice 相等性比较功能，它或许
	会比较有用，也能解决 slice 作为 map 键的问题，但是如果操作符 == 对 slice 和 数组的行为不一致，
	会带来困扰。所以最安全的方法就是不允许直接比较 slice。

 */

// slice 与 运算符号， slice 的 nil 的表现形式
func sliceCompareFunc2(){
	m := [...]string{"apple","banana","orange","grape"}
	n := [4]string{"banana","orange","apple"}
	fmt.Println(m == n)

	m1 := m[1:3]
	n1 := n[:2]
	// fmt.Println(m1 == n1)    // 编译报错： Invalid operation: m1 == n1 (operator == not defined on []string)
	//  前面说到 slice 之间不能用 == 运算符，但是 slice 唯一允许的比较操作室和 nil 做比较，例如：

	if m1 != nil {
		if n1 == nil {

		}
	}

	/*
		slice 类型的零值是 nil . 值为 nil 的 slice 没有对应的底层数组。值为 nil 的 slice 长度和容量都是 零，
		但是也有非 nil 的 slice 长度和容量是 零，例如 []int{} 或 make([]int,3)[3:]。
		对于任何类型，如果它们的值可以是 nil,那么这个类型的 nil 值可以使用一种转换表达式，例如 []int(nil)
	 */
	a := []int{}              // 长度，容量，都是 0
	b := make([]int,3)[3:]    // slice b 是在 底层数组[]int{}基础上 切片出来的。
	fmt.Println(len(b), cap(b))   // "0 0"
	fmt.Println(a,b)              // "[] []"

	var s []int              // len(s) == 0, s == nil
	s = nil                  // len(s) == 0, s == nil
	s = []int(nil)           // len(s) == 0, s == nil
	s = []int{}              // len(s) == 0, s != nil
	fmt.Println(s)           // []

	var str []string         // len(str) == 0, str == nil
	str = nil                // len(str) == 0, str == nil
	str = []string(nil)      // len(str) == 0, str == nil
	str = []string{}         // len(str) == 0, str != nil
	str = []string{""}       // len(str) == 1, str != nil
	fmt.Println(str)         // []

	/*
		所以，如果想检查一个 slice 是否是空，那么使用 len(s) == 0，而不是 s == nil,
		因为 s != nil 的情况下， slice 也有可能是空。除了可以和 nil 做比较之外，
		值为 nil 的 slice 表现和其他长度为 零的 slice 一样。例如， reverse 函数调用 reverse(nil) 也是安全的。
		除非文档上面写明了与此相反， 否则无论值是否为 nil, Go 的函数都应该以相同的方式对待所有长度为 零的 slice.
	 */
}

/*  浅拷贝与深拷贝： https://www.jianshu.com/p/35d69cf24f1f
    数据分为基本数据类型(String, Number, Boolean, Null, Undefined，Symbol)和对象数据类型。

	1、基本数据类型的特点：直接存储在栈(stack)中的数据

	2、引用数据类型的特点：存储的是该对象在栈中引用，真实的数据存放在堆内存里

	引用数据类型在栈中存储了指针，该指针指向堆中该实体的起始地址。当解释器寻找引用值时，
	会首先检索其在栈中的地址，取得地址后从堆中获得实体。
 */
func shallowCopyAndDeepCopy(){

	// 基本数据类型，直接存储在 栈 中的数据
	a := "abcdefg"
	var b string = a
	b = "xxxx"
	fmt.Println(b,a)
	c := a
	c = "yyyy"
	fmt.Println(a,b,c)

	d := &a    // 将 a 的地址赋值给 d, 修改 d 则就会 改动到 a
	fmt.Println(*d)
	*d = "afgsdfsdf"
	fmt.Println(a)


	x := []int{1,2,3,4,5}
	y := x
	/*  赋值操作：
		当我们把一个对象赋值给一个新的变量时，赋的其实是该对象的在栈中的地址，而不是堆中的数据。
		也就是两个对象指向的是同一个存储空间，无论哪个对象发生改变，其实都是改变的存储空间的内容，因此，两个对象是联动的。
	 */
	fmt.Printf("%p %p\n", &x,&y)
	y[2] = 10
	fmt.Println(x,y)
}



