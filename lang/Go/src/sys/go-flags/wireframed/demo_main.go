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

	"github.com/jessevdk/go-flags"
)

//////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
type OptsT struct {
	Host    string `short:"H" long:"host" env:"REDO_HOST" description:"host address" default:"localhost"`
	Port    int    `short:"p" long:"port" env:"REDO_PORT" description:"listening port" default:"80"`
	Force   bool   `short:"f" long:"force" env:"REDO_FORCE" description:"force start"`
	Verbflg func() `short:"v" long:"verbose" description:"Verbose mode (Multiple -v options increase the verbosity)"`
	Verbose uint
}

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
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			fmt.Println()
			parser.WriteHelp(os.Stdout)
			os.Exit(1)
		}
	}
	fmt.Println("")
}
