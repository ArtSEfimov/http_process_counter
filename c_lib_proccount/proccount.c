#include "proccount.h"
#include <windows.h>
#include <psapi.h>

int countProcessesWin(void) {
    DWORD processes[1024], needed;

    if (!EnumProcesses(processes, sizeof(processes), &needed)) {
        return -1;
    }

    const DWORD numProcesses = needed / sizeof(DWORD);
    return (int)numProcesses;
}
