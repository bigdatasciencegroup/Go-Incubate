#ifndef COMPUTEPI_H
#define COMPUTEPI_H

class computepi
{
    private:
    public:
        // Constructor
        computepi();
        // Run the code
        double runPiSerial(int numSteps);
        double runPiParallel(int numSteps, int numThreads);
};

#endif