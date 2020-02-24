package main

import (
	"fmt"
	"math"
	"unsafe"
)

const (
	deadbeef = 0xdeadbeef      // 无类型整数，值为 3735928559
	a1 = uint32(deadbeef)      // uint32, 值为 3735928559
	b1 = float32(deadbeef)     // float32, 值为 3.7359286e+09   (3735928576 向上取整)
	c1 = float64(deadbeef)     // float64, 值为 3.735928559e+09  (3735928559精确值)
	// d1 = int32(deadbeef)    // 编译报错：溢出, int32 无法容纳 deadbeef 的常量值。
	// e1 = float64(1e309)     // 编译报错：溢出, constant 1e+309 overflows float64
	// f1 = uint(-1)           // 编译报错：溢出, constant -1 overflows uint
)


// 学习第三章--3.6.2-无类型常量
/*
	Go的常量自由特别之处。虽然常量可以是任意基本数据类型，如 int 或 float64,
	也包括具名的基本类型 (如 time.Duration)，但是许多常量并不从属某一具体类型。
	编译器将这些从属类型待定的常量表示成某些值，这些值比基本类型的数字精度更高，
	且算术精度高于原生的机器精度。可以认为它们的精度至少达到 256位。从属类型待定的常量共有 6种，
	分别是 无类型布尔、无类型整数、无类型文字符号、无类型浮点数、无类型复数、无类型字符串。


	Go语言源码： 下面定义了很多常量，它们的精度大部分都是超过了 int64, float64 等范围，但常量定义可以完整的保存它们。

	// Mathematical constants.
	const (
		E   = 2.71828182845904523536028747135266249775724709369995957496696763 // https://oeis.org/A001113
		Pi  = 3.14159265358979323846264338327950288419716939937510582097494459 // https://oeis.org/A000796
		Phi = 1.61803398874989484820458683436563811772030917980576286213544862 // https://oeis.org/A001622

		Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974 // https://oeis.org/A002193
		SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931 // https://oeis.org/A019774
		SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779 // https://oeis.org/A002161
		SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038 // https://oeis.org/A139339

		Ln2    = 0.693147180559945309417232121458176568075500134360255254120680009 // https://oeis.org/A002162
		Log2E  = 1 / Ln2
		Ln10   = 2.30258509299404568401799145468436420760110148862877297603332790 // https://oeis.org/A002392
		Log10E = 1 / Ln10
	)

 */
func main() {

	fmt.Println("learn3.6.2")

	// unsureConstFunc()
	// unsureConstFunc2()
	// unsureConstFunc3()
	// unsureConstFunc4()
	// unsureConstFunc5()
	// unsureConstFunc6()
	unsureConstFunc7()
}

func unsureConstFunc(){
	/*
		借助推迟确定从属类型，无类型常量不仅能暂时维持更高的精度， 与类型已确定的常量相比，
		它们还能写进更多表达式而无需转换类型。比如，3.6.1 示例中 ZiB 和 YiB 的值过大，
		用哪种整型都无法存储，但它们都是合法常量并且可以用在下面的表达式中：
	 */
	// fmt.Println(YiB/ZiB)   // 1024

	// 再例如, 浮点型常量 math.Pi 可用于任何需要浮点值或复数的地方：
	var x float32 = math.Pi    // 无类型浮点数 赋值给  float32
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println("math.Pi 值所占字节数：",unsafe.Sizeof(math.Pi))   //  8
	/*
			8位，可能是 float64类型，但是其实 float64类型也无法完整保留 math.Pi 的精度值，
			所以 无类型浮点数 math.Pi 的作用就在这里了, 可以保留完整的精度值。
	 */
	fmt.Println("float32 x 值所占字节数：",unsafe.Sizeof(x))       //  4
	fmt.Println("float64 y 值所占字节数：",unsafe.Sizeof(y))       //  8
	fmt.Println("complex128 z 值所占字节数：",unsafe.Sizeof(z))    //  16
	fmt.Println("float32 : ",x)      // float32 :  3.1415927
	fmt.Println("float64 : ",y)      // float64 :  3.141592653589793
	fmt.Println("complex128 : ",z)   // complex128 :  (3.141592653589793+0i)

}

/*
	字面量的类型由语法决定。 0、0.0、0i 和 '\u0000' 全部表示相同的常量值，
	但类型相异，分别是： 无类型整数、无类型浮点数、无类型复数 和 无类型文字符号。
	类似地， true 和 false 是 无类型布尔值， 而在字符串字面量 则是无类型字符串。

 */
func unsureConstFunc2(){
	/*
		根据除法运算中操作数的类型，除法运算的结果可能是 整型或 浮点型。
		所以， 常量除法表达式中，操作数选择不同的字面写法会影响结果：

		除法运算符 / ,  结果是 整型还是浮点型，取决于 操作数是否都为 整型
	 */
	var f float64 = 212
	fmt.Println((f - 32)*5/9)       // "100"   //  (f-32)*5 的结果是 float64 型
	fmt.Println(5/9*(f - 32))       // "0"     //  5/9 的结果是无类型整数，0
	fmt.Println(5.0/9.0*(f - 32))   // "100"   //  5.0/9.0  的结果是无类型浮点数
}

// 只有常量才可以是无类型的。
/*
	若将无类型常量声明为 变量(如下面的第一句语句所示)，
	或在类型明确的变量赋值的右方出现无类型变量(如下面的其他三条语句所示)，
	则常量会被 隐式转换成该变量的类型。
 */
func unsureConstFunc3(){

	var f float64 = 3 + 0i    // 无类型复数  赋值给  float64
	f = 2                     // 无类型整数  赋值给  float64
	f = 1e123                 // 无类型浮点数 赋值给 float64
	f = 'a'                   // 无类型  赋值给  float64
	fmt.Println(f)

	// 上述语句与下面的语句等价：
	/*
		var f float64 = float64(3+0i)
		f = float64(2)
		f = float64(1e123)
		f = float64('a')
	*/
}

/*
	不论隐式 或 显式，常量从一种类型转换成 另一种，都要求目标类型能够表示原值。
	实数和复数允许舍入取整：
 */
func unsureConstFunc4(){

	fmt.Printf("%d\n",a1)
	fmt.Printf("%g\n",b1)
	fmt.Printf("%g\n",c1)
	//fmt.Println(d1)
	//fmt.Println(e)
	//fmt.Println(f)
}

// 变量声明 (包括短变量声明)中，假如没有显式指定类型，无类型常量会隐式转换成该变量的默认类型，如下例所示：
func unsureConstFunc5(){
	i:= 0          // 无类型整数；隐式 int(0)
	r:= '\000'     // 无类型文字字符；隐式 rune('\000')
	f:= 0.0        // 无类型浮点数；隐式 float64(0.0)
	c:= 0i         // 无类型整数；隐式 complex128(0i)

	fmt.Println(i,r,f,c)   // 0 0 0 (0+0i)
}

/*
	注意各类型的不对称性：无类型整数可以转换成 int, 其大小不确定，
	但无类型浮点数和 无类型复数被转换成大小明确的 float64 和 complex128.
	Go 语言中，只有大小不明确的 int 类型， 却不存在大小不确定的 float类型和 complex 类型，
	原因是， 如果浮点型数据的大小不明，就很难写出正确的数值算法。
 */
func unsureConstFunc6(){
	/*
		要将变量转换成 不同的类型，我们必须将无类型常量显式转换为期望的类型，
		或在声明变量时指明想要的类型，如下例所示：
			var i = int8(0)
			var i int8 = 0
	 */
}

/*
	在将无类型常量转换为 接口值时 (见第7章)，这些默认类型就分外重要，
	因为它们决定了接口值的动态类型。
 */
func unsureConstFunc7(){
	// 回顾 %T 转义符的意思： 任何值的类型 （输出 这个值是什么类的）
	fmt.Printf("%T\n", 0)       // int
	fmt.Printf("%T\n", 0.0)     // float64
	fmt.Printf("%T\n", 0i)      // complex128
	fmt.Printf("%T\n", '\000')  // int32
	fmt.Printf("%T\n", math.Pi)      // float64
}