package main

import "fmt"

func main() {
	var a = 100
	var b = "沙河哪吒"
	var c = false
	fmt.Println(a, b, c)
	// %v 俗称占位符， 任意类型， 是什么类型打印什么类型的值
	fmt.Printf("a=%v \n", a)    // "a=100"
	fmt.Printf("a的类型是%T \n", a) // "a=int"
	// %% ,  格式化输出 百分号
	fmt.Printf("%d%%\n", a) // "100%"

	// %5d : 整型长度为5， 右对齐，左边留白
	fmt.Printf("|%5d|\n", a) // "|  100|"
	// %-5d: 整型长度为5， 左边对齐，右边留白
	fmt.Printf("|%-5d|\n", a) // "|100  |"
	// %05d: 整型长度为5，右边对齐，左边补0
	fmt.Printf("|%05d|\n", a) // "|00100|"
	// %-05d: 整型长度为5，左边对齐，右边留白  （为什么不是补0？？）
	fmt.Printf("|%-05d|\n", a) // "|100  |"

	// %-5d  与   %-05d，  效果一致
	str1 := fmt.Sprintf("%-5d", a)
	str2 := fmt.Sprintf("%-05d", a)
	fmt.Println(str1 == str2) //  "true"

	f1 := 3.146564645545454

	fmt.Printf("%f\n", f1)   // "3.146565" // 进行四色五入
	fmt.Printf("%.2f\n", f1) // "3.15"     // 小数点保留两位（会进行四色五入）
	fmt.Printf("%.2g\n", f1) // "3.1"      // 总共保留几个数字

	s1 := "这是一个字符串\""
	fmt.Printf("%s\n", s1) // "这是一个字符串""    // 将转义字符转换成字符串， 双引号里的内容打印出来
	fmt.Printf("%q\n", s1) // ""这是一个字符串\"""     //  将 s1 原样输出，包括引号

	// 指定长度，字符串右对齐
	fmt.Printf("|%.20s|\n", s1) // "|            这是一个字符串"|"
	// 指定长度，字符串左对齐
	fmt.Printf("|%-20s|\n", s1) // "|这是一个字符串"            |"
	// %.ns, 如果n 大于 s1 的长度，则正常输出， 如果 n 小于 s1 的长度，例如：%.5s ： 只保留五个字符， 类似于截取字符串的功能
	fmt.Printf("|%.5s|\n", s1) // "|这是一个字|"

}
