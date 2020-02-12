package main

import "fmt"

// 学习第三章--3.4-布尔值
func main() {
	fmt.Println("learn3.4")
	boolVarFunc()
	
}

func boolVarFunc(){

	var b bool
	fmt.Println(b)  // "false"   // bool 类型默认值为 false

	fmt.Println(!true == false)   // "true"    // 一元操作符 (!) 表示逻辑相反，因此 !true 就是 false
	fmt.Println(2 < 1)  // "false"   // 比较操作符 (如 == 和 < ) 也能得出布尔值结果。

	/*
		布尔值可以由运算符 &&(AND) 以及  ||(OR) 组合运算， 这可能引起短路行为：
		如果运算符左边的操作数已经能直接确定总体结果，  则右边的操作数不会计算在内，所以下面的表达式是安全的：
	 */

	var s string
	fmt.Println("a"+s+"b")  // "ab"  // 得出了 string 默认值为 ""  (注意字符长度为 0)
	// fmt.Println(s[0])    // 运行报错， 由于s 的字符长度为 0，所以 s[0] 则无法获取字符，导致宕机行为。

	if (s != "" && s[0] == 'x'){
		fmt.Println(s[0])
	}

	/*
	 	以上运行正常， && 左边先运算， 发现 s != ""  这个条件为 false ,所以就不会执行左边的 比较。
		&& 只要有一个条件为 false ，则不会执行另外一个条件。
	 */

	// 在多个运算符中，优先级是怎样的，看下面的例子：
	var c uint8
	if 'a'<= c && c<= 'z' ||
		'A'<= c && c<= 'Z' ||
		'0'<=c && c<= '9' {

		// ... ASCII 字母或数字
	}
	/*
		因为 && 较 || 优先级更高 （助记窍门： &&表示逻辑乘法， || 表示逻辑加法），
		所以如上的条件无线加 圆括号。
	 */

	var m bool
	// if  m:=fasle {    // 编译错误， 布尔值无法隐式转换成数值 （如 0 或 1），反之也不行。如下情况就有必要使用显式 if ：
	if m {  // 显式  if 条件
		fmt.Println(m)
	}


}

// bool 类型 常用场景如下
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool {
	return i != 0    // 返回比较的结果， true 或 false
}

