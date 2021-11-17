package influx

import (
	"LogAgent/system"
	"fmt"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

var (
	cli client.Client
)

// 连接
func connFlux() (err error) {
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

// 查询
func queryDB(cmd string) (res []client.Result, err error) {
	qry := client.Query{
		Command:  cmd,
		Database: "test",
	}
	if rsp, err := cli.Query(qry); err == nil {
		if rsp.Error() != nil {
			return res, rsp.Error()
		}
		res = rsp.Results
	} else {
		return res, err
	}
	return res, nil
}

func InsertCpuInfo(info *system.CpuInfo) (err error) {
	points, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "test",
		Precision: "s", // 精度，默认ns
	})
	if err != nil {
		fmt.Printf("insert failed:, err:%v", err)
		return err
	}

	tags := map[string]string{"cpu": "cpu0"}
	fields := map[string]interface{}{
		"cpu_percent": info.CpuPercent,
	}

	point, err := client.NewPoint("cpu", tags, fields, time.Now())
	if err != nil {
		fmt.Printf("insert failed:, err:%v", err)
		return err
	}
	points.AddPoint(point)
	err = cli.Write(points)
	if err != nil {
		fmt.Printf("insert failed:, err:%v", err)
		return err
	}
	fmt.Println("insert success")
	return
}

func InsertMemInfo(info *system.MemInfo) (err error) {
	points, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "test",
		Precision: "s",
	})
	if err != nil {
		fmt.Printf("insert failed:, err:%v", err)
		return
	}

	tags := map[string]string{"mem": "mem"}
	fields := map[string]interface{}{
		"total":        info.Total,
		"available":    info.Available,
		"used":         info.Used,
		"used_percent": info.UsedPercent,
		"buffers":      info.Buffers,
		"cached":       info.Cached,
	}

	point, err := client.NewPoint("memory", tags, fields, time.Now())
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
	fmt.Println("insert success")
	return
}

func InsertDiskInfo(info *system.DiskInfo) (err error) {
	points, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "test",
		Precision: "s", // 精度，默认ns
	})
	if err != nil {
		fmt.Printf("insert failed:, err:%v", err)
		return err
	}

	tags := map[string]string{"disk": "disk"}
	fields := map[string]interface{}{}

	point, err := client.NewPoint("disk", tags, fields, time.Now())
	if err != nil {
		fmt.Printf("insert failed:, err:%v", err)
		return err
	}
	points.AddPoint(point)
	err = cli.Write(points)
	if err != nil {
		fmt.Printf("insert failed:, err:%v", err)
		return err
	}
	fmt.Println("insert success")
	return
}

func InsertNetInfo(info *system.NetInfo) (err error) {
	points, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "test",
		Precision: "s", // 精度，默认ns
	})
	if err != nil {
		fmt.Printf("insert failed:, err:%v", err)
		return err
	}

	tags := map[string]string{"net": "net"}
	fields := map[string]interface{}{}

	point, err := client.NewPoint("net", tags, fields, time.Now())
	if err != nil {
		fmt.Printf("insert failed:, err:%v", err)
		return err
	}
	points.AddPoint(point)
	err = cli.Write(points)
	if err != nil {
		fmt.Printf("insert failed:, err:%v", err)
		return err
	}
	fmt.Println("insert success")
	return
}
