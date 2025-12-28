package main

import "C"
import "fmt"

// main is required but ignored in shared libraries
func main() {}

//export HelloFromGo
func HelloFromGo() {
	fmt.Println("Go: Hello from the Go Runtime!")
}