/*

Cross Compiling Demo
http://golangcookbook.com/chapters/running/cross-compiling/
https://github.com/golang/go/wiki/WindowsCrossCompiling

$ GOOS=windows GOARCH=386 go build -o hello.exe hello.go

You can now run hello.exe on a Windows machine near you.

Note that the command above will silently rebuild most of standard library, and for this reason will be quite slow. To speed-up the process, you can install all the windows-amd64 standard packages on your system with

GOOS=windows GOARCH=amd64 go install

*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
	// NB, it’s really a compile-time value - not something determined at runtime.
}

/*

$ go run CrossCompile.go
OS: linux
Architecture: amd64

*/
