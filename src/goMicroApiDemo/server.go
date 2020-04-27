package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
	pb "goMicroApiDemo/proto"
	// pb2 "goMicroApiDemo/proto/user"

)
/*  尝试启动第二个微服务，貌似会影响第一个

type User struct {

}

func (u *User) Info(ctx context.Context, req *pb2.InfoRequest, resp *pb2.InfoResponse) error {
	fmt.Println("收到 User.Info 请求")
	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no content")
	}
	resp.Message = "RPC Info 收到了你的请求" + req.Name
	return nil
}


*/
type Example struct {

}

type Foo struct {

}

func (e *Example) Call (ctx context.Context, req *pb.CallRequest, resp *pb.CallResponse) error{

	fmt.Println("收到 Example.Call 请求")
	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no content")
	}
	resp.Message = "RPC Call 收到了你的请求" + req.Name
	return nil
}

func (f *Foo) Bar (ctx context.Context, req *pb.EmptyRequest, resp *pb.EmptyResponse) error {
	fmt.Println("收到 Foo.Bar 请求")
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),     // 由于设置了以 http请求, 所以服务名必须以左边这种格式： go.micro.api.微服务名
		)

	err := pb.RegisterExampleHandler(service.Server(), new(Example))
	if err != nil {
		fmt.Println("RegisterExampleHandler err : ", err)
		return
	}

	err = pb.RegisterFooHandler(service.Server(), new(Foo))
	if err != nil {
		fmt.Println("RegisterFooHandler err : ", err)
		return
	}

	/*  尝试启动第二个微服务，但是一旦启动， 上面第一个服务就访问不了
	service2 := micro.NewService(
			micro.Name("go.micro.api.user"),
		)

	err = pb2.RegisterUserHandler(service2.Server(), new(User))
	if err != nil {
		fmt.Println("RegisterUserHandler err : ", err)
		return
	}
	*/

	err = service.Run()   // Run() 方法会造成 主程序阻塞
	if err != nil {
		fmt.Println("service.Run err : ", err)
		return
	}

	/* 尝试启动第二个微服务，但是一旦启动， 上面第一个服务就访问不了
	err = service2.Run()
	if err != nil {
		fmt.Println("service.Run err : ", err)
		return
	}
	*/

	// 0, 进入proto 目录下 执行 protoc 命令，将 .proto 文件生成为 .pb.go 和 .pb.micro.go 文件
	//    protoc -I. --micro_out=. --go_out=. ./user.proto

	// 1, 启动 service.Run() 之前，先使用 micro 命令设置一下访问 go-micro 的方式  （设置 以http 请求方式访问）
	/*	执行命令：micro api --handler=rpc

		2020-04-27 17:53:15  level=info service=api Registering API RPC Handler at /
		2020-04-27 17:53:15  level=info service=api HTTP API Listening on [::]:8080
		2020-04-27 17:53:15  level=info service=api Starting [service] go.micro.api
		2020-04-27 17:53:15  level=info service=api Server [grpc] Listening on [::]:61423
		2020-04-27 17:53:15  level=info service=api Registry [mdns] Registering node: go.micro.api-d73f6ee6-b287-4bc3-b935-9da292c4e681
	 */

	// 由上可看出， service=api HTTP API Listening on [::]:8080

	// 2, 运行  go 语言程序 (service.Run() ),  服务启动。

	// 3，通过 postman 模拟 http 请求
		/*
			http://127.0.0.1:8080/example/call    (访问 go-micro 中的 Call() 方法)
			        // 调用Call() 携带参数，要使用 post 请求中的 json 数据格式

			示例：   参数：
					{
						"name":"liudehua"
					}
			 		返回：
					{
		    			"message": "RPC Call 收到了你的请求liudehua"
					}


			http://127.0.0.1:8080/example/foo/bar  (访问 go-micro 中的 Bar() 方法)
					// 调用 Bar() 无需携带参数

		*/
}
