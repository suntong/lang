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
	"io/ioutil"
	"os"
	"strings"
)

import "github.com/andybalholm/crlf"

func main() {

	f, err := crlf.Create("CRLF.txt")
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(f, "Hello\nWorld\n")
	f.Close()

	f = ioutil.NopCloser(crlf.NewWriter(os.Stdout))
	fmt.Fprintln(f, "==========\nGo\n environment\n variable\n demo\n==========")

	// To set a key/value pair, use `os.Setenv`. To get a
	// value for a key, use `os.Getenv`. This will return
	// an empty string if the key isn't present in the
	// environment.
	os.Setenv("FOO", "1")
	fmt.Fprintln(f, "FOO:", os.Getenv("FOO"))
	fmt.Fprintln(f, "BAR:", os.Getenv("BAR"))

	// Use `os.Environ` to list all key/value pairs in the
	// environment. This returns a slice of strings in the
	// form `KEY=value`. You can `strings.Split` them to
	// get the key and value. Here we print all the keys.
	fmt.Fprintln(f)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Fprintln(f, pair[0])
	}
}
