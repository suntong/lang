// -*- go -*-
////////////////////////////////////////////////////////////////////////////
// Program: fi
// Purpose: File input demo
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-easygen/cli"
	clix "github.com/go-easygen/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// fi

type rootT struct {
	cli.Helper
	Self *rootT       `cli:"c,config"usage:"config file" json:"-" parser:"jsonfile" dft:"fi.json"`
	Name string       `cli:"*n,name"usage:"Name (mandatory)"`
	Tag  string       `cli:"*t,tag"usage:"Tag used for record saving (mandatory)"`
	ID   string       `cli:"id"usage:"ID to use"`
	Fi   *clix.Reader `cli:"i,input"usage:"The source (or stdin if unspecified)"`
}

var root = &cli.Command{
	Name: "fi",
	Desc: "File input demo",
	Text: "File input demo with mandatory options",
	Argv: func() interface{} { t := new(rootT); t.Self = t; return t },
	Fn:   fi,

	NumOption: cli.AtLeast(1),
}

func main() {
	//NOTE: You can set any writer implements io.Writer
	// default writer is os.Stdout
	if err := cli.Root(root).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("")
}

func fi(ctx *cli.Context) error {
	ctx.JSON(ctx.RootArgv())
	ctx.JSON(ctx.Argv())
	fmt.Println()

	argv := ctx.Argv().(*rootT)
	data, err := ioutil.ReadAll(argv.Fi)
	argv.Fi.Close()
	if err != nil {
		return err
	}
	ctx.String("read from file(or stdin): %s\n", string(data))
	ctx.String("filename: %s, isStdin=%v\n", argv.Fi.Name(), argv.Fi.IsStdin())

	return nil
}
