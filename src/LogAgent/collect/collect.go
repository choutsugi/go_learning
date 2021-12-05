package collect

import (
	"LogAgent/influx"
	"LogAgent/system"
	"time"
)

func Run(interval time.Duration) {
	ticker := time.Tick(interval)
	var err error
	for _ = range ticker {
		err = influx.InsertCpuInfo(system.GetCpuInfo())
		if err != nil {
			//
		}
		err = influx.InsertMemInfo(system.GetMemInfo())
		if err != nil {
			//
		}

		err = influx.InsertDiskInfo(system.GetDiskInfo())
		if err != nil {
			//
		}

	}
}
