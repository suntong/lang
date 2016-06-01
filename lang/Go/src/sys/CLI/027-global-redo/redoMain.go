// -*- go -*-
////////////////////////////////////////////////////////////////////////////
// Program: redo
// Purpose: global option redo
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

func main() {
	cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
	//NOTE: You can set any writer implements io.Writer
	// default writer is os.Stdout
	if err := cli.Root(root, cli.Tree(buildCmd)).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("")
}

func redo(ctx *cli.Context) error {
	ctx.JSON(ctx.RootArgv())
	ctx.JSON(ctx.Argv())
	fmt.Println()

	return nil
}
