// -*- go -*-
////////////////////////////////////////////////////////////////////////////
// Program: redo
// Purpose: global option redo
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// redo

type rootT struct {
	cli.Helper
	Self *rootT `cli:"c,config"usage:"config file" json:"-" parser:"jsonfile" dft:"redo.json"`
	Host string `cli:"H,host"usage:"host address" dft:"$HOST"`
	Port int    `cli:"p,port"usage:"listening port" dft:80`
}

var root = &cli.Command{
	Name:   "redo",
	Desc:   "global option redo",
	Text:   "  redo global option via automatic code-gen",
	Global: true,
	Argv:   func() interface{} { t := new(rootT); t.Self = t; return t },
	Fn:     redo,

	NumArg: cli.ExactN(1),
}

// func redo(ctx *cli.Context) error {
// 	ctx.JSON(ctx.RootArgv())
// 	ctx.JSON(ctx.Argv())
// 	fmt.Println()

// 	return nil
// }

////////////////////////////////////////////////////////////////////////////
// build

type buildT struct {
	Dir    string `cli:"dir" usage:"source code root dir" dft:"./"`
	Suffix string `cli:"suffix" usage:"source file suffix" dft:".go,.c,.s"`
	Out    string `cli:"o,out" usage:"output filename"`
}

var buildCmd = &cli.Command{
	Name: "build",
	Desc: "Build the network application",
	Text: "Usage:\n  redo build [Options] Arch(i386|amd64)",
	Argv: func() interface{} { return new(buildT) },
	Fn:   build,

	NumArg:      cli.ExactN(1),
	CanSubRoute: true,
}

// func build(ctx *cli.Context) error {
// 	ctx.JSON(ctx.RootArgv())
// 	ctx.JSON(ctx.Argv())
// 	fmt.Println()

// 	return nil
// }

////////////////////////////////////////////////////////////////////////////
// install

type installT struct {
	Dir    string `cli:"dir" usage:"source code root dir" dft:"./"`
	Suffix string `cli:"suffix" usage:"source file suffix" dft:".go,.c,.s"`
	Out    string `cli:"o,out" usage:"output filename"`
}

var installCmd = &cli.Command{
	Name: "install",
	Desc: "Install the network application",
	Text: "Usage:\n  redo install [Options] package [package...]",
	Argv: func() interface{} { return new(installT) },
	Fn:   install,

	NumArg:      cli.AtLeast(1),
	CanSubRoute: true,
}

// func install(ctx *cli.Context) error {
// 	ctx.JSON(ctx.RootArgv())
// 	ctx.JSON(ctx.Argv())
// 	fmt.Println()

// 	return nil
// }

////////////////////////////////////////////////////////////////////////////
// publish

type publishT struct {
	Dir    string `cli:"dir" usage:"source code root dir" dft:"./"`
	Suffix string `cli:"suffix" usage:"source file suffix" dft:".go,.c,.s"`
	Out    string `cli:"o,out" usage:"output filename"`
	List   bool   `cli:"l,list" usage:"list all sub commands"`
}

var publishCmd = &cli.Command{
	Name: "publish",
	Desc: "Publish the network application",
	Argv: func() interface{} { return new(publishT) },
	Fn:   publish,
}

// func publish(ctx *cli.Context) error {
// 	ctx.JSON(ctx.RootArgv())
// 	ctx.JSON(ctx.Argv())
// 	fmt.Println()

// 	return nil
// }
