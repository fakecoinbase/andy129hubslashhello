package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpcDemo/proto"
	"net"
)


// 定义服务端实现约定的接口
type UserInfoService struct {

}

var u = UserInfoService{}

// 实现服务端需要实现的接口
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	name := req.Name
	// 在数据库查用户信息
	if name == "zs" {
		resp = &pb.UserResponse{
			Id:    1,
			Name:  name,
			Age:   22,
			// 切片字段
			Hobby: []string{"Sing", "Run"},
		}
	}
	err = nil
	return
}

// server 端
func main() {
	// 1, 监听
	address := "127.0.0.1:8080"
	lis,err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Listen err : ", err)
		return
	}
	fmt.Printf("开始监听 %s\n", address)

	// 2, 实例化
	s := grpc.NewServer()

	// 3, 在 grpc 上注册微服务
	// 参数1：grpc 服务， 参数2：接口类型的变量
	pb.RegisterUserInfoServiceServer(s, &u)

	// 4, 启动 grpc 服务的服务端
	s.Serve(lis)
}
