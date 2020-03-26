package main

import (
	"fmt"
	"math"
)

func main() {

	var a int = 10
	var b int = 077  // 八进制，  0* 8 的二次方 加上  7 * 8 的一次方 加上 7*8的 零次方 =  0+56+7 = 63
	var c int = 0xff // 16进制

	// Println 默认是  十进制输出
	fmt.Println(a, b) // "10 63"
	// %b  二进制输出
	fmt.Printf("%b\n", a) // "1010"
	// %o  八进制输出
	fmt.Printf("%o\n", b) // "77"
	// %x  十六进制输出
	fmt.Printf("%x\n", c) // "ff"
	// 默认十进制输出
	fmt.Println(c) // "255"
	// 求 c 变量的内存地址
	fmt.Printf("%p\n", &c) // "0xc0000120e0"

	// float32 与 float64 的最大值
	fmt.Println(math.MaxFloat32) // "3.4028234663852886e+38"
	fmt.Println(math.MaxFloat64) // "1.7976931348623157e+308"
	// math 包里的 Pi 常量
	fmt.Println(math.Pi) // "3.141592653589793"

	// 字符串转义
	fmt.Println("\"c:\\go\"") //""c:\go""    带双引号字符串
	var s1 = "单行文本"
	// ``  多行文本，里面的字符串 完全输出，不需要转义
	var s2 = ` 这
		是  "zemy" \n
		多行 \t
		文本
		！
	`
	fmt.Println(s1) // "单行文本"
	fmt.Println(s2)
	/*  s2 打印：
			 这
	                是  "zemy" \n
	                多行 \t
	                文本
	                ！
	*/
}
