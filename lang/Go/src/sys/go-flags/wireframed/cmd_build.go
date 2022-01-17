////////////////////////////////////////////////////////////////////////////
// Program: redo
// Purpose: global option redo
// Authors: Myself <me@mine.org> (c) 2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
)

// *** Sub-command: build ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The BuildCommand type defines all the configurable options from cli.
type BuildCommand struct {
	Dir string `long:"dir" description:"source code root dir" default:"./"`
}

var buildCommand BuildCommand

// BuildCommand implements the business logic of command `build`
func (x *BuildCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "Build the network application\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2022, Myself <me@mine.org>\n\n")
	// fmt.Printf("Doing Build, with %#v\n", args)
	// fmt.Println(x.Dir)
	// err := ...
	return nil
}

func init() {
	parser.AddCommand("build",
		"Build the network application",
		"Usage:\n  redo build [Options] Arch(i386|amd64)",
		&buildCommand)
}
