package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// //学习第三章--3.3-复数
/*   数学 范畴

	加法法则
	复数的加法法则：设z1=a+bi，z2=c+di是任意两个复数。两者和的实部是原来两个复数实部的和，
	它的虚部是原来两个虚部的和。两个复数的和依然是复数。

	乘法法则
	复数的乘法法则：把两个复数相乘，类似两个多项式相乘，结果中i2= -1，
	把实部与虚部分别合并。两个复数的积仍然是一个复数。

	除法法则
	复数除法定义：满足 的复数 叫复数a+bi除以复数c+di的商。
	运算方法：将分子和分母同时乘以分母的共轭复数，再用乘法法则运算，

	其他法则
		.....

 */
func main() {

	fmt.Println("learn3.3")
	// complexVar()

	// 曼德勃罗特集 是人类有史以来做出的最奇异，最瑰丽的几何图形，曾被称为“上帝的指纹”
	// 输出到终端上 全是乱码
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0,0,width,height))
	for py:=0; py<height; py++ {
		y := float64(py)/height*(ymax - ymin) + ymin
		for px:=0; px<width; px++ {
			x:= float64(px)/width*(xmax-xmin)+ xmin
			z:= complex(x,y)
			// 点(px,py) 表示复数值 z
			img.Set(px,py,mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)  // 在终端上全是乱码，如何解决？
}

/*
	Go具备两种大小的复数 complex64 和 complex128, 二者分别由 float32 和 float64构成。
	内置的 complex 函数根据给定的实部和虚部 创建复数， 而内置的 real 函数和 imag 函数则分别提取复数的实部和 虚部
 */
func complexVar(){
	var x complex128 = complex(1,2)   // 1 是实部， 2 是虚部
	var y complex128 = complex(3,4)   // 3 是实部， 4 是虚部
	fmt.Println(x) 	 	// "1+2i"
	fmt.Println(y)  	// "3+4i"
	fmt.Println(x*y) 	// "(-5+10i)"
	/*  x*y 的计算过程  (复数 乘法法则， i的平方 == -1)
		(1+2i)*(3+4i) == 3+6i+4i+8i的平方 == 3+10i+8*(-1) == -8+3+10i == -5+10i
	 */
	fmt.Println(real(x*y)) // "-5"     // real 函数 取实部值
	fmt.Println(imag(x*y)) // "10"     // imag 函数 取虚部值

	/*
		i 在这里是 go语言中代表虚数的标志， 源码中，如果在浮点数或 十进制整数后面紧接着写字母 i ， 如 3.141592i 或 2i ,
		它就变成一个虚数， 表示一个实部为 0 的复数。

		注意：  虚部 与 虚数的区别，  复数里 分 实部与虚部， 虚部中 有虚数。
		例如 虚数 2i 可以看成是 实部为 0 的一个复数。（0+2i）
	 */
	fmt.Println(1i * 1i)  // "(-1+0i)"  // i的平方 = -1

	/*
		根据常量运算规则， 复数常量可以和其他常量相加（整型或浮点型，实数和虚数皆可），
		折让我们可以自然地写出复数， 如 i+2i , 或 等价地， 2i+1 。 前面 x 和 y 的声明可以简写为：
	 */
	a := 1+2i
	b := 3+4i
	fmt.Println(a,b)

	/*
		可以用 == 或 != 判断复数是否等值。 若两个复数的实部和 虚部都相等， 则它们相等。
		math/cmplx 包提供了复数运算所需的库函数， 例如复数的平方根函数和复数的 幂函数。
	 */
	fmt.Println(cmplx.Sqrt(-1))    // "(0+1i)"    // 求复数的平方根，返回的是一个 complex128 类型的复数
	fmt.Println(cmplx.Sqrt(2))     // "(1.4142135623730951+0i)"
	fmt.Println(cmplx.Sqrt(-5i))   // "(1.5811388300841898-1.5811388300841898i)"
	fmt.Println(cmplx.Sqrt(3+4i))  // "(2+1i)"
}

// 曼德勃罗集
func mandelbrot(z complex128) color.Color{
	const iterations = 200
	const contrast = 15

	var v complex128
	for n:= uint8(0); n<iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}