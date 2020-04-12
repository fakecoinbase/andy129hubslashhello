package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "login 返回",
	})
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"msg": "index 返回",
	})
}

func homeHandler(c *gin.Context) {
	// "home/index.html" 指定模板名称， 然后在输出的模板文件中，定义这个名称： {{define "home/index.html"}}  html  {{end}}
	c.HTML(http.StatusOK, "home/index.html", gin.H{
		"msg": "home 返回",
	})
}

// gin -- html -- render
func main() {

	router := gin.Default()

	// gin 框架加载模板文件，底层使用的就是  go 语言中内置的 template 包，所以它支持 模板语法，我们可以在 .html 文件中使用模板语法

	// 加载多个指定 模板文件
	// router.LoadHTMLFiles("template/index.html", "template/login.html")

	// 加载 template/**/ 目录下的所有 .html文件， 不会去加载 与 ** 同级目录的 .html文件， (**代表文件夹)
	router.LoadHTMLGlob("template/**/*")

	// 加载 指定目录下的所有模板文件 (*表示.html文件)， 这样指定时，template 目录下不能有空文件夹, 必须要有 .html 文件，
	// router.LoadHTMLGlob("template/*") //

	// 指定一个名称， 以及绝对路径 (部署至 linux 服务器上时，这里要填写 /var/www 目录)
	router.Static("/staticObj", "./static") // 类似于 给绝对路径 起一个别名，当做相对路径，让模板文件根据 相对路径去寻找资源
	router.GET("/login", loginHandler)
	router.GET("/index", indexHandler)
	router.GET("/home", homeHandler)

	// r.Run("127.0.0.1:8080")   // 指定服务器地址与端口

	// 也可以这样写：
	router.Run() // 查看源码，有默认端口：127.0.0.1:8080
}
