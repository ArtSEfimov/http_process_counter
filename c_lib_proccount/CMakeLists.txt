cmake_minimum_required(VERSION 3.30)
project(proccount C)

set(CMAKE_C_STANDARD 11)

add_library(proccount STATIC proccount.c)

target_include_directories(proccount PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})