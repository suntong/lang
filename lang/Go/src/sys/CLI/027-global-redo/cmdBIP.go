////////////////////////////////////////////////////////////////////////////
// Program: cmd-bip.go
// Purpose: sample implementation of commands build, install & publish
//          in reality, each should in a separated file
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/mkideal/cli"
)

func build(ctx *cli.Context) error {
	ctx.String("%s:", ctx.Path())
	ctx.JSON(ctx.Argv())
	ctx.String("[build]: %v\n", ctx.Args())

	return nil
}

func install(ctx *cli.Context) error {
	ctx.String("%s:", ctx.Path())
	ctx.JSON(ctx.Argv())
	ctx.String("[install]: %v\n", ctx.Args())

	return nil
}

func publish(ctx *cli.Context) error {
	ctx.String("%s:", ctx.Path())
	ctx.JSON(ctx.Argv())
	ctx.String("[publish]: %v\n", ctx.Args())

	return nil
}
