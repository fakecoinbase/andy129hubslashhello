package main

import (
	"fmt"
	"math"
	"time"
)

// const pi = 3.14159     // 近似数； math.Pi 是更精准的近似值

// 与变量类似，同一个声明可以定义一系列常量，这适用于一组相关的值：
const (
	e = 2.71828182845904523536028747135266249775724709369995957496696763
	pi = 3.14159265358979323846264338327950288419716939937510582097494459
	IPv4Len = 4
	noDelay time.Duration = 0    // int64 类型
	timeout = 5 * time.Minute    // int64 类型
)

const (
	a = 1
	b
	c = 2
	d
)

// 学习第三章--3.6-常量
func main() {
	fmt.Println("learn3.6")

	//constFunc()
	//constFunc2()
	constFunc3()
}

/*
	常量是一种表达式，其可以保证在编译阶段就计算出表达式的值，并不需要等到运行时，
	    从而使编译器得以知晓其值。所有常量本质上都属于基本类型：布尔值、字符串或数字。
	常量的声明定义了具名的值，它看起来在语法上与变量类似， 但该值恒定，这防止了程序运行过程中的意外(或恶意)修改。
	    例如，要表示数学常量，像圆周率，在 Go 程序中用常量比变量更适合，因其值恒定不变：


 */
func constFunc(){

	fmt.Println(e)        // 2.718281828459045
	fmt.Println(pi)       // 3.141592653589793
	fmt.Println(math.Pi)  // 3.141592653589793

	/*
		许多针对常量的计算完全可以在编译时就完成，以减免运行时的工作量并让其他编译器优化得以实现。
		某些错误通常要在运行时才能检测到，但如果操作数是常量，编译时就会报错，例如 整数除以 0，
		字符串下标越界，以及任何产生无限大值的浮点数运算。 看下面的例子
	 */

	// e = 32423423523523    // 编译报错，常量除了初始化赋值 之外，其他情况下不能被赋值，这是常量的特性

	/*
		var a = 12
		b := a%0            // 编译不报错， 但运行报错
		fmt.Println(b)

		c := e%0            // 编译报错， 其中 e 是 常量，这里体现了 常量的特性，在编译时能及时发现一些错误
		fmt.Println(c)
	*/

	/*
		对于常量操作数，所有数学运算、逻辑运算和比较运算的结果依然是常量，
		常量的类型转换结果和某些内置函数的返回值，例如 len、cap、real、imag、complex 和 unsafe.Sizeof, 同样是常量。

		因为编译器知晓其值，常量表达式可以出现在涉及类型的声明中，具体而言就是数组类型的长度： ( IPv4Len 为定义的常量)

				func parseIPv4(s string) IP{
					var p [IPv4Len]byte
						// ...
				}
	*/
}

/*
	常量声明可以同时指定类型和值， 如果没有显示指定类型，则类型根据右边的表达式推断。
	下例中， time.Duration 是一种具名类型，其基本类型是 int64，time.Minute 也是基于 int64 的常量。以下注释为 go 语言源码

		// A Duration represents the elapsed time between two instants
		// as an int64 nanosecond count. The representation limits the
		// largest representable duration to approximately 290 years.
		type Duration int64

	下面声明的两个常量都属于 time.Duration 类型，通过 %T 展示：

 */
func constFunc2(){

	fmt.Printf("%T %[1]v\n", noDelay)       // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)       // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute)   // "time.Duration 1m0s"

	// %T  打印值所属 类型，  %[1]v ， 取用第一个参数，然后打印这个值的 内置格式

	fmt.Printf("%T %[1]d\n", timeout)   //  300000000000
	// 5 * 60 * 1000 * 1000 * 1000 =  300000000000  纳秒
	//  时，分，秒，毫秒，微妙，纳秒....
}

/*
	若同时声明一组变量，除了第一项之外， 其他项在等号右侧的表达式都可以省略，
	这意味着会复用前面一项的表达式及其类型。 例如：

		const (
			a = 1
			b
			c = 2
			d
		)

 */
// 更复杂的声明方式 请参考  3.6.1--常量生成器 iota
func constFunc3(){
	fmt.Println(a,b,c,d)    // "1 1 2 2"
}




