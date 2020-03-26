package main

import "fmt"

func main() {

	if age2 := 28; age2 > 18 {
		fmt.Println("成年了！")
	}
	// fmt.Println(age2)
	// 编译错误： undefine:  age2 ,  说 age2 未定义, 这是因为 if 条件里面定义的变量，只能用于 if 语句块里，作用域仅限于 if 里，还有 switch ,for 与 if 类似。

}
