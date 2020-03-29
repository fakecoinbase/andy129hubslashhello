package main

import (
	"fmt"
	"runtime"
)

// runtime
func main() {
	test1()
}

func test1(){
	runtime.GOMAXPROCS(1)

	fmt.Println(runtime.GOOS)  // "windows" , 根据你的运行环境，自动生成的信息  (windows平台下就是 windows )

	fmt.Println(runtime.GOARCH)// "amd64" , 根据你的运行环境，自动生成的信息 
}