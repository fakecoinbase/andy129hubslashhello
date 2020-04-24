package main

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

var (
	kafkaClient sarama.SyncProducer
	MsgChan chan *sarama.ProducerMessage
)

// kafka 服务
// InitKafkaConfig 是一个初始化方法
func InitKafkaConfig(address []string, chanSize int64) (err error){

	// 1，生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // ACK确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 分区
	config.Producer.Return.Successes = true // 确认

	// 2, 连接 kafka  (kafka 默认端口 9092)
	kafkaClient, err = sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		//logrus.Errorf("producer create failed, err : %v", err)
		return
	}

	// 初始化msg 通道
	MsgChan = make(chan *sarama.ProducerMessage, chanSize)

	//  client --从通道读取Msg并发送-> kafka (开启一个 goroutine 在后台一直从 MsgChan 通道中读取信息，并发送给 kafka)
	go sendMsg()
	return

	// 初始化一个全局的 Client , 暂时不用关闭
	// defer Client.Close()

	/*
	// 3, 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "shopping"
	msg.Value = sarama.StringEncoder("this is a test log")
	// 4, 发送消息
	pid, offset, err := Client.SendMessage(msg)
	if err != nil {
		logrus.Error("send msg failed, err : %v",err)
		return
	}
	*/
}

func sendMsg(){
	for true {
		select {
			case msg := <- MsgChan:
				pid, offset, err := kafkaClient.SendMessage(msg)
				if err != nil {
					logrus.Infof("pid : %d, offset : %d, send msg failed, err : %v", pid, offset, err)
					return
				}
		}
	}

}

