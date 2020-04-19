package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

// tail  常用于日志收集， 监听日志文件
func main() {
	filename := "./my.log"
	tailFile, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	// tail 能实现的功能是： 日志每记录一行， 我们就可以从 tailFile.Lines 这个通道里拿到 一行日志数据 (前提是日志文件里 记录一条日志后进行了 换行、保存的操作)
	for true{
		msg, ok := <- tailFile.Lines   // Lines 是一个通道类型，
		if !ok {
			fmt.Printf("tail file close reopen, filename: %s\n", tailFile.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
