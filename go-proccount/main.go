package main

import (
	"go-proccount/process_counter"
	"go-proccount/system_uptime"
)

func main() {
	process_counter.ProcessCounter()
	system_uptime.GetSystemUptime()
}
