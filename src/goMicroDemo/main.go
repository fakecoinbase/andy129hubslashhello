package main

// go-micro 学习
/*
	1.go-micro简介
		Go Micro是一个插件化的基础框架，基于此可以构建微服务，Micro的设计哲学是可插拔的插件化架构
		在架构之外，它默认实现了consul作为服务发现（2019年源码修改了默认使用mdns），通过http进行通信，通过protobuf和json进行编解码

	2.go-micro的主要功能
		a, 服务发现：自动服务注册和名称解析。服务发现是微服务开发的核心。当服务A需要与服务B通话时，它需要该服务的位置。默认发现机制是多播DNS（mdns），
			一种零配置系统。您可以选择使用SWIM协议为p2p网络设置八卦，或者为弹性云原生设置设置consul
		b, 负载均衡：基于服务发现构建的客户端负载均衡。一旦我们获得了服务的任意数量实例的地址，我们现在需要一种方法来决定要路由到哪个节点。
			我们使用随机散列负载均衡来提供跨服务的均匀分布，并在出现问题时重试不同的节点
		c, 消息编码：基于内容类型的动态消息编码。客户端和服务器将使用编解码器和内容类型为您无缝编码和解码Go类型。
			可以编码任何种类的消息并从不同的客户端发送。客户端和服务器默认处理此问题。这包括默认的protobuf和json
		d, 请求/响应：基于RPC的请求/响应，支持双向流。我们提供了同步通信的抽象。对服务的请求将自动解决，负载平衡，拨号和流式传输。
			启用tls时，默认传输为http / 1.1或http2
		e, Async Messaging：PubSub是异步通信和事件驱动架构的一流公民。
			事件通知是微服务开发的核心模式。启用tls时，默认消息传递是点对点http / 1.1或http2
		f, 可插拔接口：Go Micro为每个分布式系统抽象使用Go接口，因此，这些接口是可插拔的，并允许Go Micro与运行时无关，可以插入任何基础技术
		插件地址：https://github.com/micro/go-plugins
 */

/*
	1.go-micro安装
		查看的网址：https://github.com/micro/go-micro
		cmd中输入下面3条命令下载，会自动下载相关的很多包
		go get github.com/micro/micro
		go get github.com/micro/go-micro
		go get github.com/micro/protoc-gen-micro

	根据官方文档， go-micro 要使用最新的版本 go-micro/v2,  直接 go get github.com/micro/go-micro/v2 不能下载
	根据官方提示，需要开启 go mod


	2, go mod 开启后，依赖包都下载完毕，服务端代码也编写完毕，服务端启动成功， 如何验证 go-micro 服务ok呢
		我们使用 micro 命令， 在终端中通过命令的形式 调用 go-micro 服务中的方法 （micro.exe 如何获取：）

		a, 首先 micro.exe 如何获取： 参考文档 https://www.cnblogs.com/flash55/p/11300627.html
			根据参考进入 github.com\micro\micro 目录下，执行 go build , 会在当前目录生成  micro.exe 文件
			然后将 micro.exe 放入到 go bin 目录下, 即可在 终端中使用 micro 命令

        b, 执行 micro 命令 访问 go-micro
      		// 在终端中使用 micro 命令访问 go-micro, 如下： (命令解析：micro call 微服务名 结构体.方法 {"字段名", "字段值"})
			/*
				micro call hello Hello.Info {\"username\":\"zhangsan\"}

				返回结果：
				{
						"msg": "你好zhangsan"
				}
			*/
