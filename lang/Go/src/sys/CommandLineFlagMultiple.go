// Refs https://github.com/suntong/lang/blob/master/lang/Go/src/sys/CommandLineFlagArray.go

package main

import (
	"flag"
	"fmt"
)

// mFlags extend Go flags so that it can be specified multiple times
type mFlags int

func (f *mFlags) String() string {
	return "n.a."
}

func (f *mFlags) Set(value string) error {
	*f += 1
	return nil
}

var multiple mFlags

func main() {
	flag.Var(&multiple, "v", "Verbose, can be specified multiple times.")
	flag.Parse()
	fmt.Printf("\nAfter parsing the flags\n")
	fmt.Println("Verbose:", multiple)
}

/*

Usage Example:

var multiple mFlags

func main() {
 flag.Var(&multiple, "v", "Verbose, can be specified multiple times.")
 flag.Parse()
}

*/
