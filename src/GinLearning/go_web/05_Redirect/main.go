package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// HTTP重定向
	router.GET("/test", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.bilibili.com")
	})

	// 路由重定向
	router.GET("/test1", func(context *gin.Context) {
		// 指定重定向的URL
		context.Request.URL.Path = "/test2"
		router.HandleContext(context)
	})

	router.GET("/test2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"Hello": "Golang!",
		})
	})

	_ = router.Run(":9000")
}
