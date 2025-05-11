#include <windows.h>
#include "uptime.h"

unsigned long long getSystemUptimeSec(void) {
    return GetTickCount64() / 1000;
}
