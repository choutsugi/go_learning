package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second) // Sleep 5s，测试用
		c.String(http.StatusOK, "Hello Golang!")
	})

	// 将router注册到server。
	srv := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen:%s\n", err)
		}
	}()

	// 等待中断信号到来以优雅关闭服务器，5s超时时间。
	quit := make(chan os.Signal, 1) // 接收信号的通道
	// kill 默认发送 syscall.SIGTERM信号。
	// kill -2 发送 syscall.SIGINT 信号，同Ctrl-C。
	// kill -9 发送 syscall.SIGKILL 信号，无法被捕获。
	// signal.Notify将收到的 syscall.SIGINT 或 syscall.SIGTERM 信号转发给quit。
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 转发信号到channel，不阻塞
	<-quit                                               // 取值，阻塞，直至接收到信号
	log.Println("Shutdown Server...")
	// 创建5s超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5s内优雅关闭服务（超过5s则超时退出）
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
