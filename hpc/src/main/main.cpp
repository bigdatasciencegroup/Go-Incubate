// #include "appConfig.h"
#include <iostream>
#include <stdlib.h>
#include <math.h>
#include "foo.h"

int main(int argc, char* argv[]){

    // fprintf(stdout,"%s Version %d.%d\n",
    //         argv[0],
    //         app_VERSION_MAJOR,
    //         app_VERSION_MINOR);

    const char* foo1 = "hello";
    const char* bar1 = foo1;
 
    foo foo2("Exam");
    foo bar = foo("ple");  // move-construction
    foo bar3 = foo2;
    foo2 = foo2 + bar;            // move-assignment
    std::cout << foo2.content() << "\n";
    std::cout << foo1 << "\n";
    std::cout << bar1 << "\n";
    std::cout << bar3.content() << "\n";
    foo foo4 = duplicate(foo2);
    std::cout << foo4.content() << "\n";

}