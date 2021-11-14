package main

// 日志收集客户端：类似于filebeat
// 收集指定目录下的日志文件，发送到kafka中。

import (
	"LogAgent/kafka"
	"LogAgent/logger"
	"LogAgent/tailfile"
	"strings"
	"time"

	"github.com/Shopify/sarama"

	"gopkg.in/ini.v1"
)

type Config struct {
	KafkaConfig   `ini:"kafka"`
	CollectConfig `ini:"collect"`
}

type KafkaConfig struct {
	Address  string `ini:"address"`
	Topic    string `ini:"topic"`
	ChanSize int64  `ini:"chan_size"`
}

type CollectConfig struct {
	LogFilePath string `ini:"logfile_path"`
}

func run() (err error) {
	for {
		line, ok := <-tailfile.TailLines()
		if !ok {
			logger.Z.Warnf("tail file close reopen, filename:%s\n", tailfile.TailFilename())
			time.Sleep(1 * time.Second)
			continue
		}
		// 过滤空行
		if len(strings.Trim(line.Text, "\r")) == 0 {
			continue
		}

		// 使用channel实现异步发送
		msg := &sarama.ProducerMessage{
			Topic: "web_log",
			Value: sarama.StringEncoder(line.Text),
		}
		kafka.Write(msg)
	}
}

func main() {
	var configObj = new(Config)
	logger.Init()

	// 0. 读配置文件
	err := ini.MapTo(configObj, "./config/config.ini")
	if err != nil {
		logger.Z.Error("load config failed, err:%v", err)
		return
	}
	logger.Z.Info("read config success!")

	// 1. 初始化连接kafka
	err = kafka.Init([]string{configObj.KafkaConfig.Address}, configObj.KafkaConfig.ChanSize)
	if err != nil {
		logger.Z.Error("init kafka failed, err:%v", err)
		return
	}
	logger.Z.Debug("init kafka success!")

	// 2. 根据配置中的日志路径使用tail收集
	err = tailfile.Init(configObj.CollectConfig.LogFilePath)
	if err != nil {
		logger.Z.Error("init tailfile failed, err:%v", err)
		return
	}
	logger.Z.Debug("init tailfile success!")

	// 3. 日志发送到kafka
	err = run()
	if err != nil {
		logger.Z.Error("run failed")
		return
	}
}
