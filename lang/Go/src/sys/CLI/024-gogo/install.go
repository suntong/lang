package main

import (
	"github.com/go-easygen/cli"
)

var _ = app.Register(&cli.Command{
	Name: "install",
	Desc: "Install golang application",
	Text: "Usage:\n  gogo install [Options] package [package...]",
	Argv: func() interface{} { return new(installT) },
	Fn:   install,

	NumArg:      cli.AtLeast(1),
	CanSubRoute: true,
})

type installT struct {
	cli.Helper
	Dir    string `cli:"dir" usage:"source code root dir" dft:"./"`
	Suffix string `cli:"suffix" usage:"source file suffix" dft:".go,.c,.s"`
	Out    string `cli:"o,out" usage:"output filename"`
}

func install(ctx *cli.Context) error {
	argv := ctx.Argv().(*installT)
	ctx.String("%s: %v", ctx.Path(), jsonIndent(argv))
	ctx.String("Installing: %v\n", ctx.Args())
	return nil
}
