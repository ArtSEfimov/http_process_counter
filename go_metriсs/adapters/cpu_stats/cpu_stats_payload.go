package cpu_stats

type CPUStats struct {
	UserTime   uint64
	KernelTime uint64
	IdleTime   uint64
	Success    bool
}
