package main

import (
	"bufio"
	"fmt"
	"net"

	"code.oldboy.com/studygolang/lesson7/05socket_stick/proto"
)

// server
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("服务器监听失败, err : ", err)
		return
	}

	defer listener.Close()

	for {
		fmt.Println("-----server for")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("连接失败, err : ", err)
			continue
		}

		go process(conn)
	}

}

func process(conn net.Conn) {
	defer conn.Close()

	// var buf [1024]byte
	reader := bufio.NewReader(conn)

	/*
		// 循环读
		for {

			n, err := reader.Read(buf[:])
			if err == io.EOF {
				fmt.Println("读取数据 EOF")
				break
			}
			if err != nil {
				fmt.Println("读取数据失败, err : ", err)
				break
			}
			recvStr := string(buf[:n])
			fmt.Println("收到client发来的数据：", recvStr)
		}
	*/

	for {
		// 解决粘包， 优化后
		msg, err := proto.Decode(reader) // 调用自定义的协议 proto.Decode 去解包
		if err != nil {
			fmt.Println("decode msg failed, err : ", err)
			return
		}

		fmt.Println("收到client发来的数据：", msg)
	}

}
