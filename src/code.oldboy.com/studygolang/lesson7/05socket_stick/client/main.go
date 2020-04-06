package main

import (
	"bufio"
	"fmt"
	"net"

	"code.oldboy.com/studygolang/lesson7/05socket_stick/proto"
)

// client
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("请求连接失败, err : ", err)
		return
	}

	defer conn.Close()

	writer := bufio.NewWriter(conn)

	for i := 0; i < 20; i++ {
		msg := fmt.Sprintf("这是来自客户端的第 %d 条信息", i)

		msgByteArr, err := proto.Encode(msg) // 调用自定义的协议 proto.Encode 去封装包
		if err != nil {
			fmt.Println("encode 出错，err : ", err)
			return
		}
		writer.Write(msgByteArr)
		writer.Flush() // 通过 bufio 进行写操作，别忘了最后要 Flush 一下
	}
}
