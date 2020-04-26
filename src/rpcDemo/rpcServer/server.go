package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 服务端, 求矩形面积和周长
/*
	golang写RPC程序，必须符合4个基本条件，不然RPC用不了
	1, 结构体字段首字母要大写，可以让别人调用
	2, 函数名必须首字母大写
    3, 函数第一参数是接收参数，第二个参数是返回给客户端的参数，必须是指针类型
    4, 函数还必须有一个返回值error

*/

// 声明矩形对象
type Rect struct {

}

// 声明参数结构体
type Params struct {
	Width int
	Height int
}
// 定义求矩形面积的方法
func (r *Rect)Area(p Params, ret *int) error{
	*ret = p.Width * p.Height
	return nil
}
// 定义求矩形周长的方法
func (r *Rect)Perimeter(p Params, ret *int) error{
	*ret = (p.Width+p.Height) *2
	return nil
}

func main(){
	testJsonRpc()
}

// jsonrpc 服务  （json编码实现跨语言调用）
func testJsonRpc(){
	// 1, 注册服务
	rect := new(Rect)
	rpc.Register(rect)

	// 2, 监听服务
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net.Listen err : ", err)
		return
	}

	// 循环监听服务
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}

		// 协程
		go func(conn net.Conn){
			fmt.Println("new a client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}

}

// 传统RPC 服务 (只能适用于 go程序之间调用)
func testRpc(){
	// 1, 注册服务
	rect := new(Rect)
	rpc.Register(rect)

	// 2, 把服务处理绑定到 http 协议上
	rpc.HandleHTTP()

	// 3, 监听服务， 等待客户端调用求面积和周长的方法
	err := http.ListenAndServe(":8080", nil )
	if err != nil {
		fmt.Println("ListenAndServe err : ", err )
		return
	}
}


