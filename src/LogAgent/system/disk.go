package system

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

func GetDiskInfo() (diskInfo *DiskInfo) {
	diskInfo = &DiskInfo{
		PartitionUsageStat: make(map[string]*UsageStat, 16),
	}
	stats, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("system: get disk info failed, err:%v", err)
		return
	}

	for _, stat := range stats {
		usage, err := disk.Usage(stat.Mountpoint)
		if err != nil {
			fmt.Printf("system: get disk info failed, err:%v", err)
			continue
		}
		diskInfo.PartitionUsageStat[stat.Mountpoint] = usage
	}
	return diskInfo
}
