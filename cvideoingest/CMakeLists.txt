# cmake version to be used
cmake_minimum_required(VERSION 3.0.0)

# project name
project(cvideoingest VERSION 0.1.0)

# flags

# files
file(GLOB_RECURSE sources      ./src/*.cpp ./src/*.h)
file(GLOB_RECURSE sources_test ./test/*.cpp)

# include
include_directories( ./include ./src ./lib/*/include)

# target
add_executable(app ${sources})
target_compile_options(app PUBLIC -std=c++14 -Wall -Wfloat-conversion)

# external libs
# add_subdirectory( ./lib/Logger )
# target_link_libraries( binary logger )

# messages
message("CMAKE source directory = " ${CMAKE_SOURCE_DIR})
message("Project source directory = " ${PROJECT_SOURCE_DIR})
message("CMAKE_BUILD_TYPE is ${CMAKE_BUILD_TYPE}")