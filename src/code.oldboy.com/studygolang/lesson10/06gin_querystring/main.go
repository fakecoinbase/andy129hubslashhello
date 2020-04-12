package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserInfo 是一个用户信息的结构体
type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password" binding:"required"` // binding:"required"  意思就是绑定的数据不能为空，是必须项
}

func queryStringHandler(c *gin.Context) {

	var defaultName = "yang"
	var defaultCity = "湖北"

	// 获取 query string 参数

	// 指定 key, 查询 value
	// nameValue := c.Query("name")

	// DefaultQuery()  当查询不到 key对应的值时，会返回一个指定的默认值
	nameValue := c.DefaultQuery("name", defaultName)
	cityValue := c.DefaultQuery("city", defaultCity)

	c.JSON(http.StatusOK, gin.H{
		"name": nameValue,
		"city": cityValue,
	})
}

func formHandler(c *gin.Context) {

	var defaultCity = "湖北"
	nameValue := c.PostForm("name")
	cityValue := c.DefaultPostForm("city", defaultCity)

	c.JSON(http.StatusOK, gin.H{
		"name": nameValue,
		"city": cityValue,
	})
}

func paramsHandler(c *gin.Context) {

	// 提取路径参数 (获取具体的请求动作)
	action := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"action": action,
	})
}

/*
	ShouldBind会按照下面的顺序解析请求中的数据完成绑定：

	果是 GET 请求，只使用 Form 绑定引擎（query）。
	如果是 POST 请求，首先检查 content-type 是否为 JSON 或 XML，然后再使用 Form（form-data）。
*/
func bindHandler(c *gin.Context) {

	var userInfo UserInfo
	// ShouldBind()会根据请求的Content-Type自行选择绑定器, 可解析 form ,json, xml, querystring 等格式
	err := c.ShouldBind(&userInfo)

	/*  指定绑定 解析某个格式的数据
	c.ShouldBindJSON()
	c.ShouldBindXML()
	c.ShouldBindQuery()
	c.ShouldBindHeader()
	c.ShouldBindYAML()
	*/

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
			// 当 password 的tag  设置为 binding:"required",  如果没有获取到 password的信息，或者password 为空字符串， 则会报错：
			// "err": "Key: 'UserInfo.Password' Error:Field validation for 'Password' failed on the 'required' tag"
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"username": userInfo.Username,
		"password": userInfo.Password,
	})
}

// gin -- querystring, form
func main() {

	router := gin.Default()

	// 获取url 上携带的参数
	// query string: http://127.0.0.1:8080/search?name=小王子&city=beijing    (key=value&key=value&..形式)
	router.GET("/query_string", queryStringHandler)

	// 注意 获取form 表单数据，是 post 还是 get ，取决于html里面的 method 的定义， 一般使用 POST
	// POSTMAN 模拟网页请求，里面 form 表单数据请求是 ：执行 x-www-form-urlencoded 标准
	// 如若涉及到 文件上传，则POSTMAN 中执行的是 : form-data 标准
	router.POST("/form", formHandler)

	// URL 参数: /book/list ,  /book/new  , /book/delete
	// /posts/2019/01 ,  /posts/2019/06
	router.GET("/book/:action", paramsHandler)

	// ShouldBind()会根据请求的Content-Type自行选择绑定器，大概意思就是  支持所有请求数据格式， 内部根据不同的数据格式进行解析
	router.POST("/bind", bindHandler)

	router.Run()

}
