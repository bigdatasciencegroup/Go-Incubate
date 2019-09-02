#ifndef LEC03_01_COMPUTEPI_H
#define LEC03_01_COMPUTEPI_H

#include <random>

class computePi
{
    private:
        std::random_device rd; // Used to obtain a seed
        std::mt19937 gen{rd()}; // Standard mersenne_twister_engine seeded with rd()
        std::uniform_real_distribution<double> dist{0,1};
        int total;
    public:
        // Constructor
        computePi(int total);
        // Get value of Pi
        double get();
};

#endif