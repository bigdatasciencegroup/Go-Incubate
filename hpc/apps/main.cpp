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
    
    const char* s = getenv("PATH");
    printf("%s\n",s);
    const char* w = getenv("KITE");
    printf("%s\n",w);

    // Slide 38 - Parallel 'Hello World' program
    helloWorld helloWorldInst;
    helloWorldInst.run(); 

    // Slide 48 - Compute Pi in Serial
    computepi computePiInst;
    computePiInst.runPiSerial(1e8); 

    // Slide 52 - Compute Pi in Parallel with false sharing
    computePiInst.runPiParallel(1e8, 1); 
    computePiInst.runPiParallel(1e8, 2); 
    computePiInst.runPiParallel(1e8, 3); 
    computePiInst.runPiParallel(1e8, 4); 

    // Slide 57 - Compute Pi in Parallel with padding
    computePiInst.runPiParallelPad(1e8, 1); 
    computePiInst.runPiParallelPad(1e8, 2); 
    computePiInst.runPiParallelPad(1e8, 3); 
    computePiInst.runPiParallelPad(1e8, 4); 

    // Slide 69 - Compute Pi in Parallel with synchronisation
    computePiInst.runPiParallelSync(1e8, 3); 

    // Slide 88 - Compute Pi with WorkSharing
    computePiInst.runPiWorkSharing(1e8); 


    return 0;

}
