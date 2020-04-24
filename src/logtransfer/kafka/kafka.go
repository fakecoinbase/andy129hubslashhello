package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"logtransfer/common"
)

var (
	consumer sarama.Consumer
)

// 初始化 kafka
func Init(address []string, topic string, chanSize int64) (err error){

	// 创建一个连接 kafka 的消费者
	consumer, err = sarama.NewConsumer(address, nil)
	if err != nil {
		logrus.Error("NewConsumer failed , err : ", err)
		return
	}

	// 初始化 MsgChan
	common.MsgChan = make(chan *sarama.ConsumerMessage, chanSize)

	// 通过 topic 查找 分区列表
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		logrus.Error("find partitionList failed , err : ", err)
		return
	}
	logrus.Info(partitionList)

	// 从分区列表中依次迭代每一个分区， 从该分区中读取最新的数据
	for partition := range partitionList {
		logrus.Info("partition : ", partition)
		var partitionConsumer sarama.PartitionConsumer
		partitionConsumer, err = consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			logrus.Error("find partitionConsumer failed , err : ", err)
			return
		}
		// 调用 PartitionConsumer 自身提供的 AsyncClose() 方法 释放一些资源 (关闭通道等操作)
		// defer partitionConsumer.AsyncClose()   后台一直在运行，所以暂时不关闭， 后期在哪里关闭呢？

		// 异步从每个分区消费 中获取数据
		// 将从 kafka 中获取的数据，发送至通道
		go run(partitionConsumer)

	}
	return
}

func run(partitionConsumer sarama.PartitionConsumer){
	for msg := range partitionConsumer.Messages() {   // 这是一个通道，所以 一直从通道里读取值 ， 测试一下 for 加 select 用法与这里的不同之处

		logrus.Info("get msg from kafka, msg : ", string(msg.Value))
		common.MsgChan <- msg

		logrus.Info("Partition: %d, Offset: %d, key: %s, value: %s\n",
			msg.Partition, msg.Offset, msg.Key, msg.Value)
	}
}
