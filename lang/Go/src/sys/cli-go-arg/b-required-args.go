package main

import (
	"time"

	"github.com/alexflint/go-arg"
)

var args struct {
	ID      int `arg:"required"`
	Timeout time.Duration
}

func main() {
	arg.MustParse(&args)
}

/*

$ go run b-required-args.go
Usage: b-required-args --id ID [--timeout TIMEOUT]
error: --id is required
exit status 255


$ go run b-required-args.go -h
Usage: b-required-args --id ID [--timeout TIMEOUT]

Options:
  --id ID
  --timeout TIMEOUT
  --help, -h             display this help and exit

*/
