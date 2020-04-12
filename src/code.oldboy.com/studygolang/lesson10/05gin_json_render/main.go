package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Name string `json:"name"`
	Password string `json:"pwd"`
}


func indexHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"msg":"这是index 页面!",
	})
}

func helloHandler(c *gin.Context){
	var userInfo = user{
		Name : "阳浩", 
		Password : "123456",
	}
	c.JSON(http.StatusOK,userInfo)
}

func xmlInfoHandler(c *gin.Context){
	c.XML(http.StatusOK, gin.H{
		"msg":"这是 xml 信息", 
	})
}

func yamlInfoHandler(c *gin.Context){
	c.YAML(http.StatusOK, gin.H{
		"msg":"这是 YAML 信息",
	})
}

// gin -- json -- render
func main() {
	
	router := gin.Default()

	// 单条 json 数据
	router.GET("/index", indexHandler)
	/*
		{
			"msg": "这是index 页面!"
		}
	*/

	// json 对象序列化
	router.GET("/hello", helloHandler)
	/*
		{
			"name": "阳浩",
			"pwd": "123456"
		}
	*/

	// xml 
	router.GET("/xmlInfo", xmlInfoHandler)
	/*
		<map>
   			 <msg>这是 xml 信息</msg>
		</map>
	*/

	// YAML 类似于配置信息文档
	router.GET("/yamlInfo", yamlInfoHandler)
	/*
		msg: 这是 YAML 信息
	*/


	// protobuf 渲染后续再研究

	router.Run()
}