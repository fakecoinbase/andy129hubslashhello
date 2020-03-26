package main

import "fmt"

type student struct {
	name   string
	age    int
	gender string
	hobby  []string
}

// 自定义函数功能 实现结构体的 构造函数
/*
	简单构造函数，值传递,每次返回值类型，将结构体里的内容拷贝一份
	优化版构造含，指针传递，每次返回指针类型，无须拷贝结构体里的内容
*/
func main() {
	// test1()

	test2()
}

// 构造函数返回指针类型
func test2() {
	hobbySlice := []string{"篮球", "足球"}
	stu := newStudentByPtr("张学友", 59, "男", hobbySlice)
	fmt.Println(*stu)        // "{张学友 59 男 [篮球 足球]}"
	fmt.Println((*stu).name) // "张学友"
}

// 优化版， 返回的是指针，而不会每次构造都会拷贝。
func newStudentByPtr(name string, age int, gender string, hobby []string) *student {
	return &student{
		name:   name,
		age:    age,
		gender: gender,
		hobby:  hobby,
	}
}

// 实现简单的 构造函数
func test1() {
	hobbySlice := []string{"篮球", "足球"}
	stu := newStudent("刘德华", 59, "男", hobbySlice) // 这种构造函数，由于是 student 是一个值类型，所以返回值是 值拷贝的过程，多次创建时会消耗内存，所以要优化
	fmt.Println(stu)                              // "{刘德华 59 男 [篮球 足球]}"
}

func newStudent(name string, age int, gender string, hobby []string) student {

	return student{
		name:   name,
		age:    age,
		gender: gender,
		hobby:  hobby,
	}
}
