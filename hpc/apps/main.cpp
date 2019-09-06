#include <iostream>
#include "lec0407computePi.h"
#include "helloworld.h"
#include "computepi.h"

int main(int argc, char* argv[]){

    //Lec0407computePiSer
    // auto start = std::chrono::high_resolution_clock::now(); 
    // computePi piSerial(1000000);
    // piSerial.get();
    // auto stop = std::chrono::high_resolution_clock::now(); 
    // auto duration = std::chrono::duration_cast<std::chrono::microseconds>(stop - start); 
    // std::cout << "Time taken by function: " << duration.count() << " microseconds \n"; 
    
    // Slide 38 - Parallel 'Hello World' program
    helloWorld helloWorldInst;
    helloWorldInst.run(); 

    // Slide 48 - Compute Pi in Serial
    computepi computePiInst;
    computePiInst.runPiSerial(100000); 

    // Slide 52 - Compute Pi in Parallel
    computePiInst.runPiParallel(100000, 1); 
    computePiInst.runPiParallel(100000, 2); 
    computePiInst.runPiParallel(100000, 3); 
    computePiInst.runPiParallel(100000, 4); 
    

    return 0;

}
