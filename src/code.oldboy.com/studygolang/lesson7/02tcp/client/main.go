package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// 1, 根据地址找到 server
// 2, 向 server 端发消息

// tcp -- client
func main() {
	
	// 1, 
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("连接服务端失败, err : ", err)
		return 
	}
	defer conn.Close()

	// 2， 向 server 端发消息 (conn 也是一种 io, 所以可以直接使用 Fprintln)
	// fmt.Fprintln(conn, "你好，服务端!")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil  {
		fmt.Println("reader读取 失败， err : ",err)
		return 
	}
	_, err = conn.Write([]byte(input))

	if err != nil {
		fmt.Println("发送消息失败，err : ", err)
		return 
	}

}