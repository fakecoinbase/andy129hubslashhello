package main

import "fmt"

func main() {

	// 算术运算符
	n1 := 19
	n2 := 3
	n1++
	n2--
	fmt.Println(n1)       // "20"
	fmt.Println(n2)       // "2"
	fmt.Println(n1 == n2) // "false"
	fmt.Println(n1 > n2)  // "true"
	fmt.Println(n1 >= n2) // "true"
	fmt.Println(n1 < n2)  // "false"
	fmt.Println(n1 <= n2) // "false"
	fmt.Println(n1 != n2) // "true"

	// 逻辑运算符
	a := true
	b := false

	fmt.Println(a && b) // "false"

	fmt.Println(a || b) // "true"

	fmt.Println(!a) // "false"

	// 位运算符， 操作二进制
	fmt.Printf("13的二进制%b \n", 13) // "1101"
	fmt.Printf("3的二进制%b \n", 3)   // "11"

	// & : 按位与， 比较二进制，两个都为1，则结果为1，否则为 0
	fmt.Println(13 & 3) // "1"，      //  1101   &  0011 比较
	// | : 按位或， 比较二进制，两个其中有一个为1， 则为 1
	fmt.Println(13 | 3) // "15"       //  1101   |  0011 比较
	// ^ : 按位异或，比较二进制，两个不一致，则为 1
	fmt.Println(13 ^ 3) // "14"
	// <<: 左移,
	fmt.Println(3 << 10) // "3072",  先把 3 转换成 二进制:  11，  然后左移10位，  110000000000,  1*2的11次方 + 1*2的10次方 = 2048+1024 = 3072
	// >>: 右移,
	fmt.Println(3 >> 1) // "1",  先把 3 转换成 二进制:  11， 然后右移1位， 二进制就变成了 1， 转换为十进制就是 1

	// 赋值运算符
	num := 10
	x := 2
	num = num / x
	// 同理可简写成下面
	// num /= x
	fmt.Println(num) // "5"

}
