package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
	"time"
)

// tailTask 为每一项配置（从etcd中获取的配置）的监测日志的任务结构体
type tailTask struct {
	path    string
	topic   string
	tailObj *tail.Tail
	ctx     context.Context    // 定义 context
	cancel  context.CancelFunc // 定义 context中 cancel 方法, 方便管理 goroutine (停止 goroutine)
}

type tailTaskMgr struct {
	tailTaskMap      map[string]*tailTask
	collectEntryList []CollectEntry
}

var (
	ttMgr *tailTaskMgr
)

// 创建 tailTask 的实例
func newTailTask(path, topic string) *tailTask {
	ctx, cancel := context.WithCancel(context.Background())
	task := &tailTask{
		path:   path,
		topic:  topic,
		ctx:    ctx,
		cancel: cancel,
	}
	return task
}

// 初始化 tailTask
func (t *tailTask) init() (err error) {
	tailConfig := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	t.tailObj, err = tail.TailFile(t.path, tailConfig)
	return
}

func (t *tailTask) run() {
	for true {

		select {
		case <-t.ctx.Done():
			logrus.Infof("tailTask stopped, path : %s", t.path)
			return // 收到停止新号，则 return, 函数返回之后，该 goroutine 就会停止
		case line, ok := <-t.tailObj.Lines: // 一直从通道中读取 tail 监测的日志信息
			if !ok {
				logrus.Errorf("tail file for filename: %s, err : %v", t.path)
				time.Sleep(100 * time.Millisecond) // 读取通道时，休眠100 毫秒，避免一直从通道中取值时占用太多资源
				continue
			}
			// fmt.Println("line:", line.Text)

			// windows 下 tail 读取每一行日志，会读取到 \r,
			if runtime.GOOS == "windows" {
				line.Text = strings.Trim(line.Text, "\r")
				logrus.Info("windows trim \r!")
			}
			if runtime.GOOS == "linux" {
				line.Text = strings.Trim(line.Text, "\n")
				logrus.Info("linux trim \n!")
			}

			if len(line.Text) == 0 {
				logrus.Info("empty line!")
				continue
			}
			//// 判断是否为空行, 如果为空行，则忽略
			//if len(strings.TrimSpace(line.Text)) == 0 {
			//	logrus.Info("empty line!")
			//	continue
			//}

			// 注意 tail 监测日志文件时, line.Text 没一行的内容后面会加一个 \r, 会影响后续 json 的序列化
			fmt.Printf("line.Text : %q\n", line.Text) //  "{\"id\":8,\"name\":\"xxx\",\"age\":58}\r"

			// 利用通道将同步的发送 改为异步的 (这里不直接发送到kafka, 而是发送到 通道中，由kafka 客户端也就是producer 从通道中取值再发送给 kafka)
			// 构造一个消息
			msg := &sarama.ProducerMessage{} // 创建一个 ProducerMessage 结构体指针
			msg.Topic = t.topic
			msg.Value = sarama.StringEncoder(line.Text) // 将 line 封装成 msg.Value 类型

			fmt.Println("msg", msg.Value)

			MsgChan <- msg // 将 msg 发送到 通道中
		}

	}
	return
}

// collect 日志收集服务
// InitCollectConfig 是一个初始化方法
func InitCollectConfig(allConfig []CollectEntry) (err error) {

	// 初始化 tailTaskMgr (用于管理tailTask)
	ttMgr = &tailTaskMgr{
		tailTaskMap:      make(map[string]*tailTask, 20),
		collectEntryList: allConfig,
	}

	// 遍历从 etcd服务中获取的 日志目录以及topic 的配置项
	for _, conf := range allConfig {
		// 根据每一项配置创建一个 tailTask
		task := newTailTask(conf.Path, conf.Topic)
		err = task.init()
		if err != nil {
			// fmt.Println("tail file err:", err)
			logrus.Errorf("create tailfile = %s failed, err : %v", task.path, err)
			return
		}

		// 保存所有的 tailTask, 方便后续管理 (key, value )
		ttMgr.tailTaskMap[task.path] = task

		// 开启goroutine, tail 开始监测日志输入行为
		go task.run()
	}

	// 开启一个 goroutine 去监测 ConfigChan
	go watchConfigChan()

	return
}

func watchConfigChan() {
	for {
		newConfs := <-ConfigChan // 取不到值会一直阻塞在这里，但是一旦取到了值，则就会退出，所以也需要在外层加 for ，让它一直监听通道中的值
		logrus.Infof("get newConfs from ConfigChan : %v", newConfs)

		// 遍历新的配置项，存在的则不处理，不存在的则新添加 tailTask 开始监控新目录下的日志文件
		// 例如：新配置项为 [1,3], 现有 tailTaskMap 里为 [1,2], 则 1  不处理， 对 3 作处理(新添加一个3)， 此时tailTaskMap 为 [1,2,3]
		for _, conf := range newConfs {
			if isExsit(conf) {
				continue
			}
			task := newTailTask(conf.Path, conf.Topic)
			err := task.init()
			if err != nil {
				// fmt.Println("tail file err:", err)
				logrus.Errorf("create new tailfile = %s failed, err : %v", task.path, err)
				return
			}
			// 将新的配置项添加到集合中
			ttMgr.tailTaskMap[task.path] = task
			logrus.Infof("watchConfigChan() new tailTask : %v", conf)
			go task.run() // tailTask 开始监控新的路径下的日志文件
		}

		// 对比新的配置项与 现有的 tailTaskMap 集合， 找出 在tailTaskMap 存在，新的配置项不存在的项，则代表该项要 停止
		// 例如: 紧跟上面的操作， 新配置项为 [1,3], 现有 tailTaskMap 为 [1,2,3], 则要找出
		for key, task := range ttMgr.tailTaskMap {
			var same = false
			for _, conf := range newConfs {
				if key == conf.Path {
					same = true
					break
				}
			}
			// 根据key  找出不同的 tailTask
			if !same {
				logrus.Infof("watchConfigChan() tailTask path : %s , need stop", task.path)
				task.cancel()

				delete(ttMgr.tailTaskMap, task.path) // 将集合中的数据删除 (停止的 tailTask)
			}
		}
	}

}

// 判断 集合中是否存在 conf
func isExsit(conf CollectEntry) bool {
	_, ok := ttMgr.tailTaskMap[conf.Path]
	return ok
}
