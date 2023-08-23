cmake_minimum_required(VERSION 3.5)

project(TDClient LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

set(EXECUTABLE_OUTPUT_PATH ${PROJECT_SOURCE_DIR}/../output)
set(EXECUTABLE_OUTPUT_PATH ${PROJECT_SOURCE_DIR}/../output)

include_directories(../TDClientLib)

add_executable(${PROJECT_NAME} main.cpp)

target_link_libraries(${PROJECT_NAME} TDClientLib curl)
