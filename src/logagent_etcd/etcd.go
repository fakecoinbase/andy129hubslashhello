package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"time"
)

// EtcdClient 是一个 ectd 客户端
var (
	EtcdClient *clientv3.Client
)

// InitEtcd 初始化 etcd 服务，创建一个连接 etcd 的客户端
func InitEtcd(address []string) (err error) {
	EtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		err = fmt.Errorf("connect to etcd failed, err : %v\n", err)
		return
	}

	// 初始化新配置的通道(用于更新 etcd 中设置的配置项)
	ConfigChan = make(chan []CollectEntry) // 创建一个无缓冲区的通道

	return
}

// 从etcd 中获取 多个日志目录配置项
func GetConf(key string) (collectEntryList []CollectEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	defer cancel()

	resp, err := EtcdClient.Get(ctx, key)
	if err != nil {
		err = fmt.Errorf("get from etcd failed, err : %v\n", err)
		return
	}
	if len(resp.Kvs) == 0 {
		logrus.Warningf("get resp.Kvs == 0 from etcd key : %s\n", key)
		return
	}

	ret := resp.Kvs[0] // 获取配置数据
	// fmt.Println(ret.Value)

	// 从 etcd 中获取的配置信息(json 数据)，反序列化为 collectEntryList  (CollectEntry 类型的切片)
	err = json.Unmarshal(ret.Value, &collectEntryList)
	if err != nil {
		err = fmt.Errorf("json unmarshal failed, err : %v\n", err)
		return
	}
	return
}

func WatchEtcdConfig(key string) {

	for {
		// Watch 一次就只监听一次，为了一直监听 key ,所以我们要在外层加 for 循环，一直监听
		wCh := EtcdClient.Watch(context.Background(), key)

		// 一直从 etcd 监听通道里取值 (事件状态信息)
		for wresp := range wCh {
			logrus.Info("get new conf from etcd!")
			for _, evt := range wresp.Events {
				logrus.Infof("Type: %s, key: %s, value: %s", evt.Type, evt.Kv.Key, evt.Kv.Value)

				var newConf []CollectEntry
				// 如果是 etcd 的删除事件，代表对key 进行了删除操作，所以我们返回一个 空的 []CollectEntry 切片
				if evt.Type == clientv3.EventTypeDelete {
					ConfigChan <- newConf
					continue
				}
				// 将新的配置项信息 反序列化到 newConf 中
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					logrus.Errorf("json unmarshal new conf failed, err : %v", err)
					continue
				}
				// 将新配置信息发送到通道中，让 collect 日志收集服务更新配置信息
				ConfigChan <- newConf
			}
		}
	}

}
