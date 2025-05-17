package main

import (
	"fmt"
	"go-proccount/config"
	"go-proccount/internal/all_metrics"
	"go-proccount/internal/cpu_metrics"
	"go-proccount/internal/processes_count"
	"go-proccount/internal/system_uptime"
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

	all_metrics.NewHandler(metricMux, &all_metrics.HandlerDeps{Config: metricConfig})
	cpu_metrics.NewHandler(metricMux, &cpu_metrics.HandlerDeps{Config: metricConfig})
	processes_count.NewHandler(metricMux)
	system_uptime.NewHandler(metricMux)

	fmt.Printf("Metric server listening at localhost:%s...\n", metricConfig.Port)
	if err := metricServer.ListenAndServe(); err != nil {
		fmt.Printf("metric server err: %v\n", err)
	}

}
