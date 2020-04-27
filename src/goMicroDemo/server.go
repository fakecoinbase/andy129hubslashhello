package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"    // 现在要使用 v2 版本
	pb "goMicroDemo/proto"
)

// 声明结构体
type Hello struct {

}

// 实现接口方法
func (g *Hello) Info(ctx context.Context, req *pb.InfoRequest, rep *pb.InfoResponse) error {
	rep.Msg = "你好" + req.Username
	return nil
}
func main() {
	// 1, 得到微服务实例
	service := micro.NewService(
		// 设置微服务的名字，用来做访问用的
		micro.Name("hello"),
		)
	// 2, 初始化
	service.Init()

	// 3, 服务注册
	err := pb.RegisterHelloHandler(service.Server(), new(Hello))
	if err != nil {
		fmt.Println("RegisterHelloHandler err : ", err)
		return
	}

	// 4, 启动微服务
	err = service.Run()
	if err != nil {
		fmt.Println("service.Run() err : ", err)
		return
	}


	// 在终端中使用 micro 命令访问 go-micro, 如下： (命令解析：micro call 微服务名 结构体.方法 {"字段名", "字段值"})
	/*
		micro call hello Hello.Info {\"username\":\"zhangsan\"}

		返回结果：
		{
				"msg": "你好zhangsan"
		}
	 */
}
