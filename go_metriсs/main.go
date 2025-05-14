package main

import (
	"fmt"
	"go-proccount/config"
	"go-proccount/internal/metrics"
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
	
	metrics.NewHandler(metricMux, &metrics.HandlerDeps{Config: metricConfig})

	fmt.Printf("Metric server listening at localhost:%s...\n", metricConfig.Port)
	if err := metricServer.ListenAndServe(); err != nil {
		fmt.Printf("metric server err: %v\n", err)
	}

}
