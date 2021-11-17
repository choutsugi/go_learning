package system

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

type DiskUsageStat struct {
	Path              string  `json:"path"`
	Fstype            string  `json:"fstype"`
	Total             uint64  `json:"total"`
	Free              uint64  `json:"free"`
	Used              uint64  `json:"used"`
	UsedPercent       float64 `json:"used_percent"`
	InodesTotal       uint64  `json:"inodes_total"`
	InodesFree        uint64  `json:"inodes_free"`
	InodesUsed        uint64  `json:"inodes_used"`
	InodesUsedPercent float64 `json:"inodes_used_percent"`
}

// DiskInfo 磁盘属性
type DiskInfo struct {
	PartitionUsageStat map[string]*DiskUsageStat
}

// NetInfo 网络属性
type NetInfo struct {
}
