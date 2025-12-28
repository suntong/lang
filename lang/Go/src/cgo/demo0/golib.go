package main

/*
#include <stdlib.h>
#include <stdio.h>
*/
import "C"

// No blank lines allowed between the C comment block and import "C"

import (
	"fmt"
	"unsafe"
)

// main is required but ignored in shared libraries
func main() {}

//export HelloFromGo
func HelloFromGo() {
	fmt.Println("Go: Hello from the Go Runtime!")

	// 1. Allocate in C memory
	cstr := C.CString("Hello from C world")
	// 2. Ensure we clean up manually!
	defer C.free(unsafe.Pointer(cstr))

	// 3. Call C function
	C.puts(cstr)
}
