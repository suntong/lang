package main

import (
	"github.com/go-easygen/cli"
)

type argT struct {
	cli.Helper
	Hello string `cli:"hello" usage:"world"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		ctx.String("native args: %v\n", ctx.NativeArgs())
		ctx.String("args: %v\n", ctx.Args())
		return nil
	})
}
