add_library(trialRandom
    ./trial.c
    ./trial.h)

# Find packages
find_package(OpenMP REQUIRED) # Find OPENMP package

# Link external libraries
target_link_libraries(trialRandom OpenMP::OpenMP_C)