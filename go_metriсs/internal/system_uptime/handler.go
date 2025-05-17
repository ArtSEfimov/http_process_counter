package system_uptime

import (
	"go-proccount/adapters/system_uptime"
	"go-proccount/pkg/response"
	"net/http"
	"time"
)

const layout = "2006-01-02 15:04:05"

type Handler struct{}

func NewHandler(router *http.ServeMux) {
	handler := &Handler{}
	router.HandleFunc("GET /metrics/uptime", handler.GetUptime())

}

func (handler *Handler) GetUptime() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		su := system_uptime.GetSystemUptime()

		response.JsonResponse(w, Response{
			Timestamp: time.Now().Format(layout),
			Uptime:    su,
		}, http.StatusOK)
	}
}
