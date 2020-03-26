package main

import (
	"fmt"
	"time"
)

// nullInterface 是一个空接口
type nullInterface interface {
}

// 空接口 : 没有定义任何方法的接口。
// 因此任何类型都实现了 空接口。 也就是说  空接口可以接收任何类型的值
func main() {
	fmt.Println("空接口")

	// test1()
	// test2()
	// test3()
	test4()

}

// 使用空接口接收 任意类型
func test1() {

	//空接口 可省略名称， 对比如上:  type interface {}
	// 定义一个空接口变量
	var x interface{}

	x = 100
	x = "沙河"
	x = true
	x = 'c'
	x = 99.99

	x = struct {
		name string
	}{
		name: "花花",
	}

	x = time.Second

	// fmt.Println(x)

	showType(x)
}

// 打印空接口接收值的 类型
func showType(x interface{}) {
	fmt.Printf("%T\n", x)
}

// 空接口在 map value 中的运用
// 可以接收各种不同类型的值
func test2() {
	var stuInfo = make(map[string]interface{}, 10)

	stuInfo["yang"] = 100
	stuInfo["zhang"] = true
	stuInfo["刘德华"] = 99.999
	stuInfo["hhh"] = "哈哈哈"
	stuInfo["time"] = time.Now()

	fmt.Println(stuInfo) // "map[hhh:哈哈哈 time:2020-03-20 19:28:18.4300617 +0800 CST m=+0.003502001 yang:100 zhang:true 刘德华:99.999]"
}

// (类型 断言) 判断空接口接收的值的 类型
// x.(T) :  x 是接收值的变量 , T 为 断言的类型,  x.(T) 会返回
func test3() {
	var x interface{}

	x = 100
	checkType(x)

}

func checkType(x interface{}) {
	v, ok := x.(int)
	if !ok {
		fmt.Println("断言失败")
	} else {
		fmt.Printf("%v, %t\n", v, ok) // "100, true" ,   x 传入为 100
		fmt.Println("a 是一个int 类型")
	}

	v2, ok2 := x.(string)
	if !ok2 {
		fmt.Println("断言失败") // "断言失败",   x 传入为 100
	} else {
		fmt.Printf("%v, %t\n", v2, ok2)
		fmt.Println("a 是一个string 类型")
	}

	/*   多种不同类型，则要写多个判断
	......

	*/
}

func test4() {

	var x interface{}
	checkType2(x)
	x = 100
	checkType2(x)
	x = "沙河"
	checkType2(x)
	x = true
	checkType2(x)
	x = 'c'
	checkType2(x)
	x = 99.99
	checkType2(x)

}

// 断言 ： 死记下面这种方式
func checkType2(x interface{}) {

	// 注意，空接口接收的值的类型 也属于 interface 类型, 但这种断言 没有意义，
	// 我们使用 x.(T)的目的就是 断言出它 真正的类型，所以我们在这里 去掉 interface 的 case
	switch v := x.(type) {
	/*
		case interface{}:
			fmt.Printf("x is a interface, value is %v\n", v)
	*/
	case string:
		fmt.Printf("x is a string, value is %v\n", v)
	case int:
		fmt.Printf("x is a int, value is %v\n", v)
	case bool:
		fmt.Printf("x is a bool, value is %v\n", v)
	case float32:
		fmt.Printf("x is a float32, value is %v\n", v)
	case float64:
		fmt.Printf("x is a float64, value is %v\n", v)
	default:
		fmt.Printf("unsuport type! value is %v\n", v)
	}
}
