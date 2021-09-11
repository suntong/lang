// Refs
// https://stackoverflow.com/questions/28322997/
// https://pkg.go.dev/flag#Value

package main

import(
	"flag"
	"fmt"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var myFlags arrayFlags

func main() {
	flag.Var(&myFlags, "list1", "Some description for this param.")
	flag.Parse()

	fmt.Printf("\nAfter parsing the flags\n")
	fmt.Println("list1:", myFlags)
}

/*

$ go run CommandLineFlagArray.go --list1 value1 --list1 value2

After parsing the flags
list1: [value1 value2]

*/
