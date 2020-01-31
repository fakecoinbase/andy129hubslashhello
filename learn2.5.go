package main

import "fmt"

// 类型声明：  type name underlying-type
// type 声明定义一个新的命名类型，它和某个已有类型使用同样的底层类型。
// Celsius 与 Fahrenheit 即使使用相同的底层类型 float64, 它们也不是相同的类型，所以它们不能使用算术表达式进行比较和合并
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func main() {

	fmt.Printf("%g\n", BoilingC - FreezingC) // 类型相同，可以进行算术表达式
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF - CToF(FreezingC)) // CToF()将Celsius 类型转换为 Fahrenheit，可以进行算术表达式
	// fmt.Printf("%g\n", boilingF - FreezingC)  // Fahrenheit 与 将Celsius 类型不匹配，不能进行 算术操作

	var c Celsius
	var f Fahrenheit = 2.0
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// fmt.Println(c == f) // 编译出错，类型不匹配
	fmt.Println(c == Celsius(f)) // 将 Fahrenheit 类型转换为 Celsius,  类型相同，可以进行比较

	c = FToC(212.0)
	//fmt.Println("Celsius类型转换为字符串： "+c.String())  // 调用字符串  String(),将 Celsius 类型转换为一个字符串
	fmt.Printf("直接转义符输出： %v\n", c) // 不需要显示调用字符串，直接用 转义符输出
	fmt.Printf("字符串转义符输出： %s\n", c) // 使用 %s 转义符输出,  调用字符串  String()
	fmt.Println(c)  // 调用字符串  String()
	fmt.Printf("%g\n",c)   // 以浮点类型 打印, 不调用字符串  String()
	fmt.Println(float64(c))  // 不调用字符串 String()
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)  // Fahrenheit(t) 显式类型转换，而不是函数调用。
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32)*5 / 9)  // Celsius(t) 显式类型转换，而不是函数调用。
}

/*
通过以上两个函数， 说明 类型转换不改变类型值的表达方式，仅改变类型。
对于每个类型 T， 都有一个对应的类型转换操作 T(x) 将值 x 转换为 类型 T。
 */

// Celsius 参数 c 出现在函数名字前面，就规定了 这个函数只能被 Celsius 使用，其他类型不能调用
// 类似于 java 里面的 override 重写功能， 重新修改 String() 里面的功能
// 很多类型都声明这样一个 String 方法，在变量通过 fmt包 作为字符串输出时，它可以控制类型值的显示方式
func (c Celsius) String() string {
	return fmt.Sprintf("%g℃",c)
}
