package main

import (
	"fmt"
	"go-proccount/cpu_stats"
	"go-proccount/process_counter"
	"go-proccount/system_uptime"
)

func main() {
	process_counter.ProcessCounter()
	system_uptime.GetSystemUptime()
	cpuStatsError := cpu_stats.GetCPUStats()
	if cpuStatsError != nil {
		fmt.Println(cpuStatsError.Error())
	}
}
