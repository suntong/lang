////////////////////////////////////////////////////////////////////////////
// Program: redo
// Purpose: global option redo
// Authors: Myself <me@mine.org> (c) 2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
//	"fmt"
//	"os"

//	"github.com/go-easygen/go-flags/clis"
)

// *** Sub-command: install ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The InstallCommand type defines all the configurable options from cli.
type InstallCommand struct {
	Dir    string `short:"d" description:"source code root dir" default:"./"`
	Suffix string `long:"suffix" description:"source file suffix" default:".go,.c,.s"`
}

var installCommand InstallCommand

func init() {
	parser.AddCommand("install",
		"Install the network application",
		"The add command adds a file to the repository. Use -a to add all files",
		&installCommand)
}

func (x *InstallCommand) Execute(args []string) error {
	return x.Exec(args)
}

// Exec implements the business logic of command `install`
// func (x *InstallCommand) Exec(args []string) error {
// 	clis.Setup(fmt.Sprintf("%s::%s", progname, "install"), Opts.Verbose)
// 	fmt.Fprintf(os.Stderr, "Install the network application\n")
// 	// fmt.Fprintf(os.Stderr, "Copyright (C) 2022, Myself <me@mine.org>\n\n")
// 	// clis.Verbose(1, "Doing Install, with %+v, %+v\n", Opts, args)
// 	// fmt.Println(x.Dir, x.Suffix)
// 	// // err := ...
// 	// // clis.WarnOn("Doing Install", err)
// 	// // or,
// 	// // clis.AbortOn("Doing Install", err)
// 	return nil
// }
