package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http -- server
func main() {
	http.HandleFunc("/", sayHello)

	// 启动服务
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		panic(err)
	}
}

func sayHello(w http.ResponseWriter, t *http.Request) {
	// w.Write([]byte("<h1>hello, 沙河</h1>"))

	data, err := ioutil.ReadFile("./form.html")
	if err != nil {
		fmt.Println("读取失败, err : ", err)
		w.WriteHeader(http.StatusInternalServerError) // 返回 网络的错误信息 , 详见 http/status.go  里面的常量
		return
	}

	w.Write(data)
}
