package main

// init()  示例：
// 总结：
/*
	1, 包执行的过程:
					package main  --->
					import ""  --->

					按照导入顺序，去实例化，
					进入其它包，依旧按照 先 查看 import 包， 进入其它引用包，
					然后回到本包 声明全局变量，常量，结构体，然后 init(),
					直到所有引用的包 声明变量，常量，结构体，init()结束之后，再回到本包。  --->

					回到main 包进行全局变量声明, 常量，结构体等 --->
					init() --->
					main().

*/

// 匿名导入包 (包前面加入 _)
// import _ "包"
import (
	"fmt"

	_ "code.oldboy.com/studygolang/lesson4/02package/admin_pkg"
	_ "code.oldboy.com/studygolang/lesson4/02package/third_pkg"
)

var today = "星期天"

const Week = 7

type student struct {
	name string
	age  int
}

// init() 是一个初始化函数， 包被导入的时候会自动执行
func init() {
	fmt.Println(Week) // "7"
	fmt.Println("------init()")
}

func main() {
	fmt.Println("-----main()函数")

	/*  打印信息如下：

	这是 stu_pkg 里 stu.go 的 init() : number :  10              // stu_pkg 是 被 admin_pkg 包 引用的包，所以先执行
	这是 admin_pkg 里 admin.go 的 init() : Money :  1000         // admin_pkg 是被 main 包里 引用的包，按照引入包的顺序
	这是 third_pkg 里 0.go 的 init()         // third_pkg 是被 main 包里引用的包，所以开始初始化 third_pkg 里所有的.go 文件，0.go 是 third_pkg 里面的文件，文件名为 数字 0
	这是 third_pkg 里 01.go 的 init()
	这是 third_pkg 里 1.go 的 init()
	这是 third_pkg 里 11.go 的 init()
	这是 third_pkg 里 2.go 的 init()
	这是 third_pkg 里 a.go 的 init()         // a.go 是 third_pkg 里面的文件，对比 forth.go,third.go, 应该是按照首字母的顺序 一 一执行
	这是 third_pkg 里 forth.go 的 init()
	这是 third_pkg 里 third.go 的 init()
	7                                       // init 里面打印了 Week的值，代表全局变量的初始化要 先与 init()
	------init()
	-----main()函数
	*/
}
