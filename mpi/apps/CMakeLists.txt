# Executable 1
add_executable(tutorial ./tutorial.c)
find_package(MPI REQUIRED)
# include_directories(${MPI_C_INCLUDE_PATH}) 
target_link_libraries(tutorial PRIVATE MPI::MPI_C)
target_link_libraries(tutorial PRIVATE helloworld)
target_include_directories(tutorial PRIVATE ../libs/helloworld/)

# Exercise 1 & 2 (hello)
# Exercise 3 & 4 & 5 (commrank)
add_executable(hello ./hello.c)
find_package(MPI REQUIRED)
target_link_libraries(hello PRIVATE MPI::MPI_C)

# Exercise 6 (prime)
add_executable(prime ./prime.c)
find_package(MPI REQUIRED)
target_link_libraries(prime PRIVATE MPI::MPI_C)

# Exercise 8 & 9 (randommax)
add_executable(randommax ./randommax.c)
find_package(MPI REQUIRED)
target_link_libraries(randommax PRIVATE MPI::MPI_C)

# Exercise 10 (jordan)
add_executable(jordan ./jordan.c)
find_package(MPI REQUIRED)
target_link_libraries(jordan PRIVATE MPI::MPI_C)