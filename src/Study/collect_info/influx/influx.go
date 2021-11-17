package influx

import (
	"collect_info/logger"
	"collect_info/system"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

var (
	cli client.Client
)

// Init 初始化
func Init() (err error) {
	cli, err = client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
		//Username: "admin",
		//Password: "",
	})
	if err != nil {
		return err
	}
	return
}

func insertCpuInfo(info *system.CpuInfo) (err error) {
	points, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "collect",
		Precision: "s", // 精度，默认ns
	})
	if err != nil {
		logger.Z.Printf("influx:InsertCpuInfo failed, err:%v\n", err)
		return err
	}

	tags := map[string]string{"cpu": "cpu0"}
	fields := map[string]interface{}{
		"cpu_percent": info.CpuPercent,
	}

	point, err := client.NewPoint("cpu", tags, fields, time.Now())
	if err != nil {
		logger.Z.Printf("influx:InsertCpuInfo failed, err:%v\n", err)
		return err
	}
	points.AddPoint(point)
	err = cli.Write(points)
	if err != nil {
		logger.Z.Printf("influx:InsertCpuInfo failed, err:%v\n", err)
		return err
	}
	logger.Z.Println("influx:InsertCpuInfo success")
	return
}

func insertMemInfo(info *system.MemInfo) (err error) {
	points, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "collect",
		Precision: "s",
	})
	if err != nil {
		logger.Z.Printf("influx:InsertMemInfo failed, err:%v\n", err)
		return
	}

	tags := map[string]string{"mem": "mem"}
	fields := map[string]interface{}{
		"total":        int64(info.Total),
		"available":    int64(info.Available),
		"used":         int64(info.Used),
		"used_percent": int64(info.UsedPercent),
		//"buffers":      info.Buffers,
		//"cached":       info.Cached,
	}

	point, err := client.NewPoint("memory", tags, fields, time.Now())
	if err != nil {
		logger.Z.Printf("influx:InsertMemInfo failed, err:%v\n", err)
		return
	}
	points.AddPoint(point)
	err = cli.Write(points)
	if err != nil {
		logger.Z.Printf("influx:InsertMemInfo failed, err:%v\n", err)
		return
	}
	logger.Z.Println("influx:InsertMemInfo success")
	return
}

func insertDiskInfo(info *system.DiskInfo) (err error) {
	points, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "collect",
		Precision: "s", // 精度，默认ns
	})
	if err != nil {
		logger.Z.Printf("influx:InsertDiskInfo failed, err:%v\n", err)
		return err
	}

	for tag, stat := range info.PartitionUsageStat {
		tags := map[string]string{"path": tag}
		fields := map[string]interface{}{
			"total":               int64(stat.Total),
			"free":                int64(stat.Free),
			"used":                int64(stat.Used),
			"used_percent":        stat.UsedPercent,
			"inodes_total":        int64(stat.InodesTotal),
			"inodes_used":         int64(stat.InodesUsed),
			"inodes_free":         int64(stat.InodesFree),
			"inodes_used_percent": stat.InodesUsedPercent,
		}
		point, err := client.NewPoint("disk", tags, fields, time.Now())
		if err != nil {
			logger.Z.Printf("influx:InsertDiskInfo failed, err:%v\n", err)
			continue
		}
		points.AddPoint(point)
	}

	err = cli.Write(points)
	if err != nil {
		logger.Z.Printf("influx:InsertDiskInfo failed, err:%v\n", err)
		return err
	}
	logger.Z.Println("influx:InsertDiskInfo success")
	return
}

func insertNetInfo(info *system.NetInfo) (err error) {
	points, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "collect",
		Precision: "s", // 精度，默认ns
	})
	if err != nil {
		logger.Z.Printf("influx:InsertNetInfo failed, err:%v\n", err)
		return err
	}

	tags := map[string]string{"net": "net"}
	fields := map[string]interface{}{}

	point, err := client.NewPoint("net", tags, fields, time.Now())
	if err != nil {
		logger.Z.Printf("influx:InsertNetInfo failed, err:%v\n", err)
		return err
	}
	points.AddPoint(point)
	err = cli.Write(points)
	if err != nil {
		logger.Z.Printf("influx:InsertNetInfo failed, err:%v\n", err)
		return err
	}
	logger.Z.Println("influx:InsertNetInfo success")
	return
}

func setCpuInfo() {
	cpuInfo := system.GetCpuInfo()
	err := insertCpuInfo(cpuInfo)
	if err != nil {
		logger.Z.Errorf("influx:setCpuInfo failed, err:%v\n", err)
	}
}

func setMemInfo() {
	memInfo := system.GetMemInfo()
	err := insertMemInfo(memInfo)
	if err != nil {
		logger.Z.Errorf("influx:setMemInfo failed, err:%v\n", err)
	}
}

func setDiskInfo() {
	diskInfo := system.GetDiskInfo()
	err := insertDiskInfo(diskInfo)
	if err != nil {
		logger.Z.Errorf("influx:setDiskInfo failed, err:%v\n", err)
	}
}

func CollectSysInfo(interval time.Duration) {
	ticker := time.Tick(interval)
	for range ticker {
		setCpuInfo()
		setMemInfo()
		setDiskInfo()
	}
}
