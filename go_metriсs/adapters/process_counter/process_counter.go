package process_counter

/*
#cgo CFLAGS:  -I${SRCDIR}/c_process_counter
#cgo LDFLAGS: -L${SRCDIR}/c_process_counter -lproccount
#include "proccount.h"
*/
import "C"

func ProcessCounter() uint64 {
	count := uint64(C.countProcessesWin())
	return count
}
