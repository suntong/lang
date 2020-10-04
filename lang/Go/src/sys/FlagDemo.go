// These examples demonstrate more intricate uses of the flag package.
// From https://golang.org/pkg/flag/#example_
//  and https://golang.org/doc/go1.5#flag

package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

// Example 1: A single string flag called "specie" with default value "gopher".
var species = flag.String("specie", "gopher", "the specie `type` we are studying")

// Example 2: Two flags sharing a variable, so we can have a shorthand.
// The order of initialization is undefined, so make sure both use the
// same default value. They must be set up with an init function.
var gopherType string

func init() {
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")
}

// Example 3: A user-defined flag type, a slice of durations.
type interval []time.Duration

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (i *interval) String() string {
	return fmt.Sprint(*i)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (i *interval) Set(value string) error {
	// If we wanted to allow the flag to be set multiple times,
	// accumulating values, we would delete this if statement.
	// That would permit usages such as
	//	-deltaT 10s -deltaT 15s
	// and other combinations.
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

// Define a flag to accumulate durations. Because it has a special type,
// we need to use the Var function and therefore create the flag during
// init.

var intervalFlag interval

// go1.5+, the default is listed only when it is not the zero value for the type
var write = flag.Bool("w", false, "write result back instead of stdout")
var Help bool

func init() {
	// Tie the command-line flag to the intervalFlag variable and
	// set a usage message.
	flag.Var(&intervalFlag, "deltaT", "comma-separated list of `intervals` to use between events")
	flag.BoolVar(&Help, "help", false, "show usage help")
}

func main() {
	// All the interesting pieces are with the variables declared above, but
	// to enable the flag package to see the flags defined there, one must
	// execute, typically at the start of main (not init!):
	flag.Parse()

	flag.PrintDefaults()
	fmt.Println(*write, Help)
}

/*

Output:

$ go run FlagDemo.go -help
  -deltaT intervals
        comma-separated list of intervals to use between events (default [])
  -g string
        the variety of gopher (shorthand) (default "pocket")
  -gopher_type string
        the variety of gopher (default "pocket")
  -specie type
        the specie type we are studying (default "gopher")
  -w    write result back instead of stdout
false true

*/
