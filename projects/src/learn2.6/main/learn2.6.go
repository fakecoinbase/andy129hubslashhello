package main


import (
	"fmt"
	 "learn2.6/tempconv"
)

/* 目前总结初始化顺序 ？？？  待后续验证
	包级别变量以及依赖的函数或其他
	导入包
	包 (init)
	包 (main)
 */

/*  包的初始化
包的初始化按照在程序中导入的顺序来进行，依赖顺序优先，每次初始化一个包。
因此，如果包 p 导入了包 q, 可以确保 q 在 p 之前已完全初始化。初始化过程是自下向上的，
main 包最后初始化。在这种方式下，在程序的 main函数开始执行前，所有的包已初始化完毕。
 */

/*  导包的路径问题
之所以一直找不到 导入的包  tempconv ，是因为需要在 golang 编辑器里设置 GOPATH
举个例子：  GOPATH 为 G:\Goworkspace\projects
设置完 GOPATH 之后，需要在 projects 目录下建立 src 目录， src 目录下再创建一个又一个单独的项目, 例如： learn2.6
 */

/*  导包的名称冲突的问题, 假如要导入多个包，包名有冲突的情况下
	import (
		"learn/tempconv"
		m "learn2.6/tempconv"   // 同样都是 tempconv , 则需要在其中一个前面指定一个 替代名称
	)
 */

// 包级别变量
var a= b+c
var b= f()
var c =1

func main() {

	// 常量名字也大写字母开头的 都是 包级别的常量，可以直接访问。
	fmt.Printf("Brrrr!  %v\n",tempconv.AbsoluteZeroC) // "Brrrr! -273.15℃"

	// 通过包名去调用包里面的函数
	fmt.Println(tempconv.CToF(tempconv.BoilingC))  // "212℉"

	// 包的初始化从初始化包级别的变量开始, 这些变量按照声明的顺序开始初始化，在依赖已解析完毕的情况下，根据依赖的顺序进行。
	/*  先按顺序初始化，一但发现有依赖关系，例如， var a= b+c ，依赖于 b,c ，所以就先初始化 b,c
		var b = f(),  b 又依赖于 c ，所以先初始化 var c =1，所以顺序是先 c 到 b 再到 a
		// 包级别变量
		var a= b+c
		var b= f()
		var c =1

	 */
	fmt.Printf("a: %d, b: %d, c: %d\n",a, b, c) // "Brrrr! -273.15℃"
}

func f() int {
	fmt.Printf("-------包级别变量依赖 \n")
	return c+1
}

// 包的初始化函数，可以用来进行一些 预处理的操作
func init(){
	fmt.Printf("-------init \n") // 包初始化，可以在这里进行一些预处理的操作， 例如数组，集合等
}
