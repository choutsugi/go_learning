package main

import (
	"fmt"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 设置默认值
	viper.SetDefault("fileDir", "./")
	// 读取配置文件
	//viper.SetConfigFile("config.yaml")
	viper.SetConfigName("config") // 配置文件名（不包括后缀）
	viper.SetConfigType("yaml")   // 配置文件类型
	// 查找配置文件所在的路径：多次调用以添加多个搜索路径
	viper.AddConfigPath("/etc/appname/")  // 指定目录查找
	viper.AddConfigPath("$HOME/.appname") // 用户目录下查找
	viper.AddConfigPath(".")              // 工作目录中查找配置

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file:%s\n", err))
	}

	// 实时监控配置文件的变化
	viper.WatchConfig()
	// 当配置文件变化后调用回调函数
	viper.OnConfigChange(func(in fsnotify.Event) {
		// 配置文件发生变化后调用的回调函数
		fmt.Println("config file changed!", in.Name)
	})

	r := gin.Default()
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("version"))
	})

	r.Run(":9000")
}
