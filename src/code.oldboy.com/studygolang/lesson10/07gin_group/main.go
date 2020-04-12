package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func liveIndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "live index",
	})
}

func liveLoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "live login",
	})
}

func shopIndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "shop index",
	})
}

func shopLoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "shop login",
	})
}

// gin -- group  路由分组
func main() {
	router := gin.Default()

	// 不同业务分组 (可实现业务分离， 请求分离)
	shoppingRouterGroup := router.Group("/shopping")
	{
		shoppingRouterGroup.GET("/index", shopIndexHandler)
		shoppingRouterGroup.GET("/login", shopLoginHandler)
	}
	liveRouterGroup := router.Group("/live")
	{
		liveRouterGroup.GET("/index", liveIndexHandler)
		liveRouterGroup.GET("/login", liveLoginHandler)
	}

	// 不同版本分组， v1 版本，与 v2 版本 分组 (版本更新，业务升级等)
	v1 := router.Group("/v1")
	{
		v1.GET("/index", shopIndexHandler)
		v1.GET("/login", shopLoginHandler)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/index", liveIndexHandler)
		v2.GET("/login", liveLoginHandler)
	}

	// 还有一张不经常用的 分组嵌套
	v3 := router.Group("/v3")
	{
		shopRouter := v3.Group("/shopping")
		{
			shopRouter.GET("/index", shopIndexHandler)
			shopRouter.GET("/login", shopLoginHandler)
		}
	}

	router.Run()
}
