package main
/*
探讨 作用域
声明将名字和程序实体关联起来，如一个函数或一个变量。声明的作用域是指用到声明时所声明名字的源代码段

 */
import (
	"fmt"
	"log"
	"os"
)

func main() {
	//example1()
	//example2()
	//example3()
	//example4()
	// example5()

	log.Printf("------onCreate = %s", cwd)   // 通过正确的 init(), cwd 已被初始化
}

/*
在函数里面，词法块可能嵌套很深， 所以一个局部变量声明可能覆盖另一个。
例如  example1() 中的 x 变量

for 循环创建了两个词法块： 一个是循环体本身的显式块， 以及一个隐式块，
它包含了一个闭合结构，其中就有初始化语句中声明的变量，如变量 i.
隐式块 中声明的变量的作用域包括条件、后置语句( i++ ),以及 for 语句体本身。

 */
func example1() {
	x:= "hello!" // 函数体中的变量声明 x
	for i:= 0;i< len(x);i++ {  // 使用外层的 x
		x:= x[i] // 又定义了一个 x 与 使用外层 x
		if x!= '!' { // 使用上一层的 x
			x:= x + 'A' - 'a' // 又定义了一个x 与使用上一层的 x
			// 应该是对字符进行了 ASCII 码数值的操作，使其进行了 大小写的转换，然后以 %c 字符形式输出，则成为了字符
			fmt.Printf("%c", x) //使用if语句创建的词法块儿里的 x,  依次打出每一个字符，在终端中最终效果： HELLO
		}
	}
}

/*
	变量在不同的词法块中声明，以及 显示与隐式的词法块

	知识点回顾：  _ (下划线) 为 空标识符
 */
func example2() {
	x:= "hello"   // x 在函数体中的 声明   (显式块儿)
	for _, x:= range x { // x 在for 语句块中的声明  (隐式块儿)
		x:= x + 'A' - 'a'  // x 在循环体中的声明  (显式块儿)
		fmt.Printf("%c\n", x)
	}
}

/*
像 for循环一样，除了本身的主体块之外， if 和 switch 语句还会创建隐式的词法块。

第二个 if语句嵌套在第一个中，所以 第一个语句的初始化部分声明的变量在 第二个语句中是可见的。
 */
func example3() {
	if x:= f(); x == 0 {  // 隐式块
		fmt.Println(x)
	}else if y:= g(x); x== y { // 隐式块
		fmt.Println(x,y)
	}else {
		fmt.Println(x,y)
	}
	// fmt.Println(x,y)  //  编译错误: x 与 y 在这里不可见
}

func f() int {
	return 1
}
func g(x int) int {
	return 0
}

func example4() error{

	// 写法1 (错误)
	/*
	if f,err := os.Open(""); err != nil {  // 编译错误：Unused variable 'f'   没有使用f 这个变量
		return err
	}
	f.Stat()   // 由于 f 在 if 语句中，所以它的作用域也就仅仅存在于 f 这个语句中，外面无法调用 f
	f.Close()
	*/

	// 写法2 (正确)
	/*
	f, err:= os.Open("")
	if err!= nil {
		return err
	}
	f.Stat()
	f.Close()

	return nil
	*/

	// 注意与第一种错误的写法 做比较
	// 写法3 (正确)
	if f,err := os.Open(""); err!= nil {
		return err
	}else {
		// if-else 词块儿， f 在这里可以被使用
		f.Stat()
		f.Close()
	}
	return nil
}

var cwd string

/*
 同 init() 一样，函数中的cwd 在这里因为用了 赋值符号，所以它不是给全局变量 cwd 赋值，而是重新在函数中定义了一个同名变量
 */
func example5() {
	cwd,  err:= os.Getwd()
	if err!= nil {
		log.Fatalf("os.Getwd failed : %v", err)
	}
	log.Printf("working directory = %s", cwd)  // 将 cwd以 log 信息打印出来，避免编译报错： cwd 未被使用
}


/*
 // 函数中的cwd 在这里因为用了 赋值符号，所以它不是给全局变量 cwd 赋值，而是重新在函数中定义了一个同名变量
所以 在 init() 中并没有做到给 全局变量 cwd 赋值的功能
 */
/*
func init() {
	log.Printf("----------------init()")
	cwd,  err:= os.Getwd()
	if err!= nil {
		log.Fatalf("os.Getwd failed : %v", err)
	}
	log.Printf("working directory = %s", cwd)
}
*/

func init() {
	log.Printf("----------------正确初始化")
	var err error
	cwd,  err = os.Getwd()   // 解除 短变量声明，正确赋值， 此时 cwd 所指的是定义的 全局变量
	if err!= nil {
		log.Fatalf("os.Getwd failed : %v", err)
	}
	log.Printf("working directory = %s", cwd)
}



