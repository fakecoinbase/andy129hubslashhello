package main

import (
	"fmt"
	"strings"
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
	stringSplit()
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