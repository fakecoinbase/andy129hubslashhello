package main

import "fmt"

// for 循环
func main() {

	// for 循环可以通过 break, goto, return , panic 语句强制退出循环
	age := 18
	for age > 0 {
		fmt.Println(age)
		age--
	}

	// 无线循环
	/*
		for {
			// 循环语句
		}
	*/

	// for range (键值循环)
	for k, v := range "hello沙河" {
		fmt.Printf("%d\t%c\n", k, v)
	}
}
