#include "foo.h"
#include <iostream> 

void foo::print() const{
    std::cout << "Hello World from foo!\n" << num << std::endl;
}