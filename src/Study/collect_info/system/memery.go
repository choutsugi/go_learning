package system

import (
	"collect_info/logger"

	"github.com/shirou/gopsutil/mem"
)

func GetMemInfo() (memInfo *MemInfo) {
	memInfo = new(MemInfo)
	info, err := mem.VirtualMemory()
	if err != nil {
		logger.Z.Printf("system: get mem info failed, err:%v", err)
	}

	memInfo.Total = info.Total
	memInfo.Available = info.Available
	memInfo.Used = info.Used
	memInfo.UsedPercent = info.UsedPercent
	memInfo.Buffers = info.Buffers
	memInfo.Cached = info.Cached
	return
}
