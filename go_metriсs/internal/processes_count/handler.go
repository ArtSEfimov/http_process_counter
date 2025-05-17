package processes_count

import (
	"go-proccount/adapters/processes_counter"
	"go-proccount/pkg/response"
	"net/http"
	"time"
)

const layout = "2006-01-02 15:04:05"

type Handler struct{}

func NewHandler(router *http.ServeMux) {
	handler := &Handler{}
	router.HandleFunc("GET /metrics/processes", handler.GetProcessesCount())

}

func (handler *Handler) GetProcessesCount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		pc := processes_counter.ProcessesCounter()

		response.JsonResponse(w, Response{
			Timestamp: time.Now().Format(layout),
			Processes: pc,
		}, http.StatusOK)
	}
}
