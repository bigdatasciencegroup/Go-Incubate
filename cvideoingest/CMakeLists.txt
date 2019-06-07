# cmake version 
CMAKE_MINIMUM_REQUIRED(VERSION 3.0.0)

# project
PROJECT(cvideoingest VERSION 0.1.0)

# flags

# files
FILE(GLOB_RECURSE sources      ./src/*.cpp ./src/*.h ./include/*.h)
FILE(GLOB_RECURSE sources_test ./test/*.cpp)

# include
INCLUDE_DIRECTORIES( ./include ./src ./lib)

INCLUDE_DIRECTORIES(/usr/local/lib)  
INCLUDE_DIRECTORIES(/usr/include/glib-2.0)     #gstreamer
INCLUDE_DIRECTORIES(/usr/include/gstreamer-0.10) 


find_package(GStreamer)
if(GStreamer_FOUND)
    message(STATUS "FOUND!")
endif()

# library
# SHARED for dynamic library
# STATIC for static library
# ADD_LIBRARY(foo ./src/foo/*.cpp)

# Build settings
SET(EXECUTABLE_OUTPUT_PATH build/bin)
SET(LIBRARY_OUTPUT_PATH build/lib)


ADD_EXECUTABLE(app ${sources})
# TARGET_LINK_LIBRARIES(app foo) 
# INSTALL(TARGETS app DESTINATION "/build/${CMAKE_BUILD_TYPE}/")
# TARGET_COMPILE_OPTIONS(app PUBLIC -std=c++14 -Wall -Wfloat-conversion)

# external libs
# add_subdirectory( ./lib/Logger )
# target_link_libraries( binary logger )

# messages
MESSAGE("CMAKE source directory = " ${CMAKE_SOURCE_DIR})
MESSAGE("Project source directory = " ${PROJECT_SOURCE_DIR})
MESSAGE("CMAKE_BUILD_TYPE is ${CMAKE_BUILD_TYPE}")