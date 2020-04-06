package main

import (
	"fmt"
	"io"
	"net"
)

// http -- client
func main() {
	conn, err := net.Dial("tcp", "www.zhcw.com:80")   // 某些网站 https (使用 443端口),  有安全证书，不能直接 tcp 请求连接
	if err != nil {
		fmt.Println("访问 失败")
		return 
	}

	defer conn.Close()

	// 发送请求  (HTTP GET请求格式)
	// "GET / HTTP/1.0\r\n\r\n"  注意这种格式，每个字段都是有含义的
	/*
		请求方法 URL 协议版本回车符换行符    --- 请求行
		头部字段名:值回车符换行符
				.....                      --- 请求头部
		头部字段名:值回车符换行符
		回车符换行符
		                                   --- 请求数据

	*/
	conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))   // HTTP/1.1 暂不支持   

	// 接收数据 
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[:])
		if err == io.EOF {
			fmt.Print(string(buf[:n]))
			return 
		}
		if err != nil {
			fmt.Println("接收数据失败, err : ", err)
			return 
		}
		fmt.Print(string(buf[:n]))
	}
}