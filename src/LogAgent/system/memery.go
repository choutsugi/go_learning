package system

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

func GetMemInfo() (memInfo *MemInfo) {
	memInfo = new(MemInfo)
	info, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("system: get mem info failed, err:%v", err)
		return
	}

	memInfo.Total = info.Total
	memInfo.Available = info.Available
	memInfo.Used = info.Used
	memInfo.UsedPercent = info.UsedPercent
	memInfo.Buffers = info.Buffers
	memInfo.Cached = info.Cached

	return memInfo
}
