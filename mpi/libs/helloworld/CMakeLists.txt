add_library(helloworld 
    ./helloworld.c
    ./helloworld.h)

# MPI
find_package(MPI REQUIRED)
target_link_libraries(helloworld PRIVATE MPI::MPI_C)