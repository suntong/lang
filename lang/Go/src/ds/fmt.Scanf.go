package main

import (
	"fmt"
	"io"
)

func main() {
	var d int
	var s string
	for {
		_, err := fmt.Scanln(&d, &s)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		fmt.Printf("(%d) '%s'\n", d, s)
	}
}

/*

// http://golang.org/pkg/fmt/#Scanf

fmt.Scanf deals with *space delimited* items:

Scan scans text read from standard input, storing successive space-separated values into successive arguments.

$ seq 6 | xargs -n 2  | go run fmt.Scanf.go
(1) '2'
(3) '4'
(5) '6'

$ seq 12 | xargs -n 3 | go run fmt.Scanf.go
panic: expected newline

*/
