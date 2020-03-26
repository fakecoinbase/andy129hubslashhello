package main

import "fmt"

type people struct {
}

// interface 接口值
func main() {
	// test1()
	test2()
}

func test1() {
	var x interface{} // 接口值分为两部分：<Type, Value>

	var a int64 = 100

	x = a                        // <int64,100>
	fmt.Printf("%T, %v\n", x, x) // "int64, 100"

	v, ok := x.(string) // 如果类型猜对了，则会返回这个类型具体的值，如果猜错了，则会返回 你猜的这个类型的默认值，例如：string 类型默认值为 ""
	// 类型断言
	if ok { // 不能这样写，因为 x.(T) 返回的有两个值
		fmt.Println("x 是字符串, v : ", v)
	} else {
		fmt.Println("x 不是字符串, v : ", v) // "x 不是字符串, v :"
	}

	v2, ok2 := x.(int64)
	if ok2 {
		fmt.Println("x 是int, v2 : ", v2) // x 是int, v2 :  100
	} else {
		fmt.Println("x 不是int, v2 : ", v2)
	}

	v3, ok3 := x.(bool) // 猜对了，则返回 x 的值，猜错了，则返回 你猜测的 bool 类型的默认值 false
	if ok3 {
		fmt.Println("x 是bool, v3 : ", v3)
	} else {
		fmt.Println("x 不是bool, v3 : ", v3) // x 不是bool, v3 :  false
	}

}

// 类型断言， switch
func test2() {

	var a = 100
	var a2 int32 = 345
	var a3 int64 = 534534
	var b = true
	var c = "donghan"
	var d = 23.43
	var p = people{}

	checkType(a)  // x 是一个int  (var a = 100, 没指定类型，则默认为 int类型，不是int32 也不是 int64)
	checkType(a2) // x 是一个int32
	checkType(a3) // x 是一个int64
	checkType(b)  // x 是一个bool
	checkType(c)  // x 是一个string
	checkType(d)  // x 是一个float64
	checkType(p)  // x 是一个people, Type : main.people
}

func checkType(x interface{}) {

	// v := x.(type) 返回值 v 为 x 的值

	switch x.(type) {
	case string:
		fmt.Println("x 是一个string")
	case int:
		fmt.Println("x 是一个int")
	case int32:
		fmt.Println("x 是一个int32")
	case int64:
		fmt.Println("x 是一个int64")
	case bool:
		fmt.Println("x 是一个bool")
	case float64:
		fmt.Println("x 是一个float64")
	case people:
		fmt.Printf("x 是一个people, Type : %T\n", x)
	default:
		fmt.Println("x 无法识别类型")
	}
}
