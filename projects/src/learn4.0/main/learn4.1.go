package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int
const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

/*
	第3章讨论了 Go程序中的基础数据类型；它们就像宇宙中的原子一样。
	本章介绍 复合数据类型，复合数据类型是由基本数据类型以各种方式组合而构成的，
	就像分子由原子构成一样。我们将重点讲解四种复合数据类型，分别是 数组、slice、map和结构体。
	另外本章末尾将演示如何将使用这些数据类型构成的结构化数据编码为 JSON 数据，从 JSON 数据转换为结构化数据，
	以及从模板生成 HTML 页面。

	数组和结构体都是 聚合类型，它们的值由内存中的一组变量构成。数组的元素具有相同的类型，
	而结构体中的元素数据类型则可以不同。数组和结构体的长度都是固定的。
	反之，slice 和 map 都是动态数据结构，它们的长度在元素添加到结构中时可以动态增长。
 */

// 学习第四章-复合数据类型--4.1-数组
func main() {

	fmt.Println("learn4.1")
	// arrayFunc()
	// arrayFunc2()
	// arrayFunc3()
	// arrayFunc4()
	// arrayFunc5()
	// arrayFunc6()
	// arrayFunc7()
	// arrayFunc8()
	// arrayFunc9()
	arraySliceFunc()
}

/*
	数组是具有固定长度且拥有 零个 或者多个 相同数据类型元素的序列。
	由于数组的长度固定，所以在 Go 里面很少直接使用。 slice 的长度可以增长或缩短，在很多场合下使用得更多。
	然而，在理解 slice 之前，我们必须现理解数组。
 */
func arrayFunc(){
	/*
		数组中的每个元素都是通过索引来访问的，索引从 0 到数组长度减 1。
		Go 内置的函数 len 可以返回数组中的元素个数。
	 */
	var a [3]int                   // 3 个整数的数组，在没有初始值的情况下，int 类型默认值为 0
	fmt.Println(a[0])              // 输出数组的第一个元素
	fmt.Println(a[len(a) - 1])     // 输出数组的最后一个元素，即 a[2]

	// 输出索引(i) 和 元素(v) 值
	for i, v:= range a {
		fmt.Printf("%d %d\n", i , v)
	}
	// 仅输出元素(v) 值，  索引返回值使用 空标识符 (_)代替，因为后面用不上。
	for _, v:= range a {
		fmt.Printf("%d\n", v)
	}
}


func arrayFunc2(){
	/*
		默认情况下，一个新数组中的元素初始值为元素类型的零值，对于数字来说，就是 0.
		也可以使用 数组字面量 根据一组值来初始化一个数组。
	*/
	var q [3]int = [3]int{1,2,3}
	var r [3]int = [3]int{1,2}
	fmt.Println(q[2])            // "3"
	fmt.Println(r[2])            // "0"

	/*
		在元素字面量中，如果省略号 "..." 出现在数组长度的位置，那么数组的长度由初始化数组的元素个数决定。
		以上 数组 q 的定义可以简化为：
	 */
	b := [...]int{1,2,3}
	p := []int{1,2,3}
	fmt.Println(len(b) == len(p))   // "true"
	fmt.Println(q == b)             // "true"
	//fmt.Println(b == p)           // 编译报错。。。
	/*
		编译报错，b是 [3]int 类型， p是[]int 类型 (定义时未在[] 指明长度)，所以b,p 类型不匹配 不能比较
		数组的长度是 数组类型的一部分，所以 [3]int 和 []int 是两种不同的数组类型。
	 */
	fmt.Printf("%d %T\n", len(b), b)  // "3 [3]int"
	fmt.Printf("%d %T\n", len(p), p)  // "3 []int"  // 观察p 数组与 b 数组打印区别，[]里面加 省略号的区别

	/*
		数组的长度是数组类型的一部分，所以 [3]int 和 []int 以及 [4]int 是 不同的数组类型。
		数组的长度必须是常量表达式，也就是说， 这个表达式的值在程序编译时就可以确定。
	 */
	k:= [3]int{1,2,3}
	// k = [4]int{1,2,3,4}      // 编译报错，不可以将 [4]int 赋值给 [3]int
	fmt.Printf("%d %T\n", len(k), k)

	// 看以下两种错误写法：
	// var j [3]int = []int{1,2,3}  // 编译报错：Cannot use '[]int{1,2,3}' (type []int) as type [3]int in assignment
	// var k []int = [3]int{1,2,3}  // 编译报错：Cannot use '[3]int{1,2,3}' (type [3]int) as type []int in assignment
}

/*
	如我们所见，数组、slice、map 和结构体的字面语法都是相似的。
	上面的例子是按顺序给出一组值；也可以像这样给出一组值，这一组值同样具有索引和索引对应的值：
 */
func arrayFunc3(){

	symbol:= [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB,symbol[RMB])    // "3 ¥"
	fmt.Printf("%d %T\n", len(symbol), symbol)

	//	对比上面的写法，我能不能写成下面这样子？

	var a,b,c,d int
	a,b,c,d = 0,1,2,3
	fmt.Println(a,b,c,d)
	// change2:= [...]string{a: "$", b: "€", c: "£", d: "¥"}
	/*
		上列代码 编译报错：Index a,b,c,d must be a non-negative integer constant
		 翻译过来就是： index 必须是 无符号整型常量
	*/
	change2:= [...]string{0: "$", 1: "€", 2: "£", 3: "¥"}   // 直接将索引值赋予 常量(无符号整型)
	fmt.Printf("%d %T\n", len(change2), change2)   // "4 [4]string"


	/*
		在这种情况下，索引可以按照任意顺序出现，并且有的时候还可以省略。
		和上面一样，没有指定值的索引位置的元素默认被赋予数组元素类型的零值。例如,

		定义了一个拥有 100 个元素的数组 r, 除了最后一个元素值是 -1 外， 该数组中的其他元素值都是 0.
	 */
	r:= [...]int{99:-1}
	// 可以这样看： index 为99 的位置上（也就是该数组最后一个元素的位置）赋值为 -1 ，前面 0-98位默认赋值为 0

	fmt.Println(r[99])                         // "-1"
	fmt.Printf("%d %T\n", len(r), r)   // "100 [100]int"

	z:= [...]int{99: -1, 103: 2}
	// 这说明了，99,103，以最大数来初始化数组的长度，并且 99, 103 之间的 100,101,102 省略的空位 都存在，并且默认值都为 0
	fmt.Println(z[99])     // "-1"
	fmt.Println(z[103])    // "2"
	fmt.Println(z[100])    // "0"
	fmt.Printf("z:  %d %T\n", len(z), z)   // "104 [104]int"

	// 同上，string 与  int ，类似
	x:= []string{99: "99", 103: "103"}   // [] 中未加 ...  也不影响它的初始化规律
	fmt.Println(x[99])     // "99"
	fmt.Println(x[103])    // "103"
	fmt.Println("a"+x[100]+"b")    // "ab"  // 说明 x[100] 值为 ""
	fmt.Printf("x:  %d %T\n", len(x), x)  // "104 []string"

	change:= [...]string{"$", "€", "£", "¥"}   // 去掉 常量index , 也能正常声明赋值，这种更简化，使用更多。
	fmt.Println(RMB,change[RMB])   //  "3 ¥"
	fmt.Printf("%d %T\n", len(change), change)

	fmt.Println(symbol == change)   // "true"
}

/*
	如果一个数组的元素类型是 可比较的， 那么这个数组也是可比较的，这样我们就可以直接使用 == 操作符来比较两个数组，
	比较的结果是两边元素的值是否完全相同。使用 != 来比较两个数组是否不同。
 */
func arrayFunc4(){
	a:= [2]int{1,2}
	b:= [...]int{1,2}
	c:= [2]int{1,3}

	fmt.Println(a == b, a == c, b == c)  // "true false false"
	d:= [3]int{1,2}
	fmt.Printf("%d %T\n", len(d), d)  // "3 [3]int"
	// fmt.Println(a == d )  // 编译报错：Invalid operation: a == d (mismatched types [2]int and [3]int)

	f:= []int{1,2}
	fmt.Printf("%d %T\n", len(f), f)  // "2 []int"
	// fmt.Println(a == f)  // 编译报错：Invalid operation: a == f (mismatched types [2]int and []int)

	/*
		数组的长度是数组类型的一部分，所以 [2]int 和 []int 以及 [3]int 是 不同的数组类型。
		数组的长度必须是常量表达式，也就是说， 这个表达式的值在程序编译时就可以确定。

		[]int 与 []int 也不能进行 == 运算，如下例：
	*/
	m:= []int{1,2}
	fmt.Printf("%d %T\n", len(m), m)  // "2 []int"
	// fmt.Println(f == m)  // 编译报错：Invalid operation: f == m (operator == not defined on []int)

}

// 数组可以比较，那它们的地址相同吗？
func arrayFunc5(){

	a:= [2]int{1,2}     // 注意，数组与数组能用 == 比较时，必须 [] 里面有指定长度，用 ... 的也行。
	b:= [...]int{1,2}

	fmt.Println(a == b)            // "true"    // 数组之间比较可以直接用 == 就可以对比其数组里面的值
	fmt.Println(&a == &b)     // "false"   // 虽然数组 a == b ,但是 这两个数组的地址不同。

	fmt.Printf("%p\n",&a)  // "0xc00006e0a0"    // 数组要使用 %p 来打印 地址，如左。
	fmt.Println(&a[0])             // "0xc00006e0a0"    // %p &a 数组的地址，就是 &a[0] 数组首个元素的地址
	fmt.Println(&a[1])             // "0xc00006e0a8"
	fmt.Printf("%p\n",&b)  // "0xc00006e0b0"    // 虽然数组 a == b ,但是 这两个数组的地址不同。
	fmt.Println(&b[0])             // "0xc00006e0b0"    // 数组的地址与 数组首个元素的地址相同
	fmt.Println(&b[1])             // "0xc00006e0b8"

	fmt.Println(&a)                // "&[1,2]"      // 它不像其他指针类型，&可以取地址，&对数组时会 出现左边的格式 &[]

	// string 就可以用 & 取地址，不需要用到 %p
	m := "abc"
	n := "abc"
	fmt.Println(m == n)            // "true"
	fmt.Println(&m == &n)     // "false"
	fmt.Println(&m)                // "0xc0000401f0"
	fmt.Println(&n)                // "0xc000040200"
}

/*
	举一个更有意义的例子，crypto/sha256 包里面的函数 Sum256
	用来为存储在任意字节 slice 中的消息使用 SHA256 加密散列算法生成一个摘要。
	摘要信息是 256位，即 [32]byte。如果两个摘要信息相同，那么很有可能这两条原始消息就是相同的；
	如果这两个摘要信息不同，那么这两条原始消息就是不同的。
	下面的程序输出并比较了 "x"和 "X" 的 SHA256 散列值：
 */
func arrayFunc6(){

	c1:= sha256.Sum256([]byte("x"))
	c2:= sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1,c2,c1 == c2,c1)

	/*  打印信息如下： %x 输出十六进制，  %t 输出布尔值， %T 输出一个值的类型
		2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
		4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
		false
		[32]uint8

	*/

	/*
		这两个原始消息仅有一位 (bit) 之差，但是它们生成的摘要消息有将近一半的位不同。
		注意，上面的格式化字符串 %x 表示将一个数组或者 slice 里面的字节按照十六进制的方式输出，
		%t 表示输出一个布尔值， %T 表示输出一个值的类型。
	 */
}

/*
	当调用一个函数的时候，每个传入的参数都会创建一个副本，然后赋值给对应的函数变量，
	所以函数接受的是一个副本，而不是原始的参数。使用这种方式传递大的数组会变的低效，
	并且在函数内部对数组的任何修改都仅影响副本，而不是原始数组。这种情况下，
	Go 把数组和其他的类型都看成值传递。而在其他的语言中，数组是 隐式地使用引用传递。
 */
func arrayFunc7(){

	a := [3]int{1,2,3}   // 数组作为参数时的 值 传递
	fmt.Printf("%p\n",&a)
	fmt.Println("值传递，处理前：", a)
	// handleArrValueFunc(a)      // "[1,2,3]"
	handleArrValueFunc2(a)   // "[1,2,3]"
	fmt.Println("值传递，处理后：", a)

	c := []int{1,2,3}
	fmt.Printf("%p\n",&c)
	fmt.Println("引用传递，处理前：", c)
	// handleArrReferFunc(c)    // "[2 2 2]"
	handleArrReferFunc2(c) // "[10 10 3]"
	fmt.Println("引用传递，处理后：", c)

	/*  疑问： 值传递，函数中对数组的操作不会影响到 原数组
		引用传递，函数中对数组的操作，就会影响到 原数组，那为什么
		数组定义，数组作为参数，数组进入函数赋值给另一个数组， 这三个变量的地址都不相同？
		当地址都不相同的时候，在函数中对数组的操作，那又是如何影响到 最开始定义的原数组呢？
	 */
}
// 数组作为参数时的 值传递
func handleArrValueFunc(a [3]int) {
	a[0] = 2
	a[2] = 2
}
func handleArrValueFunc2(a [3]int){
	var b [3]int
	b = a

	fmt.Printf("%t %t\n", b == a, &b == &a)   // "true" "false"
	b[0] = 10
	b[1] = 10
}

// 数组作为参数时的 引用传递
func handleArrReferFunc(c []int){
	c[0] = 2
	c[2] = 2
}

func handleArrReferFunc2(c []int){
	var b []int
	b = c

	// fmt.Println(b == c)    // 编译报错： Invalid operation: b == c (operator == not defined on []int)
	fmt.Printf("%t\n",&b == &c)

	b[0] = 10
	b[1] = 10
}

/*
	当然，也可以显式地传递一个数组的指针给函数，这样在函数内部对数组的任何修改都会反映到原始数组上面。
	下面的程序演示如何将一个数组 [32]byte 的元素清零：
*/
func arrayFunc8(){
	var ptr [32]byte = [32]byte{'a','b','c'}
	zero(&ptr)
	fmt.Println(ptr)   // "[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]"

	var ptr2 [32]byte = [32]byte{'a','b','c'}
	zero(&ptr2)
	fmt.Println(ptr2)  // "[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]"

	var ptr4 []byte = []byte{'a','b','c'}
	// 注意指定长度，与不指定长度的数组，进行指针操作时的区别。原则上没有指定长度的数组，少用指针操作
	zero4(&ptr4)
	fmt.Println(len(ptr4))   // "0"
	fmt.Println(ptr4)  // "[]"

}

func zero(ptr *[32]byte) {
	for i:= range ptr {
		ptr[i] = 0
	}
}
/*
	数组字面量 [32]byte{} 可以生成一个拥有 32个字节元素的数组。
	数组中每个元素的值都是 字节类型的零值，即 0. 可以利用这一点来写另一个版本的数组清零程序：
 */
func zero2(ptr *[32]byte){
	*ptr = [32]byte{}
}

/*  写法错误， 数组未指定长度，所以不能被  range
// 编译报错信息： Cannot range over 'ptr' (type *[]byte)
func zero3(ptr *[]byte){
	for i:= range ptr {
		ptr[i] = 0
	}
}
*/

// 未指定长度的数组，进行指针操作时，无法用指针去访问 单个元素
func zero4(ptr *[]byte){
	// ptr[0] = 'x'   // 编译报错：Invalid operation: ptr[0] (type *[]byte does not support indexing)
	*ptr = []byte{}   // 不仅将原数组里面的内容清零，还将原数组的长度也清零了。
}

/*
	使用数组指针是高效的，同时允许被调函数修改调用方数组中的元素，
	但是因为数组长度是固定的，所以数组本身是不可变的。例如上面的zero 函数不能接受一个 [16]byte 这样的数组指针，
	同样，也无法为数组添加或者删除元素。由于数组的长度不可变的特性，除了在特殊的情况下之外，我们很少使用数组。
	上面关于 SHA256 的例子中，摘要的结果拥有固定的长度，我们可以使用数组作为函数参数或结果，但是更多的情况下，
	我们使用 slice。 示例见 4.2 slice
 */

// 数组的截取 与 len(), cap() 函数使用
func arrayFunc9(){
	a:= [5]int{1,2,3,4,5}
	b:= a[2:]   // 截取数组 a ，从下标为2 的元素开始截取到最后
	c:= a[:]    // 截取数组 a ，从默认开始位置到结束位置，也就是全部数组元素
	fmt.Println(c)   // "[1 2 3 4 5]"
	fmt.Println(b)   // "[3 4 5]"

	len:= len(a)    // 数组长度
	cap:= cap(a)    // 数组容量
	fmt.Printf("cap: %d , len: %d\n",cap,len)   // "cap: 5 , len: 5"
}

// 深入理解数组截取后子串的 容量和长度关系
func arraySliceFunc(){
	a:= [10]int{1,2,3,4,5,6,7,8,9,10}
	b:= a[3:7]   //  截取数组 a[i:j] 并赋值给 b
	fmt.Println(b)                // "[4 5 6 7]"
	fmt.Println(cap(b), len(b))   // "7 4"
	/*  数组b 属于 数组a 的子串
	数组b 的长度 len(b)，就是数组b 里面实际的数组个数，a[i:j],  数组长度： j-i == 7-3 == 4
	数组b 的容量 cap(b), 就是从数组b 中的第一个元素的下标 i 到 母串(原始 a 的数组)的结尾，数组容量：cap(a)-i == 10-3 == 7
	*/

	fmt.Println(b[:4])            // "[4 5 6 7]"     // 取 子串b 中的下标为0 到下标为 4-1 的所有元素，元素个数为： 4-0
	fmt.Println(b[:6])            // "[4 5 6 7 8 9]"
	/*
		从 cap(b) 得出b 的容量为 7，len(b) 得出 b 的长度为 4, 所以 b[:6] ， 6 超过 长度但没有 超过容量，所以

		b[:6] 会得到一个长度大于 b 的数组。
	*/

	// fmt.Println(b[5:8])   // 运行报错： panic: runtime error: slice bounds out of range [:8] with capacity 7
	/*
		b[5:8] , b[i:j]
		j 大于 cap(b)== 7, 所以导致程序宕机。
	*/

	// 将子串 b 再分割，遵循以上所有规则， cap, len
	c:= b[2:6]
	fmt.Println(c)                 // "[6 7 8 9]"
	fmt.Println(cap(c), len(c))    // "5 4"
	fmt.Println(c[:5])             // "[6 7 8 9 10]"

	/*
		总结：如果 slice 的引用超过了被引用对象的容量，即 cap(s), 那么会导致程序宕机；
		但是如果slice 的引用超出了被引用对象的长度，即 len(s)， 那么最终 slice 会比原 slice 长：
	*/
}