#include <iostream>
#include "lec04_07_computePi.h"

computePi::computePi(int total){
    total = total;
}

double computePi::get(){
    double x, y;
    double circle = 0;
    for (int n = 0; n < computePi::total; n++) {
        x = computePi::dist(computePi::gen);
        y = computePi::dist(computePi::gen);
        if (x*x + y*y <= 1.0){
            circle ++;
        }
    }
    double pi = 4*circle/computePi::total;
    std::cout << "lec03_01_computePi - Value of Pi "<< pi << '\n';
    return pi;
}