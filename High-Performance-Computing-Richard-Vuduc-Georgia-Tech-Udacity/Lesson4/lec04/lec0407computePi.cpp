#include <iostream>
#include "lec0407computePi.h"

computePi::computePi(int n){
    computePi::nIter = n;
}

double computePi::get(){
    double x, y;
    double circle = 0;
    for (int j = 0; j < computePi::nIter; j++) {
        x = computePi::dist(computePi::gen);
        y = computePi::dist(computePi::gen);
        if (x*x + y*y <= 1.0){
            circle ++;
        }
    }
    double pi = 4*circle/computePi::nIter;
    std::cout << "lec0407computePiSerial. Pi = "<< pi << '\n';
    return pi;
}