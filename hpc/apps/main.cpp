#include <iostream>
#include <random>
#include <omp.h>
#include "mpi.h"
#include "lec0407computePi.h"

int main(int argc, char* argv[]){

    //Lec04_07_computePiSer
    // auto start = std::chrono::high_resolution_clock::now(); 
    // computePi piSerial(1000000);
    // piSerial.get();
    // auto stop = std::chrono::high_resolution_clock::now(); 
    // auto duration = std::chrono::duration_cast<std::chrono::microseconds>(stop - start); 
    // std::cout << "Time taken by function: " << duration.count() << " microseconds \n"; 
  
    #pragma omp parallel num_threads(7)
    {
        int id = omp_get_thread_num();
        int data = id;
        int total = omp_get_num_threads();
        printf("Greetings from process %d out of %d with Data %d\n", id, total, data);
    }
    printf("parallel for ends.\n");

    #pragma omp parallel
    {
        int ID = omp_get_thread_num();
        printf("hello(%d)",ID);
        printf("world(%d)\n",ID);
    }



    return 0;

}
