package main

import (
	"fmt"
)

// 结构体中字段大写开头 表示可公开访问，小写表示私有(仅在定义当前接头的包中可访问)
type Animal struct {
	name string
}

type Dog struct {
	feet int
	Animal
}

// 结构体内嵌 模拟"继承”
func main() {
	test1()
}

func (a *Animal) move() {
	fmt.Printf("%s会动~\n", a.name)
}

func (d *Dog) wangwang() {
	fmt.Printf("%s 在叫：汪汪汪~\n", d.name)
}

func test1() {
	var d = Dog{
		feet: 4,
		Animal: Animal{
			name: "旺财",
		},
	}
	d.move()     // "旺财会动~"
	d.wangwang() // "旺财 在叫：汪汪汪~"

}
