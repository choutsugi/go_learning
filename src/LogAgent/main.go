package main

// 日志收集客户端：类似于filebeat
// 收集指定目录下的日志文件，发送到kafka中。

import (
	"LogAgent/collect"
	"LogAgent/etcd"
	"LogAgent/kafka"
	"LogAgent/logger"
	"LogAgent/tailfile"
	"fmt"
	"time"

	"gopkg.in/ini.v1"
)

type Config struct {
	KafkaConfig `ini:"kafka"`
	EtcdConfig  `ini:"etcd"`
}

type KafkaConfig struct {
	Address  string `ini:"address"`
	ChanSize int64  `ini:"chan_size"`
}

type EtcdConfig struct {
	Address    string `ini:"address"`
	CollectKey string `ini:"collect_key"`
}

//func run() {
//	select {}
//}

func main() {
	var configObj = new(Config)
	logger.Init()

	// 0. 读配置文件
	err := ini.MapTo(configObj, "./config/config.ini")
	if err != nil {
		logger.Z.Errorf("load config failed, err:%v", err)
		return
	}
	logger.Z.Info("read config success!")

	// 1. 初始化连接kafka
	err = kafka.Init([]string{configObj.KafkaConfig.Address}, configObj.KafkaConfig.ChanSize)
	if err != nil {
		logger.Z.Errorf("init kafka failed, err:%v", err)
		return
	}
	logger.Z.Debug("init kafka success!")

	// 2.初始化Etcd连接
	err = etcd.Init([]string{configObj.EtcdConfig.Address})
	if err != nil {
		logger.Z.Errorf("init etcd failed, err:%v", err)
		return
	}
	logger.Z.Info("init etcd success!")

	// 3.从etcd拉取配置
	allConf, err := etcd.GetConf(configObj.EtcdConfig.CollectKey)
	if err != nil {
		logger.Z.Errorf("get conf failed, err:%v", err)
		return
	}
	fmt.Println(allConf)

	// 4.监视etcd
	go etcd.WatchConf(configObj.EtcdConfig.CollectKey)

	err = tailfile.Init(allConf)
	if err != nil {
		logger.Z.Errorf("init tailfile failed, err:%v", err)
		return
	}
	logger.Z.Debug("init tailfile success!")

	collect.Run(time.Second)
}
