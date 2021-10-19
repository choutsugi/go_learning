package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 绑定到JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"user" json:"password" binding:"required"`
}

func main() {
	// Default返回默认路由引擎
	router := gin.Default()
	// 获取querystring参数
	router.GET("/user/search", func(context *gin.Context) {
		// DefaultQuery获取不到值时，返回默认值。
		//username := context.DefaultQuery("username", "stone")
		username := context.Query("username")
		address := context.Query("address")

		// 输出json返回到调用方
		context.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	// 获取form参数
	router.POST("/user/search", func(context *gin.Context) {
		router.POST("/user/search", func(context *gin.Context) {
			// DefaultPostForm获取不到值时，返回默认值。
			//username := context.DefaultPostForm("username","stone")
			username := context.PostForm("username")
			address := context.PostForm("address")

			// 输出json返回到调用方
			context.JSON(http.StatusOK, gin.H{
				"message":  "ok",
				"username": username,
				"address":  address,
			})
		})
	})

	// 获取JSON参数
	router.POST("/json", func(context *gin.Context) {
		// 从context.Request.Body中读取数据。
		data, err := context.GetRawData()
		if err != nil {
			return
		}
		var m map[string]interface{}
		// 反序列化
		err = json.Unmarshal(data, &m)
		if err != nil {
			return
		}

		context.JSON(http.StatusOK, m)
	})

	// 获取path参数
	router.GET("/user/search/:username/:address", func(context *gin.Context) {
		username := context.Param("username")
		address := context.Param("address")

		// 输出json返回到调用方
		context.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	// 参数绑定：识别请求数据类型，利用反射机制自动提取到结构体中。
	// 绑定JSON示例：({"user": "stone", "password": "123456"})
	router.POST("/loginJSON", func(context *gin.Context) {
		var login Login
		// ShouldBind()根据请求的Content-Type自行选择绑定器。
		err := context.ShouldBind(&login)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Printf("login info:%#v\n", login)
		context.JSON(http.StatusOK, gin.H{
			"user":     login.User,
			"password": login.Password,
		})
	})

	// 绑定form示例：(user=stone&password=123456)
	router.POST("/loginForm", func(context *gin.Context) {
		var login Login
		// ShouldBind()根据请求的Content-Type自行选择绑定器。
		err := context.ShouldBind(&login)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Printf("login info:%#v\n", login)
		context.JSON(http.StatusOK, gin.H{
			"user":     login.User,
			"password": login.Password,
		})
	})

	// 绑定querystring示例：(/loginQuery?user=stone&password=123456)
	router.GET("/loginForm", func(context *gin.Context) {
		var login Login
		// ShouldBind()根据请求的Content-Type自行选择绑定器。
		err := context.ShouldBind(&login)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Printf("login info:%#v\n", login)
		context.JSON(http.StatusOK, gin.H{
			"user":     login.User,
			"password": login.Password,
		})
	})

	err := router.Run(":9000")
	if err != nil {
		return
	}
}
