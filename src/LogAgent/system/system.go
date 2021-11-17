package system

import "time"

func run(interval time.Duration) {
	ticker := time.Tick(interval)
	for _ = range ticker {
		GetCpuInfo()
		GetMemInfo()
	}
}
