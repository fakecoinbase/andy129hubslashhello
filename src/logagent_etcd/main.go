package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

var configObj = new(Config)

// Config 是一个整个config 文件的结构体
type Config struct {
	KafkaConfig   `ini:"kafka"` // 定义tag, 指定 ini 中对应的 名称
	CollectConfig `ini:"collect"`
	EtcdConfig    `ini:"etcd"`
}

// KafkaConfig 是一个 kafka 配置信息的结构体
type KafkaConfig struct {
	Address  string `ini:"address"` // 定义tag, 指定 ini 中对应的 名称
	ChanSize int64  `ini:"chan_size"`
}

// CollectConfig 是一个 collect 配置信息的结构体
type CollectConfig struct {
	LogFilePath string `ini:"logfile_path"` // 定义tag, 指定 ini 中对应的 名称
}

// EtcdConfig 是一个 etcd 配置信息的结构体
type EtcdConfig struct {
	Address    string `ini:"address"`
	CollectKey string `ini:"collect_key"` // 对应多个 日志文件路径 (后期新添加一个项目或者 多个topic的日志收集)
}

// logagent 开发,  etcd 使用
func main() {

	// 0, 获取本地IP， 初始化 collect_key
	ip, err := GetOutboundIP()
	if err != nil {
		logrus.Errorf("GetOutboundIP err :%v", err)
		return
	}

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

	// 2.2.1 初始化 etcd 服务
	err = InitEtcd([]string{configObj.EtcdConfig.Address})
	if err != nil {
		logrus.Errorf("InitEtcd err :%v", err)
		return
	}

	logrus.Info("InitEtcd success!")

	// 2.2.2从 etcd 中 获取日志路径以及topic的配置项
	collectKey := fmt.Sprintf(configObj.EtcdConfig.CollectKey, ip) // 最终效果：collect_log_%s_config  ---> collect_log_192.168.1.2_config
	logrus.Infof("collectKey : %s", collectKey)
	allConf, err := GetConf(collectKey)
	if err != nil {
		logrus.Errorf("etcd GetConf err :%v", err)
		return
	}
	fmt.Println(allConf)

	// 2.2.3 监听 etcd 服务，获取配置项的改变
	go WatchEtcdConfig(collectKey)

	// 2.3 从 etcd 服务中拿到配置项的数据(allConf)，初始化 collect 日志收集服务
	err = InitCollectConfig(allConf)
	if err != nil {
		logrus.Errorf("InitCollectConfig err :%v", err)
		return
	}
	logrus.Info("InitCollectConfig success!")

	// 采用 以下语句可造成 主线程阻塞
	select {}

}
