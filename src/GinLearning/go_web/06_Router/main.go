package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// 普通路由
	router.GET("/index", func(context *gin.Context) {
		// ...
	})

	router.GET("/login", func(context *gin.Context) {
		// ...
	})

	router.POST("/login", func(context *gin.Context) {
		// ...
	})

	// 匹配所有请求方法
	router.Any("/test", func(context *gin.Context) {
		// ...
	})

	// 为没有配置处理函数的路由添加处理程序
	router.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "views/404.html", nil)
	})

	// 路由组
	// 用户组
	userGroup := router.Group("/user")
	{
		userGroup.POST("/user/add", func(context *gin.Context) {

		})
		userGroup.DELETE("/user/delete", func(context *gin.Context) {

		})
	}
	// 文章组
	articleGroup := router.Group("/article")
	{
		articleGroup.GET("/article/search", func(context *gin.Context) {

		})

		articleGroup.PUT("/article/update", func(context *gin.Context) {

		})
	}

	_ = router.Run(":9000")
}
