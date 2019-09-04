# CMake version 
CMAKE_MINIMUM_REQUIRED(VERSION 3.10 FATAL_ERROR)

# Project name
PROJECT(app VERSION 0.1.0)

# Include CMAKE
INCLUDE(${CMAKE_ROOT}/Modules/FindPkgConfig.cmake)

# Build settings
set(CMAKE_CXX_STANDARD 17)
SET(EXECUTABLE_OUTPUT_PATH ${CMAKE_SOURCE_DIR}/bin/)
SET(LIBRARY_OUTPUT_PATH ${CMAKE_SOURCE_DIR}/build/lib/)


# Set config parameters
# set (app_VERSION_MAJOR 1)
# set (app_VERSION_MINOR 0)

# Configure a header file to pass some of the CMake settings to the source code
# CONFIGURE_FILE ("${PROJECT_SOURCE_DIR}/appConfig.h.in" "${PROJECT_BINARY_DIR}/appConfig.h")

# Files
# FILE(GLOB_RECURSE sources      ./src/*.cpp)
# FILE(GLOB_RECURSE headers      ./src/*.h ./include/*.h)
# FILE(GLOB_RECURSE sources_test ./test/*.cpp)

# Find packages
FIND_PACKAGE(OpenMP REQUIRED) # Find OPENMP package
FIND_PACKAGE(MPI REQUIRED) # Find OPENMPI package

# Library
# SHARED for dynamic library
# STATIC for static library
# ADD_LIBRARY(foo ./src/foo/*.cpp)

# Internal includes
INCLUDE_DIRECTORIES(./include/)

# External includes
INCLUDE_DIRECTORIES(${MPI_CXX_INCLUDE_PATH}) 

# Static compile
# SET(CMAKE_FIND_LIBRARY_SUFFIXES ".a")
# SET(BUILD_SHARED_LIBS OFF)
# SET(CMAKE_EXE_LINKER_FLAGS "-static")

# Executable
# ADD_EXECUTABLE(app ${sources} ${headers})
ADD_EXECUTABLE(${PROJECT_NAME} ./src/main/main.cpp)

# Internal libraries
ADD_SUBDIRECTORY(./libs/lec04)

# Link internal libraries
TARGET_LINK_LIBRARIES(${PROJECT_NAME} lec04)

# Link external libraries
# TARGET_LINK_LIBRARIES(${PROJECT_NAME} ${OpenMP_CXX_LIBRARIES}) # Link against it for C++
# TARGET_LINK_LIBRARIES(${PROJECT_NAME} ${MPI_CXX_LIBRARIES}) # Link against it for C++
TARGET_LINK_LIBRARIES(${PROJECT_NAME} OpenMP::OpenMP_CXX)
TARGET_LINK_LIBRARIES(${PROJECT_NAME} MPI::MPI_CXX)

# Testing
# ENABLE_TESTING()

# Test
# ADD_TEST(appTest app)

# Messages
MESSAGE("CMAKE source directory = " ${CMAKE_SOURCE_DIR})
MESSAGE("Project source directory = " ${PROJECT_SOURCE_DIR})
MESSAGE("Project binary directory = " ${PROJECT_BINARY_DIR})
MESSAGE("CMAKE_BUILD_TYPE is ${CMAKE_BUILD_TYPE}")
# MESSAGE("MPI_INCLUDE_PATH = " ${MPI_INCLUDE_PATH})