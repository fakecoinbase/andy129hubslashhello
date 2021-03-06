package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

const sample = ` .bd.b2.3d.bc.20.e2.8c.98`    // 网上参考的示例 未达到效果
// 学习第三章--3.5-字符串--字符串和字节 slice
/*   4个标准包对字符串操作特别重要： bytes、strings、strconv 和 unicode.

	strings 包提供了许多函数，用于搜索、替换、比较、修整、切分与连接字符串。
	bytes 包也有类似的函数，用于操作字节 slice([]byte 类型，其某些属性和字符串相同)。
	由于字符串不可变，因此按增量方式构建字符串会导致多次内存分配和复制。
	这种情况下，使用 bytes.Buffer类型会更高效，范例见后。
	strconv 包具备的函数，主要用于转换布尔值、整数、浮点数为与之对应的字符串形式，
	或者把字符串转换为布尔值、整数、浮点数，另外还有为字符串添加 / 去除引号的函数。
	unicode 包备有判别文字符号值特性的函数， 如 IsDigit、IsLetter、IsUpper 和 IsLower。
	每个函数以单个文字符号值作为参数， 并返回布尔值。若文字符号值是英文字母，转换函数（如 ToUpper 和 ToLower）
	将其转换成指定的大小写。 上面所有函数都遵循 Unicode 标准对字母数字等的分类原则。 strings 包也有类似的函数，
	函数名也是 ToUpper 和 ToLower ， 它们对原字符串的每个字符做指定变换， 生成并返回一个新字符串。

 */
func main() {
	fmt.Println("learn3.5.4")

	//stringPrint()
	//stringPrint2()
	//stringPrint3()
	//stringReverse()
	//stringSplit()
	//strAndByteArr()
	//intsToStringFunc()
	runeFunc()
}

func stringPrint(){

	fmt.Printf("%x\n", sample)

	b := 'A'
	fmt.Printf("%c\n", b)   // "A"
	fmt.Printf("%q\n", b)   // "'A'"
	fmt.Printf("%d\n", b)   // "65"
	fmt.Printf("%b\n", b)   // "1000001"

}

func stringPrint2(){

	s := "hello, 世界"
	fmt.Println(s, &s)  // hello, 世界 0xc0000401f0
	str := s[7:]   // 创建了一个新的变量，并分配一个新的地址
	fmt.Println(str, &str)  // 世界 0xc000040210
	str = str[0:3]   // 未改变 str 地址，只改变了 str 地址所指向的值
	fmt.Println(str, &str)  // 世 0xc000040210

}

func stringPrint3(){

	strResult := NumberFormat("1234567898.55")
	fmt.Println(strResult)
}

//格式护数值    1,234,567,898.55
func NumberFormat(str string) string {
	length := len(str)
	if length < 4 {
		return str
	}
	arr := strings.Split(str, ".") //用小数点符号分割字符串,为数组接收
	length1 := len(arr[0])  // "1234567898"
	if length1 < 4 {
		return str
	}
	count := (length1 - 1) / 3
	for i := 0; i < count; i++ {
		// 将整段数据分成 3个为一组处理。
		arr[0] = arr[0][:length1-(i+1)*3] + "," + arr[0][length1-(i+1)*3:]   // 这里处理是核心关键
		/*  注意这里 arr[0]的写法
			arr[0] 返回的是 string 类型，然后 string 类型又可以进行字符串拆分  [:], 所以 arr[0][:] 操作是正确的
		 */
		fmt.Println(arr[0])
	}
	return strings.Join(arr, ".")
	//最后将 arr[0] 与 arr 里面其它数组元素进行拼接，将一系列字符串连接为一个字符串，之间用sep来分隔。
}

func stringReverse(){

	a := "Hello, 世界"
	println(a)
	println(Reverse(a))   // "界世 ,olleH"
}

//  反转字符串
func Reverse(s string) string {
	r := []rune(s)  // 为了能够处理字符串里面包含中文的情况，这里使用rune进行处理。
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]   // 首位和最后一位替换，以此类推。。。
	}
	return string(r)  // 将字节数组转换为 string
}

/*
	下例中，basename 函数模仿 UNIX shell 中的同名实用程序。只要 s 的前缀看起来像是文件系统路径(各部分由斜杠分隔),
	该版本的 basename(s) 就将其移出，貌似文件类型的后缀也被移除；

	basename2() 为 简化版利用库函数 strings.LastIndex:

	path 包和 path/filepath 包提供了一组更加普遍适用的函数，用来操作文件路径等具有层次结构的名字。
	path 包处理以斜杠'/' 分段的路径字符串， 不分平台。它不适合用于处理文件名，却适合其他领域，
	像 URL 地址的路径部分。 相反地， path/filepath 包根据宿主平台 (host platform) 的规则处理文件名，
	例如 POSIX 系统使用 /foo/bar, 而 Microsoft Windows 系统使用 c:\foo\bar

 */
func stringSplit(){

	fmt.Println(basename("a/b/c.go"))  // "c"
	fmt.Println(basename("c.d.go"))    // "c.d"
	fmt.Println(basename("abc"))       // "abc"

	fmt.Println(basename2("a/b/c.go"))  // "c"
	fmt.Println(basename2("c.d.go"))    // "c.d"
	fmt.Println(basename2("abc"))       // "abc"

}

// 初版的 basename 独自完成全部工作，并不依赖任何库：
// basename 移除路径部分和 . 后缀
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string{
	// 将最后一个 '/' 和之前的部分全部舍弃
	for i:=len(s)-1; i>=0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// 保留最后一个 '.' 之前的所有内容
	for i:=len(s)-1; i>=0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// 简化版利用库函数 strings.LastIndex:
func basename2(s string) string{
	slash:= strings.LastIndex(s, "/") // 如果没找到 "/"，则 slash 取值 -1
	s = s[slash+1:]
	/*
		无论有没有 "/" 都要 +1, 为什么呢？ 因为没有的话，-1+1 为0 ，
		从字符串首位元素开始取到最后，相当与整个字符串
		如果有呢， 例如2，  2是 "/"的位置，我们取的是 2 以后的所有元素，
		所以 要从 2+1 位开始取值，所以不论能不能找到"\" 都需要 +1 的操作
	 */
	if dot:= strings.LastIndex(s, "."); dot>=0 {
		s = s[:dot]
	}
	return s
}

// 字符串可以和字节 slice 相互转换：
func strAndByteArr(){
	s:= "abc"
	b:= []byte(s)   // 如有中文，可以使用 []rune(s)
	s2:= string(b)
	fmt.Println(s2)  // "abc"
}
/*
	概念上，[]byte(s) 转换操作会分配新的字节数组，拷贝填入 s 含有的字节，并生成一个 slice 引用，指向整个数组。
	具备优化功能的编译器在某些情况下可能会避免分配内存和复制内容，但一般而言，复制有必要确保 s 的字节维持不变
	(即使b 的字节在转换后发生改变)。反之，用 string(b) 将字节slice 转换成字符串也会产生一份副本，保证 s2 也不可变。

	为了避免转换和不必要的内存分配， bytes包和 strings 包都预备了许多对应的实用函数 (utility function)，
	它们两两相对应。例如，strings 包具备下面 6 个函数：

	func Contains(s, substr string) bool
	func Count(s, sep string) int
	func Fields(s string) []string
	func HasPrefix(s, prefix string) bool
	func Index(s, sep string) int
	func Join(a []string, sep string) string

	bytes 包里面的对应函数为：

	func Contains(b, subslice []byte) bool
	func Count(s, sep []byte) int
	func Fields(s []byte) [][]byte
	func HasPrefix(s, prefix []byte) bool
	func Index(s, sep []byte) int
	func Join(s [][]byte, sep []byte) []byte

	唯一的不同是，操作对象由字符串变为字节 slice.

 */

func intsToStringFunc(){

	a:= []int{1,2,3}
	fmt.Println(intsToString(a))   // "[1,2,3]"
}

/*
	bytes 包为高效处理字节 slice 提供了 Buffer 类型。 Buffer 起初为空，其大小随着各种类型数据的写入而增长，
	如 string、byte 和 []byte。 如下例所示， bytes.Buffer 变量无须初始化，原因是 零值 本来就有效：

	若要在 bytes.Buffer 变量后面添加任意文字符号的UTF-8 编码，最好使用 bytes.Buffer 的 WriteRune 方法，
	而追加 ASCII 字符， 如 '[' 和 ']'，则使用 WriteByte 亦可。

	bytes.Buffer 类型用途极广， 在第7章讨论接口的时候， 假若 I/O 函数需要一个字节接收器（io.Writer）或字节发生器(io.Reader),
	我们将看到能如何用其来代替文件，其中接收器的作用就如上例中的 Fprintf 一样。
 */

// intsToString 与 fmt.Sprint(values)类似，但插入了逗号
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	//buf.WriteString("开始")   // 也可以插入中文等 Unicode 码
	//buf.WriteString("こんにちは")
	// range 用来遍历数组和切片的时候返回索引和元素值
	// 如果我们不要关心索引可以使用一个下划线(_)来忽略这个返回值, 例如可以将下面的 i ，替换成 _
	for i,v:=range values {
		if i>0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	//buf.WriteString("结束") // 也可以插入中文等 Unicode 码
	//buf.WriteString("さようなら")
	return buf.String()
}

/*   了解 rune 类型
	// rune is an alias for int32 and is equivalent to int32 in all ways. It is
	// used, by convention, to distinguish character values from integer values.

	//int32的别名，几乎在所有方面等同于int32
	//它用来区分字符值和整数值

	type rune = int32

	golang中还有一个byte数据类型与rune相似，它们都是用来表示字符类型的变量类型。它们的不同在于：

	byte 等同于int8，常用来处理ascii字符
	rune 等同于int32,常用来处理unicode或utf-8字符

 */

func runeFunc(){

	var str = "hello 你好"

	//golang中string底层是通过byte数组实现的，座椅直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	fmt.Println("len(str):", len(str))  // "12"

	//以下两种都可以得到str的字符串长度

	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))   // "8"

	//通过rune类型处理unicode字符
	fmt.Println("rune:", len([]rune(str)))   // "8"
}