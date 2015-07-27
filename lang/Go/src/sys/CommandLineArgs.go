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

	println("")
	for i, v := range os.Args {
		fmt.Printf("Args %d: %s\n", i, v)
	}
}

/*
Ref:

http://golang.org/pkg/os/
http://stackoverflow.com/questions/3356011/whats-gos-equivalent-of-argv0

*/
