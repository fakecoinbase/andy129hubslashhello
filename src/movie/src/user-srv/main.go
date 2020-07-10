package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
	"movie/src/share/config"
	"movie/src/share/pb"
	"movie/src/user-srv/db"
	"movie/src/user-srv/handler"
)

// 主程序入口
func main() {

	// 1, 创建 service
	// etcdRegistry := etcd.NewRegistry(registry.Addrs("127.0.0.1:2380"))   // 2379?
	//
	//mySelector := selector.NewSelector(
	//	selector.Registry(etcdRegistry),
	//	selector.SetStrategy(selector.RoundRobin),
	//)


	service := micro.NewService(
		//micro.Selector(mySelector),
		// micro.Registry(etcdRegistry),
		micro.Name(config.Namespace+"user"),
		//micro.Address("127.0.0.1:8888"),
		micro.Version("latest"),
		)

	// 2, 初始化 service
	service.Init(

		// init 行为
		micro.Action(func(c *cli.Context) error {

			fmt.Println("service.Init ...")
			// 初始化 db
			db.Init(config.MysqlDNS)
			// 注册服务
			err := pb.RegisterUserServiceHandler(service.Server(), handler.NewUserHandler(), server.InternalHandler(true))
			if err != nil {
				fmt.Println("RegisterUserServiceHandler err : ", err)
				return err
			}
			return nil
		}),

		// 服务停止之后触发的事件
		micro.AfterStop(func() error {
			fmt.Println("AfterStop ...")
			return nil
		}),

		// 服务启动之后触发的事件
		micro.AfterStart(func() error{
			fmt.Println("AfterStart ...")
			return nil
		}),

	)

	fmt.Println("service.Run()")
	// 4, 启动服务
	err := service.Run()
	if err != nil {
		fmt.Println("service.Run err : ", err)
		return
	}
}
