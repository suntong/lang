package main

import (
	"fmt"
)

type AddCommand struct {
	All bool `short:"a" long:"all" description:"Add all files"`

	// Example of positional arguments
	Args struct {
		ID   string
		Num  int
		Rest []string
	} `positional-args:"yes" required:"yes"`
}

var addCommand AddCommand

func (x *AddCommand) Execute(args []string) error {
	fmt.Printf("Adding (all=%v): %v %#v\n", x.All, x.Args, args)
	return nil
}



func init() {
	parser.AddCommand("add",
		"Add a file",
		"The add command adds a file to the repository. Use -a to add all files.",
		&addCommand)
}
