package main

import (
	"fmt"
	"strings"
)

func main() {

	// fuzhi()
	// multiFuzhi()
	fuzhiStr()

}

// 普通赋值
func fuzhi(){
	var x int = 1
	fmt.Printf("%d\n",x)
	x++
	x--

	p := new(bool)
	*p = true
}

// 多重赋值
func multiFuzhi(){
	var x,y int
	x,y = y,x

	// a[i],a[j] = a[j],a[i]   // 数组或 slice 或 map的元素

	var a,b = 560,720
	num := gcd(a,b)
	fmt.Printf("%d, %d 的最大公约数是： %d\n",a,b,num)

	// var f, err = os.Open("filename.txt")
	//通常函数使用额外的返回值来指示一些错误情况，例如通过 os.open 返回的 error 类型
	// 或者一个通常叫 ok 的 bool 类型变量
	/*
		v,ok = m[key]
		v,ok = x.(T)
		v,ok = <-ch
	 */

	// 像变量声明一样，可以将不需要的值赋给空标识符 _,
	/*
		_, err = io.Copy(dst,src)
		_, ok = x.(T)

	 */

}

func gcd(x,y int) int {
	for y != 0 {
		x,y = y, x%y
		fmt.Printf("%d, %d\n",x,y)
	}
	return x
}

func fuzhiStr(){

	// 可赋值性
	var medals = []string{"gold", "silver", "bronze"}   // 隐式赋值

	medals[0] = "a"  // 显式赋值
	medals[1] = "b"
	medals[2] = "c"
	fmt.Printf("字符串 medals : %s\n",medals)  // 直接打印数组，则打印结果是 [a b c]
	var str = strings.Join(medals, "/")  // 将字符串数组 转换为一个字符串，并以 自定义分割符来 分割
	fmt.Printf("字符串 medals : %s\n",str) // 转换后的字符串打印为： a/b/c
}



