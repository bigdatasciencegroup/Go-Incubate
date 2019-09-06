# CMake version cmake@v3.10.x
cmake_minimum_required(VERSION 3.10 FATAL_ERROR)

# Project name and a few useful settings. Other commands can pick up the results
project(hpc 
    VERSION 0.1.0 
    LANGUAGES CXX)

# Must use GNUInstallDirs to install libraries into correct locations on all platforms.
include(GNUInstallDirs)

# Include CMAKE
include(${CMAKE_ROOT}/Modules/FindPkgConfig.cmake)

# Build settings
set(CMAKE_CXX_STANDARD 14)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF) # Let's ensure -std=c++xx instead of -std=g++xx
set(EXECUTABLE_OUTPUT_PATH ${CMAKE_SOURCE_DIR}/bin/)
# set(LIBRARY_OUTPUT_PATH ${CMAKE_SOURCE_DIR}/build/lib/)

# Set config parameters
# set (app_VERSION_MAJOR 1)
# set (app_VERSION_MINOR 0)

# Configure a header file to pass some of the CMake settings to the source code
# CONFIGURE_FILE ("${PROJECT_SOURCE_DIR}/appConfig.h.in" "${PROJECT_BINARY_DIR}/appConfig.h")

# Docs only available if this is the main app
find_package(Doxygen)
# if(Doxygen_FOUND)
#     add_subdirectory(./docs)
# else()
#     message(STATUS "Doxygen not found, not building docs")
# endif()

# Find packages
# find_package(OpenMP REQUIRED) # Find OPENMP package
# find_package(MPI REQUIRED) # Find OPENMPI package

# Internal includes
include_directories(./include/)

# External includes
# include_directories(${MPI_CXX_INCLUDE_PATH}) 

# Executable
# add_executable(${PROJECT_NAME} ./apps/main.cpp)
source_group("Integration" FILES main.cpp)

# Add subdirectories - executable
add_subdirectory(./apps/)

# Add subdirectories - libraries
add_subdirectory(./libs/lec04/)
add_subdirectory(./libs/helloworld/)
add_subdirectory(./libs/computepi/)

# add_subdirectory(./libs/*)
# target_link_libraries(${PROJECT_NAME} slide38)

# Link external libraries
# target_link_libraries(${PROJECT_NAME} OpenMP::OpenMP_CXX)
# target_link_libraries(${PROJECT_NAME} MPI::MPI_CXX)

# Test
include(CTest)
# add_test(NAME appTest COMMAND app)

# Messages
message("CMAKE source directory = " ${CMAKE_SOURCE_DIR})
message("CMAKE binary directory = " ${CMAKE_BINARY_DIR})
message("Project source directory = " ${PROJECT_SOURCE_DIR})
message("Project binary directory = " ${PROJECT_BINARY_DIR})
message("CMAKE_BUILD_TYPE is ${CMAKE_BUILD_TYPE}")
message("MPI_INCLUDE_PATH = " ${MPI_INCLUDE_PATH})