package main

import "fmt"

// 总结：
/*
	1, test1() 接口的接受者 是值类型  (可以有语法糖结构)
	2, test2() 接口的接受者是 指针类型 (必须要使用指针类型)
	3, test3() 一个结构体使用 值类型和 指针类型分别实现了 多个接口

	// 我们一般使用 指针类型作为接口的接受者

*/

// Animal 是一个接口
type Animal interface {
	speak()
	move()
}

// Cat 是一个结构体
type Cat struct {
	name string
}

// 方法一：使用值接受者
// Cat 类型实现 Animal 的接口
func (c Cat) speak() {
	fmt.Println("喵喵喵")
}

func (c Cat) move() {
	fmt.Println("猫会动")
}

// Human 是一个接口
type Human interface {
	work()
	eat()
}

// Student 是一个结构体
type Student struct {
	name string
	age  int
}

// 方法二：使用指针接受者
func (s *Student) work() {
	fmt.Println("人会干活")
}

func (s *Student) eat() {
	fmt.Println("人会吃饭")
}

// 方法三，Student 在方法二中用 指针类型实现了 接口 Human, 然后在这里又通过 值类型 实现了 Animal.
func (s Student) speak() {
	fmt.Println("哒哒哒")
}

func (s Student) move() {
	fmt.Println("学生会动")
}

// 实现接口时 使用 指针接受者 和使用 值接受者 的区别
func main() {

	//test1()
	// test2()
	test3()
}

// 实现接口用 值类型的接收者， 如果传入一个 指针类型，可否调用实现的方法 (可以，因为 go 语言中的语法糖)
func test1() {
	var x Animal
	fmt.Printf("%T\n", x) // "<nil>"

	// 实例化一个 Cat
	hh := Cat{"花花"}
	x = hh
	fmt.Printf("%T\n", x) // "main.Cat"
	fmt.Println(x)        // "{花花}"

	// 实例化一个 Cat 的指针类型
	tom := &Cat{"Tom"}
	x = tom  // 当指针类型作为接口的接受者时，go语言里面会自动判断，并做 *tom 的操作把值取出来。 （go语言中的 语法糖）
	x.move() // 实际： (*x).move(),   (*x).speak()
	x.speak()
	fmt.Printf("%T\n", x) // "*main.Cat"
	fmt.Println(x)        // "&{Tom}"
}

// 当使用指针作为接口的接受者时， 接口类型是否还能接受 值类型的变量 (不行，接口类型只认准 指针类型)
func test2() {
	var h Human
	fmt.Printf("%T\n", h) // "<nil>"

	/*
		ss := Student{
			"阳",
			25,
		}
		h = ss // ss 是 Student 类型 (一个值类型)，无法赋值给 接口 h , 因为实现 接口h 的是 指针类型，这里类型不匹配
		fmt.Printf("%T\n", h)
	*/

	zhang := &Student{
		"章",
		33,
	}
	h = zhang
	h.eat()
	h.work()
	fmt.Printf("%T\n", h) // "*main.Student"
	fmt.Println(h)        // "&{章 33}"
}

// 值类型与 指针类型混合， 一个Student 结构体类型 通过 值类型 和  指针类型 实现了两个 接口 (Animal , Human )
func test3() {

	// Student 即实现了 Animal  又实现了 Human ，所以 Student 的实例化可以调用 它实现的所有方法
	yang := &Student{
		"张",
		33,
	}
	yang.speak()
	yang.move()
	yang.eat()
	yang.work()

	// 值类型，实现了 Animal 接口， 赋值给 Animal 接口变量a 之后， a 只能调用 Animal 接口里的方法
	var a Animal
	a = yang              // 语法糖
	fmt.Printf("%T\n", a) // "*main.Student"
	a.move()
	a.speak()

	// 指针类型 实现了 Human 接口， 赋值给 Human 接口变量a 之后， a 只能调用 Human 接口里的方法
	var h Human
	h = yang
	fmt.Printf("%T\n", h) // "*main.Student"
	h.eat()
	h.work()

}
