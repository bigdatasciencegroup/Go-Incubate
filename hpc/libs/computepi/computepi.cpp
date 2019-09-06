#include <iostream>
#include <omp.h>
#include "computepi.h"

computepi::computepi(){
}

double computepi::runPiSerial(int numSteps){
    double step;
    double x, pi, sum = 0.0;
    step = 1.0/double(numSteps);

    double startTime = omp_get_wtime();
    for (int ii = 0; ii < numSteps; ii++){
        x = (ii + 0.5)*step;
        sum = sum + 4.0/(1.0+x*x);
    }
    double totalTime = omp_get_wtime() - startTime;
    pi = step * sum;
    
    printf("computePi.runPiSerial(). Pi = %f. Time = %f\n", pi, totalTime);
    return pi;
}

double computepi::runPiParallel(int numSteps, int nThreadsInput){
    double x;
    double step = 1.0/double(numSteps);
    double sum[nThreadsInput];
    double sumTotal;
    int nThreadsMaster;

    omp_set_num_threads(nThreadsInput);
    double startTime = omp_get_wtime();
    #pragma omp parallel
    {
        int id = omp_get_thread_num(); //Thread ID
        int nThreadsSlave = omp_get_num_threads();
        if (id == 0){
            nThreadsMaster = nThreadsSlave;
        }
        int ii;
        for (ii = id, sum[id]=0; ii < numSteps; ii=ii+nThreadsSlave){
            x = (ii + 0.5)*step;
            sum[id] = sum[id] + 4.0/(1.0+x*x);
        }
    }
    for (int ii = 0, pi=0.0; ii<nThreadsMaster; ii++){
        sumTotal = sumTotal + sum[ii];
    }
    double totalTime = omp_get_wtime() - startTime;
    double pi = sumTotal*step;
    
    printf("computePi.runPiParallel(). Pi = %f. Time = %f\n", pi, totalTime);
    return pi;
}