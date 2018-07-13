package main

import (
	"github.com/go-easygen/cli"
	clix "github.com/go-easygen/cli/ext"
)

type argT struct {
	Content clix.File `cli:"f,file" usage:"read content from file or stdin"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		ctx.String(argv.Content.String())
		return nil
	})
}
