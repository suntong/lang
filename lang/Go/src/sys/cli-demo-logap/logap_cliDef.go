////////////////////////////////////////////////////////////////////////////
// Program: logap
// Purpose: log gap finder
// Authors: Tong Sun (c) 2020, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	//  	"fmt"
	//  	"os"

	"github.com/mkideal/cli"
	//  	"github.com/mkideal/cli/clis"
	clix "github.com/mkideal/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

//==========================================================================
// logap

type rootT struct {
	cli.Helper
	Reader  clix.Reader `cli:"*i,input" usage:"log file to check"`
	Gap     int         `cli:"*g,Gap" usage:"minimum gap between log entries to report (in minutes)" dft:"$LG_GAP"`
	Keyword string      `cli:"*k,kw" usage:"keyword for the log entries to concentrate on" dft:"$LG_KW"`
	Verbose cli.Counter `cli:"v,verbose" usage:"Verbose mode (Multiple -v options increase the verbosity)\n"`
}

var root = &cli.Command{
	Name: "logap",
	Desc: "log gap finder\nVersion " + version + " built on " + date +
		"\nCopyright (C) 2020, Tong Sun",
	Text: "Tool to find gaps in log file entries that are over the given limit",
	Argv: func() interface{} { return new(rootT) },
	Fn:   logap,

	NumOption: cli.AtLeast(1),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
//  type OptsT struct {
//  	Reader	clix.Reader
//  	Gap	int
//  	Keyword	string
//  	Verbose	cli.Counter
//  	Verbose int
//  }

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "logap"
//          version   = "0.1.0"
//          date = "2020-07-18"

//  	rootArgv *rootT
//  	// Opts store all the configurable options
//  	Opts OptsT
//  )

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.DenseNormalStyle) // left-right, for up-down, use ManualStyle
//  	//NOTE: You can set any writer implements io.Writer
//  	// default writer is os.Stdout
//  	if err := cli.Root(root,).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  		os.Exit(1)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Dumb root handler

//  func logap(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here
