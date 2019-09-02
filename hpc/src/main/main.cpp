#include <iostream>
#include <random>
#include <omp.h>
#include "mpi.h"
#include "lec04_07_computePi.h"

int main(int argc, char* argv[]){

    //Lec04_07_computePi
    computePi pi(1000000);
    pi.get();

    //Lec04_09_computePiOmp


}
