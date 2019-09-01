include "computePi.h"

computePi::computePI(){
    std::random_device seed;
    std::mt19937 gen{seed()}; // seed the generator
    std::uniform_real_distribution<> dist(0.0,1.0);

}