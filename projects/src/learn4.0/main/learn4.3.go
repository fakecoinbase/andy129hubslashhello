package main

import (
	"fmt"
	"math"
	"strconv"
)

// 学习第四章-复合数据类型--4.3-map
func main() {
	fmt.Println("learn4.3")

	NaNTest()   // 回顾之前学习的 NaN 知识点

}

// 回顾前面学到的 NaN 知识点
func NaNTest(){
	var a int
	fmt.Println(a)         // 0
	fmt.Println(-a)        // 0
	//fmt.Println(1/a)       // 运行报错：panic: runtime error: integer divide by zero
	//fmt.Println(-1/a)      // 运行报错：panic: runtime error: integer divide by zero
	//fmt.Println(a/a)       // 运行报错：panic: runtime error: integer divide by zero

	fmt.Println("------------------------float运算, +Inf, -Inf, NaN-----------------------------")
	// 对比发现，只有 float 运算结果才会出现 +Inf , -Inf,  NaN 这些特殊值。

	var z float64
	fmt.Println(z)        // 0
	fmt.Println(-z)       // -0
	fmt.Println(+z)       // 0
	fmt.Println(1/z)      // +Inf
	fmt.Println(-1/z)     // -Inf
	fmt.Println(z/z)      // NaN


	fmt.Println(math.IsInf(-1/z, -1))    // true, sign < 0 ,则返回是否为 负无穷大
	fmt.Println(math.IsInf(1/z, 1))      // true, sign > 0 ,则返回是否为 正无穷大
	fmt.Println(math.IsInf(1/z, 0))      // true, sign == 0, 无论是正无穷大或者是 负无穷大，只要是无穷大则返回 true
	fmt.Println(math.IsInf(z, 0))        // false

	// 下面为 Go语言 内部对于  math.IsInf(f float64, sign int) bool 的解释说明

	/*
		// IsInf reports whether f is an infinity, according to sign.
		// If sign > 0, IsInf reports whether f is positive infinity.
		// If sign < 0, IsInf reports whether f is negative infinity.
		// If sign == 0, IsInf reports whether f is either infinity.
		func IsInf(f float64, sign int) bool {
			// Test for infinity by comparing against maximum float.
			// To avoid the floating-point hardware, could use:
			//	x := Float64bits(f);
			//	return sign >= 0 && x == uvinf || sign <= 0 && x == uvneginf;
		return sign >= 0 && f > MaxFloat64 || sign <= 0 && f < -MaxFloat64
	}

	 */

	fmt.Println("------------------------NaN-----------------------------")

	fmt.Println(math.IsNaN(z/z))      // true


	nan := math.NaN()    //   NaN returns an IEEE 754 ``not-a-number'' value.
	fmt.Println(nan)     // "NaN" , "not-a-number"
	// fmt.Println(nan == "NaN")  // 编译错误，nan 是 float 类型，所以不能与 string 字符串做比较

	fmt.Println(nan == nan, nan < nan, nan > nan)   // false false  false
	/*  Go语言 对于 NaN() 函数的解释
	// NaN returns an IEEE 754 ``not-a-number'' value.
	func NaN() float64 { return Float64frombits(uvnan) }
	*/

	/*
		math.IsNaN 函数判断其参数是否是 非数值， math.NaN 函数则返回 非数值(NaN).
		在数字运算中，我们倾向于将 NaN 当做信号值 (sentinel value), 但直接判断具体的
		计算结果是否为 NaN 可能导致潜在错误，因为 与 NaN 的比较总不成立 (除了 !=, 它总是与 == 相反)
	 */

	// 下面纯属自己突发奇想，并不是教材里面要求，所以最好还是不用  nan 做比较，就算比较，它除了 != ，其他总是与 == 相反。
	// 如果要强行来判断，则可以进行下面这种操作，先将 float 转换为 string ,然后再与 "NaN" 对比。
	str := strconv.FormatFloat(nan,'g',-1,64)
	fmt.Println(str)              // "NaN"
	fmt.Println(str == "NaN")     // true
}
