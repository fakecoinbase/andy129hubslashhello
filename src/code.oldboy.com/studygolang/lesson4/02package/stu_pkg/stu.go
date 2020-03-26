package stu_pkg

// 这是一个第三方包，用于测试 init() 函数的调用顺序
import "fmt"

var number = 10

// 初始化
func init() {
	fmt.Println("这是 stu_pkg 里 stu.go 的 init() : number : ", number)
}
