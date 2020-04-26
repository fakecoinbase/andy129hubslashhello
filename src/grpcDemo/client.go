package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpcDemo/proto"
)

// gRPC 使用
/*
	gRPC与Protobuf介绍

	a, 微服务架构中，由于每个服务对应的代码库是独立运行的，无法直接调用，彼此间的通信就是个大问题
	b, gRPC可以实现微服务，将大的项目拆分为多个小且独立的业务模块，也就是服务，各服务间使用高效的protobuf协议进行RPC调用，
		gRPC默认使用protocol buffers，这是google开源的一套成熟的结构数据序列化机制（当然也可以使用其他数据格式如JSON）
	c, 可以用proto files创建gRPC服务，用message类型来定义方法参数和返回类型


	安装gRPC和Protobuf
	1, go get github.com/golang/protobuf/proto
	2, go get google.golang.org/grpc（无法使用，用如下命令代替）
		a, git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
		b, git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
		c, git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
		d, go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
		e, git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
		f, cd $GOPATH/src/
		g, go install google.golang.org/grpc
	3, go get github.com/golang/protobuf/protoc-gen-go
	4, 上面安装好后，会在GOPATH/bin下生成protoc-gen-go.exe
	5, 但还需要一个protoc.exe，windows平台编译受限，很难自己手动编译，直接去网站下载一个，地址：https://github.com/protocolbuffers/protobuf/releases/tag/v3.9.0 ，同样放在GOPATH/bin下

	6, 使用 protoc 命令生成 .pb.go 文件  (进入proto文件夹下)

		protoc -I . --go_out=plugins=grpc:. ./user.proto

	执行命令时出现以下问题：
	2020/04/26 18:06:26 WARNING: Missing 'go_package' option in "user.proto", please specify:
        option go_package = ".;proto";

	此时我们需要按照提示 更改 user.proto 文件，添加以下：
	package proto;
	option go_package = ".;proto";


	编译时，会提示一系列的包找不到 , 那么要怎么办呢？ 只有一个办法，使用 letsvpn 软件，翻墙下载, 确定可以翻墙之后，输入以下命令， 会将一些依赖包全部下载下来
		go get github.com/golang/protobuf/proto  （下载补全 google.golang.org\protobuf 包 ）

	..\github.com\golang\protobuf\proto\buffer.go:11:2: cannot find package "google.golang.org/protobuf/encoding/prototext" in any of:
        G:\Go\src\google.golang.org\protobuf\encoding\prototext (from $GOROOT)
        G:\Goworkspace\src\google.golang.org\protobuf\encoding\prototext (from $GOPATH)
..\github.com\golang\protobuf\proto\buffer.go:12:2: cannot find package "google.golang.org/protobuf/encoding/protowire" in any of:
        G:\Go\src\google.golang.org\protobuf\encoding\protowire (from $GOROOT)
        G:\Goworkspace\src\google.golang.org\protobuf\encoding\protowire (from $GOPATH)
..\github.com\golang\protobuf\proto\extensions.go:13:2: cannot find package "google.golang.org/protobuf/proto" in any of:
        G:\Go\src\google.golang.org\protobuf\proto (from $GOROOT)
        G:\Goworkspace\src\google.golang.org\protobuf\proto (from $GOPATH)
..\github.com\golang\protobuf\proto\defaults.go:8:2: cannot find package "google.golang.org/protobuf/reflect/protoreflect" in any of:
        G:\Go\src\google.golang.org\protobuf\reflect\protoreflect (from $GOROOT)
        G:\Goworkspace\src\google.golang.org\protobuf\reflect\protoreflect (from $GOPATH)
..\github.com\golang\protobuf\proto\extensions.go:15:2: cannot find package "google.golang.org/protobuf/reflect/protoregistry" in any of:
        G:\Go\src\google.golang.org\protobuf\reflect\protoregistry (from $GOROOT)
        G:\Goworkspace\src\google.golang.org\protobuf\reflect\protoregistry (from $GOPATH)
..\github.com\golang\protobuf\proto\extensions.go:16:2: cannot find package "google.golang.org/protobuf/runtime/protoiface" in any of:
        G:\Go\src\google.golang.org\protobuf\runtime\protoiface (from $GOROOT)
        G:\Goworkspace\src\google.golang.org\protobuf\runtime\protoiface (from $GOPATH)
..\github.com\golang\protobuf\proto\buffer.go:13:2: cannot find package "google.golang.org/protobuf/runtime/protoimpl" in any of:
        G:\Go\src\google.golang.org\protobuf\runtime\protoimpl (from $GOROOT)
        G:\Goworkspace\src\google.golang.org\protobuf\runtime\protoimpl (from $GOPATH)
..\github.com\golang\protobuf\ptypes\any\any.pb.go:9:2: cannot find package "google.golang.org/protobuf/types/known/anypb" in any of:
        G:\Go\src\google.golang.org\protobuf\types\known\anypb (from $GOROOT)
        G:\Goworkspace\src\google.golang.org\protobuf\types\known\anypb (from $GOPATH)
..\github.com\golang\protobuf\ptypes\duration\duration.pb.go:9:2: cannot find package "google.golang.org/protobuf/types/known/durationpb" in any of:
        G:\Go\src\google.golang.org\protobuf\types\known\durationpb (from $GOROOT)
        G:\Goworkspace\src\google.golang.org\protobuf\types\known\durationpb (from $GOPATH)
..\github.com\golang\protobuf\ptypes\timestamp\timestamp.pb.go:9:2: cannot find package "google.golang.org/protobuf/types/known/timestamppb" in any of:
        G:\Go\src\google.golang.org\protobuf\types\known\timestamppb (from $GOROOT)
        G:\Goworkspace\src\google.golang.org\protobuf\types\known\timestamppb (from $GOPATH)

*/

/*  尝试GO mod  （在没有翻墙之前，下面是无用的，确定可以翻墙了，还没有测试 go mod）
			require (
		    github.com/golang/protobuf/proto latest
		    google.golang.org/grpc latest
		)

		require google.golang.org/grpc v1.29.1

		replace google.golang.org/genproto => github.com/googleapis/go-genproto v0.0.0-20200424135956-bca184e23272

		require github.com/golang/protobuf/protoc-gen-go latest
*/

func main() {

	// 1, 创建与 grpc 服务端的连接
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())   // WithInsecure 建立安全的连接
	if err != nil {
		fmt.Println("dial err : ", err)
		return
	}

	defer conn.Close()

	// 2, 实例化 grpc 客户端
	client := pb.NewUserInfoServiceClient(conn)

	// 3, 组装参数
	req := new(pb.UserRequest)
	req.Name = "zs"
	// 4, 调用接口
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Println("GetuserInfo err : ", err)
		return
	}
	fmt.Printf("响应结果：%v\n", resp)
}

