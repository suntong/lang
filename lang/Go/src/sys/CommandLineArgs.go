////////////////////////////////////////////////////////////////////////////
// Porgram: CommandLineArgs
// Purpose: Go command line arguments demo
// Authors: Tong Sun (c) 2013-2014, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2 -w

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	println("I am", os.Args[0])

	baseName := filepath.Base(os.Args[0])
	println("The base name is", baseName)

	// The length of array a can be discovered using the built-in function len
	println("Argument # is", len(os.Args))

	// the first command line arguments
	if len(os.Args) > 1 {
		println("The first command line argument: ", os.Args[1])
	}

	// there is only one command line argument
	if len(os.Args) == 2 {
		println("The dup first command line argument as 2nd: ", os.Args[1])
		os.Args = append(os.Args, os.Args[1])
	}

	println("")
	for i, v := range os.Args {
		fmt.Printf("Args %d: %s\n", i, v)
	}
}

/*

Ref:

http://golang.org/pkg/os/
http://stackoverflow.com/questions/3356011/whats-gos-equivalent-of-argv0

Run:

$ go run CommandLineArgs.go
I am /tmp/go-build303379025/b001/exe/CommandLineArgs
The base name is CommandLineArgs
Argument # is 1

Args 0: /tmp/go-build303379025/b001/exe/CommandLineArgs

$ go run CommandLineArgs.go abc
I am /tmp/go-build619429067/b001/exe/CommandLineArgs
The base name is CommandLineArgs
Argument # is 2
The first command line argument:  abc
The dup first command line argument as 2nd:  abc

Args 0: /tmp/go-build619429067/b001/exe/CommandLineArgs
Args 1: abc
Args 2: abc

$ go run CommandLineArgs.go abc def
I am /tmp/go-build136951689/b001/exe/CommandLineArgs
The base name is CommandLineArgs
Argument # is 3
The first command line argument:  abc

Args 0: /tmp/go-build136951689/b001/exe/CommandLineArgs
Args 1: abc
Args 2: def

*/
