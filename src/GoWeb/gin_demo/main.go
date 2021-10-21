package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "OK")
	})
	_ = r.Run()

}
