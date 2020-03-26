package main

import (
	"fmt"
	"math"
	"strings"
	str "strings"

	"code.oldboy.com/studygolang/lesson4/01package/math_pkg"
	mpk "code.oldboy.com/studygolang/lesson4/01package/math_pkg"
)

/*  总结：

1，引用包路径： GOPATH 配置的路径下,  从 src 目录下开始（不包含 src 目录）,  路径使用 "/"
2, math_pkg 目录下面的 .go 文件里面命名的 package ，必须统一为  package math_pkg
3, import 导入多个包时，  内置包与 引用第三方包 之间要 空一行
4, 可以给 包(内置包或 其他第三方包) 定义一个别名
5，包的别名定义 与 原生包的引用，可以同时存在 (前提是 代码引用时分别用到了 别名和 原生包名)
	，例如 import "strings" ,  import str "strings"


	/*  如下：
			import (
				"fmt"
				"math"

				"code.oldboy.com/studygolang/lesson4/01package/math_pkg"

				"strings"
			)

		或者：

		import (
				"fmt"
				"math"
				"strings"

				"code.oldboy.com/studygolang/lesson4/01package/math_pkg"
			)

*/

// package 包的使用
func main() {
	fmt.Println("package")

	// 对比 test1(), test2() ，看看 总结3
	//test1()
	//test2()
	test3()
}

// 调用 非go语言内置导入包 的方法
// import "code.oldboy.com/studygolang/lesson4/01package/math_pkg"
func test1() {
	ret := math_pkg.Add(4, 5)
	fmt.Println(ret)
}

// 调用 go语言内置包 的方法
// (并给 strings 定义一个别名)
func test2() {

	fmt.Println(math.Pi)

	// 给内置 strings 包 定义一个别名: str
	fmt.Println(str.Contains("stttsf", "s"))
	// 如何又用到了 strings 包，则还会导入一次 strings 包
	fmt.Println(strings.Contains("hhfdf", "d"))

	/*
		import (
				...
			"strings"
			str "strings"
				...
	*/
}

// 给第三方包定义一个 别名，使用别名调用
// import mpk "code.oldboy.com/studygolang/lesson4/01package/math_pkg"
func test3() {

	stu := mpk.Student{}
	stu.Name = "yang"
	stu.Age = 44
	fmt.Println(stu.Name)
	fmt.Println(stu.Age)

}
