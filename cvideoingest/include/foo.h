#ifndef FOO_H
#define FOO_H

#include <string>

class foo
{
        int num;
        string* ptr;
    public:
        // constructor
        foo(); 
        foo(const string& x);
        // destructor
        ~foo();
        // copy contructor
        foo(const foo& x);
        // copy assignment
        foo& operator= (const foo&);
        // move constructor
        foo(foo&& x) : ptr(x.ptr) {x.ptr=nullptr;}
        // move assignment
        foo& operator= (foo&& x) {
        delete ptr; 
        ptr = x.ptr;
        x.ptr=nullptr;
        return *this;
        }
        // access content:
        const string& content() const;
};

#endif