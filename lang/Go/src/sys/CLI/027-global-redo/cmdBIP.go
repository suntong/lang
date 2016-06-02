////////////////////////////////////////////////////////////////////////////
// Program: cmd-bip.go
// Purpose: sample implementation of commands build, install & publish
//          in reality, each should in a separated file
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"

	"github.com/mkideal/cli"
)

func build(ctx *cli.Context) error {
	rootArgv := ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*buildT)
	jR, _ := json.Marshal(*rootArgv)
	jC, _ := json.Marshal(*argv)
	ctx.String("[build]:\n  %v\n  %v\n  %v\n", string(jR), string(jC), ctx.Args())
	return nil
}

func install(ctx *cli.Context) error {
	rootArgv := ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*buildT)
	jR, _ := json.Marshal(*rootArgv)
	jC, _ := json.Marshal(*argv)
	ctx.String("[install]:\n  %v\n  %v\n  %v\n", string(jR), string(jC), ctx.Args())
	return nil
}

func publish(ctx *cli.Context) error {
	rootArgv := ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*buildT)
	jR, _ := json.Marshal(*rootArgv)
	jC, _ := json.Marshal(*argv)
	ctx.String("[publish]:\n  %v\n  %v\n  %v\n", string(jR), string(jC), ctx.Args())
	return nil
}
