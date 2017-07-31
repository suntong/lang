package main

import (
	"io/ioutil"

	"github.com/mkideal/cli"
	clix "github.com/mkideal/cli/ext"
)

type argT struct {
	Reader *clix.Reader `cli:"r,reader" usage:"read from file, stdin or any io.Reader"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		data, err := ioutil.ReadAll(argv.Reader)
		argv.Reader.Close()
		if err != nil {
			return err
		}
		ctx.String("reade from reader: '%s'\n", string(data))
		return nil
	})
}
