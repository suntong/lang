////////////////////////////////////////////////////////////////////////////
// Program: logap
// Purpose: log gap finder
// Authors: Tong Sun (c) 2020, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

//go:generate sh -v logap_cliGen.sh

import (
	"fmt"
	"os"
	"time"

	"github.com/mkideal/cli"
	"github.com/mkideal/cli/clis"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "logap"
	version  = "0.1.0"
	date     = "2020-07-01"

	rootArgv *rootT
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	cli.SetUsageStyle(cli.DenseNormalStyle) // left-right, for up-down, use ManualStyle
	//NOTE: You can set any writer implements io.Writer
	// default writer is os.Stdout
	if err := cli.Root(root).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("")
}

//==========================================================================
// Dumb root handler

func logap(ctx *cli.Context) error {
	ctx.JSON(ctx.RootArgv())
	fmt.Println()

	rootArgv = ctx.RootArgv().(*rootT)
	clis.Setup(progname, rootArgv.Verbose.Value())
	allowedGap := time.Second * 60 * time.Duration(rootArgv.Gap)
	clis.Verbose(2, "%v\n", allowedGap)

	return nil
}

/*

touch /tmp/abc

$ cli-demo-logap -g 3 -k kw -i /tmp/abc -vv
{"Reader":{},"Gap":3,"Keyword":"kw","Verbose":{}}
[logap] 3m0s

*/
