// https://github.com/jessevdk/go-flags/issues/223

package main

import (
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

type args struct {
	Verbose    bool `short:"v" long:"verbose" description:"verbose output"`
	Positional struct {
		Dir string `positional-arg-name:"DIRECTORY"`
	} `positional-args:"true" required:"true"`
}

func main() {
	var args args
	_, err := flags.Parse(&args)

	if e, ok := err.(*flags.Error); ok {
		if e.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	if err != nil {
		log.Fatal(err)
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
