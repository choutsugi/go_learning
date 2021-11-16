package main

import (
	"fmt"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/shirou/gopsutil/cpu"
)

var (
	cli client.Client
)

// 连接
func initConnFlux() (err error) {
	cli, err = client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
		//Username: "admin",
		//Password: "",
	})
	if err != nil {
		return
	}
	return
}

// 插入
func insertPoints(percent int64) {
	points, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "monitor",
		Precision: "s", // 精度，默认ns
	})
	if err != nil {
		fmt.Printf("insert failed:, err:%v", err)
		return
	}
	tags := map[string]string{"cpu": "cpu0"}
	fields := map[string]interface{}{
		"cpu_percent": percent,
	}

	point, err := client.NewPoint("cpu_percent", tags, fields, time.Now())
	if err != nil {
		fmt.Printf("insert failed:, err:%v", err)
		return
	}
	points.AddPoint(point)
	err = cli.Write(points)
	if err != nil {
		fmt.Printf("insert failed:, err:%v", err)
		return
	}
	//fmt.Println("insert success")
}

func getCpuInfo() {
	// 获取CPU信息
	percent, _ := cpu.Percent(time.Second, false)
	// 写入influxDB
	fmt.Printf("%s:%v\n", time.Now().String(), percent)
	insertPoints(int64(percent[0]))
}

func main() {
	err := initConnFlux()
	if err != nil {
		fmt.Printf("init connect influxDB failed, err:%v\n", err)
		return
	}

	ticker := time.Tick(time.Millisecond * 1000)
	for {
		select {
		case <-ticker:
			fmt.Printf("%s\n", time.Now().String())
			getCpuInfo()
		}
	}
}
