package processes_counter

/*
#cgo CFLAGS:  -I${SRCDIR}/c_processes_counter
#cgo LDFLAGS: -L${SRCDIR}/c_processes_counter -lproccount
#include "proccount.h"
*/
import "C"

func ProcessesCounter() uint64 {
	count := uint64(C.countProcessesWin())
	return count
}
