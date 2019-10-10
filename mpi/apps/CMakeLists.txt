find_package(MPI REQUIRED)

# Exercise 1 & 2 (hello)
# Exercise 3 & 4 & 5 (commrank)
add_executable(hello ./hello.c)
target_link_libraries(hello PRIVATE MPI::MPI_C)

# Exercise 6 (prime)
add_executable(prime ./prime.c)
target_link_libraries(prime PRIVATE MPI::MPI_C)

# Exercise 8 & 9 (randommax)
add_executable(randommax ./randommax.c)
target_link_libraries(randommax PRIVATE MPI::MPI_C)

# Exercise 10 (jordan)
add_executable(jordan ./jordan.c)
target_link_libraries(jordan PRIVATE MPI::MPI_C)

# Exercise 13 (pingpong)
add_executable(pingpong ./pingpong.c)
target_link_libraries(pingpong PRIVATE MPI::MPI_C)

# Exercise 14 (inprod aka sumsquares)
add_executable(sumsquares ./sumsquares.c)
target_link_libraries(sumsquares PRIVATE MPI::MPI_C)