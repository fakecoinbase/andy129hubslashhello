package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// flag 包
func main() {

	// test1() // 打印命令行参数

	// test2()

	test3()
}

// os.Args  获取命令行参数
func test1() {
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}

/*	执行：go run main.go rename a.txt b.txt
	打印：
	args[0]=C:\Users\andy\AppData\Local\Temp\go-build702668834\b001\exe\main.exe
	args[1]=rename
	args[2]=a.txt
	args[3]=b.txt
*/

// flag.Type()
func test2() {
	// flag.String() ， flag.Int() 返回的都是 指针类型
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")

	// 需要注意的是，此时 name、age、 married 、delay 均为对应类型的指针。
	fmt.Printf("%T\t%v\n", name, *name)       // "*string 张三"
	fmt.Printf("%T\t%v\n", age, *age)         // "*int    18"
	fmt.Printf("%T\t%v\n", married, *married) // "*bool   false"
	fmt.Printf("%T\t%v\n", delay, *delay)     // "*time.Duration  0s"

}

// flag.TypeVar()
func test3() {
	var name string
	var age int
	var married bool
	var delay time.Duration

	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")

	flag.Parse() // 执行到这一步， 可以使用  go run main.go -help

	fmt.Println(name, age, married, delay)
	//  执行到这一步，命令行输入指令： go run main.go -name 杨浩 -d 1h -age 30 -married=false
	//  可打印： 杨浩 30 false 1h0m0s

	// 返回命令行参数后的其它参数
	fmt.Println(flag.Args()) // 例如：go run main.go -name 杨浩 -d 1h -age 30 -married=false xxxx ,    xxxx 是属于可解析范围之外的其它参数，所以打印 [xxxx]

	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg()) // 例如： 上面这个 xxxx,   只有一个其它参数，所以是 1
	// 返回使用的命令行参数个数
	fmt.Println(flag.NFlag()) // 例如：go run main.go -name 杨浩 -d 1h -age 30 -married=false ,   name, d, age, married , 共4个参数，所以返回 4

	// 命令行什么参数都没有， 则会直接打印 默认值
	/*
		张三 18 false 0s
		[]
		0
		0
	*/

}

/*  命令行参数格式有如下几种：

通过以上两种方法定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析。

	支持的命令行参数格式有以下几种：

	-flag xxx （使用空格，一个-符号）
	--flag xxx （使用空格，两个-符号）
	-flag=xxx （使用等号，一个-符号）
	--flag=xxx （使用等号，两个-符号）
	其中，布尔类型的参数必须使用等号的方式指定。

	Flag解析在第一个非flag参数（单个”-“不是flag参数）之前停止，或者在终止符”–“之后停止。
*/
