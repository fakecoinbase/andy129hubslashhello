package main

import (
	"fmt"
	"strings"
)

// 学习第三章--3.4-字符串
func main() {
	fmt.Println("learn3.5")

	// stringVarFunc()
	stringApendFunc()
}

func stringVarFunc(){

	s:= "hello, world"
	lenth := len(s)
	fmt.Println(lenth)  //"12"   // 字符串 s 的字符长度为 12
	fmt.Println(s[0], s[7])  // "104 119"  // 'h' 和 'w' 的 ASCII 码
	fmt.Printf("s[0] : %c, s[7] : %c\n", s[0], s[7])  // s[0] : h, s[7] : w    //  以字符的形式打出来


	// c := s[len(s)]   // 宕机： 下标越界 ,     下标 index 的取值范围： 0<=index< len(s)

	/*
		注意： 字符串的第 i 个字节不一定就是第 i 个字符， 因为非 ASCII 字符的 UTF-8 码点需要两个字节或多个字节。
		稍后将讨论如何使用字符。
	 */

	/*  子串生成操作
		s[i:j] 产生一个新字符串， 内容取自原字符穿的字节，下标从 i (含边界值)开始，
		直到 j (不含边界值)。 结果的大小 是 j-i 个字符。
	 */
	fmt.Println(s[0:5])   // "hello"

	/*   再次强调， 若 下标越界， 或者 j 的值小于 i, 将触发宕机异常。
		操作数 i 与 j 的默认值分别是 0 (字符串起始位置) 和 len(s) (字符串终止位置)，
		若省略 i 或 j ， 或 两者， 则取默认值。
	 */
	fmt.Println(s[:5])   // "hello"     // 省略 i 值， s[:5] == s[0:5]
	fmt.Println(s[7:])   // "world"     // 省略 j 值， s[7:12]
	fmt.Println(s[:])    // "hello, world"    // 省略 i , j 两个值,  s[:] == s[0:12]

	// fmt.Println("goodbye"+s[5])
	//编译报错， s[5] 取单个字符，类型为 uint8,   string 与 uint8 属于不同类型，故不能通过 + 连接 。

	fmt.Println("goodbye"+s[5:])  // "goodbye, world"
	//编译通过， 说明 s[5:] 得出的是一个 字符串类型的变量

}

func stringApendFunc(){

	/*
		字符串可以通过比较运算符 做比较， 如 == 和 < ； 比较运算按字节进行， 结果服从本身的字典排序。
		尽管肯定可以将 新值赋予字符串变量， 但是字符串值无法改变： 字符串值本身所包含的字节序列永不可变。
		要在一个字符串后面添加 另一个字符串， 可以这样编写代码：
	*/

	s := "left foot"
	fmt.Println("left foot字符串 地址： ", &s)   // "0xc0000401f0"
	t := s
	fmt.Println("s 值赋值给 t, s的地址： ", &s)  // "0xc0000401f0"
	fmt.Println("s 值赋值给 t, t的地址： ", &t)  // "0xc000040200"
	/*
		从这里可以看出 t 是新变量，新分配了一个地址，并且将 s 的值赋予了过去，
		而 s 的地址不变，包括下面 追加了一个字符串之后，还是没变。
	 */
	s += ", right foot"
	fmt.Println("s 追加字符串后的地址： ", &s)   // "0xc0000401f0"


	/*
		这并不改变 s 原有的字符串值， 只是将 += 语句生成的新字符串赋予 s。
		同时， t 仍然持有旧的字符串值。
	*/
	fmt.Println(s)    	// "left foot, right foot"
	fmt.Println(t)		// "left foot"

	// 因为字符串不可改变， 所以字符串内部的数据不允许修改；
	// s[0] = 'L'    // 编译报错，不能赋值

	/*  string 字符串的底层内存 特性。
		不可变意味着 两个字符串能安全地共用同一段 底层内存， 使得复制任何长度字符串的开销都低廉。
		类似地，字符串 s 及其子串 (如 s[7:]) 可以安全地共用数据， 因此子串生成操作的开销低廉。
		这两种情况下 都没有分配新内存。
	*/

	// 举个 字符串数组的例子 做对比， 数组里面的每个元素是可以 赋值修改的。
	arr := []string{"gold", "yellow", "red"}
	fmt.Println(arr[0])  // "gold"
	arr[0] = "green"    // 赋值给 arr[0] , 数组中的第一个元素
	fmt.Println(arr)  	// "[green yellow red]"   // 数组中的原始数据

	str := strings.Join(arr, "/")  // 将字符串数组 转换成一个字符串，数组中的元素 以 "/"  分割
	fmt.Println(str)  //  "green/yellow/red"


}
