package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/net"

	"github.com/shirou/gopsutil/disk"

	host2 "github.com/shirou/gopsutil/host"

	"github.com/shirou/gopsutil/mem"

	"github.com/shirou/gopsutil/load"

	"github.com/shirou/gopsutil/cpu"
)

// GetCpuInfo 获取CPU使用率
func GetCpuInfo() {
	infos, err := cpu.Info()
	if err != nil {
		return
	}
	for _, info := range infos {
		fmt.Println(info)
	}

	// CPU使用率
	for {
		percents, err := cpu.Percent(time.Second, true)
		if err != nil {
			return
		}
		for id, percent := range percents {
			fmt.Printf("cpu id:%d , percent:%v\n", id, percent)
		}

	}
}

// GetCpuLoad 获取CPU负载
func GetCpuLoad() {
	info, err := load.Avg()
	if err != nil {
		return
	}
	fmt.Println(info)
}

// GetMemInfo 获取内存信息
func GetMemInfo() {
	info, err := mem.VirtualMemory()
	if err != nil {
		return
	}
	fmt.Println(info)
}

// GetHostInfo host信息
func GetHostInfo() {
	host, err := host2.Info()
	if err != nil {
		return
	}
	fmt.Println(host)
}

// GetDiskInfo 磁盘信息
func GetDiskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		return
	}
	fmt.Println(parts)
	for _, part := range parts {
		stat, err := disk.Usage(part.Mountpoint)
		if err != nil {
			return
		}
		fmt.Println(stat)
	}

	// 磁盘IO信息
	ioStats, _ := disk.IOCounters()
	for key, stat := range ioStats {
		fmt.Printf("DiskName:%v, IOStatus:%v\n", key, stat)
	}
}

// GetNetInfo 网络信息
func GetNetInfo() {
	netIOs, err := net.IOCounters(true)
	if err != nil {
		return
	}

	for _, netIO := range netIOs {
		fmt.Println(netIO)
	}
}

func main() {
	//GetCpuInfo()
	//GetCpuLoad()
	//GetMemInfo()
	//GetHostInfo()
	//GetDiskInfo()
	GetNetInfo()
}
