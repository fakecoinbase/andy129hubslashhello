package main

import "fmt"

func main() {
	fmt.Println("char")

	s1 := "Golang"
	c1 := 'G'           // ASCII 码下占一个字节 (8位，8bit)
	fmt.Println(s1, c1) // "Golang 71"

	s2 := "中国"
	c2 := '中'           //  UTF-8 编码下 一个中文占用3个字节
	fmt.Println(s2, c2) // "中国 20013"

	s3 := "hello沙河"
	fmt.Println(len(s3)) // "11"

	for i := 0; i < len(s3); i++ {

		fmt.Printf("%c\n", s3[i]) // 中文会打印出乱码
	}

	fmt.Println()
	// for range 循环是 按照 rune 类型去遍历的
	for k, v := range s3 {
		fmt.Printf("%d\t%c\n", k, v)
	}
	/*  打印内容如下：
	0       h
	1       e
	2       l
	3       l
	4       o
	5       沙
	8       河
	*/

	// 强制类型转换
	s5 := "big"
	byteArray := []byte(s5)
	fmt.Println(byteArray)
	byteArray[0] = 'p' // 修改值
	// 将字节数组强制转换为 string
	s5 = string(byteArray)

	fmt.Println(s5) // "pig"

	s6 := "hello"

	reverseStr([]byte(s6))
	fmt.Println()
}

func reverseStr(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
