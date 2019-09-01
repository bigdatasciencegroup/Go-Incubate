#ifndef COMPUTEPI_H
#define COMPUTEPI_H

#include <random>

using namespace std;

class computePi
{
        random_device seed;
        std::mt19937 gen{seed()}; // seed the generator
    std::uniform_real_distribution<> dist(0.0,1.0);

        string* ptr;
    public:
        // constructor
        computePi(); 
        foo(const string& x);
        // destructor
        ~foo();
        // copy contructor
        // foo(const foo& x)=delete;
        foo(const foo& x);
        // copy assignment
        // foo& operator=(const foo&)=delete;
        foo& operator=(const foo&);
        // move constructor
        // foo(foo&& x)=delete;
        foo(foo&& x);
        // move assignment
        // foo& operator=(foo&& x)=delete;
        foo& operator=(foo&& x);

        // const function
        const string& content() const;
        void print() const;

        double getPi(int x);

        // non-const function
        foo operator+(const foo& rhs);

};

foo duplicate (const foo& param);

#endif