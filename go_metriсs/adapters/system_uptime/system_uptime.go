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

func GetSystemUptime() string {
	rowSeconds := uint64(C.getSystemUptimeSec())

	duration := time.Duration(rowSeconds) * time.Second

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)

}
