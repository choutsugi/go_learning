package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 处理multipart form提交文件时默认内存限制32Mib，通过以下方式修改：
	router.MaxMultipartMemory = 8 << 20 // 8Mib
	router.POST("/upload", func(context *gin.Context) {
		// 单个文件
		file, err := context.FormFile("file")
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Println(file.Filename)
		dst := fmt.Sprintf("../tmp/%s", file.Filename)

		// 上传文件到指定目录
		err = context.SaveUploadedFile(file, dst)
		if err != nil {
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})

	err := router.Run(":9000")
	if err != nil {
		return
	}
}
