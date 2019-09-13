# Executable 1
add_executable(tutorial ./tutorial.cpp)
target_link_libraries(tutorial helloworld)
target_link_libraries(tutorial computepi)

# Executable 2
add_executable(mandel ./mandel.c)
find_package(OpenMP REQUIRED) # Find OPENMP package
target_link_libraries(mandel OpenMP::OpenMP_C)