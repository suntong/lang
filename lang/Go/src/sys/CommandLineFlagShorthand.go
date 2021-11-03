////////////////////////////////////////////////////////////////////////////
// Purpose: Go command line flag shorthand demo
// Authors: Tong Sun (c) 2021, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"flag"
	"fmt"
	"os"
)

// From https://golang.org/src/flag/example_test.go
// Example 2: Two flags sharing a variable, so we can have a shorthand.
// The order of initialization is undefined, so make sure both use the
// same default value. They must be set up with an init function.
var gopherType string

func init() {
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")
}

func usage() {
	// Fprintf allows us to print to a specifed file handle or stream
	fmt.Fprintf(os.Stderr, "\nUsage: %s [flags] file [path ...]\n\n",
		"CommandLineFlag") // os.Args[0]
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {

	flag.Usage = usage
	flag.Parse()

	// able to show the usage text
	if len(flag.Args()) < 1 {
		usage()
	}

	fmt.Printf("Gopher type: '%s'\n", gopherType)

}

/*

$ go run CommandLineFlagShorthand.go

Usage: CommandLineFlag [flags] file [path ...]

  -g string
        the variety of gopher (shorthand) (default "pocket")
  -gopher_type string
        the variety of gopher (default "pocket")


$ go run CommandLineFlagShorthand.go .
Gopher type: 'pocket'

$ go run CommandLineFlagShorthand.go -g foo .
Gopher type: 'foo'

$ go run CommandLineFlagShorthand.go -gopher_type bar .
Gopher type: 'bar'

*/
