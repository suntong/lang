package main

import (
	"flag"
	"fmt"
)

var (
	wordPtr = flag.String("word", "foo", "a string")
	numbPtr = flag.Int("numb", 42, "an int")
	forkPtr = flag.Bool("fork", false, "a bool")

	svar string
)

func init() {
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()
}

func main() {
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

/*

$ go run CommandLineFlagInit.go
word: foo
numb: 42
fork: false
svar: bar
tail: []

$ go run CommandLineFlagInit.go -svar fur -fork -numb 24 a b c
word: foo
numb: 24
fork: true
svar: fur
tail: [a b c]

*/
