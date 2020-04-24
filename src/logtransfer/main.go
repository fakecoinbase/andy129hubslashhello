package main

import (
	"logtransfer/es"
	"logtransfer/kafka"
	"logtransfer/model"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

// logtransfer
/*
	0, 初始化 kafka 和 es (连接哪个 kafka以及读取 kafka 中哪个topic,  和 插入哪个 es中的哪张表)
	1，创建一个消费者 从kafka 中获取日志信息
	2, 将日志信息插入 es 中

	进入 kibana 主页检索 es 中的数据
 */

/*   项目遗留问题：
	1， 为什么 MsgChan 通道中的数据(从 kafka 中读取的)会自动带上 \r   (sarama.ConsumerMessage ？？)
	2， 反序列化 给 map 时，注意也要给 map 取地址，虽然说  map 是引用类型
	3,  添加性能优化测试，看看系统运行性能指标

 */
func main() {

	// 解析 .ini 配置文件
	var cfg = new(model.Config)
	err := ini.MapTo(cfg, "./config/logtransfer.ini")
	if err != nil {
		logrus.Error("ini parse error : ", err)
		return
	}
	logrus.Info("parse .ini success!")

	// 连接 kafka
	err = kafka.Init([]string{cfg.KafkaConfig.Address}, cfg.KafkaConfig.Topic, cfg.KafkaConfig.ChanSize)
	if err != nil {
		logrus.Error("kafka.Init error : ", err)
		return
	}
	logrus.Info("kafka init success!")

	// 连接 es
	err = es.Init(cfg.ESConfig.Address, cfg.ESConfig.Index, cfg.ESConfig.GoroutineNum)
	if err != nil {
		logrus.Error("es.Init error : ", err)
		return
	}
	logrus.Info("es init success!")

	select {

	}

}
