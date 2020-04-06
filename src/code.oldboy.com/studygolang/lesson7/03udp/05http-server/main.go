package main

import (
	"fmt"
	"net/http"
)

// http -- server 
func main() {
	
	// 必须放在  http.ListenAndServe() 前面执行
	http.HandleFunc("/", sayHello)   // 注册路由：当你访问 /  就执行 sayHello 函数
	// :9090  --->  127.0.0.1:9090 的简写
	err := http.ListenAndServe(":9090", nil)  // 建立监听
	if err != nil {
		fmt.Printf("http server failed, err : %v\n", err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w, "<h1 style:'color:red'>Hello 沙河!</h1>")

}