#include <stdlib.h>
#include <stdio.h>
#include <mpi.h>

int main(int argc,char **argv) {

    int ierr;
    ierr = MPI_Init(&argc,&argv);

    // Get the number of processes
    int world_size;
    MPI_Comm_size(MPI_COMM_WORLD, &world_size);

    // Get the rank of the process
    int world_rank;
    MPI_Comm_rank(MPI_COMM_WORLD, &world_rank);

    // Get the name of the processor
    char processor_name[MPI_MAX_PROCESSOR_NAME];
    int name_len;
    MPI_Get_processor_name(processor_name, &name_len);

    // Print off a hello world message
    printf("Hello world from processor %s, rank %d out of %d processors\n", processor_name, world_rank, world_size);
    
    // Let only processs zero print out the total number
    if (world_rank == 0){
        printf("There are %d processes. Printed by rank %d\n", world_size, world_rank);
    }

    //
    FILE *fp;
    char text[] = "", 
    fp = fopen (["../data/data.txt" world_rank], "w");
    fclose(fp);

    ierr = MPI_Finalize();
    
    return 0;
}