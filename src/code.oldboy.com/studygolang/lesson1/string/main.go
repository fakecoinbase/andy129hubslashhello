package main

import (
	"fmt"
	"strings"
)

// 字符串操作
func main() {

	fmt.Println("string")

	s1 := "alexdsb"
	fmt.Println(len(s1)) // "7"

	// 字符串拼接
	s2 := "Python"
	fmt.Println(s1 + s2) // "alexdsbPython"

	// Sprintf, 格式化拼接然后返回值
	s3 := fmt.Sprintf("%s---%s", s1, s2)
	fmt.Println(s3) // "alexdsb---Python"

	// 字符串分割
	ret := strings.Split(s1, "x")
	fmt.Println(ret) // "[ale dsb]"
	// 判断是否包含
	ret2 := strings.Contains(s1, "dsb")
	fmt.Println(ret2) // "true"

	// 判断前缀和后缀
	ret3 := strings.HasPrefix(s1, "alex")
	ret4 := strings.HasSuffix(s1, "sb")
	fmt.Println(ret3, ret4) // "true true"

	// 求子串的位置
	s4 := "applepen"
	fmt.Println(strings.Index(s4, "p"))     // "1"
	fmt.Println(strings.LastIndex(s4, "p")) // "5"

	// Join
	a1 := []string{"Python", "PHP", "JavaScript", "Ruby", "Golang"}
	fmt.Println(strings.Join(a1, "-")) // "Python-PHP-JavaScript-Ruby-Golang"
}
