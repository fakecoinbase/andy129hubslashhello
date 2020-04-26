package main

import (
	"encoding/json"
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 声明参数结构体
type Params struct {
	Width int
	Height int
}

type People struct {
	User
	Student
}

type User struct {

}
type Student struct {

}


// client
func main() {

	// testNil()
	// testRpc()

	testJsonRpc()

}

// 客户端使用 json 编码rpc
func testJsonRpc(){
	// 1, 连接远程的 RPC 服务  (与传统相比，差异在这里)
	rpcClient, err := jsonrpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("DialHTTP err : ", err)
		return
	}

	// 2, 调用远程的方法
	// 定义接收服务端传回来的计算结果的变量
	var ret int
	// 求面积
	err = rpcClient.Call("Rect.Area", Params{50,100}, &ret)
	if err != nil {
		fmt.Println("rpc call Rect.Area err : ", err)
		return
	}

	fmt.Println("json 编码，面积：", ret)

	// 求周长
	err = rpcClient.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err != nil {
		fmt.Println("rpc call Rect.Perimeter err : ", err)
		return
	}

	fmt.Println("json 编码，周长：", ret)
}

// 传统客户端 调用 rpc
func testRpc(){
	// 1, 连接远程的 RPC 服务
	rpcClient, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("DialHTTP err : ", err)
		return
	}

	// 2, 调用远程的方法
	// 定义接收服务端传回来的计算结果的变量
	var ret int
	// 求面积
	err = rpcClient.Call("Rect.Area", Params{50,100}, &ret)
	if err != nil {
		fmt.Println("rpc call Rect.Area err : ", err)
		return
	}

	fmt.Println("面积：", ret)

	// 求周长
	err = rpcClient.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err != nil {
		fmt.Println("rpc call Rect.Perimeter err : ", err)
		return
	}

	fmt.Println("周长：", ret)
}

//
func testNil(){


	var b People
	fmt.Printf("%p\n", &b)  // 0xaf1548
	fmt.Println(b) // {{} {}}
	fmt.Println(&b)  // &{{} {}}

	var p Params
	fmt.Printf("%p\n", &p)  // 0xc0000ee020
	// p := Params{}
	fmt.Println(p)
	fmt.Println(&p)

	p.Height = 10    // 可以修改


	var m map[string]string

	fmt.Printf("m未申请内存：%p\n", m)  // m未申请内存：0x0
	fmt.Println(m == nil)  // true
	fmt.Println(&m)

	// 针对未申请内存空间的 map，如上， 不能进行 key,value 赋值操作，但是却可以 json 反序列化，如下 (取地址传入，不会报错)
	// m["name"] = "zhang"   // 运行报错： panic: assignment to entry in nil map
	// m["pwd"] = "333"

	data := `{"name":"yang", "pwd":"123"}`
	json.Unmarshal([]byte(data), &m)   // 即使 m 为 nil ,依旧可以传入
	fmt.Printf("m 进行反序列化后：%p\n", m)  // m 进行反序列化后：0xc0000d8330
	fmt.Println(m)

}
