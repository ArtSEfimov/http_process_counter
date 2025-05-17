package processes_count

type Response struct {
	Timestamp string `json:"timestamp"`
	Processes uint64 `json:"processes"`
}
