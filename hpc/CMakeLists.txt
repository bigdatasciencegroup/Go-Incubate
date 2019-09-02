# CMake version 
CMAKE_MINIMUM_REQUIRED(VERSION 3.0.0 FATAL_ERROR)

# Project name
PROJECT(app VERSION 0.1.0)

# Include CMAKE
INCLUDE(${CMAKE_ROOT}/Modules/FindPkgConfig.cmake)

# Build settings
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

# Library
# SHARED for dynamic library
# STATIC for static library
# ADD_LIBRARY(foo ./src/foo/*.cpp)

# Internal includes
INCLUDE_DIRECTORIES(./include/)

# External includes
# INCLUDE_DIRECTORIES(/usr/include/gstreamer-1.0) 
# INCLUDE_DIRECTORIES(/usr/include/glib-2.0)  
# INCLUDE_DIRECTORIES(/usr/lib/x86_64-linux-gnu/glib-2.0/include)  

# Static compile
# SET(CMAKE_FIND_LIBRARY_SUFFIXES ".a")
# SET(BUILD_SHARED_LIBS OFF)
# SET(CMAKE_EXE_LINKER_FLAGS "-static")

# Executable
# ADD_EXECUTABLE(app ${sources} ${headers})
ADD_EXECUTABLE(app ./src/main/main.cpp)

# Internal libraries
ADD_SUBDIRECTORY(./libs/lec03/)

# Link internal libraries
TARGET_LINK_LIBRARIES(app lec03)

# # Link external libraries
FIND_PACKAGE(OpenMP REQUIRED) # Find the package
TARGET_LINK_LIBRARIES(app ${OpenMP_CXX_LIBRARIES}) # Link against it for C++

# Testing
ENABLE_TESTING()

# Test
# ADD_TEST(appTest app)

# Messages
MESSAGE("CMAKE source directory = " ${CMAKE_SOURCE_DIR})
MESSAGE("Project source directory = " ${PROJECT_SOURCE_DIR})
MESSAGE("Project binary directory = " ${PROJECT_BINARY_DIR})
MESSAGE("CMAKE_BUILD_TYPE is ${CMAKE_BUILD_TYPE}")