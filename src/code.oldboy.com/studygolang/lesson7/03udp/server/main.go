package main

import (
	"fmt"
	"net"
)

// udp -- server
func main() {
	// 与 tcp server 相比, upd listen方法直接返回 updConn, 而不是一个 listener
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		// IP :    net.IPv4(0,0,0,0)  == net.ParseIP("127.0.0.1")
		IP:   net.IPv4(0, 0, 0, 0), // 监听本机
		Port: 30000,
	})
	if err != nil {
		fmt.Println("启动server 失败，err : ", err)
	}
	// (待研究)这里不能使用 go process(udpConn)  , 否则，程序就直接结束了？
	process(udpConn)
}

// 循环收发数据
func process(udpConn *net.UDPConn) {

	defer udpConn.Close() // 最后关闭 udpConn

	// 循环收发数据
	for {
		var buf [1024]byte
		// 接收消息
		n, addr, err := udpConn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("接收消息失败, err : ", err)
			return
		}
		fmt.Printf("接收到来自 %v 的消息: %v\n", addr, string(buf[:n]))

		// 回复消息
		_, err = udpConn.WriteToUDP([]byte("收到"), addr)
		if err != nil {
			fmt.Println("回复消息失败, err : ", err)
			return
		}
	}
}

//  开启一个 udp server, 开启两个 udp client, 两个 client 给 udp 发送消息,  最终效果:

/*
	客户端1 收到消息：  收到回复:  收到
	客户端2 收到消息：  收到回复:  收到


	服务器端收到消息：    接收到来自 127.0.0.1:55870 的消息: 你好, udp server
						接收到来自 127.0.0.1:55667 的消息: 你好, udp server

*/
