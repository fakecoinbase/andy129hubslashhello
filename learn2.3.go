package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// 每一个声明有一个通用的形式
	// var name type = expression
	// varDemo()
	// shortVarDemo()
	// ptVarDemo()
	// ptVarDemo2()
	// ptVarDemo3()
	// ptVarDemo4()

	newFuncDemo()
}

/*
插播一条，转义字符打印，请用 fmt.Printf()，例如：

	fmt.Printf("i=%d,j=%d,k=%d\n", i,j,k)

	fmt.Printf("第一次赋值： a = %d\n",a)
*/

// 标准变量声明
func varDemo(){
	var st string
	fmt.Println("a"+st+"b") // "ab", st默认字符为""（注意中间没有空字符，空字符也代表一个字符）

	var i,j,k int    // int,int,int  声明时指定 int 类型,默认值都为 0
	var b,f,s = true, 2.3, "four" // bool, float64, string 未指定类型，直接赋值
							// 转义符 %t,   %g,      %s
	fmt.Printf("i=%d,j=%d,k=%d\n", i,j,k)
	fmt.Printf("b=%t,f=%g,s=%s\n", b,f,s)

	// 变量还可以通过调用返回多个值的函数进行初始化
	// var r ,err = os.Open(name)
}

// 短变量声明
func shortVarDemo(){

	// 短变量的特点就是  前面不用加 var 标识，变量类型由 := 后面的 expression 决定
	// name := expression
	/*
		anim := gif.GIF{LoopCount: nframes}
		freq := rand.Float64()*3.0
		t := 0.0
	 */

	// 与 var 声明一样，多个变量可以以短变量声明的方式声明和初始化
	i,j := 0,1

	// := 表示声明， 而 = 表示赋值, 下例代表 将右边的值赋给左边对应的变量
	i,j = j,i

	// 变量 a,b 声明并赋值
	a,b := j,i

	fmt.Printf("第一次赋值： a = %d\n",a)   // a 值为 0
	// 变量 a,c 声明并赋值， a 被重复赋值，由于下面这行代码中 由一个新的变量 c 所以编译通过，不报错
	// go语言中，短变量声明最少声明一个新变量，否则编译不通过，
	// 例如： a,b := j,i     a,b := 0,1  错误, a,b 已声明，第二行再声明一次，编译会提示错误：没有新变量
	a,c := 1,0
	fmt.Printf("第二次赋值： a = %d\n",a)  // a 值为 1
	// a,b := 0,1    // 编译错误：没有新的变量，可以这样修改：  a,b = 0,1

	fmt.Println(a+b+c)
}

// 指针变量 声明及 赋值
func ptVarDemo(){

		x := 1       // 短变量声明 x, 并初始化
		p := &x      // &x 代表指向 x 的地址，并赋值给 p, 然后我们说 p 就指向 x 或者 p 包含 x的地址
		fmt.Printf("通过指针变量，获取 变量x的值：%d\n",*p) // p 指向的变量写成 *p, 表达式 *p 获取变量x 的值。
		*p = 2  // p 指向的变量写成 *p, 因为 *p代表一个变量，所以它也可以 出现在赋值操作符左边。
		fmt.Printf("通过指针变量，修改 变量x的值：%d\n",x)

		fmt.Printf("x的地址：%x\n",&x)   // 以16进制输出 变量x 的地址
		fmt.Println(&x)

		// 指针类型的 零值是 nil
		var a,b int
		fmt.Println(&a == &a, &a == &b, &a == nil)
}

// 指针变量的函数调用，观察函数中的临时变量的值 的改变
func ptVarDemo2(){
	var p1 = f()
	var p2 = f()
	fmt.Printf("第一次调用f(), 返回p1指向的值为 : %d\n", *p1)
	*p1 = 500
	fmt.Printf("修改p1指向的值后 : %d\n", *p1)
	fmt.Printf("第一次调用f(), 返回&v : %x\n", p1)  // fmt.Printf("第一次调用f(), 返回&v : %x\n", p1)
	fmt.Printf("第二次调用f(), 返回&v : %x\n", p2)
	fmt.Printf("第三次调用f(), 返回&v : %x\n", f())
	fmt.Printf("第四次调用f(), 返回&v : %x\n", f())
	fmt.Printf("第五次调用f(), 返回&v : %x\n", f())
	fmt.Printf("第六次调用f(), 返回&v : %x\n", f())
	fmt.Printf("第七次调用f(), 返回&v : %x\n", f())
	fmt.Printf("第八次调用f(), 返回&v : %x\n", f())

	/*  机器运行结果

	第一次调用f(), 返回&v : c0000140b0
	第二次调用f(), 返回&v : c0000140b8
	第三次调用f(), 返回&v : c0000140d0
	第四次调用f(), 返回&v : c0000140d8
	第五次调用f(), 返回&v : c0000140e0
	第六次调用f(), 返回&v : c0000140e8
	第七次调用f(), 返回&v : c0000140f0
	第八次调用f(), 返回&v : c0000140f8

	*/

	fmt.Println(p1 == p2)  // false

	//
	/*
	总结： 函数中的变量，每一次调用时系统内存给分配的地址 会不一样，可以理解为 函数调用完毕就销毁的临时变量 ？？？？
	添加注释： 以上说法并不准确，函数中 v := 1 ，每执行一次函数，系统会给 v 重新分配一个地址，但该函数 v 的地址则是一直有效的，
	可以在函数外读取，赋值等一系列操作
	 */
}

func f() *int{
	v := 1
	return &v
}

// 指针变量的 函数调用，对指针地址进行操作，观察其值的 变化
func ptVarDemo3(){
	v :=1
	incr(&v)
	fmt.Printf("第一次调用incr()后， v = %d\n",v)  // fmt.Printf("第一次调用incr()后， v = %d\n",v)
	incr(&v)
	fmt.Printf("第二次调用incr()后， v = %d\n",v)
	incr(&v)
	fmt.Printf("第三次调用incr()后， v = %d\n",v)
	incr(&v)
	fmt.Printf("第四次调用incr()后， v = %d\n",v)
	incr(&v)
	fmt.Printf("第五次调用incr()后， v = %d\n",v)

	/*

	第一次调用incr()后， v = 2
	第二次调用incr()后， v = 3
	第三次调用incr()后， v = 4
	第四次调用incr()后， v = 5
	第五次调用incr()后， v = 6

	*/

	// 总结： 将指针通过参数传入函数中，并在函数中对其进行处理后，会直接改变指针所指向变量的值
}

func incr(p *int) int{
	*p++   // 等同于 *p = *p +1
	return *p
}

// 下面这个例子展示的是， 指定一个指针变量，由它取访问指针所指向的值，下面的 nt ,sep 都是指针变量，访问指针的值只有 *nt ,*sep

// usage：使用方法的描述说明，可以使用 -help 命令在命令行上显示出来
var nt = flag.Bool("n",false,"omit trailing newline") // 类似于 var n bool = false  *nt == n
var sep = flag.String("s", " ", "separator") // 类似于 var s string = " "   *sep == s
// separator  分割，设置分隔符

func ptVarDemo4(){
	flag.Parse() // 当程序运行时，在使用标识前，必须调用 flag.Parse 来更新标识变量的默认值
	fmt.Print(strings.Join(flag.Args(),*sep))
	if !*nt {  // 当 n 为 false 时，执行以下命令，换行，当为 true 时，则不换行，即 忽略正常输出时结尾的换行符
		fmt.Println()
	}
}

/*
执行以上函数，需要这样操作 :     ./文件 参数1 参数2 参数3...， 运行文件时就设置参数

注意 ，"/"分隔符，git 命令行里的 应打成 "//", 这是 windows下与 linux 平台不同的区别

$go build learn2.3.go

./learn2.3 a bc def
a bc def

$ ./learn2.3 -s // a bc def
a/bc/def

$ ./learn2.3 -help
Usage of G:\Goworkspace\learn2.3.exe:
  -n    omit trailing newline
  -s string
        separator (default " ")

 */

func newFuncDemo(){
	// 表达式 new(T), new 只是语法上的便利，不是一个基础概念。
	p := new(int)  // 等同于  var p int = 0
	fmt.Printf("*p 默认值 : %d\n", *p)
	*p = 2
	fmt.Printf("*p 赋值后 : %d\n", *p)

	// 下面两个自定义的函数  newInt(), newInt2() 返回的是一个指针，
	// 所指的地址不同，但是地址所指的变量的值都是一样的，int 变量的默认值为 0
	a := newInt()
	b := newInt2()

	fmt.Printf("值比较：*a == *b : %t\n", *a == *b )
	fmt.Printf("地址比较：a == b : %t\n", a == b)

	/*
	插播一条，使用 new 函数时要注意一个 特殊情况，
	在 go 语言中 两个变量的类型不携带任何信息且是 零值，例如 struct{} 或 [0]int， 当前的实现里面，它们有相同的地址。
	因为最常见的未命名变量都是结构体类型，它的语法（参考4.4.1节）比较复杂，所以 new 函数使用的相对较少。
	 */

	// new 是一个预声明的函数，不是一个关键字，所以它可以重定义为另外的其他类型，例如： delta()
	delta(3,5)

}

func newInt() *int{
	return new(int)
}

func newInt2() *int{
	var v int
	return &v
}
// new 是一个预声明的函数，不是一个关键字，所以它可以重定义为另外的其他类型，例如下面的 int 类型的变量: new
// 但 函数中当有变量的名称声明为 new 时，  new(T) 函数在这里就不能用了，导致了冲突
func delta(old, new int) int {

	// v := new(int)  //  new(T) 不能再使用，因为与 new 变量冲突
	return new - old
}
