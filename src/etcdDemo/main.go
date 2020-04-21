package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)
/*  linux 版

etcd 启动
/home/andy/Goworkspace/tools/etcd-v3.4.7-linux-amd64 目录下：

启动命令： （默认：127.0.0.1:2379）
./etcd

操作etcd (如果提示你没有 put 命令，则要 export ETCDCTL_API=3)
/home/andy/Goworkspace/tools/etcd-v3.4.7-linux-amd64 目录下：

put 值操作
./etcdctl --endpoints=http://127.0.0.1:2379 put class "三年二班"

get 值操作
./etcdctl get class

 */

//etcd 示例
func main() {

	// testEtcdPutGet()

	// testEtcdWatcher()

	testEtcdDel()

}

// etcd watcher 监听者
func testEtcdWatcher(){
	client, err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:time.Second*5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err : %v", err)
		return
	}
	defer client.Close()

	// watch
	watchChan := client.Watch(context.Background(), "name")   // 监听  name 这个 key 对应 value 的改变
	// 一直从 通道里取值 (事件状态信息)
	for  wresp := range watchChan {
		for _, evt := range wresp.Events {
			fmt.Printf("Type: %s, key: %s, value: %s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
		}
	}

}

// 初步学习 etcd 的连接 以及put,get 方法
func testEtcdPutGet(){

	// 创建连接 etcd 的客户端
	client, err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:time.Second*5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err : %v", err)
		return
	}
	defer client.Close()

	// put
	key := "collect_log_config"
	// confStr := `[{"path":"E:/logs/s4.log","topic":"s4_log"}]`
	// confStr := `[{"path":"E:/logs/s4.log","topic":"s4_log"},{"path":"G:/logs/mylog.log","topic":"web_log"}]`
	 confStr := `[{"path":"E:/logs/s4.log","topic":"s4_log"},{"path":"G:/logs/mylog.log","topic":"web_log"},{"path":"C:/logs/test.log","topic":"test"} ]`
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = client.Put(ctx, key, confStr)
	if err != nil {
		fmt.Printf("put to etcd failed, err : %v", err)
		return
	}

	cancel()

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	gr, err := client.Get(ctx,key)
	if err != nil {
		fmt.Printf("get from etcd failed, err : %v", err)
	}

	for _, ev := range gr.Kvs {
		fmt.Printf("key: %s, value : %s\n", ev.Key, ev.Value)
	}

	cancel()
}

func testEtcdDel(){
	// 创建连接 etcd 的客户端
	client, err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:time.Second*5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err : %v", err)
		return
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	key := "collect_log_config"
	_, err = client.Delete(ctx, key)
	if err != nil {
		fmt.Printf("put to etcd failed, err : %v", err)
		return
	}

	cancel()

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	gr, err := client.Get(ctx,key)
	if err != nil {
		fmt.Printf("get from etcd failed, err : %v", err)
	}

	for _, ev := range gr.Kvs {
		fmt.Printf("key: %s, value : %s\n", ev.Key, ev.Value)
	}

	cancel()
}

