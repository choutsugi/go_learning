package system

import (
	"collect_info/logger"

	"github.com/shirou/gopsutil/disk"
)

func GetDiskInfo() (diskInfo *DiskInfo) {
	diskInfo = new(DiskInfo)
	diskInfo = &DiskInfo{PartitionUsageStat: make(map[string]*DiskUsageStat, 16)}
	parts, err := disk.Partitions(true)
	if err != nil {
		logger.Z.Printf("system: get disk info failed, err:%v", err)
		return
	}
	for _, part := range parts {
		stat, err := disk.Usage(part.Mountpoint)
		if err != nil {
			logger.Z.Printf("system: get disk info failed, err:%v", err)
			continue
		}
		diskInfo.PartitionUsageStat[part.Mountpoint] = &DiskUsageStat{
			Path:              stat.Path,
			Fstype:            stat.Fstype,
			Total:             stat.Total,
			Free:              stat.Free,
			Used:              stat.Used,
			UsedPercent:       stat.UsedPercent,
			InodesTotal:       stat.InodesTotal,
			InodesFree:        stat.InodesFree,
			InodesUsed:        stat.InodesUsed,
			InodesUsedPercent: stat.InodesUsedPercent,
		}
	}
	return
}
