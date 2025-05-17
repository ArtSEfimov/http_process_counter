package all_metrics

import "go-proccount/pkg/di"

type Response struct {
	Timestamp string     `json:"timestamp"`
	Processes uint64     `json:"processes"`
	Uptime    string     `json:"system_uptime"`
	CPU       di.CPUInfo `json:"cpu_metrics"`
}
