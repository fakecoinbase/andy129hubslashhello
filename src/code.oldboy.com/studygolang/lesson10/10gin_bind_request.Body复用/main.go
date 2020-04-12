package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type formA struct {
	Name string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type formB struct {
	Title string `form:"title" binding:"required"`
	Price float64 `form:"price" binding:"required"`
}

/*
	只有某些格式需要此功能，如 JSON, XML, MsgPack, ProtoBuf。 对于其他格式, 
	如 Query, Form, FormPost, FormMultipart 可以多次调用 c.ShouldBind() 而不会造成任任何性能损失
*/
func formHandler(c *gin.Context){

	//var f1 formA
	//var f2 formB

	f1 := formA{}
	f2 := formB{}

	// 当提交的数据格式为 JSON, XML, MsgPack, ProtoBuf时， c.ShouldBind() 使用了 c.Request.Body，不可重用。 
	// 意思就是 在 if 语句中，已经获取了一次 Request.Body, 请求的数据就丢失了， 
	// 后面 else if 里面是拿不到请求数据的，因为现在 c.Request.Body 是 EOF，所以也就没法进行判断

	// 解决办法: 使用 c.ShouldBindBodyWith() 
	// c.ShouldBindBodyWith 会在绑定之前将 body 存储到上下文中。 这会对性能造成轻微影响，如果调用一次就能完成绑定的话，那就不要用这个方法。
	
	// if err1 := c.ShouldBind(&f1); err1 == nil {
	if err1 := c.ShouldBindBodyWith(&f1, binding.JSON); err1 == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "是formA 提交",
			"Name" : f1.Name,
			"Password" : f1.Password,
		})
		return 
	}else if err2 := c.ShouldBindBodyWith(&f2, binding.JSON); err2 == nil{
	//else if err2 := c.ShouldBind(&f2); err2 == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "是formB 提交",
			"Title" : f2.Title,
			"Price" : f2.Price,
		})
		return
	}else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "nothing",
		})
	}
	
}

// gin -- bind 之 request.Body 复用
func main() {

	router := gin.Default()

	router.Any("/some", formHandler)

	router.Run()
	
}