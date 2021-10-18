package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {
	// 200：状态码。
	// gin.H：map结构，key为string，value为interface。
	c.JSON(200, gin.H{
		"message": "Hello Golang!",
	})
}

func main() {
	// 使用默认路由引擎
	r := gin.Default()

	// 指定用户使用GET请求访问/hello时，执行sayHello函数。
	r.GET("/hello", sayHello)

	// RESTful风格
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	// 启动服务：指定端口，默认8080。
	err := r.Run(":9090")
	if err != nil {
		return
	}
}
