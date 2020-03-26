package main

import "fmt"

func main() {
	fmt.Println("复习第一节")

	s := "hello中国"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) // 默认按照 ASCII 码去打印
		//fmt.Printf("%c\n", s[i]) // 打印字符，中文会乱码
	}
	fmt.Println("------------------------------------------------------")
	for i, char := range s {
		fmt.Printf("%d\t%c\n", i, char)
	}

}
