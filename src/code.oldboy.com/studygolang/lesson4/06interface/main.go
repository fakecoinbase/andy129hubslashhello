package main

import "fmt"

type speaker interface {
	speak()
}

type mover interface {
	move()
}

// 接口的嵌套 （接口内部嵌套其他接口） 示例：
type animal interface {
	speaker
	mover
}

type cat struct {
	name string
}

func (c cat) move() {
	fmt.Println("喵喵喵")
}

func (c cat) speak() {
	fmt.Println("猫会动")
}

// 接口的嵌套
func main() {

	var a animal
	a = cat{
		name: "花花",
	}

	a.move()
	a.speak()

	// 其他测试：
	var m mover
	m = cat{
		name: "小黑",
	}
	m.move()
	// m.speak()   // 编译报错：m.speak undefined (type mover has no field or method speak)
	// 虽然 cat 实现了 mover 和 speaker 两个接口里面的方法，但是 你用 mover 来接收 cat 的实例化之后， m 也只能调用 mover 里面定义的方法
}
