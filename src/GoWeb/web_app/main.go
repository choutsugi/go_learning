package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/routes"
	"web_app/settings"

	"go.uber.org/zap"
)

// Go Web开发通用的脚手架模板

func main() {

	// 1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}

	// 2.初始化日志
	if err := logger.Init(settings.Conf.Log); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer logger.Sync()

	// 3.初始化MySQL
	if err := mysql.Init(settings.Conf.MySQL); err != nil {
		zap.L().Fatal("init mysql failed", zap.Error(err))
		return
	}
	defer mysql.Close()

	// 4.初始化Redis
	if err := redis.Init(settings.Conf.Redis); err != nil {
		zap.L().Fatal("init redis failed", zap.Error(err))
		return
	}
	defer redis.Close()

	// 5.注册路由
	router := routes.Setup()

	// 6.启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.App.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
