#include "cpu_stats.h"

static ULONGLONG toULL(const FILETIME *ft) {
    return (ULONGLONG) ft->dwHighDateTime << 32 | ft->dwLowDateTime;
}


CPUStats GetRawCPUStats(void) {
    FILETIME idle, kernel, user;
    CPUStats s = {0, 0, 0, 0};

    if (GetSystemTimes(&idle, &kernel, &user)) {
        s.idleTime = toULL(&idle);
        s.kernelTime = toULL(&kernel);
        s.userTime = toULL(&user);
        s.success = 1;
    }
    return s;
}
