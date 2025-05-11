package process_counter

/*
#cgo CFLAGS:  -I${SRCDIR}/c_process_counter
#cgo LDFLAGS: -L${SRCDIR}/c_process_counter -lproccount
#include "proccount.h"
*/
import "C"

import "fmt"

func ProcessCounter() {
	count := C.countProcessesWin()
	fmt.Printf("Processes: %d\n", count)
}
