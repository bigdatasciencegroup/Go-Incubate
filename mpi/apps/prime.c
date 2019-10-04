#include <stdlib.h>
#include <stdio.h>
#include <mpi.h>

int main(int argc,char **argv) {
  MPI_Comm comm = MPI_COMM_WORLD;
  int nprocs, procno;
  int bignum = 2000000111, maxfactor = 45200;
  
  MPI_Init(&argc,&argv);
  MPI_Comm_size(comm,&nprocs);
  MPI_Comm_rank(comm,&procno);

  // Exercise:
  // -- Parallelize the do loop so that each processor
  //    tries different candidate numbers.
  // -- If a processors finds a factor, print it to the screen.

  // 1. Set loop bounds
  // Use cyclic distribution in the for loop for task assignment
  
  // 2. Fill in loop header
  for ( int myfactor=procno; myfactor <= maxfactor; myfactor = myfactor+nprocs) {
    if (myfactor == 0 || myfactor == 1){
      continue;
    }
    if (bignum%myfactor==0) {
      printf("Processor %d found factor %d\n",procno,myfactor);
    }
  }
  
  MPI_Finalize();
  return 0;
}
