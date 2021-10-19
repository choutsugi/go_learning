package main

import (
	"net/http"

	"github.com/gin-gonic/gin/testdata/protoexample"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// JSON渲染
	r.GET("/someJson", func(context *gin.Context) {
		// 方式一：自行拼接Json
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello Golang!",
		})
	})

	r.GET("/moreJson", func(context *gin.Context) {
		// 方式二：使用结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}

		msg.Name = "stone"
		msg.Message = "Hello Golang!"
		msg.Age = 18
		context.JSON(http.StatusOK, msg)
	})

	// XML渲染
	r.GET("/someXML", func(context *gin.Context) {
		// 方式一：自行拼接
		context.XML(http.StatusOK, gin.H{
			"message": "Hello Golang!",
		})
	})

	r.GET("/moreXML", func(context *gin.Context) {
		// 方式二：使用结构体
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}

		var msg MessageRecord
		msg.Name = "stone"
		msg.Message = "Hello Golang!"
		msg.Age = 18
		context.XML(http.StatusOK, msg)
	})

	// YAML渲染
	r.GET("someYAML", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{
			"message": "ok",
			"status":  http.StatusOK,
		})
	})

	// protobuf渲染
	r.GET("someProtobuf", func(context *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 输出被序列化的二进制数据。
		context.ProtoBuf(http.StatusOK, data)
	})

	err := r.Run(":60000")
	if err != nil {
		return
	}
}
