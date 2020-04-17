package main

import "fmt"

// test 
func main() {
	testNew()
}

// 使用 new 对基本数据类型进行初始化 内存申请
func testNew(){
	// ------------------bool ---------------------------

	b := new(bool) // new 针对基本数据类型进行申请内存的操作，返回的是类型的指针

	*b = true  //  赋值操作注意
	fmt.Printf("%v, %T\n", b, b)  // "0xc0000120c0, *bool"

	fmt.Println(b)  // "0xc00006e068"


	// ------------------string ---------------------------

	s := new(string)

	*s = "abc"  //  赋值操作注意
	fmt.Printf("%v, %T\n", s, s)  // "0xc0000561e0, *string"

	fmt.Println(s)  // "0xc0000561e0"
}