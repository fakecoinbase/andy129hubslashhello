package main

import (
	"fmt"
	"strings"
)

const sample = ` .bd.b2.3d.bc.20.e2.8c.98`    // 网上参考的示例 未达到效果
// 学习第三章--3.5-字符串--字符串和字节 slice
func main() {
	fmt.Println("learn3.5.4")

	//stringPrint()
	//stringPrint2()
	//stringPrint3()
	stringReverse()
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
