package system_uptime

type Response struct {
	Timestamp string `json:"timestamp"`
	Uptime    string `json:"system_uptime"`
}
