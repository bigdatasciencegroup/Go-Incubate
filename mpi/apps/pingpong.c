#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "mpi.h"

int main(int argc,char **argv) {

  MPI_Comm comm = MPI_COMM_WORLD;
  int nprocs, procno;
  
  MPI_Init(&argc,&argv);

  MPI_Comm_size(comm,&nprocs);
  MPI_Comm_rank(comm,&procno);

  #define NEXPERIMENTS 10000
  // Exercise:
  // -- set source and target processors two ways:
  //    close together and far apart
  // -- run the experiment both ways.
  int processA, processB;
  processA = 0;
  processB = nprocs-1;
  double t, send[10000],recv[10000]; 
  send[0] = 1.1;

  if (procno==processA) {
    t = MPI_Wtime();
    for (int n=0; n<NEXPERIMENTS; n++) {
      MPI_Send(send,5000,MPI_DOUBLE, processB, 0, comm);
      MPI_Recv(recv,5000,MPI_DOUBLE, processB, MPI_ANY_TAG, comm, MPI_STATUS_IGNORE);
    }
    t = MPI_Wtime()-t; t /= NEXPERIMENTS;
    {
      int nanosec = t*1000000000;
      printf("Time for pingpong: %7.3f (microsec)\n",nanosec*1.e-3);
    }
  } else if (procno==processB) {
    for (int n=0; n<NEXPERIMENTS; n++) {
      MPI_Recv(recv,5000,MPI_DOUBLE, processA, MPI_ANY_TAG, comm,MPI_STATUS_IGNORE);
      MPI_Send(send,5000,MPI_DOUBLE, processA, 0, comm);
    }
  }

  MPI_Finalize();
  return 0;
}
