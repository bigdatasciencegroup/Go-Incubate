#include "foo.h"
#include <iostream> 
#include <string>

using namespace std;

// constructor
foo::foo(){
    ptr = new string("Empty");
}

foo::foo(const string& x){
    ptr = new string(x);
}

// destructor
foo::~foo(){
    delete ptr;
}

// copy contructor
foo::foo(const foo& x){
    ptr = new string(x.content());
}

// copy assignment
foo& foo::operator= (const foo& x) {
    delete ptr;                     // delete currently pointed string
    ptr = new string(x.content());  // allocate space for new string, and copy
    return *this;
}

// move constructor
foo::foo(foo&& x) : ptr(x.ptr) {
    x.ptr=nullptr;
}

// move assignment
foo& foo::operator=(foo&& x) {
    delete ptr; 
    ptr = x.ptr;
    x.ptr=nullptr;
    return *this;
}

// const function
const string& foo::content() const{
    return *ptr;
}

// non-const function
foo foo::operator+(const foo& rhs){
    return foo(content()+rhs.content());
}

foo duplicate (const foo& x){
    return foo("duplicated");
}