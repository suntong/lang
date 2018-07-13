package main

import (
	"fmt"
	"os"

	"github.com/go-easygen/cli"
)

type argT struct {
	cli.Helper
	Hello string `cli:"hello" usage:"world"`
}

var app = &cli.Command{
	Name:     "app",
	NeedArgs: true,
	Argv:     func() interface{} { return new(argT) },
	Fn: func(ctx *cli.Context) error {
		return nil
	},
}

func main() {
	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
