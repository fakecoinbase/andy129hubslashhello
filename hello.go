package main

// import 以下为导入包的两种形式(单行定义与 括号内统一定义)

// 导入包的 写法1
//import "fmt"
//import "os"

// 导入包的 写法2
import (
	"fmt"
	"os"
	"strings"
)

func main(){
	// 对 test1()等其他示例函数调用出现在其 声明之前。 函数和其他包级别的实体可以以任意次序声明。
	// test1()
	// test2()
	// test3()
	test4()
}

// hello world 学习程序
func test1(){
	 fmt.Println("Hello,世界")
}

// for 循环练习
func test2(){
	var s,sep string

	for i:=1; i<len(os.Args);i++{
		s += sep + os.Args[i]
		sep = " "

	}
	fmt.Println("打印"+ s)
	fmt.Println("打印"+ os.Args[0])
}

// for 循环练习 空标识符等
func test3(){
	s,sep := "",""
	for _, arg := range os.Args[1:]{
		s+= sep+ arg
		sep = " "
	}
	fmt.Println(s)
}

// 使用 string.Join
func test4(){
	fmt.Println("打印"+strings.Join(os.Args[1:]," "))
}

