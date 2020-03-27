package main

import "fmt"

// switch
// Go语言规定，每个 switch 只能有一个 default 分支。
func main() {
	//switchDemo1(3) // "中指"

	//switchDemo2(2) // "偶数"

	//switchDemo3()
	// switchDemo4()

	testSwitch()
}

// switch 的坑 (一个奇怪的问题)
func testSwitch() {
	f := func() bool {
		return false
	}

	// 情景1： 打印 "假"
	switch f() {
	case false:
		fmt.Println("假")
	case true:
		fmt.Println("真")
	}

	// 情景2： 打印 "真" ?  为什么呢
	switch f(); {
	// f(); 写法不会报错，那它代表什么意思呢?    相当于： if _ = f();  相当于把 f()的返回值取出来丢弃了 ,
	// _ = f(); 此行代码 不再是 switch 的判断条件,  所以 switch 的判断条件就没了，那就是默认 为 true ,  所以就匹配到  case true :
	case false:
		fmt.Println("假")
	case true:
		fmt.Println("真")
	}

}

// 多个不同的判断条件时，可以使用 switch
func switchDemo1(finger int) {

	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}

// case 多种条件（满足其中一种条件即可）
func switchDemo2(number int) {

	// switch number = 7; number {    // 也可以直接在 switch 判断语句里面定义
	switch number {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(number)
	}
}

// switch 后面无须添加 变量判断
func switchDemo3() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}

// switch 里面的 fallthrough 用法
// fallthrough 语法可以执行满足条件的 case 的下一个 case, 是为了兼容 C 语言中的 case 设计的。

// 以下打印 a, 接着又打印了 b,  满足 第一个条件之后， fallthrough 关键字决定了立即执行 紧接着的 fmt.Println("b"),  不管 s == "b" 是否成立。
func switchDemo4() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
