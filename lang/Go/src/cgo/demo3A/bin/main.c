#include <stdio.h>

// Forward declaration of the function inside libx.so
// In production, you would provide a handwritten 'libx_api.h'
void LogString(const char* msg);

int main() {
    printf("--- Main C App Starting ---\n");

    // Pass standard C strings. No Go types visible here.
    LogString("Server initialized");
    LogString("Connection received from 192.168.1.1");
    LogString("Error: Low memory");

    printf("--- Main C App Finished ---\n");
    return 0;
}
