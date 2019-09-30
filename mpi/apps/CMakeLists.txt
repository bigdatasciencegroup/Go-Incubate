# Executable 1
add_executable(tutorial ./tutorial.c)
find_package(MPI REQUIRED)
target_link_libraries(tutorial PRIVATE MPI::MPI_C)
target_link_libraries(tutorial PRIVATE helloworld)
target_include_directories(tutorial PRIVATE ../libs/helloworld/)
