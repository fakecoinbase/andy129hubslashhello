package main

import (
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/config/cmd"
	"io/ioutil"
	"movie/src/share/config"
	"movie/src/share/utils/path"
	"net/http"
	"strconv"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerRPC)
	fmt.Println("Listen on:9095")
	err := http.ListenAndServe(":9095", mux)
	if err != nil {
		fmt.Println("ListenAndServe 9095 err : ", err)
		return
	}
}

func handlerRPC(w http.ResponseWriter, r*http.Request){
	fmt.Println("handlerRPC ...")

	// 1, 正常请求
	if r.URL.Path == "/" {
		_, err := w.Write([]byte("server ..."))
		if err != nil {
			fmt.Println("handlerRPC write err : ", err)
			return
		}
		return
	}

	// 2, RPC请求， 跨域请求 参数设置
	if origin := r.Header.Get("Origin"); true {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}

	w.Header().Set("Access-Control-Allow-Methods", "POST,GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length,Accept-Encoding,X-Token,X-Client")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	handlerJSONRPC(w, r)
	return
}

func handlerJSONRPC(w http.ResponseWriter, r *http.Request) {
	service, method := path.PathToReceiver(config.Namespace, r.URL.Path)
	fmt.Println("handlerJSONRPC service : ", service)
	fmt.Println("handlerJSONRPC method : ", method)

	// 读取请求体
	br, _ := ioutil.ReadAll(r.Body)
	fmt.Println("data : ", string(br))
	request := json.RawMessage(br)
	var response json.RawMessage

	ctx := path.RequestToContext(r)

	/*  尝试修改 content-type
	requestOption := func(reqo *client.RequestOptions){
		reqo.ContentType = "application/json"
		reqo.Stream = false
		reqo.Context = ctx
	}

	req := (*cmd.DefaultOptions().Client).NewRequest(service, method, &request, requestOption) // 这里没有 NewJsonRequest方法
	*/

	// fmt.Printf("options : %#v\n", (*cmd.DefaultOptions().Client).Options())

	// req := client.NewRequest(service, method, &request)

	// client.WithContentType("application/json") 来设置 content-type
	req := (*cmd.DefaultOptions().Client).NewRequest(service, method, &request, client.WithContentType("application/json"))

	fmt.Println("req.ContentType : ", req.ContentType())

	// ctx := path.RequestToContext(r)
	err := (*cmd.DefaultOptions().Client).Call(ctx, req, &response)
	if err != nil {
		fmt.Println("(*cmd.DefaultOptions().Client).Call err : ", err)
		return
	}
	b, _:= response.MarshalJSON()
	// 设置响应头
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	_, err = w.Write(b)
	if err != nil {
		fmt.Println("Write err : ", err )
	}
}
