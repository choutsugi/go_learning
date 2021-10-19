package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// StatCost 统计请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		// 利用context.Set()在请求上下文中设置值，供后续处理函数调用处理。
		context.Set("name", "stone")
		context.Next() // 调用该请求的剩余处理程序。
		//context.Abort() // 不调用该请求的剩余处理程序。
		cost := time.Since(start)
		log.Println(cost)
	}
}

func main() {
	router := gin.Default()
	// 为全局路由注册中间件
	router.Use(StatCost())

	router.GET("/test1", func(context *gin.Context) {
		name := context.MustGet("name")
		log.Println(name)
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello Golang!",
		})
	})

	// 为单个路由注册中间件
	router.GET("/test2", StatCost(), func(context *gin.Context) {
		name := context.MustGet("name")
		log.Println(name)
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello Golang!",
		})
	})

	// 为路由组注册中间件
	userGroup := router.Group("/user")
	userGroup.Use(StatCost())
	{
		userGroup.GET("/login", func(context *gin.Context) {

		})
	}

	_ = router.Run()
}
