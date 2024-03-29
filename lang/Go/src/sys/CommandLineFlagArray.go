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
	fmt.Println("list1:", len(myFlags), myFlags)

	for ii, val := range myFlags {
		fmt.Println(ii, val)
	}

}

/*

$ go run CommandLineFlagArray.go --list1 value1 --list1 value2

After parsing the flags
list1: 2 [value1 value2]

$ go run CommandLineFlagArray.go --list1 value1 --list1 value2 --list1 value3

After parsing the flags
list1: 3 [value1 value2 value3]
0 value1
1 value2
2 value3


*/
