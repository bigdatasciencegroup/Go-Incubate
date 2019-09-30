add_library(helloworld 
    ./helloworld.c
    ./helloworld.h)

# Find packages
find_package(MPI REQUIRED)

# Link external libraries
target_link_libraries(helloworld PRIVATE MPI::MPI_C)