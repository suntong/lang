////////////////////////////////////////////////////////////////////////////
// Porgram: OsEnv
// Purpose: Go environment variable demo
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: https://gobyexample.com/environment-variables
//
////////////////////////////////////////////////////////////////////////////

// [Environment variables](http://en.wikipedia.org/wiki/Environment_variable)
// are a universal mechanism for [conveying configuration
// information to Unix programs](http://www.12factor.net/config).
// Let's look at how to set, get, and list environment variables.

package main

import (
	"fmt"
	"os"
	"strings"
)

var date = "1970-01-01"

func main() {

	fmt.Println("==========\nGo\n environment\n variable\n demo\n==========")

	// To set a key/value pair, use `os.Setenv`. To get a
	// value for a key, use `os.Getenv`. This will return
	// an empty string if the key isn't present in the
	// environment.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"), os.Getenv("FOO") == "1")
	fmt.Println("BAR:", os.Getenv("BAR"), os.Getenv("BAR") == "")

	bar, exists := os.LookupEnv("BAR")
	if exists {
		// Print the value of the environment variable
		fmt.Println(bar)
	} else {
		fmt.Println("BAR not defined in system environment variable")
	}

	// Use `os.Environ` to list all key/value pairs in the
	// environment. This returns a slice of strings in the
	// form `KEY=value`. You can `strings.Split` them to
	// get the key and value. Here we print all the keys.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0][0] == 'H' {
			fmt.Println(pair[0])
		}
	}
	fmt.Println()
	fmt.Println(date)
}
