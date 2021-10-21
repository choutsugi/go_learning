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
		// 多个文件 Multipart form
		form, _ := context.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("../tmp/%s_%d", file.Filename, index)
			// 上传文件到指定目录
			_ = context.SaveUploadedFile(file, dst)
		}
		context.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", len(files)),
		})
	})

	_ = router.Run()
}
