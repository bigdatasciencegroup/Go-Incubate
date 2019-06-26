#include "foo.h"
#include <iostream> 
#include <string>

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

const string& foo::content() const{
    return *ptr;
}