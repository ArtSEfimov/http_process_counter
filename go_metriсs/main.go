package main

import (
	"fmt"
	"go-proccount/cpu_stats"
	"go-proccount/process_counter"
	"go-proccount/system_uptime"
	"time"
)

func main() {
	process_counter.ProcessCounter()
	system_uptime.GetSystemUptime()
	cpuStatsError := cpu_stats.GetCPUStats(500 * time.Millisecond)
	if cpuStatsError != nil {
		fmt.Println(cpuStatsError.Error())
	}
}
