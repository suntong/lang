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
	Port int    `cli:"p,port"usage:"listening port" dft:"80"`
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

// func main() {
// 	cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
// 	//NOTE: You can set any writer implements io.Writer
// 	// default writer is os.Stdout
// 	if err := cli.Root(root,
// 		cli.Tree(buildCmd),
// 		cli.Tree(installCmd),
// 		cli.Tree(publishCmd)).Run(os.Args[1:]); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}
// 	fmt.Println("")
// }

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
// 	rootArgv := ctx.RootArgv().(*rootT)
// 	argv := ctx.Argv().(*buildT)
// 	fmt.Printf("[build]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
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
// 	rootArgv := ctx.RootArgv().(*rootT)
// 	argv := ctx.Argv().(*installT)
// 	fmt.Printf("[install]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
// 	return nil
// }

////////////////////////////////////////////////////////////////////////////
// publish

type publishT struct {
	Dir    string `cli:"*d,dir" usage:"publish dir"`
	Suffix string `cli:"suffix" usage:"source file suffix" dft:".go,.c,.s"`
	Out    string `cli:"o,out" usage:"output filename"`
	List   bool   `cli:"l,list" usage:"list all sub commands"`
}

var publishCmd = &cli.Command{
	Name: "publish",
	Desc: "Publish the network application",
	Argv: func() interface{} { return new(publishT) },
	Fn:   publish,

	NumOption: cli.AtLeast(1),
}

// func publish(ctx *cli.Context) error {
// 	rootArgv := ctx.RootArgv().(*rootT)
// 	argv := ctx.Argv().(*publishT)
// 	fmt.Printf("[publish]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
// 	return nil
// }
