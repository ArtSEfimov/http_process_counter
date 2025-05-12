cmake_minimum_required(VERSION 3.30)
project(c_lib_cpu_load_time C)

set(CMAKE_C_STANDARD 11)

add_library(c_lib_cpu_load_time STATIC cpu_stats.c
        cpu_stats_struct.h)
