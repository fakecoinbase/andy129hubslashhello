package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)


// S 是一个结构体
type S struct {
	data map[string]interface{}
}

// go 内置包 gob
func main() {

	// jsonDemo()
	gobDemo()
}

// 通过 gob 编解码之后，数据类型可以得到正确的转换
// 注意： 标准库gob是golang提供的“私有”的编解码方式，它的效率会比json，xml等更高，特别适合在Go语言程序间传递数据。
func gobDemo(){
	var s1 = S{
		data: make(map[string]interface{}, 8),
	}
	s1.data["age"] = 18

	// encode 编码
	buf := new(bytes.Buffer)    // new 针对基本数据类型进行申请内存的操作，返回指针
	enc := gob.NewEncoder(buf)  // 造一个编码器对象
	err := enc.Encode(s1.data)  // 编码
	if err != nil {
		fmt.Println("gob encode failed, err : ", err)
		return 
	}
	b := buf.Bytes()   // 拿到编码之后的字节数据

	var s2 = S{
		data: make(map[string]interface{}, 8),
	}

	// 将编码之后字节数组，进行解码操作
	dec := gob.NewDecoder(bytes.NewBuffer(b))  // 造一个解码器对象
	err = dec.Decode(&s2.data)  // 解码
	if err != nil {
		fmt.Println("gob decode failed, err : ", err)
		return 
	}
	fmt.Printf("%#v\n", s2.data)   // "map[string]interface {}{"age":18}"

	for _, v := range s2.data {
		fmt.Printf("value : %v, type : %T\n", v, v) // "value : 18, type : int"
	}

}


// json 序列化与反序列化的问题
// 注意： Go语言中的json包在序列化空接口存放的数字类型（整型、浮点型等）都序列化成float64类型。
func jsonDemo() {

	var s1 = S{
		data: make(map[string]interface{}, 8),
	}

	s1.data["age"] = 12

	// json 序列化之前， interface{} 对应的值是 1， 类型为  int
	for _, v := range s1.data {
		fmt.Printf("value : %v, type : %T\n", v, v) // "value : 12, type : int"
	}
	// json 序列化操作
	ret, err := json.Marshal(&s1.data)
	if err != nil {
		fmt.Println("Marshall err : ", err)
		return
	}
	// 序列化为 json 格式的数据
	fmt.Printf("%v\n", string(ret)) // "{"age":12}"

	// 定义一个 S 类型的结构体，用于接收 把序列化的数据，进行反序列化后的数据
	var s2 = S{
		data: make(map[string]interface{}, 8),
	}

	// json 反序列化操作
	err = json.Unmarshal(ret, &s2.data)
	if err != nil {
		fmt.Println("Unmarshal err : ", err)
		return
	}

	// json 反序列化之后， interface{} 对应的值是 1， 但是类型变成为 float64， 这是为什么呢？
	// 这是因为 Go语言中的json包在序列化空接口存放的数字类型（整型、浮点型等）都序列化成float64类型。
	for _, v := range s2.data {
		fmt.Printf("value : %v, type : %T\n", v, v) // "value : 12, type : float64"
	}

	fmt.Printf("%#v\n", s2)
}