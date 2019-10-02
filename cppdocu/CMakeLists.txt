cmake_minimum_required (VERSION 3.10)

project ("cppdocu")

# Add the cmake folder so the FindSphinx module is found
set(CMAKE_MODULE_PATH "${PROJECT_SOURCE_DIR}/cmake" ${CMAKE_MODULE_PATH})
# set(EXECUTABLE_OUTPUT_PATH ${CMAKE_SOURCE_DIR}/bin/)
# set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)

# message("CMAKE source directory = " ${CMAKE_SOURCE_DIR})
# message("CMAKE binary directory = " ${CMAKE_BINARY_DIR})
# message("Project source directory = " ${PROJECT_SOURCE_DIR})
# message("Project binary directory = " ${PROJECT_BINARY_DIR})

#Internal libraries
add_subdirectory ("./libs/tutorial")

#Documentation
add_subdirectory ("./docs")
