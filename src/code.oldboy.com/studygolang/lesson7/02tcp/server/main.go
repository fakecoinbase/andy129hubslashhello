package main

import (
	"fmt"
	"net"
)

// 1, 监听端口
// 2, 接收客户端请求建立连接
// 3, 创建 goroutine 处理连接

// tcp -- server
func main() {
	// 1, 监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("tcp 监听端口失败")
		return 
	}
	defer listener.Close()  // 程序退出时，释放 20000 端口的监听

	// 2, 一直等待 客户端的请求连接
	for {
		conn, err := listener.Accept()   // 如果没有客户端连接 就阻塞， 一直在等待
		
		if err != nil {
			fmt.Println("连接失败， err : ",err)
			continue
		}
		
		// 3, 处理连接
		go process(conn)
	}
}

// 单独处理每个连接的函数
func process(conn net.Conn) {

	defer conn.Close()    // 处理完毕，需要关闭连接

	// 从连接中接收数据
	var buf [1024]byte
	// bufio.NewReader(conn)
	n, err := conn.Read(buf[:])   // n 表示读了多少数据
	if err != nil {
		fmt.Println("接收客户端发来的消息失败了, err : ", err)
		return 
	}
	fmt.Println("接收客户端发来的消息：", string(buf[:n]))
}