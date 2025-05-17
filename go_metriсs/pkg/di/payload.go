package di

type CPUInfo struct {
	TotalLoad             float64 `json:"total_load"`
	UserLoad              float64 `json:"user_load"`
	KernelLoad            float64 `json:"kernel_load"`
	AverageLoadSinceStart float64 `json:"average_load_since_start"`
}
