package main

import (
	"fmt"
	"math"
	"strconv"
)

//学习第三章--3.2-浮点数
func main() {

	fmt.Println("learn3.2")
	// floatVarFunc()
	// floatPrintFunc()
	floatPrintFunc2()
}

func floatVarFunc(){
	var f32Max = math.MaxFloat32
	var f64Max = math.MaxFloat64

	fmt.Println("float32 最大值：", f32Max)  // 3.4028234663852886e+38
	fmt.Println("float64 最大值：", f64Max)  // 1.7976931348623157e+308

	/*
		十进制下， float32 的有效数字大约是 6位， float64 的有效数字大约是15位。
		绝大多数情况下，应优先选用 float64，因为除非格外小心，否则 float32的运算会迅速积累误差。
		另外，float32 能精确表示的正整数范围有限，例如以下情况
	 */
	var f float32 = 16777216  // float32 能精确表示的正整数范围有限。
	m := f+1
	fmt.Println(f)  // 1.6777216e+07
	fmt.Println(m)  // 1.6777216e+07
	fmt.Println(f == m) // "true"
}

func floatPrintFunc() {

	const a = 3.1415923587555456545554646549645654654654654
	fmt.Println(a)  // 3.1415923587555454

	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34
	fmt.Println(Avogadro)  // 6.02214129e+23
	fmt.Printf("%g\n", Avogadro)  // 6.02214129e+23
	fmt.Println(Planck) // 6.62606957e-34
	fmt.Printf("%g\n", Planck)  // 6.62606957e-34
	fmt.Printf("%e\n", Planck)  // 6.62606957e-34

	/*
		浮点值能方便地通过 Printf 的谓词 %g 输出， 该谓词会自动保持足够的精度，
		并选择最简洁的表达方式，但是对于数据表， %e(有指数) 或 %f(无指数)的形式可能更合适。
		这三个谓词都能掌握输出宽度和数值精度。
	 */
	var temp = 23.4559489439
	fmt.Printf("%g\n", temp)  // 23.4559489439
	fmt.Printf("%e\n", temp)  // 2.345595e+01
	fmt.Printf("%f\n", temp)  // 23.455949
	/*
		%e 与 %f 的区别(有指数与无指数的区别是否是 科学记数法 e 或 E) ?
	 */
}

func floatPrintFunc2(){
	for a:= 0;a<8;a++ {

		var exp float64 = math.Exp(float64(a))
		var str = floatConvertStr(exp)
		fmt.Printf("exp : %g , len : %d\n", exp, len(str))
		fmt.Printf("a = %d   e^x = %8.3f\n", a, exp)
		/* 解释说明
		%8.3f 的意思：  输出8个字符的宽度，并且保留小数点后3位，例如 下面的打印信息 :

			a = 0   e^x =    1.000       // 这条信息， 1.000前面还有三个空格位，为的是补齐要求输出的 8位。
			a = 7   e^x = 1096.633       // 这条信息，刚好8位 (小数点也占 一位)

		 */
	}

	/*	打印信息：

		exp : 1 , len : 1
		a = 0   e^x =    1.000
		exp : 2.718281828459045 , len : 17
		a = 1   e^x =    2.718
		exp : 7.38905609893065 , len : 16
		a = 2   e^x =    7.389
		exp : 20.085536923187668 , len : 18
		a = 3   e^x =   20.086
		exp : 54.598150033144236 , len : 18
		a = 4   e^x =   54.598
		exp : 148.4131591025766 , len : 17
		a = 5   e^x =  148.413
		exp : 403.4287934927351 , len : 17
		a = 6   e^x =  403.429
		exp : 1096.6331584284585 , len : 18
		a = 7   e^x = 1096.633
	*/

}

// 将 float 转换为字符串
func floatConvertStr(a float64) string{
	return strconv.FormatFloat(a, 'f', -1, 64)
}
