package main

import (
	"fmt"
	"go-proccount/config"
	"go-proccount/cpu_stats"
	"go-proccount/process_counter"
	"go-proccount/system_uptime"
	"net/http"
)

func main() {

	// get config
	metricConfig := config.NewConfig()

	// create new mux
	metricMux := http.NewServeMux()

	// create new server
	metricServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", metricConfig.Port),
		Handler: metricMux,
	}

	fmt.Printf("Metric server listening at localhost:%s\n", metricConfig.Port)
	if err := metricServer.ListenAndServe(); err != nil {
		fmt.Printf("metric server err: %v\n", err)
	}

	// get metrics
	process_counter.ProcessCounter()
	system_uptime.GetSystemUptime()

	duration := metricConfig.TimeDuration
	cpuStatsError := cpu_stats.GetCPUStats(duration)
	if cpuStatsError != nil {
		fmt.Println(cpuStatsError.Error())
	}
	cpuAvgLoadError := cpu_stats.GetAverageCPULoad()
	if cpuAvgLoadError != nil {
		fmt.Println(cpuAvgLoadError.Error())
	}
}
