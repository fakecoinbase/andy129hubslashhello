package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// 学习第四章-复合数据类型--4.2.2-slice 就地修改
func main() {
	fmt.Println("learn4.2.2")

	// nonemptyTest()
	// sliceImplStackTest()
	// sliceRemoveTest()
	// reverseByPtrTest()
	// rotateTest()
	// removeNeiTest()
	// changeBlankTest()
	reverseUTF8Test()
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

func reverseByPtrTest(){
	a := []int{0,1,2,3,4,5}
	reverseByPtr(&a)
	fmt.Println(a)    // "[5 4 3 2 1 0]"
}
// 练习4.3： 重写函数 reverse, 使用数组指针作为参数而不是 slice .
func reverseByPtr(ptr *[]int){

	for i,j := 0,len(*ptr)-1; i < j; i,j = i+1, j-1 {
		(*ptr)[i], (*ptr)[j] = (*ptr)[j],(*ptr)[i]
	}
}

func rotateTest(){
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	r := rotate(a, 3)
	fmt.Println(r)  //  "[3 4 5 6 7 2 1 0]"
	r[2] = 124
	fmt.Println(a)   // r 扩容，底层数组不再是 a,所以 修改 r 中的值 不会影响到 a
}
// 练习4.4：编写一个函数 rotate, 实现一次遍历就可以完成元素旋转。
func rotate(s []int, position int) []int {
	r := s[position:]  // [3,4,5,6,7]
	fmt.Println(len(r),cap(r))
	for i := position - 1; i >= 0; i-- {
		r = append(r, s[i])         // 进行了扩容，导致 r 的底层数组不再是 s , 而是另外开辟了一个空间。
		fmt.Println(len(r),cap(r))
		// [3,4,5,6,7] append [2]
		// [3,4,5,6,7,2] append [1]
		// [3,4,5,6,7,2,1] append [0]
		// [3,4,5,6,7,2,1,0]
	}
	return r
}
func removeNeiTest(){
	a := []string{"aaa","b","cc","cc","cc","b"}
	a = removeNei(a)

	fmt.Println(a)   // "[aaa b cc b]"

	b := []string{"aaa","aaa"}
	b = removeNei(b)

	fmt.Println(b)

	fmt.Println("-------------------网上参考方法测试（已修复代码）--------------------")
	c := []string{"aaa","b","cc","cc","cc","b"}
	removeMultiple(&c)
	fmt.Println(c)    // "[aaa b cc b]"， 代码修复后，输出正常
	fmt.Println("-------------------网上参考方法测试（已修复代码）--------------------")

	d := []string{"s", "a", "a", "s", "d", "z", "a", "z", "v", "w", "w","w","a", "a"}
	// fmt.Println(removeNei(d))   // "[s a s d z a z v w a]"
	removeMultiple(&d)
	fmt.Println(d)                 // "[s a s d z a z v w a]"
}

// 练习4.5: 编写一个就地处理函数，用于去除 []string slice 中相邻的重复字符串元素.
func removeNei(strings []string) []string{

	str := strings[:]
	for i,j:=0,1;j<len(str);i,j=i+1,j+1 {
		if str[i] == str[j] {
			copy(str[i:], str[j:])    // 参考 remove()函数中，高位索引的元素向前移动来覆盖被移除元素所在位置的 原理
			return removeNei(str[:len(str)-1])
			// 当发现有相邻元素一致的情况时：采取了以下策略：
			/*
				1, 参考remove()函数，str[:len(str)-1] 移除slice 里面最后一个多余元素
				2，采用回调函数，将更新的 str 继续执行遍历，继续判断里面是否存在 相邻元素重复的情况。
			 */
		}
	}
	return str
}
// 练习4.5:  网上参考方法
func removeMultiple(a *[]string) {     // "aaa","b","cc","cc","cc","b"
	A := *a
	l := len(A)                    // l == 6
	for i := 0; i < l-1; i++ {    // i == 2时，
		prev := A[i]
		next := A[i+1]
		if prev == next {
			/*
				i == 2时，开始有相同元素出现，然后删除一个相同元素，
				进行 append()函数，追加完之后再复制给 A, A 此时的元素少了一个，所以 长度 l--，然后进行下一轮循环
			问题就出现在下一个循环里：
				当 i == 3时，由于 A 里的元素已更新 "[aaa b cc cc b]]", 所以此时 A[3] 与 A[4]对比，对比的是 "cc" 和 "b",
				所以就漏掉了 A[2]与 A[3] ， "cc" 与 "cc" 之间的比较。
			*/
			A = append(A[:i], A[i+1:]...)   //  append(["aaa","b"], ["cc","cc","b"])
			l--                             // l-- == 6-1 == 5
			i--                             // &&&&& 漏洞修复 &&&&&
		}
	}
	*a = A
}

func changeBlankTest(){
	// a := "a刘德华f    s章f  sd f a   s！d    a"
	a := "asdf    sadf  sd f a   sfd    a"
	runes := []rune(a)
	changeBlankByASCII(runes)    // 自定义函数，有待完善，关键是 还未理解练习题目的意思。

	removeEmpty(&a)   // 结果有待完善，既然说明了是 UTF-8 编码的字节，如果 a 里面有中文，则输出有问题。
	fmt.Println(a)

	fmt.Println("a\u0020b")  // a b         // 半角空格(英文符号)\u0020,代码中常用的;       ASCII
	fmt.Println("a\u3000b")  // a　b        // 全角空格(中文符号)\u3000,中文文章中使用      Unicode
}

// 练习4.6： 编写一个就地处理函数，用于将一个 UTF-8 编码的字节 slice 中
//	所有相邻的 Unicode 空白字符 (查看 unicode.IsSpace) 缩减为一个 ASCII 空白字符。
func changeBlankByASCII(runes []rune){

}
// 练习4.6:  网上参考方法 （有待确认?）
func removeEmpty(a *string) {
	S := *a
	str := string(S[0])
	end := 0
	for _, s := range S {
		last := rune(str[end])
		//fmt.Printf("s:  %q  %t\n",s, unicode.IsSpace(s))
		//fmt.Printf("---l:  %q  %t\n",last, unicode.IsSpace(last))
		if unicode.IsSpace(last) && unicode.IsSpace(s) {
			continue
		}
		str += string(s)
		//fmt.Println("str : ",str)
		end++
	}
	*a = str[1:]
}

// 练习4.7: 修改函数 reverse, 来翻转一个 UTF-8 编码的字符串中的字符元素，传入参数是 该字符串对应的字节 slice 类型([]byte)。
// 你可以做到不需要重新分配内存就实现该功能吗？
func reverseUTF8Test(){
	str := "fdsf我是fdfd 刘德华df f"
	b := []byte(str)
	reverseUTF8(&b)           // 进行指针操作
	fmt.Println(string(b))    // "f fd华德刘 dfdf是我fsdf"

	b[0] = 'a'
	fmt.Println(string(b))    // "a fd华德刘 dfdf是我fsdf"

	fmt.Println("-------------------网上参考方法测试之 不使用指针操作--------------------")
	str3 :=  "fsdf中华人民fsdfd共和国f"
	b3 := []byte(str3)
	reverse_byteNoPtr(b3)
	fmt.Println(string(b3))   // "f国和共dfdsf民人华中fdsf"

	b3[0] = 'a'
	fmt.Println(string(b3))   // "a国和共dfdsf民人华中fdsf"

	/*
		//  utf8.DecodeRuneInString(string(slice[0:]))  每次循环取 slice里面第一个元素

		//  slice 首个元素为"中" "中华人民共和国" "华人民共和国"  "华人民共和国国"   替换："华人民共和国中"
		//  slice 首个元素为"华" "华人民共和国"   "人民共和国"   "人民共和国国中"   替换："人民共和国华中"
		//  slice 首个元素为"人" "人民共和国"   "民共和国"   "民共和国国华中"       替换："民共和国人华中"
		//  slice 首个元素为"民" "民共和国"   "共和国"   "共和国国人华中"           替换："共和国民人华中"
		//  slice 首个元素为"共" "共和国"   "和国"   "和国国民人华中"               替换："和国共民人华中"
		//  slice 首个元素为"和" "和国"   "国"   "国国共民人华中"                   替换："国和共民人华中"
	 */
}

// 练习4.7：个人实现方法
// 对比 reverse_byteNoPtr()，我这里使用了 rune 迭代，核心代码过程简单，但却比 它多了一个指针操作，后续再研究优化。
func reverseUTF8(b *[]byte){
	str := string(*b)
	runes := []rune(str)
	len := len(runes)

	for i,_ := range runes {
		// fmt.Printf("%q\n", v)
		if (i < len/2) {
			runes[i],runes[len-i-1] = runes[len-i-1],runes[i]
		}
	}
	*b = []byte(string(runes))    // 将最后的结果赋予 指针变量，从而更新传入的 b *[]byte
	//fmt.Println(string(runes))
}

// 对比以上的方法，这种方法就算是没有使用指针，依然能达到就地修改 slice 的目的，原因是？
// 原因是  utf8.DecodeRuneInString(string(slice[0:]))  从这里开始一直使用的是 slice[:]
// 修改 slice 片段能够达到更新 底层数组的目的，所以针对 slice 的修改会更新到 []byte 里。

// 练习4.7： 网上参考方案  (是以单个 UTF-8的字节数为间隔，依次循环迭代由高位索引的元素向前移动，后面再补充首个被覆盖的元素)
func reverse_byteNoPtr(slice []byte) {

	//fmt.Println(string(*slice))

	for l := len(slice); l > 0; {
		r, size := utf8.DecodeRuneInString(string(slice[0:]))   // 每次都输出 slice 里面第一个元素
		//fmt.Println(string(r))
		copy(slice[0:l], slice[0+size:l])
		copy(slice[l-size:l], []byte(string(r)))
		l -= size
	}
	// fmt.Println(string((slice)))
}