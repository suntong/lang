#include <stdio.h>
#include "_cgo_export.h" // Auto-generated link to Go functions

// This is the clean C API we expose to the world
void LogString(const char* msg) {
    // We can add pure C logic here (pre-processing, validation)
    // printf("C Wrapper: Passing '%s' to Go...\n", msg);
    
    // Call the Go function (defined in _cgo_export.h)
    // Cast const char* to char* to match Go's generated signature
    InternalGoLogger((char*)msg);
}
