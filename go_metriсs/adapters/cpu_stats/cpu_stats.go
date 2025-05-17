package cpu_stats

/*
#cgo CFLAGS: -I${SRCDIR}/c_cpu_stats
#include "cpu_stats.h"
#include "cpu_stats.c"
*/
import "C"
import (
	"go-proccount/pkg/di"
	"go-proccount/pkg/utils"
	"log"
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

func GetCPUStats(duration time.Duration, CPUPayload *di.CPUInfo) error {
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

	CPUPayload.TotalLoad = utils.Round(total, 2)
	CPUPayload.UserLoad = utils.Round(user, 2)
	CPUPayload.KernelLoad = utils.Round(kernel, 2)

	avg, cpuAvgLoadError := getAverageCPULoad()
	if cpuAvgLoadError != nil {
		log.Println(cpuAvgLoadError.Error())
	}

	CPUPayload.AverageLoadSinceStart = utils.Round(avg, 2)

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

func getAverageCPULoad() (float64, error) {
	startCPUStats := NewCPUStats()
	if !startCPUStats.Success {
		return 0, BadRequestToCPUStatistic
	}
	busy := startCPUStats.UserTime + startCPUStats.KernelTime - startCPUStats.IdleTime
	avg := float64(busy) / float64(busy+startCPUStats.IdleTime) * 100

	return avg, nil
}
