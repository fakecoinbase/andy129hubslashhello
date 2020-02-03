package main

import "fmt"

func main() {
	example1()
}

func example1() {
	x:= "hello!"
	for i:= 0;i< len(x);i++ {
		x:= x[i]
		if x!= '!' {
			x:= x + 'A' - 'a'
			// 应该是对字符进行了 ASCII 码数值的操作，使其进行了 大小写的转换，然后以 %c 字符形式输出，则成为了字符
			fmt.Printf("%c", x) // 依次打出每一个字符，在终端中最终效果： HELLO
		}
	}
}


