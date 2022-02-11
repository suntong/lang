////////////////////////////////////////////////////////////////////////////
// Porgram: embedding.go
// Purpose: Go:embed and command line flags demo
// Authors: Tong Sun (c) 2021-22, All rights reserved
////////////////////////////////////////////////////////////////////////////


package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
)

//go:embed cpuArch.txt
// see https://pkg.go.dev/embed
var cpuArch string

func usage() {
	// Fprintf allows us to print to a specifed file handle or stream
	fmt.Fprintf(os.Stderr, "\nUsage: %s [flags] file [path ...]\n\n",
		"CommandLineFlag") // os.Args[0]
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	fmt.Printf("Before parsing the flags\n")
	fmt.Printf("CPU: '%s'\n", cpuArch)

	flag.StringVar(&cpuArch, "cpu", "AMD64", "CPU Arch")

	flag.Usage = usage
	flag.Parse()

	// There is also a mandatory non-flag arguments
	if len(flag.Args()) < 1 {
		usage()
	}

	fmt.Printf("\nAfter parsing the flags\n")
	fmt.Printf("CPU: '%s'\n", cpuArch)

}

/*

$ go run embedding.go .
Before parsing the flags
CPU: 'Intel32'

After parsing the flags
CPU: 'AMD64'


$ go run embedding.go -cpu AMD32 .
Before parsing the flags
CPU: 'Intel32'

After parsing the flags
CPU: 'AMD32'

*/

