package third_pkg

// 这是一个第三方包，用于测试 init() 函数的调用顺序
import "fmt"

// 初始化
func init() {
	fmt.Println("这是 third_pkg 里 third.go 的 init()")
}
