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

// *** Sub-command: publish ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The PublishCommand type defines all the configurable options from cli.
type PublishCommand struct {
	Dir    string   `short:"d" long:"dir" description:"publish dir" required:"true"`
	Suffix []string `long:"suffix" description:"source file suffix for publish" choice:".go" choice:".c" choice:".h"`
	Out    string   `short:"o" long:"out" description:"output filename"`

	// Example of positional arguments
	Args struct {
		ID   string
		Num  int
		Rest []string
	} `positional-args:"yes" required:"yes"`
}

var publishCommand PublishCommand

func (x *PublishCommand) Execute(args []string) error {
	return x.Exec(args)
}

// Exec implements the business logic of command `publish`
// func (x *PublishCommand) Exec(args []string) error {
// 	fmt.Fprintf(os.Stderr, "Publish the network application\n")
// 	// fmt.Fprintf(os.Stderr, "Copyright (C) 2022, Myself <me@mine.org>\n\n")
// 	// fmt.Printf("Doing Publish, with %+v, %+v\n", Opts, args)
// 	// fmt.Println(x.Dir, x.Suffix, x.Out, x.Args)
// 	// err := ...
// 	return nil
// }

func init() {
	parser.AddCommand("publish",
		"Publish the network application",
		"Publish the built network application to central repo",
		&publishCommand)
}
