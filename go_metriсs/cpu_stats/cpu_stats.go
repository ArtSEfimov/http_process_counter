package cpu_stats

/*
#cgo CFLAGS: -I${SRCDIR}/c_cpu_stats
#include "cpu_stats.h"
#include "cpu_stats.c"
*/
import "C"
import (
	"fmt"
	"time"
)

func NewCPUStats() CPUStats {
	rawData := C.GetRawCPUStats()
	return CPUStats{
		UserTime:   uint64(rawData.userTime),
		KernelTime: uint64(rawData.kernelTime),
		IdleTime:   uint64(rawData.idleTime),
		Success:    rawData.success != 0,
	}
}

func GetCPUStats(duration time.Duration) error {
	startCPUStats := NewCPUStats()
	if !startCPUStats.Success {
		return BadRequestToCPUStatistic
	}

	time.Sleep(duration)

	currentCPUStats := NewCPUStats()

	if !currentCPUStats.Success {
		return BadRequestToCPUStatistic
	}
	total, user, kernel := calculateUsage(startCPUStats, currentCPUStats)

	fmt.Printf("CPU Usage — total: %.1f%%, user: %.1f%%, kernel: %.1f%%\n", total, user, kernel)
	return nil
}

func calculateUsage(start, current CPUStats) (totalPct, userPct, kernelPct float64) {

	idleDelta := float64(current.IdleTime - start.IdleTime)
	kernelDelta := float64(current.KernelTime - start.KernelTime)
	userDelta := float64(current.UserTime - start.UserTime)

	busy := (kernelDelta + userDelta) - idleDelta
	total := busy + idleDelta

	if total > 0 {
		totalPct = busy / total * 100
	}
	if busy > 0 {
		userPct = userDelta / busy * 100
		kernelPct = (kernelDelta - idleDelta) / busy * 100
	}
	return
}
