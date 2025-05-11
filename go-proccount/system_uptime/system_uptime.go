package system_uptime

/*
#cgo CFLAGS: -I${SRCDIR}/c_system_uptime
#include "uptime.h"
#include "uptime.c"
*/
import "C"

import (
	"fmt"
	"time"
)

func GetSystemUptime() {
	rowSeconds := uint64(C.getSystemUptimeSec())

	duration := time.Duration(rowSeconds) * time.Second

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	fmt.Printf("System Uptime: %02d:%02d:%02d\n", hours, minutes, seconds)

}
