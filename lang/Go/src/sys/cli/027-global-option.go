package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/mkideal/cli"
)

func main() {
	if err := cli.Root(root, cli.Tree(sub)).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("")
}

// root command
type rootT struct {
	cli.Helper
	Self *rootT `json:"-" cli:"c,config" usage:"config" parser:"jsoncfg" dft:"027-global-option.json"`
	Host string `cli:"H,host" usage:"host addr" dft:"$HOST"`
	Port int    `cli:"p,port" usage:"listening port"`
}

var root = &cli.Command{
	Name:   "app",
	Desc:   "application",
	Global: true,
	Argv: func() interface{} {
		t := new(rootT)
		t.Self = t
		return t
	},
	Fn: func(ctx *cli.Context) error {
		fmt.Println("root")
		ctx.JSON(ctx.RootArgv())
		ctx.JSON(ctx.Argv())
		fmt.Println()

		// https://play.golang.org/p/GHg_i0vM4O
		re := regexp.MustCompile(ctx.RootArgv().(*rootT).Host)
		fmt.Println(re.FindStringIndex("A regexp \t test string"))

		return nil
	},
}

// sub command
type subT struct {
	World string `cli:"w" usage:"world is a sub flag"`
}

var sub = &cli.Command{
	Name: "sub",
	Desc: "subcommand",
	Argv: func() interface{} { return new(subT) },
	Fn: func(ctx *cli.Context) error {
		ctx.JSON(ctx.RootArgv())
		ctx.JSON(ctx.Argv())
		ctx.String("\n")

		var argv = &subT{}
		var parentArgv = &rootT{}
		if err := ctx.GetArgvList(argv, parentArgv); err != nil {
			return err
		}
		fmt.Println("sub")
		ctx.JSON(parentArgv)
		ctx.JSON(argv)
		fmt.Println()
		return nil
	},
}
