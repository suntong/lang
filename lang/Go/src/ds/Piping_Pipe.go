////////////////////////////////////////////////////////////////////////////
// Porgram: Piping_Pipe.go
// Purpose: Go Internal piping with io.Pipe
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Howard C. Shaw III
//					https://groups.google.com/d/msg/golang-nuts/snoIyANd-8c/fqUZ5FWxAwAJ
////////////////////////////////////////////////////////////////////////////

/*

Q: I want to architect my go program just like you can build piping in shell,
but inside go with different existing building block instead.

What I want is to build the middle-wares as go functions, and am able to
build middle-wares so that they can easily be chained together. What passes
between the middle-wares are simply strings.

A: Here is an example that comes pretty close to your example using io.Reader and io.Pipe;

http://play.golang.org/p/j8kpfk8-l3

The chaining would be like doFurtherMore(doMore(doThat(doThis(Something))))
though, using function composition.

I wrote the string pump just as another example of the style. There is
actually already a StringReader that would serve the same purpose, it just
wouldn't have demonstrated the structure. And in actual usage, you would
probably have started it off with an io.pipe's Reader, so that you could
push items into the Writer to pull them out of the filter pipe.

Howard

*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// This function just serves as a way to get an io.Reader with s in the buffer
// There are other ways to accomplish
func StringPump(s string) io.Reader {
	r, w := io.Pipe()
	go func(writer *io.PipeWriter) {
		io.WriteString(writer, s)
		writer.Close()
	}(w)
	return r
}

// This is an example of an in-place filter using io.Reader assuming string-compatible data
// and no concern about whitespace losses
func Capitalize(r io.Reader) io.Reader {
	nr, w := io.Pipe()
	go func(reader io.Reader, writer *io.PipeWriter) {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			io.WriteString(writer, strings.ToUpper(scanner.Text()))
		}
		writer.Close()
	}(r, w)
	return nr
}

// This is an example of a helper function that automates some of the
// boilerplate from the above, on the assumption that the passed func can
// operate on a per-word basis. go func needs to close the writer it is passed
func WrapStringWordFilter(filter func(string) string) func(io.Reader) io.Reader {
	return func(r io.Reader) io.Reader {
		nr, w := io.Pipe()
		go func(reader io.Reader, writer *io.PipeWriter) {
			scanner := bufio.NewScanner(reader)
			for scanner.Scan() {
				io.WriteString(writer, filter(scanner.Text()))
			}
			writer.Close()
		}(r, w)

		return nr
	}
}

func main() {
	// EXAMPLE 1
	reader := StringPump("Test")

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Printf("Received: %s \n", scanner.Text())
	}

	// EXAMPLE 2
	reader = Capitalize(StringPump("Test"))

	scanner = bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Printf("Received: %s \n", scanner.Text())
	}

	// EXAMPLE 3
	ToLower := WrapStringWordFilter(strings.ToLower)
	reader = ToLower(Capitalize(StringPump("Test THIS bit")))

	scanner = bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Printf("Received: %s \n", scanner.Text())
	}

	fmt.Printf("Finished.\n")
}

/*

Output :

Received: Test
Received: TEST
Received: test this bit
Finished.

*/

/*

NB, FTA, another way to architect this is,

    func chain(f0 func(io.Reader, io.Writer), fs func(io.Reader, io.Writer)...) func(io.Reader, io.Writer) {
        f := f0
        return func(r io.Reader, w io.Writer) {
            for _, fn := range fs {
                pr, pw := io.Pipe()
                go f(r, pw)
                r = pr
                f = fn
            }
            f(r, w)
        }
    }

Then you can do:

    chain(doThis, doThat, doMore, doFurtherMore)(os.Stdin, os.Stdout)

This by the way is very similar to http://labix.org/pipe, which works for
more than just commands; you can implement your steps in Go functions as
well.

Matt Harden
https://groups.google.com/d/msg/golang-nuts/snoIyANd-8c/6DKWnh4tAwAJ

*/
