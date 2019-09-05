#include <iostream>
#include <omp.h>
#include "slide38.h"

slide38::slide38(){
}

void slide38::run(){
    #pragma omp parallel
    {
        int ID = omp_get_thread_num();
        printf("hello(%d)",ID);
        printf("world(%d)\n",ID);
    }    
}