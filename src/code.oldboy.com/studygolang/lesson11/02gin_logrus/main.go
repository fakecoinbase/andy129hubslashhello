package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init(){
	log.SetFormatter(&logrus.JSONFormatter{})

	file,err := os.OpenFile("./gin.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if err != nil {
		panic(err.Error())
	}
	log.Out = file

	log.SetLevel(logrus.DebugLevel)

	// 上线模式 (发布版本)  
	// 设置为 上线模式， 则 gin 框架不会打印 内部设置的debug级别的log 信息 
	// gin.SetMode(gin.ReleaseMode)   
	/*
		- using env:	export GIN_MODE=release
 		- using code:	gin.SetMode(gin.ReleaseMode)
	*/

	// 指定gin 框架 把它的日志也记录到 .log文件里
	// gin.DefaultWriter = log.Out     // gin.DefaultWriter 默认输出方式为 os.Stdout

	// 既输出到文件中，又输出到 终端中， 设置 io.MultiWriter()
	 gin.DefaultWriter = io.MultiWriter(log.Out, os.Stdout)
	
}

// 处理 index 请求
func indexHandler(c *gin.Context){

	// 记录一条 json 格式的 warn 级别的日志信息
	log.WithFields(logrus.Fields{
		"msg":"index log msg",
	}).Warn("json warn log info")

	c.JSON(http.StatusOK, gin.H{
		"msg":"index msg",
	})
}

// gin -- logrus 操作
func main() {

	router := gin.Default()

	router.GET("/index", indexHandler)
	router.Run()
}