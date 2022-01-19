////////////////////////////////////////////////////////////////////////////
// Program: redo
// Purpose: global option redo
// Authors: Myself <me@mine.org> (c) 2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

//go:generate sh -v demo_cliGen.sh

import (
	"fmt"
	"os"

	"github.com/go-easygen/go-flags"
)

//////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "redo"
	version  = "0.1.0"
	date     = "2022-01-17"

	// Opts store all the configurable options
	Opts OptsT
)

var parser = flags.NewParser(&Opts, flags.Default)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	Opts.Verbflg = func() {
		Opts.Verbose++
	}

	if _, err := parser.Parse(); err != nil {
		fmt.Println()
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	fmt.Println()
}
