package main

import "fmt"

// goto 的使用
func main() {

	// 常规 跳出 for 循环
	flag := false
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			//
			if i == 2 && j == 2 {
				flag = true
				break // break 跳出单层循环
				// continue   // continue 继续下一个循环
			}
			fmt.Printf("%d--%d\n", i, j)
		}
		// 通过标志位来判断是否跳出外层 for 循环
		if flag {
			break
		}
	}
	fmt.Println("两层for循环结束")

	fmt.Println("-----------------------------------")

	// goto 跳出for 循环
	// 与 上面 break 加 标志位 flag 效果一致
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			//
			if i == 2 && j == 2 {
				goto label // 跳转到指定的 代码块
			}
			fmt.Printf("%d--%d\n", i, j)
		}
	}

label: // 定义一个标签 label
	fmt.Println("两层for循环结束")

	// break 语句可以结束 for, switch 和 select 的代码块。 break 语句还可以在语句后面添加标签

	fmt.Println("-----------------------------------")

	// break 后面的指定的标签，必须要在 break 语句之前定义，不能像 goto 的 label 一样，定义在后面
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {

			if j == 2 {
				break BREAKDEMO1 // break 后面加标签，也可指定跳转到 标签指定的位置
			}
			fmt.Printf("%v--%v\n", i, j)
		}
	}

	fmt.Println("................")

	// continue 标签 与  break 标签 用法一致

	//

}
