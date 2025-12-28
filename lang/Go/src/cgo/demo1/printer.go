package main

import "C"
import "fmt"

func main() {}

//export PrintMessage
func PrintMessage(msg *C.char) {
    // Convert C string to Go string (Safe Copy)
    goStr := C.GoString(msg)
    fmt.Printf("Go received: %s\n", goStr)
}