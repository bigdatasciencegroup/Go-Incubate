# Executable 1
add_executable(tutorial ./tutorial.c)
find_package(MPI REQUIRED)
# include_directories(${MPI_C_INCLUDE_PATH}) 
target_link_libraries(tutorial PRIVATE MPI::MPI_C)
target_link_libraries(tutorial PRIVATE helloworld)
target_include_directories(tutorial PRIVATE ../libs/helloworld/)

# Exercise 1 & 2 (hello)
# Exercise 3 (commrank)
add_executable(hello ./hello.c)
find_package(MPI REQUIRED)
target_link_libraries(hello PRIVATE MPI::MPI_C)
