package main

import (
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

var (
	CollectTailFile *tail.Tail
)


// collect 日志收集服务
// InitCollectConfig 是一个初始化方法
func InitCollectConfig(fileName string) (err error){
	// filename := "./my.log"

	CollectTailFile, err = tail.TailFile(fileName, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		// fmt.Println("tail file err:", err)
		logrus.Errorf("create tailfile = %s failed, err : %v\n", fileName, err)
		return
	}

	return
}
