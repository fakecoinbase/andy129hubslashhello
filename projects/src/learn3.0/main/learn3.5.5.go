package main

import (
	"fmt"
	"strconv"
)

// 学习第三章--3.5-字符串--字符串和数字的相互转换
func main() {
	fmt.Println("learn3.5.5")

	x:= 123
	y:= fmt.Sprintf("%d", x)    // 注意区别于 fmt.Printf(),  Sprintf 函数的作用就是将 Printf 打印的类容给转换为字符串
	y2:= strconv.Itoa(x)   // 将整型转换为 字符串
	/*  go语言 官方 Itoa() 函数的实现方式，其实是调用  FormatInt(int64(i), 10)
		// Itoa is equivalent to FormatInt(int64(i), 10).
		func Itoa(i int) string {
			return FormatInt(int64(i), 10)
		}
	 */
	fmt.Println(y, y2)   // "123 123"

	// FormatInt 和 FormatUint 可以按不同的进位制格式化数字：
	fmt.Println(strconv.FormatInt(int64(x), 2))    // "1111011"    // 整型 123 的 二进制形式

	/*
		fmt.Println 里的谓词 %b、%d、%o 和 %x 往往比 Format 函数方便， 若要包含数字以外的附加信息， 它就尤其有用：
	 */
	s:= fmt.Sprintf("x=%b", x)    // %b ，二进制转义符
	fmt.Println(s)   // "x=1111011"

	/*
		strconv 包内的 Atoi 函数或 ParseInt 函数用于解释表示整数的字符串， 而 ParseUint 用于无符号整数：
	 */
	a,err := strconv.Atoi("123")    //  x 是整型
	b,err2 := strconv.ParseInt("123", 10,64)  //  十进制，最长为 64位。
	fmt.Println(a,err)     // "123 <nil>"
	fmt.Println(b,err2)    // "123 <nil>"

	k,err3 := strconv.Atoi("fsdfe")  // 尝试转换无效的字符串会 提示错误
	if err3 != nil {  // 一般通过判断 err 是否为 nil ，来判断转换成功还是失败
		fmt.Println("转换失败了")
	}
	fmt.Println(k,err3)     // "0 strconv.Atoi: parsing "fsdfe": invalid syntax"     无效的语法
	/*	当转换失败时，返回值 k, err3 的值是 ：(当转换失败时，返回的值 默认为 0)
		k 值为 0
		err3 值为： strconv.Atoi: parsing "fsdfe": invalid syntax
	 */

	/*
		ParseInt 的第三个参数指定结果必须匹配何种大小的整型；例如， 16 表示 int16，
		而 0 作为特殊值表示 int. 任何情况下，结果 y 的类型总是 int64，可将他另外转换成较小的类型。
		有时候， 单行输入有字符串和数字依次混合构成，需要用 fmt.Scanf 解释，可惜 fmt.Scanf 也许不够灵活，
		处理不完整或 不规则输入时尤甚。
	 */

}

