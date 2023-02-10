package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
)

var args struct {
	Foo string
	Bar bool
}

func main() {
	arg.MustParse(&args)
	fmt.Println(args.Foo, args.Bar)
}

/*

$ go run a-foobar.go
 false

$ go run a-foobar.go --foo=hello --bar
hello true

$ go run a-foobar.go --foo='Hello, it is' --bar
Hello, it is true

$ go run a-foobar.go -h
Usage: a-foobar [--foo FOO] [--bar]

Options:
  --foo FOO
  --bar
  --help, -h             display this help and exit

*/
