#pragma once

#include <windows.h>

typedef struct {
    ULONGLONG userTime;
    ULONGLONG kernelTime;
    ULONGLONG idleTime;
    char success;
} CPUStats;

CPUStats GetRawCPUStats(void);
