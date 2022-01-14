// https://github.com/jessevdk/go-flags/issues/223

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jessevdk/go-flags"
)

type args struct {
	Verbose    bool `short:"v" long:"verbose" description:"verbose output"`
	Positional struct {
		Dir string `positional-arg-name:"DIRECTORY"`
	} `positional-args:"true" required:"true"`
}

func main() {
	var args args
	p := flags.NewParser(&args, 0)
	_, err := p.Parse()

	if e, ok := err.(*flags.Error); ok {
		if e.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			fmt.Println(err)
			fmt.Println()
			p.WriteHelp(os.Stdout)
			os.Exit(0)
		}
	}

	log.Println(args)
}

/*

$ go run positional.go
the required argument `DIRECTORY` was not provided
exit status 1

$ go run positional.go -h
Usage:
  positional [OPTIONS] DIRECTORY

Application Options:
  -v, --verbose    verbose output

Help Options:
  -h, --help       Show this help message

$ go run positional.go path/to
2022/01/13 10:51:15 {false {path/to}}

$ go run positional.go path/to trailing
2022/01/13 10:51:24 {false {path/to}}

*/
