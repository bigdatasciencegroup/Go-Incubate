#ifndef FOO_H
#define FOO_H

#include <string>
using namespace std;

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
        foo& operator=(const foo&);
        // move constructor
        foo(foo&& x);
        // move assignment
        foo& operator=(foo&& x);

        // const function
        const string& content() const;
        void print() const;

        // non-const function
        foo operator+(const foo& rhs);

};

#endif