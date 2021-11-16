package system

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

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
		// 单个CPU
		percents, err := cpu.Percent(time.Second, true)
		if err != nil {
			return
		}

		for id, percent := range percents {
			fmt.Printf("cpu id:%d , percent:%v\n", id, percent)
		}

	}
}
