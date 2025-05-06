package main

/*
#cgo CFLAGS:  -I./proccount
#cgo LDFLAGS: -L./proccount -lproccount
#include "proccount.h"
*/
import "C"
import "fmt"

func main() {
	count := C.countProcessesWin()
	fmt.Printf("Processes: %d\n", count)
}
