package main

import (
	"collect_info/influx"
	"collect_info/logger"
	"time"
)

func main() {
	logger.Init()
	err := influx.Init()
	if err != nil {
		logger.Z.Printf("init influxDB failed, err:%v\n", err)
		return
	}

	influx.CollectSysInfo(time.Second)
}
