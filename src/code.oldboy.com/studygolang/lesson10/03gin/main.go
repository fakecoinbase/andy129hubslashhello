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
func indexHandler(c *gin.Context) {
	// c.JSON(arg1,arg2),  arg1 : 状态码， arg2 : 返回内容(这是一个 map[string]interface{} 类型)
	c.JSON(http.StatusOK, gin.H{
		"msg": "这是index 页面",
	})
}

// gin 框架
// gin 框架路由的原理， 用的是 httprouter 包，使用的是 前缀树结构实现请求的检索
func main() {
	// 启动一个默认的路由
	router := gin.Default()
	// 给 /hello 请求配置一个处理函数
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello 沙河！",
		})
	})

	// 给 /index 请求配置一个处理函数
	router.GET("/index", indexHandler)

	/*  router 处理其他请求的方法
	router.POST()
	router.PUT()
	router.DELETE()
	*/

	// router 批处理
	// router.Any("/index", anyHandler) // 可以处理任何http请求

	/*  router.Any() 源码分析： (批量注册了所有的 http 请求), 所以需要在 anyHandler 处理请求的函数中，判断是属于哪一种请求,  c.Request.Method == http.MethodGet

	// Any registers a route that matches all the HTTP methods.
	// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
	func (group *RouterGroup) Any(relativePath string, handlers ...HandlerFunc) IRoutes {
		group.handle(http.MethodGet, relativePath, handlers)
		group.handle(http.MethodPost, relativePath, handlers)
		group.handle(http.MethodPut, relativePath, handlers)
		group.handle(http.MethodPatch, relativePath, handlers)
		group.handle(http.MethodHead, relativePath, handlers)
		group.handle(http.MethodOptions, relativePath, handlers)
		group.handle(http.MethodDelete, relativePath, handlers)
		group.handle(http.MethodConnect, relativePath, handlers)
		group.handle(http.MethodTrace, relativePath, handlers)
		return group.returnObj()
	}

	*/

	// http.MethodGet 等常量定义
	/*	http/method.go

			// Common HTTP methods.
		//
		// Unless otherwise noted, these are defined in RFC 7231 section 4.3.
		const (
			MethodGet     = "GET"
			MethodHead    = "HEAD"
			MethodPost    = "POST"
			MethodPut     = "PUT"
			MethodPatch   = "PATCH" // RFC 5789
			MethodDelete  = "DELETE"
			MethodConnect = "CONNECT"
			MethodOptions = "OPTIONS"
			MethodTrace   = "TRACE"
		)
	*/

	// 启动 webserver
	router.Run(":8080")

}
