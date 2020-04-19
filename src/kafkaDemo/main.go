package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)


// kafka 日志收集, 集群管理
// 参考资料：https://www.liwenzhou.com/posts/Go/go_kafka/
// Go语言中连接kafka使用第三方库:github.com/Shopify/sarama。
/*
	sarama v1.20之后的版本加入了zstd压缩算法，需要用到cgo，在Windows平台编译时会提示类似如下错误：

	# github.com/DataDog/zstd
	exec: "gcc":executable file not found in %PATH%
	所以在Windows平台请使用v1.19版本的sarama。
*/

/*   kafka  linux 下配置

server.properties     kafka 配置文件 （可配置端口，数据存放目录等）

	log.dirs=/tmp/kafka-logs

zookeeper.properties  zookeeper 配置文件（可配置端口，数据存放目录等）

	dataDir=/tmp/zookeeper


第一步：启动 zookeeper 和 kafka 服务

启动zookeeper

./zookeeper-server-start.sh ../config/zookeeper.properties

启动 kafaka

./kafka-server-start.sh ../config/server.properties

第二步：使用 sarama 包编写go 语言程序 来向 kafka 发送 消息
      使用 tail包监听日志文件的写入，每读到一条日志就将信息发送到 kafka


第三步: 使用 kafka 提供的 消费者终端工具查看 kafka 里面的消息
./kafka-console-consumer.sh --bootstrap-server 127.0.0.1:9092 --topic shopping --from-beginning

 */
func main() {
	// 1，生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // ACK确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 分区
	config.Producer.Return.Successes = true // 确认

	// 2, 连接 kafka  (kafka 默认端口 9092)
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err : ", err)
		return
	}
	defer client.Close()

	// 3, 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "shopping"
	msg.Value = sarama.StringEncoder("this is a test log")
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err : ",err)
		return
	}

	fmt.Println("pid:%v , offset:%v\n", pid, offset)

}
