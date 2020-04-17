package main

import (
	"fmt"

	"github.com/vmihailenco/msgpack"
)

// Person 是一个结构体
type Person struct {
	Name   string
	Age    int
	Gender string
}

// msgpack 
// MessagePack是一种高效的二进制序列化格式。它允许你在多种语言(如JSON)之间交换数据。但它更快更小。
// msgpack 包即能解决 json反序列化的问题，还能允许在多种语言之间交换数据 (例如 与前端交换 json数据),  并且还拥有 gob 二进制处理高效的特点
func main() {
	p1 := Person{
		Name:   "沙河娜扎",
		Age:    18,
		Gender: "男",
	}
	// marshal
	b, err := msgpack.Marshal(p1)
	if err != nil {
		fmt.Printf("msgpack marshal failed,err:%v", err)
		return
	}

	// unmarshal
	var p2 Person
	err = msgpack.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("msgpack unmarshal failed,err:%v", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2) // p2:main.Person{Name:"沙河娜扎", Age:18, Gender:"男"}
	fmt.Printf("%v, %T\n", p2.Age,p2.Age)  // "18, int"
}