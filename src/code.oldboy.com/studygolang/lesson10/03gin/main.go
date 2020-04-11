package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//
/*  1.11和1.12版本
将下面两个设置添加到系统的环境变量中

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct


1.13版本之后

需要注意的是这种方式并不会覆盖之前的配置，有点坑，你需要先把系统的环境变量里面的给删掉再设置
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

*/

// 处理 /index 请求
func indexHandler(c *gin.Context){
	// c.JSON(arg1,arg2),  arg1 : 状态码， arg2 : 返回内容(这是一个 map[string]interface{} 类型)
	c.JSON(http.StatusOK, gin.H{
		"msg": "这是index 页面",
	})
}

// gin
func main() {
	// 启动一个默认的路由
	router := gin.Default()
	// 给 /hello 请求配置一个处理函数
	router.GET("/hello", func(c *gin.Context){
		c.JSON(200, gin.H{
			"msg": "Hello 沙河！",
		})
	})

	// 给 /index 请求配置一个处理函数
	router.GET("/index", indexHandler)

	// 启动 webserver
	router.Run(":8080")
	
}