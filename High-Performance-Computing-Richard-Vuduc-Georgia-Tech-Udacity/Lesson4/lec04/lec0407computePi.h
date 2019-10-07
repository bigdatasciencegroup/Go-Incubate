#ifndef LEC0407COMPUTEPI_H
#define LEC0407COMPUTEPI_H

#include <random>

class computePi
{
    private:
        std::random_device rd; // Used to obtain a seed
        std::mt19937 gen{rd()}; // Standard mersenne_twister_engine seeded with rd()
        std::uniform_real_distribution<double> dist{0,1};
        int nIter;
    public:
        // Constructor
        computePi(int n);
        // Get value of Pi
        double get();
};

#endif