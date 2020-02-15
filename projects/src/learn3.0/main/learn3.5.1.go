package main

import "fmt"

// 学习第三章--3.4-字符串--字符串字面量

/*	 ASCII 与  Unicode 的区别

		1.什么是 ASCII

		 ASCII 是用来表示英文字符的一种编码规范。每个ASCII字符占用1 个字节，
		因此，ASCII 编码可以表示的最大字符数是255（00H—FFH）。这对于英文而言，是没有问题的，
		一般只什么用到前128个(00H--7FH,最高位为0)。而最高位为1 的另128 个字符（80H—FFH）被称为“扩展ASCII”，
		一般用来存放英文的制表符、部分音标字符等等的一些其它符号。

		2.什么是 UNICODE

      	Unicode与ASCII一样也是一种字符编码方法，它占用两个字节（0000H—FFFFH）,容纳65536 个字符，
		这完全可以容纳全世界所有语言文字的编码。在Unicode 里，所有的字符都按一个字符来处理， 它们都有一个唯一的Unicode 码。

 */

/*
	字符串的值可以直接写成字符串字面量 (string literal)， 形式上就是带双引号的字节序列：
		"Hello,世界"

	因为 Go 的源文件总是按 UTF-8 编码，并且习惯上 Go 的字符串会按 UTF-8解读，
	所以在源码中我们可以将 Unicode 码点写入字符串 字面量。
 */

func main() {
	fmt.Println("learn3.5.1")

	/*
		在带双引号的字符串字面量中， 转义序列以 反斜杠(\) 开始，可以将任意值的字节插入字符串中。
		下面是一组转义符， 表示 ASCII 控制码，如 换行符、回车符 和 制表符。

		\a		"警告"或 响铃
		\b		退格符
		\f		换页符
		\n		换行符
		\r		回车符(指返回行首)
		\t		制表符
		\v		垂直制表符
		\'		单引号 (仅用于文字字符字面量 '\'')
		\"		双引号 (仅用于 "..." 字面量内部)
		\\      反斜杠

	 */

	fmt.Printf("警告或 响铃\a")
	fmt.Printf("\n")
	fmt.Printf("退格符a\b"+"abc")
	/*
		"退格符abc"  // 目前测试来看，退格符会将字符串倒退一个字节，
		如遇到中文则会倒退一个中文字符, (顺便提一下，go 语言中一个中文占 3个字节)
		fmt.Printf("退格符\b"+"abc")   // "退格abc"
	 */
	fmt.Printf("\n")
	fmt.Printf("换行符\n")            //  这个很好理解，经常用
	fmt.Printf("\n")
	fmt.Printf("回车符\r")
	fmt.Printf("\n")
	fmt.Printf("制表符\t")
	fmt.Printf("\n")
	fmt.Printf("垂直制表符\t")
	fmt.Printf("\n")
	b := '\''     //  单引号 仅用于 文字字符字面量 '\'' (仅作用与  单引号里面)
	fmt.Printf("单引号 : %c ",b)  //  '
	fmt.Printf("\n")
	str := string(b)
	fmt.Printf("将单引号字符强制转换为字符串形式 : "+str)
	fmt.Printf("\n")
	// fmt.Printf("\'单引号\'")  // 编译报错，  单引号 仅用于 文字字符字面量 '\'' (仅作用与  单引号里面)
	fmt.Printf(str+"单引号"+str) // '单引号'
	fmt.Printf("\n")
	fmt.Printf("打印出带引号的字符 : %q ", b)  // '\''
	fmt.Printf("\n")
	fmt.Printf("\"双引号\"")  // "双引号"   // 仅用于 "..."  （双引号中的字符串打印）
	fmt.Printf("\n")

	/*
		源码中的字符串也可以包含 十六进制 或 八进制的 任意字节。十六进制的 转义字符写成 \xhh 的形式，
		h 是 十六进制数字 (大小写皆可)， 且必须是 两位。 八进制 的转义字符 写成 \000 的形式，、
		必须使用 三位八进制数字 (0 ~7)，且不能超过 \377。 这两者都表示单个字节， 内容是给定值。
		后面， 我们将看到 如何将数值形式的 Unicode 码点 嵌入字符串字面量。
	 */

	fmt.Printf("十六进制的转义字符 : \x45")     //  E
	fmt.Printf("\n")
	fmt.Printf("八进制的转义字符 : \124")       //  T
	fmt.Printf("\n")

	/*
		原生的字符串字面量 的书写形式是 `.....`, 使用反引号而不是 双引号，也不是单引号。
		原生的字符串字面量内， 转义序列不起作用； 实质内容与字面写法 严格一致，包括 反斜杠和 换行符，
		因此，在程序源码中， 原生的字符串字面量可以展开多行。 唯一的特殊处理是 回车符会被删除 （换行符会保留），
		使得同一字符串在所有平台上的值 都有相同， 包括习惯在文本文件存入 换行符的系统。

		正则表达式往往含有 大量反斜杠， 可以方便地写成原生的字符串字面量。
		原生的字面量也适用于 HTML 模板、JSON 字面量、命令行提示信息， 以及需要多行文本表达的场景。
	 */

	const GoUsage = `Go is a tool for managing Go source code. \n
			Usage:
				go command [arguments]
					.....  \n
			`
	fmt.Printf(GoUsage)

	/*
			Go is a tool for managing Go source code. \n
	                        Usage:
	                                go command [arguments]
	                                        .....  \n

				注意，在以上 原生的字符串字面量 中， \n 不起作用，所以就按照 原样输出了。
	 */

}
