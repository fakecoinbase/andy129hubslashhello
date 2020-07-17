package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)
/*
"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"

 */
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


/*
	注意：put 是 clientV3 版本的命令!
	如果使用 etcdctl.exe 来操作 etcd 的话，记得要设置环境变量:
	SET ETCDCTL_API=3

	Mac & linux :
	export ETCDCTL_API=3
*/

//etcd 示例
func main() {

	// testEtcdPutGet()

	// testEtcdWatcher()

	// testEtcdDel()

	 testEtcdLock()
}

var n = 0

// 使用worker模拟锁的抢占
func worker(key string, id int) error {
	endpoints := []string{"127.0.0.1:2379"}

	cfg := clientv3.Config{
		Endpoints:            endpoints,
		DialTimeout:          3 * time.Second,
	}

	cli, err := clientv3.New(cfg)
	if err != nil {
		log.Println("new cli error:", err)
		return err
	}

	sess, err := concurrency.NewSession(cli)
	if err != nil {
		return err
	}

	m := concurrency.NewMutex(sess, "/"+key)

	err = m.Lock(context.TODO())

	/*   /clientv3/concurrency/mutex.go 分布锁实现分析    参考文档： https://www.cnblogs.com/jiujuan/p/12147809.html

		m.Lock() 内部

	    m.myKey = fmt.Sprintf("%s%x", m.pfx, s.Lease())

		m.pfx是前缀，比如"/lockname/"
		s.Lease()是一个64位的整数值，etcd v3引入了lease（租约）的概念，concurrency包基于lease封装了session，
		每一个客户端都有自己的lease，也就是说每个客户端都有一个唯一的64位整形值

	排队等待, /lockname  为 key, /694d735b1c70d81c 为排队序号 (每一个客户端都会分配一个 int64位 全局唯一的 序号值)
		示例：m.myKey :
	    m.myKey : /lockname/694d735b1c70d81c
		m.myKey : /lockname/694d735b1c70d81d
		m.myKey : /lockname/694d735b1c70d81e

		或

		m.myKey : /lockname/694d735b1c70d842
		m.myKey : /lockname/694d735b1c70d846
		m.myKey : /lockname/694d735b1c70d844

			m.myRev = resp.Header.Revision   // revision是etcd一个全局的序列号,全局唯一且递增，每一个对etcd存储进行改动都会分配一个这个序号，在v2中叫index

		//如果上面的code操作成功了，则myRev是当前客户端创建的key的revision值。
	    //waitDeletes等待匹配m.pfx （"/lockname/"）这个前缀（可类比在这个目录下的）并且createRivision小于m.myRev-1的所有key被删除
	    //如果没有比当前客户端创建的key的revision小的key，则当前客户端者获得锁
	    //如果有比它小的key则等待，比它小的被删除
	    hdr, werr = waitDeletes(ctx, client, m.pfx, m.myRev-1)

	*/

	if err != nil {
		log.Println("lock error:", err)
		return err
	}

	defer func() {
		err = m.Unlock(context.TODO())   // 释放当前锁，让下一个排队者获取锁
		if err != nil {
			log.Println("unlock error:", err)
		}
	}()

	log.Println("get lock: ", n)
	n++
	time.Sleep(time.Second) // 模拟执行代码


	return nil
}

func testEtcdLock() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		err := worker("lockname", 1)
		if err != nil {
			log.Println(err)
		}
	}()


	go func() {
		defer wg.Done()
		err := worker("lockname",2)
		if err != nil {
			log.Println(err)
		}
	}()

	go func() {
		defer wg.Done()
		err := worker("lockname",3)
		if err != nil {
			log.Println(err)
		}
	}()

	wg.Wait()

	/*  程序运行结果：

	2020-07-17 12:56:59.455970 I | get lock:  0
	2020-07-17 12:57:00.462585 I | get lock:  1
	2020-07-17 12:57:01.469713 I | get lock:  2
	*/
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

	ip, err := GetOutboundIP()
	if err != nil {
		fmt.Printf("------GetOutboundIP err :%v", err)
		return
	}
	// put
	key := fmt.Sprintf("collect_log_%s_config", ip)
	fmt.Println("key : ", key)
	 confStr := `[{"path":"G:/logs/mylog.log","topic":"web_log"}]`
	// confStr := `[{"path":"E:/logs/s4.log","topic":"s4_log"},{"path":"G:/logs/mylog.log","topic":"web_log"}]`
	// confStr := `[{"path":"E:/logs/s4.log","topic":"s4_log"},{"path":"G:/logs/mylog.log","topic":"web_log"},{"path":"C:/logs/test.log","topic":"test"} ]`
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

// 通过取巧的方式获取 本地连接外网的IP
// Get preferred outbound ip of this machine
func GetOutboundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// fmt.Println(localAddr.String()) // 在本机中打印： 192.168.1.2:63095

	ip = strings.Split(localAddr.IP.String(), ":")[0]
	logrus.Infof("GetOutboundIP(), ip : %s", ip)
	return
}

