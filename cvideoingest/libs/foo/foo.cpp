#include "foo.h"
#include <iostream> 

foo::foo(int n){
    num = n;
}

foo::~foo(){
    
}

void foo::print(){
    std::cout << "Hello World from foo!\n" << num << std::endl;
}