package metrics

import "go-proccount/pkg/di"

type GetAllMetricsResponse struct {
	Timestamp   string     `json:"timestamp"`
	Processes   uint64     `json:"processes"`
	Uptime      string     `json:"uptime"`
	CPU         di.CPUInfo `json:"cpu"`
	AverageLoad float64    `json:"average_load"`
}
