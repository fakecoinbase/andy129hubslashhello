package main

import (
	"fmt"
	 "learn2.6/tempconv"
)

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

func main() {

	// 常量名字也大写字母开头的 都是 包级别的常量，可以直接访问。
	fmt.Printf("Brrrr!  %v\n",tempconv.AbsoluteZeroC) // "Brrrr! -273.15℃"

	// 通过包名去调用包里面的函数
	fmt.Println(tempconv.CToF(tempconv.BoilingC))  // "212℉"

}
