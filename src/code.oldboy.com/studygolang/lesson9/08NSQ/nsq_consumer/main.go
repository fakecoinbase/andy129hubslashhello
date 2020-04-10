package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

// windows 下开启 nsqd 服务 顺序： (可以先尝试 nsqlookupd 查看 tcp ,http 协议的端口，然后关闭 nsqlookupd, 最后按照顺序执行如下：)
/*
	nsqd -broadcast-address=127.0.0.1 -lookupd-tcp-address=127.0.0.1:4160
	nsqadmin -lookupd-http-address=127.0.0.1:4161
	nsqlookupd
*/

// NSQ Consumer Demo

// MyHandler 是一个消费者类型
type MyHandler struct {
	Title string
}

// HandleMessage 是需要实现的处理消息的方法
// 类似于回掉函数，当有消息发送过来时，会调用这个方法， 用户就可以得到信息
func (m *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	fmt.Printf("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	return
}

// 初始化消费者
func initConsumer(topic string, channel string, address string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second     // 每15秒查询一次，有没有新的 nsqd 节点加入进来
	c, err := nsq.NewConsumer(topic, channel, config) // 创建一个 nsq 的 consumer 对象，由它来管理 具体的 用户对象
	if err != nil {
		fmt.Printf("create consumer failed, err:%v\n", err)
		return
	}
	// 创建一个 用户实体
	consumer := &MyHandler{
		Title: "沙河1号",
	}
	c.AddHandler(consumer) // 将用户加入一个队列（类似于等待取消息的队列）

	// if err := c.ConnectToNSQD(address); err != nil { // 直接连NSQD   (模型： producer --传值--> nsqd --> consumer)
	// 非集群模式下直接连接指定 nsqd, 集群模式下通过 nsqlookupd 查询到 nsqd 地址再连接。
	if err := c.ConnectToNSQLookupd(address); err != nil { // 通过lookupd查询  (模型: producer ---> nsqd <----> nsqlookupd <--- consumer)
		return err
	}
	return nil

}

func main() {
	// 创建一个 consumer 从 指定的 topic, 指定的通道，指定的 nsqd 或者 nsqlookupd中取值
	// 如果有 first 通道则直接取值，如果没有 则创建。
	err := initConsumer("topic_demo", "first", "127.0.0.1:4161")
	if err != nil {
		fmt.Printf("init consumer failed, err:%v\n", err)
		return
	}
	c := make(chan os.Signal) // 定义一个信号的通道
	// 这里进行一个阻塞主程序的操作， 监听一个 ctrl+c 键盘输入
	signal.Notify(c, syscall.SIGINT) // 转发键盘中断信号到c
	<-c                              // 阻塞  ( 没有监听到键盘输入就阻塞在这里)
}
