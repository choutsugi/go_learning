package main

import (
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"time"
)

// 连接
func connFlux() (cli client.Client, err error) {
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
func queryDB(cli client.Client, cmd string) (res []client.Result, err error) {
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

// 插入
func insertPoints(cli client.Client) {
	points, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "test",
		Precision: "s", // 精度，默认ns
	})
	if err != nil {
		fmt.Printf("insert failed:, err:%v", err)
		return
	}
	tags := map[string]string{"cpu": "ih-cpu"}
	fields := map[string]interface{}{
		"idle":   201.1,
		"system": 43.3,
		"user":   86.6,
	}

	point, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
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
}

func main() {
	conn, err := connFlux()
	if err != nil {
		return
	}

	fmt.Println(conn)

	insertPoints(conn)

	qstr := fmt.Sprintf("SELECT * FROM %s LIMIT %d", "cpu_usage", 10)
	res, err := queryDB(conn, qstr)
	if err != nil {
		fmt.Printf("query failed:, err:%v", err)
		return
	}
	for _, row := range res[0].Series[0].Values {
		for j, value := range row {
			fmt.Printf("index: %d, value: %v\n", j, value)
		}
	}
}
