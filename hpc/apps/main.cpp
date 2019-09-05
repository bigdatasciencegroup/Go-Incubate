#include <iostream>
#include "lec0407computePi.h"
#include "slide38.h"

int main(int argc, char* argv[]){

    //Lec0407computePiSer
    // auto start = std::chrono::high_resolution_clock::now(); 
    // computePi piSerial(1000000);
    // piSerial.get();
    // auto stop = std::chrono::high_resolution_clock::now(); 
    // auto duration = std::chrono::duration_cast<std::chrono::microseconds>(stop - start); 
    // std::cout << "Time taken by function: " << duration.count() << " microseconds \n"; 
    
    // Slide 38 - Parallel 'Hello World' program
    slide38 slide38Inst;
    slide38Inst.run(); 

    // Slide 38 - Compute Pi
    slide38 slide38Inst;
    slide38Inst.run(); 

    return 0;

}
