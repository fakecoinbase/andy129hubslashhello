package admin_pkg

// 这是一个第三方包，用于测试 init() 函数的调用顺序
// 匿名导入包 (包前面加入 _) , 不需要调用 第三包里的字段，但需要第三方包初始化，可以使用 _, 例如：数据库连接, init()
import (
	"fmt"

	_ "code.oldboy.com/studygolang/lesson4/02package/stu_pkg"
)

const Money = 1000

// 初始化
func init() {
	fmt.Println("这是 admin_pkg 里 admin.go 的 init() : Money : ", Money)
}
