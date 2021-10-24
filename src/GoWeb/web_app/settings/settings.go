package settings

import (
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	App   *AppConfig   `mapstructure:"app"`
	Log   *LogConfig   `mapstructure:"log"`
	MySQL *MySQLConfig `mapstructure:"mysql"`
	Redis *RedisConfig `mapstructure:"redis"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {

	// 方式一：指定配置文件（相对于执行目录）
	//viper.SetConfigFile("./conf/config.yaml")
	// 方式二：自动查找配置文件（相对于执行目录）
	//viper.SetConfigName("config")
	//viper.AddConfigPath(".")
	// 方式三：通过命令行参数指定
	filePath := flag.String("config", "./config.yaml", "web_app -config=\"./config.yaml\"")
	flag.Parse()
	viper.SetConfigFile(*filePath)

	// 读取配置信息
	err = viper.ReadInConfig()
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("Viper.ReadInConfig() failed, err:%v\n", err)
		return
	}
	// 配置信息反序列化到全局变量Conf
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("Viper.Unmarshal() failed, err:%v\n", err)
	}
	// 监控配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config.yaml has been modified.")
		// 配置文件发生变化后，反序列化新的配置信息到全局变量Conf
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("Viper.Unmarshal() failed, err:%v\n", err)
		}
	})

	return
}
