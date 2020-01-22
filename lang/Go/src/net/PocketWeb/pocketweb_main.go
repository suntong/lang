////////////////////////////////////////////////////////////////////////////
// Porgram: PocketWeb
// Purpose: Pocket web server
// Authors: Tong Sun (c) 2020, All rights reserved
////////////////////////////////////////////////////////////////////////////

//go:generate sh -x pocketweb_cli.sh

////////////////////////////////////////////////////////////////////////////
// Program start

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	version = "1.0.00"
	date    = "2020-01-20"
)

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	flag.Usage = Usage
	flag.Parse()

	if Opts.PrintV {
		fmt.Fprintf(os.Stderr, "%s\nVersion %s built on %s\n", progname, version, date)
		os.Exit(0)
	}

	if Opts.Help {
		Usage()
		os.Exit(0)
	}

	pocketweb()

}

////////////////////////////////////////////////////////////////////////////
// Function definitions

//==========================================================================
// Support functions

func debug(input string, threshold int) {
	if !(Opts.Debug >= threshold) {
		return
	}
	print("] ")
	print(input)
	print("\n")
}

func checkError(err error) {
	if err != nil {
		log.Printf("%s: Fatal error - %s", progname, err.Error())
		os.Exit(1)
	}
}
