package main

import "fmt"

// scan
func main() {

	// test1()
	// test2()
	test3()
}

// fmt.Scan() 学习
func test1() {
	var (
		name    string
		age     int
		married bool
	)
	// Scan 从标准输入扫描文本，读取由空白分隔的值保存到传递给本函数的参数中，换行符视为空白符。(换行符，空格，tab,均可以)
	fmt.Scan(&name, &age, &married)
	fmt.Println(name, age, married)

	// 示例1：（换行符视为 空白符来进行分隔）
	/*
		终端输入：	zhang
					33
					true
		打印：   zhang 33 true
	*/

	// 示例2：（多个空格，tab键 ，同样视为空白符来进行分隔）
	/*
		终端输入: 章              33 false
		打印   ： 章 33 false
	*/

	// 示例3（错误）：(输入类型中间有不匹配时会影响后面字段的值，即使终端输入时 married:true, 但是age:z 类型不匹配，导致 married 属性为默认值)
	/*
		终端输入： zhang z true
		打印   ： zhang 0 false
	*/

}

// fmt.Scanf() 学习
func test2() {
	var (
		name    string
		age     int
		married bool
	)
	// Scanf()  按照指定的格式进行输入
	// 1, 空格， 字段与字段之间，无论多少空格，tab 都行
	// 2, 空格，但是字段与 格式化字符 之间的空格，就必须严格执行，例如： name:   %s (name 与  %s 之间的空格，在终端输入时，不能省)

	fmt.Scanf("name:%s 			age:%d married:%t", &name, &age, &married)
	fmt.Println(name, age, married)

	// 示例1：(错误) // 输入时没有严格按照指定格式输入， 指定为 空格分隔， 但输入时却是 逗号
	/*
		终端输入: name:liu,age:13,married:false

		打印   ： liu,age:13,married:false 0 false

		// 打印信息解析：
		// 当我按下换行键时，程序退出, 解析时将 换行符当做 空格符。
		// 解析时是以空格键为分隔，所以就把 liu ---> 最后的信息 全部当成了 name 的值
	*/

	// 示例2： // 输入时加入多个空格，tab 键时， 不影响解析
	/*
		终端输入: name:liu age:14                         married:true
		打印   ： liu 14 true
	*/

	// 示例3： // 输入时
	/*
		终端输入: name:zhang age:ss married:true
		打印   ： zhang 0 false

		// 打印信息解析：
		// 输入类型中间有不匹配时会影响后面字段的值，即使终端输入时 married:true, 但是age:ss 类型不匹配，导致 married 字段为默认值)
	*/

	// 示例4：// 再次修改Scanf()里面的格式,  name 与 %s 之间，添加  3个空格
	// fmt.Scanf("name:   %s 			age:%d married:%t\n", &name, &age, &married)
	/*
		终端输入：name:jiu age:21 married:true
		打印    :  0 false

		// 打印信息解析：
		// 由于 指定格式为： name:   %s  所以就要严格执行， name 与 %s 有 三个空格就 必须要输入三个空格
		// 所以一开始就解析失败，导致后面 age , married 的属性都归为默认值, name 为string，所以默认值为 空字符串 ""

		// 所以，正确的输入应该是这样的：
		终端输入：name:   jiu age:333 married:true
		打印   : jiu 333 true

	*/
}

// fmt.Scanln
func test3() {
	var (
		name    string
		age     int
		married bool
	)
	// Scanln()  识别到回车 就结束
	// 1, 空格， 字段与字段之间，无论多少空格，tab 都行

	// 尝试在 Scanln() 函数里面指定格式，结果： 程序一运行，就退出。
	// fmt.Scanln("name:%s age:%d married:%t\n", &name, &age, &married)

	fmt.Scanln(&name, &age, &married)
	fmt.Println(name, age, married)

	// 示例1(错误): (输入一个字符串，然后回车，程序就结束了)
	/*
		终端输入：zhang
		打印   ： zhang 0 false
	*/

	// 示例2(错误): (中间输入错误类型，后面字段的结果就为默认值了)
	/*
		终端输入：zhang s true
		打印   ： zhang 0 false
	*/

	// 示例3(错误): (中间输入多个空格，正常解析)
	/*
		终端输入：                章      2323    false
		打印   ： 章 2323 false
	*/
}
