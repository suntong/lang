#include <iostream>

// extern "C" is mandatory for linking Go libraries
extern "C" {
    #include "libgo.h"
}

int main() {
    std::cout << "C++: Starting application..." << std::endl;
    
    // Pass a C-style string literal
    PrintMessage((char*)"Greetings from C++");
    
    std::cout << "C++: Done." << std::endl;
    return 0;
}