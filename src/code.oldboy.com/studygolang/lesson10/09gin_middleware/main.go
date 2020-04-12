package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func shopIndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "shop index",
	})
}

func shopLoginHandler(c *gin.Context) {

	// 获取中间件里设置的值， .(string) 类型断言，如果猜错会 panic ，但是由于 gin 框架中默认开启了 Recovery() 异常处理，所不会让程序崩溃
	value := c.MustGet("name").(string)
	fmt.Println("value : ", value)

	c.JSON(http.StatusOK, gin.H{
		"msg": "shop login",
	})
}

// 计算请求耗时
func castTime(c *gin.Context) {
	now := time.Now()

	// 在中间件中设置值，可以在其他业务逻辑中获取  (前提是这个中间件作用的业务逻辑)
	c.Set("name", "castTime 中间件")

	c.Next()

	cast := time.Since(now)
	fmt.Println("耗时：", cast)
}

// gin -- middleware  中间件
func main() {
	router := gin.Default() // Default() 方法 默认使用了 两个中间件 engine.Use(Logger(), Recovery()) :  Logger()记日志， Recovery() 异常处理
	// router := gin.New()  // 也可以直接调用 New()方法初始化 路由

	/*
		gin默认中间件
		gin.Default()默认使用了Logger和Recovery中间件，其中：

		Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
		Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
		如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。
	*/

	// Use(middleware ...gin.HandlerFunc)
	router.Use(castTime) // 设置整个路由的中间件， 用于整个业务的中间件 （常用于 登陆请求等）

	// 不同业务分组 (可实现业务分离， 请求分离)
	shoppingRouterGroup := router.Group("/shopping")
	// shoppingRouterGroup.Use(castTime)   // 组路由 也可以设置 中间件， 让中间件只作用与 这个组的业务逻辑
	{
		// GET() 里面可以传入多个 handler 函数(func(*Contxt)), 执行顺序按照传入的顺序 (先执行 castTime, 再执行 shopIndexHandler)
		//  castTime 针对单独 shopping 业务的中间件
		// shoppingRouterGroup.GET("/index", castTime, shopIndexHandler)

		shoppingRouterGroup.GET("/index", shopIndexHandler)
		shoppingRouterGroup.GET("/login", shopLoginHandler)
	}
	router.Run()
}
