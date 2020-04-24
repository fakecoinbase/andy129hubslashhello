package es

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"logtransfer/common"
	"time"
)

var (
	client *elastic.Client
)

// 初始化 es  (address : es服务器地址，index: 写入es中的哪张表, goroutineNum: 启动多少个goroutine 去从通道中取值)
func Init(address string, index string, goroutineNum int) (err error){

	// 创建连接 elastic 的客户端
	client, err = elastic.NewClient(elastic.SetURL("http://"+address))
	if err != nil {
		logrus.Error("elastic.NewClient err : ", err)
		return
	}
	logrus.Info("elastic.NewClient success!")

	// 从通道中取值并写入 es
	// 启用多少个 goroutine 去通道中取值，并将数据写入 es 中 index 表中
	for i:=0;i<goroutineNum;i++{
		go run(index)
	}
	return
}

func run(index string){
	// logrus.Info("run : ", index)
	for {
		select {
		case  msg:= <-common.MsgChan:
			writeMsgToES(msg, index)
		default :
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func writeMsgToES(msg *sarama.ConsumerMessage, index string){

	logrus.Infof("writeMsgToES, index : %v, msg : %v", index, string(msg.Value))
	var msgMap map[string]interface{}
	// 将 msg 反序列化为 map (json 数据 ---> map集合)
	/* 我们将 根据不同平台，去除 \r 的操作放在了 kafka 的生产者那边，消费者这里不需要再处理 \r 的问题
	jsonStr := strings.Trim(string(msg.Value), "\r") // // 去除 \r
	logrus.Info("writeMsgToES, trim \r : ", jsonStr)
	json.Unmarshal([]byte(jsonStr), &msgMap)   // 反序列化：将 json 数据([]byte类型) 反序列化为一个 对象里面 (传入该对象的指针类型)
	 */
	json.Unmarshal(msg.Value, &msgMap)

	// 将 map 数据写入 指定的index 表中
	put1, err := client.Index().Index(index).BodyJson(msgMap).Do(context.Background())   // Do 代表执行
	if err != nil {
		logrus.Error("putMsgMap err : ", err)
		return
	}
	logrus.Infof("putMsgMap success! user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
