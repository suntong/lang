////////////////////////////////////////////////////////////////////////////
// Porgram: CRLF
// Purpose: Go environment variable demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: github.com/andybalholm/crlf
////////////////////////////////////////////////////////////////////////////

// [Environment variables](http://en.wikipedia.org/wiki/Environment_variable)
// are a universal mechanism for [conveying configuration
// information to Unix programs](http://www.12factor.net/config).
// Let's look at how to set, get, and list environment variables.

package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/andybalholm/crlf"
)

func main() {
	test1()
	f, err := crlf.Create("Env.txt")
	if err != nil {
		panic(err)
	}
	test2(f)
	test3(f)
	// test4(f)
}

func test1() {
	f, err := crlf.Create("CRLF.txt")
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(f, "Hello\nWorld\n")
	f.Close()

}

func test2(f io.Writer) {
	// X	f = ioutil.NopCloser(crlf.NewWriter(os.Stdout))
	fmt.Fprintln(f, "==========\nGo\n environment\n variable\n demo\n==========")
}

func test3(f io.Writer) {
	// To set a key/value pair, use `os.Setenv`. To get a
	// value for a key, use `os.Getenv`. This will return
	// an empty string if the key isn't present in the
	// environment.
	os.Setenv("FOO", "1")
	fmt.Fprintln(f, "FOO:", os.Getenv("FOO"))
	fmt.Fprintln(f, "BAR:", os.Getenv("BAR"))

}

func test4(f io.Writer) {
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

/*

GOOS=windows GOARCH=amd64 go build -v CRLF.go
go run CRLF.go

# clean up
rm CRLF.exe CRLF.txt Env.txt

*/
