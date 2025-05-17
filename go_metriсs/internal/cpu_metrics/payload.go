package cpu_metrics

import "go-proccount/pkg/di"

type Response struct {
	Timestamp string     `json:"timestamp"`
	CPU       di.CPUInfo `json:"cpu_info"`
}
