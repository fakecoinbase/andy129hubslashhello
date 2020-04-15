package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserInfo 是一个用户信息的结构体
type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"pwd"`
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func loginHandler(c *gin.Context) {
	if c.Request.Method == http.MethodPost {

		var user UserInfo
		err := c.ShouldBind(&user)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"err": "内部错误",
			})
			return
		}

		// 登陆成功：
		if user.Username == "yang" && user.Password == "123" {
			/*  maxAge 设置 cookie 有效时间，单位是 秒
			func (*gin.Context).SetCookie(name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool)
			*/
			c.SetCookie("username", user.Username, 30, "/", "127.0.0.1", false, true)
			c.Redirect(http.StatusMovedPermanently, "/home")
		} else {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"err": "用户名或密码错误",
			})
		}

	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func cookieMiddleWare(c *gin.Context) {
	username, err := c.Cookie("username")
	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"err": "请先登陆!",
		})
		return
	}
	c.Set("username", username)
	c.Next() //?   好像用不用都行
}

func homeHandler(c *gin.Context) {

	// username := c.MustGet("username").(string) // MustGet,  一旦获取不到值就直接 报错
	// username := c.GetString("username") // c.GetString("username") 获取不到值时会返回 空字符串
	username, ok := c.Get("username") // 判断当前值是否存在
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"username": username,
	})

}

// 处理 不经过路由处理的 请求
func notFoundHandler(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}

// cookie
func main() {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/index", indexHandler)

	router.Any("/login", loginHandler)

	// 执行顺序是 先 cookieMiddleWare()  再  homeHandler(),  不会因为 第一个方法返回了， 就不执行第二个
	router.GET("/home", cookieMiddleWare, homeHandler)

	// 处理 不经过路由处理的 请求
	router.NoRoute(notFoundHandler)

	router.Run()
}
