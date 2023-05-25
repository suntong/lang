////////////////////////////////////////////////////////////////////////////
// Porgram: CommandLineFlag
// Purpose: Go command line flags/switches/arguments demo
// Authors: Tong Sun (c) 2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2 -w

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	// main operation modes
	write = flag.Bool("w", false, "write result back instead of stdout\n\t\tDefault: No write back")

	// layout control
	tabWidth = flag.Int("tabwidth", 8, "tab width\n\t\tDefault: Standard")

	batch = flag.Duration("batch", 120*time.Second, "batch interval")

	// debugging
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to this file\n\t\tDefault: no default")

	svar string
)

func usage() {
	// Fprintf allows us to print to a specifed file handle or stream
	fmt.Fprintf(os.Stderr, "\nUsage: %s [flags] file [path ...]\n\n",
		"CommandLineFlag") // os.Args[0]
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	fmt.Printf("Before parsing the flags\n")
	fmt.Printf("T: %d\nW: %s\nC: '%s'\nS: '%s'\n",
		*tabWidth, strconv.FormatBool(*write), *cpuprofile, svar)

	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Usage = usage
	flag.Parse()

	// There is also a mandatory non-flag arguments
	if len(flag.Args()) < 1 {
		usage()
	}

	fmt.Printf("\nAfter parsing the flags\n")
	fmt.Printf("T: %d\nB: %v\nW: %s\nC: '%s'\nS: '%s'\n",
		*tabWidth, *batch, strconv.FormatBool(*write), *cpuprofile, svar)

	fmt.Println()
	for index, element := range flag.Args() {
		fmt.Printf("I: %d C: '%s'\n", index, element)
	}
}

/*

Usage: CommandLineFlag [flags] file [path ...]

  -batch duration
    	batch interval (default 2m0s)
  -cpuprofile string
    	write cpu profile to this file
    			Default: no default
  -svar string
    	a string var (default "bar")
  -tabwidth int
    	tab width
    			Default: Standard (default 8)
  -w	write result back instead of stdout
    			Default: No write back

To Test:

  go run CommandLineFlag.go
  go run CommandLineFlag.go -tabwidth=6 aa bb
  go run CommandLineFlag.go -batch 0.51h a b

*/

/*
Ref:

http://golang.org/pkg/flag/
http://golang.org/src/cmd/gofmt/gofmt.go
https://github.com/lanep/golang-me/blob/master/misc_examples/flag.go
*/
