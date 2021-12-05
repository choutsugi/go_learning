package system

import "github.com/shirou/gopsutil/disk"

const (
	CpuInfoType  = "cpu"
	MemInfoType  = "mem"
	DiskInfoType = "disk"
	NetInfoType  = "net"
)

type SysInfo struct {
	IP   string
	Type string
	Data interface{}
}

// CpuInfo CPU属性
type CpuInfo struct {
	CpuPercent float64 `json:"cpu_percent"`
}

// MemInfo 内存属性
type MemInfo struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
	Buffers     uint64  `json:"buffers"`
	Cached      uint64  `json:"cached"`
}

type UsageStat = disk.UsageStat

// DiskInfo 磁盘属性
type DiskInfo struct {
	PartitionUsageStat map[string]*UsageStat
}

// NetInfo 网络属性
type NetInfo struct {
}
