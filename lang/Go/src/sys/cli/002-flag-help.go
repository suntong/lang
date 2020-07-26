package main

import (
	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
	Port int  `cli:"p,port" usage:"port # (should > 1024)"`
	X    bool `cli:"x" usage:"boolean type"`
	Y    bool `cli:"y" usage:"boolean type, too"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		if argv.Port <= 1024 {
			ctx.String("Invalide port value --\n\n")
			ctx.WriteUsage()
			return nil
		}
		ctx.String("port=%d, x=%v, y=%v\n", argv.Port, argv.X, argv.Y)
		return nil
	})
}

/*

$ go run 002-flag-help.go
Invalide port value --

Options:

  -h, --help   display help information
  -p, --port   port # (should > 1024)
  -x           boolean type
  -y           boolean type, too

$ go run 002-flag-help.go -p 2222
port=2222, x=false, y=false

$ go run 002-flag-help.go -p 2222 -x
port=2222, x=true, y=false

*/
