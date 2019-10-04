#include <stdlib.h>
#include <stdio.h>
#include <math.h>
#include <mpi.h>

int main(int argc,char **argv) {
  // Initialize MPI
  MPI_Comm comm = MPI_COMM_WORLD;
  int nprocs, procno;
  MPI_Init(&argc,&argv);
  MPI_Comm_size(comm,&nprocs);
  MPI_Comm_rank(comm,&procno);

  // Initialize the random number generator
  srand(procno*(double)RAND_MAX/nprocs);

  // Initialize random number array
  int count = 2;
  float myrandom[count];
  float globalrandom[count];
  printf("Process %3d has myrandom array ",procno);
  for(int ii = 0; ii < count; ii++)
  {
    myrandom[ii] = rand()/(double)RAND_MAX;
    printf("- %f ", myrandom[ii]);
  }
  printf("\n");

  /*
   * Exercise part 1:
   * -- compute the sum of the values, everywhere
   * -- scale your number by the sum
   * -- check that the sum of the scaled values is 1
   */
  float sumrandom[count], scaled_random[count], sum_scaled_random[count];
  MPI_Allreduce(myrandom, sumrandom,	count, MPI_FLOAT, MPI_SUM, comm);
  for(int ii = 0; ii < count; ii++)
  {
    scaled_random[ii] = myrandom[ii] / sumrandom[ii];
  }
  MPI_Allreduce(scaled_random, sum_scaled_random,	count, MPI_FLOAT, MPI_SUM, comm);

  /*
   * Correctness test
   */
  int error = nprocs; 
  int errors;
  if ( abs(sum_scaled_random[0]-1.)>1.e-5 ) {
    printf("Suspicious sum %7.5f on process %3d\n",sum_scaled_random[0],procno);
    error = procno;
  }
  MPI_Allreduce(&error,&errors,1,MPI_INT,MPI_MIN,comm);
  if (procno==0) {
    if (errors==nprocs) 
      printf("Part 1 finished; all results correct\n");
    else
      printf("Part 1: first error occurred on rank %d\n",errors);
  }

#if 1
  // Exercise part 2:
  // -- compute the maximum random value on process zero
  MPI_Reduce(myrandom, globalrandom, count, MPI_FLOAT, MPI_MAX, 0, comm);
  if (procno==0){
    printf("Part 2: The maximum value of each column in myrandom array is ");
    for(int ii = 0; ii < count; ii++) {
      printf("- %f ", globalrandom[ii]);
    }
    printf("\n");
  }
#endif

  MPI_Finalize();
  return 0;
}
