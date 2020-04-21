package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

var configObj = new(Config)

// Config 是一个整个config 文件的结构体
type Config struct {
	KafkaConfig `ini:"kafka"`    // 定义tag, 指定 ini 中对应的 名称
	CollectConfig `ini:"collect"`
}

// KafkaConfig 是一个 kafka 配置信息的结构体
type KafkaConfig struct {
	Address string `ini:"address"`  // 定义tag, 指定 ini 中对应的 名称
	Topic string `ini:"topic"`
	ChanSize int64 `ini:"chan_size"`
}

// CollectConfig 是一个 collect 配置信息的结构体
type CollectConfig struct {
	LogFilePath string `ini:"logfile_path"` // 定义tag, 指定 ini 中对应的 名称
}

// logagent 开发,  etcd 使用
func main() {

	// 1，解析配置文件
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		logrus.Errorf("load config.ini err :%v", err)
		return
	}

	// 方法1：指定 section, 根据 key 获取其 value 值， section 对应 config.ini 中 [] 内容
	// kafkaAdress := cfg.Section("kafka").Key("address").String()

	/*
	 如果没有配置 [] 则可以 .Section("").Key("address").String()，
	 如果没有配置 [], 然后config.ini 里面又定义了多个 address, 则会获取最后一个定义 address 的值
	*/
	// kafkaAdress := cfg.Section("").Key("address").String()
	// fmt.Println(kafkaAdress)

	// 将 configObj 定义为全局对象
	// var configObj = new(Config)   // 创建一个结构体指针
	// 方法2：直接指定配置文件，映射到结构体中
	// err = ini.MapTo(configObj, "./config/config.ini")

	// 方法3：解析映射到 结构体
	err = cfg.MapTo(configObj)
	if err != nil {
		logrus.Errorf("MapTo err :%v", err)
		return
	}
	logrus.Info(configObj)

	// 2.1 拿到配置文件的数据，初始化 kafka
	err = InitKafkaConfig([]string{configObj.KafkaConfig.Address}, configObj.KafkaConfig.ChanSize)
	if err != nil {
		logrus.Errorf("InitKafkaConfig err :%v", err)
		return
	}

	logrus.Info("InitKafkaConfig success!")

	// 2.2 拿到配置文件的数据，初始化 collect
	err = InitCollectConfig(configObj.CollectConfig.LogFilePath)
	if err != nil {
		logrus.Errorf("InitCollectConfig err :%v", err)
		return
	}
	logrus.Info("InitCollectConfig success!")

	// 3 Collect 开始监测日志文件，实时读取每一条日志信息 (CollectTailFile -监测读取-> log --封装成Msg--> MsgChan
	err = run()
	if err != nil {
		logrus.Errorf("run err :%v", err)
		return
	}
}
// 监测日志文件，并将每一条日志信息 发送到 MsgChan 通道里
func run() (err error){

	for true{
		// 一直从通道中读取 tail 监测的日志信息
		line, ok := <- CollectTailFile.Lines   // Lines 是一个通道类型，
		if !ok {
			err = fmt.Errorf("tail file for filename: %s, err : %v\n", CollectTailFile.Filename, err)
			time.Sleep(100 * time.Millisecond)    // 读取通道时，休眠100 毫秒，避免一直从通道中取值时占用太多资源
			continue
		}
		// fmt.Println("line:", line.Text)

		// 判断是否为空行, 如果为空行，则忽略
		if len(strings.TrimSpace(line.Text)) == 0 {
			logrus.Info("empty line!")
			continue
		}

		// 利用通道将同步的发送 改为异步的 (这里不直接发送到kafka, 而是发送到 通道中，由kafka 客户端也就是producer 从通道中取值再发送给 kafka)
		// 构造一个消息
		msg := &sarama.ProducerMessage{}   // 创建一个 ProducerMessage 结构体指针
		msg.Topic = configObj.KafkaConfig.Topic   // 获取 kafka topic
		msg.Value = sarama.StringEncoder(line.Text) // 将 line 封装成 msg.Value 类型

		fmt.Println("msg", msg.Value)

		MsgChan <- msg   // 将 msg 发送到 通道中
	}
	return
}

