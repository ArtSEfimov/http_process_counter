package all_metrics

import (
	"go-proccount/adapters/cpu_stats"
	"go-proccount/adapters/processes_counter"
	"go-proccount/adapters/system_uptime"
	"go-proccount/config"
	"go-proccount/pkg/di"
	"go-proccount/pkg/response"
	"log"
	"net/http"
	"time"
)

const layout = "2006-01-02 15:04:05"

type HandlerDeps struct {
	Config *config.Config
}

type Handler struct {
	timeDuration time.Duration
}

func NewHandler(router *http.ServeMux, deps *HandlerDeps) {
	handler := &Handler{
		deps.Config.TimeDuration,
	}
	router.HandleFunc("GET /metrics", handler.GetAllMetrics())

}

func (handler *Handler) GetAllMetrics() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		pc := processes_counter.ProcessesCounter()
		su := system_uptime.GetSystemUptime()

		duration := handler.timeDuration
		var CPUPayload di.CPUInfo

		cpuStatsError := cpu_stats.GetCPUStats(duration, &CPUPayload)
		if cpuStatsError != nil {
			log.Println(cpuStatsError.Error())
		}

		response.JsonResponse(w, Response{
			Timestamp: time.Now().Format(layout),
			Processes: pc,
			Uptime:    su,
			CPU:       CPUPayload,
		}, http.StatusOK)
	}
}
