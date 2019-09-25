add_library(computepi 
    ./computepi.cpp
    ./random.c)

# Find packages
find_package(OpenMP REQUIRED) # Find OPENMP package

# Link external libraries
target_link_libraries(computepi OpenMP::OpenMP_CXX)