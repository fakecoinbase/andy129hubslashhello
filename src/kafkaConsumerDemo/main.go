package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var wg sync.WaitGroup

// kafka 消费者示例
func main() {

	// 创建一个连接 kafka 的消费者
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Println("NewConsumer failed , err : ", err)
		return
	}

	// 通过 topic 查找 分区列表
	partitionList, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Println("find partitionList failed , err : ", err)
		return
	}
	fmt.Println(partitionList)

	// 从分区列表中依次迭代每一个分区， 从该分区中读取最新的数据
	for partition := range partitionList {
		fmt.Println("--------------partition : ", partition)
		partitionConsumer, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("find partitionConsumer failed , err : ", err)
			return
		}

		// 调用 PartitionConsumer 自身提供的 AsyncClose() 方法 释放一些资源 (关闭通道等操作)
		defer partitionConsumer.AsyncClose()

		// 异步从每个分区消费 中获取数据
		wg.Add(1)
		go func(consumer2 sarama.PartitionConsumer) {
			for msg := range partitionConsumer.Messages() {
				fmt.Printf("Partition: %d, Offset: %d, key: %s, value: %s\n",
						msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
			wg.Done()
		}(partitionConsumer)
	}

	wg.Wait()
}
