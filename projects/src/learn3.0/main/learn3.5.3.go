package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// 学习第三章--3.4-字符串--UTF-8
func main() {
	fmt.Println("learn3.5.3")

	/*  UTF-8
		UTF-8 以字节为单位对 Unicode 码点作为变长编码。 UTF-8 是现行的一种 Unicode 标准，
		由 Go 的两位创建者 Ken Thompson 和 Rob Pike 发明。每个文字符号用 1 ~ 4 个字节表示，
		ASCII 字符的编码仅占 1个字节， 而其他常用的文书字符的编码 只是2 或 3 个字节。
		一个文字符号编码的首字节的高位指明了 后面还有多少字节。若 最高位为 0， 则标示着它是
		7位的 ASCII 码，其文字符号的编码仅占 1 字节， 这样就与传统的 ASCII 码一致。
		若 最高几位是 110， 则文字符号的编码占用 2个字节， 第二个字节以 10 开始。 更长的编码以此类推

			0xxxxxxx                               文字符号 0 ~ 127         (ASCII)
			110xxxxx 10xxxxxx                      128 ~ 2047               少于128个未使用的值
			1110xxxx 10xxxxxx 10xxxxxx             2048 ~ 65535             少于2048个未使用的值
			11110xxx 10xxxxxx 10xxxxxx 10xxxxxx    65536 ~ 0x10ffff         其它未使用的值

		变长编码的字符串无法按下标 直接访问第 n 个字符， 然后有失有得，UTF-8 换来许多有用的特性。
		UTF-8 编码紧凑， 兼容ASCII， 并且自同步：最多追溯 3字节，就能定位一个字符的起始位置。
		UTF-8 还是前缀编码，因此它能从左向右解码而 不产生歧义， 也无需超前预读。
		于是查找文字符号仅 须搜索它自身的字节，不必考虑前文内容。
		文字符号的字典字节顺序与 Unicode 码点顺序一致 (Unicode设计如此)，因此按 UTF-8 编码排序自然就是对文字符号排序。
		UTF-8 编码本身不会嵌入 NUL 字节(0 值), 这便于某些程序语言用 NUL 标记字符串结尾。

		Go 的源文件总是以 UTF-8 编码，同时，需要用 Go 程序操作的文本字符串也优先采用UTF-8 编码。
		Unicode 包具备针对单个文字符号的函数 (例如 区分字母和数字，转换大小写)，
		而 unicode/utf8 包则提供了按 UTF-8编码和解码文字符号的函数。

		许多 Unicode 字符难以直接从键盘输入；有的看起来十分相似几乎无法分辨；有些甚至不可见。
		Go语言中，字符串字面量的转义让我们得以用 码点的值来指明 Unicode 字符。
		有两种形式， \uhhhh 表示 16位码点值， \Uhhhhhhhh 表示 32位码点值，其中每个 h 代表有一个十六进制数字；
		32位形式的码点值几乎不需要用到。 这两种形式都以 UTF-8 编码表示出给定的码点。
		因此， 下面几个字符串字面量都表示长度为 6 字节的相同串。

			"世界"
			"\xe4\xb5\x96\xe7\x95\x8c"
			"\u4e16\u754c"
			"\U00004e16\U0000754c"

		后面三行的转义序列用不同形式表示第一行的字符串， 但实质上它们的字符串值都一样。

		Unicode 转义符也能用于文字符号。下列字符是等价的：
			'世'   '\u4e16'   '\U00004e16'

		码点值小于256的文字符号可以写成单个十六进制数转义的形式， 如 'A' 写成 '\x41'，
		而更高的码点值则必须使用\u 或 \U 转义。 这就导致，'\xe4\xb8\x96' 不是合法的文字符号，
		虽然这三个字节构成某个有效的 UTF-8 编码码点。

	 */

	var c uint8 = 'A'
	var c2 uint8 = '\x41'   // 16进制转换为 十进制值为  65
	fmt.Printf("ASCII : %d\n", c)    // 65
	fmt.Printf("ASCII : %c\n", c2)   // A      // 将 '\x41' 以字符的形式打出来就是  'A'

	var substr string = "商品"
	var containStr string = "特惠商品"
	var containStr2 string = "价格"
	var str string = "商品-数量-价格-备注"
	var suffixStr string = "备注"
	isPrefix := strings.HasPrefix(str,substr)  // 是否为 str 的前缀字符串。
	isSuffix := strings.HasSuffix(str, suffixStr)  // 是否为 str 的后缀字符串
	isContain := strings.Contains(str, containStr)  // 是否包含相同的字符串 (将 containStr 整体比较)
	isContain2 := strings.Contains(str, containStr2)
	fmt.Println("isPrefix : ", isPrefix)   // "true"
	fmt.Println("isSuffix : ", isSuffix)   // "true"
	fmt.Println("isContain : ", isContain)   // "false"
	fmt.Println("isContain2 : ", isContain2) // "true"

	// lenStr()
	lenStr2()

}

/*
	由于 UTF-8的优良特性， 许多字符串操作都无须解码。我们可以直接判断某个字符串是否为另一个的前缀：
	以下为 Go语言 strings.HasPrefix(...),   strings.HasSuffix(...)  源码


// HasPrefix tests whether the string s begins with prefix.
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// 或许它是否为另一个的子字符串：  ( Go 语言中 String.Contains(...) 里面的具体实现并不是如下)
func Contains(s, substr string) bool{
	for i:=0; i<len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

按 UTF-8 编码的文本的逻辑同样也适用 原生字节序列，但其他编码则无法如此。
(上面的函数取自 strings 包，其实Contains 函数的具体实现适用了 散列方法让搜索更高效。)

 */


/*
	另一方面，如果我们真的要逐个逐个处理 Unicode 字符， 则必须使用其他编码机制。
	考虑我们第一个例子的字符串 (见 3.5.1节)，它包含两个东亚字符。
 */

func lenStr() {

	s:= "Hello, 世界"
	fmt.Println(len(s))   // 13
	// 字符串中的 字节长度  （中文 “世界” 占6个字节）

	fmt.Println(utf8.RuneCountInString(s))  // 9
	// 按照UTF-8 解读, 本质是 9个码点或文字符号的编码 （中文 “世界” 就2个码点）

	// 我们需要 UTF-8 解码器来处理这些字符， unicode/utf8 包就具备一个：

	for i:=0; i<len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t%d\n", i , r , size)   // 对应的下标值， 字符， 所占字节
		i += size
	}

	/*   打印信息如下：

		0       H       1
		1       e       1
		2       l       1
		3       l       1
		4       o       1
		5       ,       1
		6               1
		7       世      3
		10      界      3

	*/

	// 通过下标取值
	fmt.Println(s[7:10])  // "世"
	fmt.Println(s[10:])   // "界"

	/*
		每次 DecodeRuneInString 的调用都返回 r (文字符号本身) 和 一个值 (表示 r 按 UTF-8 编码所占用的字节数)。
		这个值用来更新下标 i ， 定位字符串内的下一个文字符号。 可是按此方法，我们总是需要使用上例中的循环模式。
		所幸， Go 的 range 循环也适用于字符串，按UTF-8 隐式解码。 注意，对于非 ASCII 文字字符，下标增量大于1.
	 */

	for i,r:= range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	/*  打印信息如下：

		0       'H'     72
		1       'e'     101
		2       'l'     108
		3       'l'     108
		4       'o'     111
		5       ','     44
		6       ' '     32
		7       '世'    19990
		10      '界'    30028

	 */

	// 我们可用简单的 range 循环统计字符串中的文字符号数目， 如下所示：
	n :=0
	for _, _ = range "Hello, 世界" {
		n++
	}
	fmt.Printf("%d\n", n)    //  9   // 以 UTF-8编码解读

	// 与其他形式的 range 循环一样， 可以忽略没用的变量:
	f := 0
	for range s {
		f++
	}
	fmt.Printf("%d\n", f)    //  9   // 以 UTF-8编码解读

	// 或者，直截了当地调用 utf8 的函数
	count := utf8.RuneCountInString("Hello, 世界")
	fmt.Printf("%d\n", count)   //  9
}


func lenStr2() {
	s:= "プログラム"    // 中文“程序” 的日语，  unicode
	for i,r:= range s {  //  以 UTF-8 编码 遍历 字符串， 同 中文操作一样
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	/*  打印信息如下：

		0       'プ'    12503
		3       'ロ'    12525
		6       'グ'    12464
		9       'ラ'    12521
		12      'ム'    12512

	*/

	/*
		以下 Printf 里的谓词 %x (注意，% 和 x 之间有空格) 以十六进制数形式输出， 并在每两个数位间插入空格。
	 */
	fmt.Printf("% x\n", s)   //  "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"

	r:=[]rune(s)
	fmt.Printf("%x\n", r)    //  "[30d7 30ed 30b0 30e9 30e0]"  // UTF-8 编码, 与上面对比着来看

	/*
		如果把文字符号类型的 slice 转换成一个字符串， 它会输出各个文字符号的 UTF-8 编码拼接结果：
	*/
	fmt.Println(string(r))   //  "プログラム"

	// 若将一个整数值转换成 字符串， 其值按文字符号类型解读，并且产生代表该文字符号值的 UTF-8 码：
	fmt.Println(string(65))   // "A"    // 而不是 "65"
	fmt.Println(string(0x4eac))  // "京"

	// 如果文字符号值非法， 将被专门的替换字符取代
	fmt.Println(string(1234567))   // �

	/*
		之前提到过，文本字符串作为按 UTF-8 编码的 Unicode 码点序列解读，很大程度是出于习惯，
		但为了确保使用 range 循环能正确处理字符串，则必须要求而不仅仅是按照习惯。
		如果字符串含有任意二进制数，也就是说， UTF-8 数据出错，而我们对它做 range 循环，会发生什么？

		每次 UTF-8 解码器读入一个不合理的字节，无论是显式调用 utf8.DecodeRuneInString,
		还是在 range 循环内隐式读取，都会产生一个专门的 Unicode 字符 '\uFFFD' 替换它，
		其输出通常是个黑色六角形或 类似砖石的形状， 里面有个白色问号。
		如果程序碰到这个文字符号值，通常意味着，生成字符串数据的系统上游部分在处理文本编码方面存在瑕疵。

		UTF-8是一种分外便捷的交互格式， 而在程序内部使用文字字符类型可能更加方便，
		因为它们大小一致，便于在数组和 slice 中用下标访问。
	 */

}

