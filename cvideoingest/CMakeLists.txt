# cmake version 
CMAKE_MINIMUM_REQUIRED(VERSION 3.0.0)

# project
PROJECT(cvideoingest VERSION 0.1.0)



# set config parameters
set (app_VERSION_MAJOR 1)
set (app_VERSION_MINOR 0)

# configure a header file to pass some of the CMake settings to the source code
CONFIGURE_FILE ("${PROJECT_SOURCE_DIR}/appConfig.h.in" "${PROJECT_BINARY_DIR}/appConfig.h")

# files
FILE(GLOB_RECURSE sources      ./src/*.cpp)
FILE(GLOB_RECURSE headers      ./src/*.h ./include/*.h)
FILE(GLOB_RECURSE sources_test ./test/*.cpp)

# include
INCLUDE_DIRECTORIES("${PROJECT_BINARY_DIR}")

INCLUDE_DIRECTORIES("${PROJECT_SOURCE_DIR}/src/foo")
ADD_SUBDIRECTORY(./src/foo)

INCLUDE_DIRECTORIES(./include/)
INCLUDE_DIRECTORIES(/usr/lib/x86_64-linux-gnu/glib-2.0/include/)  
INCLUDE_DIRECTORIES(/usr/include/glib-2.0/)     #gstreamer
INCLUDE_DIRECTORIES(/usr/include/gstreamer-1.0/) 

# library
# SHARED for dynamic library
# STATIC for static library
# ADD_LIBRARY(foo ./src/foo/*.cpp)

# Build settings
SET(EXECUTABLE_OUTPUT_PATH ${CMAKE_SOURCE_DIR}/build/bin/)
SET(LIBRARY_OUTPUT_PATH ${CMAKE_SOURCE_DIR}/build/lib/)


# ADD_SUBDIRECTORY(./src)
# ADD_EXECUTABLE(app ${sources} ${headers})
ADD_EXECUTABLE(app ./src/main)
# ADD_EXECUTABLE(app ${sources} ${headers})
TARGET_LINK_LIBRARIES(app foo)
# TARGET_LINK_LIBRARIES(app foo) 
# INSTALL(TARGETS app DESTINATION "/build/${CMAKE_BUILD_TYPE}/")
# TARGET_COMPILE_OPTIONS(app PUBLIC -std=c++14 -Wall -Wfloat-conversion)

# external libs
# add_subdirectory( ./lib/Logger )
# target_link_libraries( binary logger )

# messages
MESSAGE("CMAKE source directory = " ${CMAKE_SOURCE_DIR})
MESSAGE("Project source directory = " ${PROJECT_SOURCE_DIR})
MESSAGE("Project binary directory = " ${PROJECT_BINARY_DIR})
MESSAGE("CMAKE_BUILD_TYPE is ${CMAKE_BUILD_TYPE}")