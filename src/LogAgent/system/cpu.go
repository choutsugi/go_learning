package system

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func GetCpuInfo() (cpuInfo *CpuInfo) {
	cpuInfo = new(CpuInfo)
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Printf("system: get cpu info failed, err:%v", err)
		return
	}

	cpuInfo.CpuPercent = percent[0]
	return cpuInfo
}
