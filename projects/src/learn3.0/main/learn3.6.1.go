package main

import (
	"fmt"
	"runtime"
	"strconv"
	"unsafe"
)

/*
	常量的声明可以使用常量生成器 iota，它创建一系列相关值，而不是逐个值显式写出。
	常量声明中，iota 从 0开始取值，逐项加 1

	下例取自 time 包，它定义了 Weekday 的具名类型，并声明每周的7天为该类型的常量，
	从 Sunday 开始，其值为 0. 这种类型通常称为 枚举型(enumeration，或 缩写成 enum)。

	以下为 Go 语言中对 iota 的定义：

	// iota is a predeclared identifier representing the untyped integer ordinal
	// number of the current const specification in a (usually parenthesized)
	// const declaration. It is zero-indexed.
	const iota = 0 // Untyped int.

 */

type Weekday int    // 指定类型名字
const(
	Sunday Weekday = iota   // 初始化值，iota 从 0开始取值。
	Monday                  // 根据 iota 常量生成器，后面以此类推依次加 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)


/*   位移符号 扩展学习：https://blog.csdn.net/catoop/article/details/100716948

	<<（向左位移）
	将数字先转换为二进制，然后向左移动N位，右边补0

	2 << 3 = 16
	先将数字转换为二进制
	2 = 0010
	然后 0010 向左移动3位后右边补 0 结果为 10000（0010 中的第一个1前面只有2个0所以整体左移3位需要在右边补1个0，得10000）
	将 10000 转换为十进制后为 16
 */
type Flags uint
const(
	FlagUp Flags = 1 << iota     // 1 << 0
	FlagBroadcast                // 1 << 1
	FlagLoopback                 // 1 << 2
	FlagPointToPoint             // 1 << 3
	FlagMulticast                // 1 << 4
)

// 下面更复杂，声明的常量表示 1024 的幂。
const(
	start = 1 << (10 * iota)    // 1 << (10 * 0) == 1 << 0
	KiB                         // 1 << (10 * 1) == 1 << 10
	MiB                         // 1 << (10 * 2) == 1 << 20
	GiB                         // 1 << 30
	TiB                         // 1 << 40
	PiB                         // 1 << 50
	EiB                         // 1 << 60
	ZiB                         // 1 << 70
	YiB                         // 1 << 80
)

// 学习第三章--3.6.1-常量生成器 iota
func main() {
	fmt.Println("learn3.6.1")

	//iotaFunc()
	//iotaFunc2()
	//iotaFunc3()
	checkIntB()
}

func iotaFunc(){

	fmt.Println(Sunday)       // 0
	fmt.Println(Monday)       // 1
	fmt.Println(Tuesday)      // 2
	fmt.Println(Wednesday)    // 3
	fmt.Println(Thursday)     // 4
	fmt.Println(Friday)       // 5
	fmt.Println(Saturday)     // 6
}

func iotaFunc2(){

	fmt.Println(FlagUp)              // 1
	fmt.Println(FlagUp == 1 << 0)    // "true"
	fmt.Println(FlagBroadcast)       // 2
	fmt.Println(FlagLoopback)        // 4
	fmt.Println(FlagPointToPoint)    // 8
	fmt.Println(FlagMulticast)       // 16
}

func iotaFunc3(){
	// cpu()
	fmt.Println(start)
	fmt.Println("默认int占有字节数：",unsafe.Sizeof(start))   // unsafe.Sizeof()  查询数据类型所占用的字节数
	fmt.Println(KiB)
	fmt.Println(MiB)
	fmt.Println(GiB)
	fmt.Println(TiB)      // 1 << 40
	/*
		从这里开始超过  1 << 32,  32位 int 的最大取值范围，那这里为什么没有运行报错，唯一的解释是 默认为  int64 类型，
		那为什么默认为 int64 呢?  看以下分析：

		查询网上资料，发现，int的位数是跟cpu有关系的。由于当前编译环境是 64位，所以默认分配 int64

			func cpu() {
				fmt.Println(runtime.GOARCH)    // amd64
				fmt.Println(strconv.IntSize)   // 64
			}
	 */
	fmt.Println(PiB)
	fmt.Println(EiB)
	//fmt.Println(ZiB)    // 1 << 70          //  已超过 1 << 64 ，  64位int 的最大值范围，运行报错
	//fmt.Println(YiB)    // 1 << 80          //  已超过 1 << 64 ，  64位int 的最大值范围，运行报错
}
/*
	查询网上资料，发现，int的位数是跟cpu有关系的。
	32位系统，对应着int为4个字节，
	64位系统，对应着int为8个字节。

	go语言对 GOARCH 的定义 ：

		// GOARCH is the running program's architecture target:
		// one of 386, amd64, arm, s390x, and so on.
		const GOARCH string = sys.GOARCH
 */
func cpu() {
	fmt.Println(runtime.GOARCH)    // amd64
	fmt.Println(strconv.IntSize)   // 64
}

// 查询各种int 类型在本机编译环境下的 占位 数,
func checkIntB(){
	var x int = 100
	fmt.Println(unsafe.Sizeof(x)) // 8
	/*
		当这种没有指定是 int32 还是 int64时，它会根据编译环境是多少位而改变，
		由于本机是在 64位(amd64) 的编译环境下，如果换做 32位的 编译系统，则会改变

			查询网上资料，发现，int的位数是跟cpu有关系的。
			32位系统，对应着int为4个字节，
			64位系统，对应着int为8个字节。
	 */

	var y int64 = 1
	fmt.Println(unsafe.Sizeof(y)) // 8
	var y1 int32 = 1
	fmt.Println(unsafe.Sizeof(y1)) // 4
	var z uint64 = 1
	fmt.Println(unsafe.Sizeof(z)) // 8
	var z1 uint32 = 1
	fmt.Println(unsafe.Sizeof(z1)) // 4
}
