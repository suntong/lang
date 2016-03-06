////////////////////////////////////////////////////////////////////////////
// Porgram: Piping_String.go
// Purpose: Go Internal piping with string
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Howard C. Shaw III
//					https://groups.google.com/d/msg/golang-nuts/snoIyANd-8c/V_IC57y4AwAJ
////////////////////////////////////////////////////////////////////////////

/*

Q: I want to architect my go program just like you can build piping in shell,
but inside go with different existing building block instead.

What I want is to build the middle-wares as go functions, and am able to
build middle-wares so that they can easily be chained together. What passes
between the middle-wares are simply strings.

A: The architecture necessary for a character filter, a word filter, a line
filter, or a whole document filter may differ. If you have to have the whole
document as context, then there is little point in the streaming methods,
while they work very well by character or word filters. If you are actually
working on strings... As in, if you were parsing a set of lines from
something, and selecting which filter stack to hand them off to, so that the
point was really to create a composable set of filters you could keep around
and apply to a single string at a time, you might do it like this:

http://play.golang.org/p/1e4xJjACFV

This example works more like the one you posted. E.g.

  Filter := NewFilter().StageOne().StageTwo().StageThree()
  ResultOne := Filter.Apply("To this text")
  ResultTwo := Filter.Apply("To some other text")

Behind the scenes, it is still building a composition of functions, but by
currying that composition into a new function and holding the state of the
stack in a struct, it restores the left-right ordering of the filters.

Note that the filter stages are basically just compile-time freezes of a
call to the currying function; you could also call the function manually
passing in a custom processing function onto the stack.

  NewFilter().StageOne().AddFilter(func(s string)string {return "I'm in your
function replacing your strings with hardcoded text."})

This version needs neither goroutines nor channels.

You could potentially use this same trick with my earlier example, or with a
struct that embeds an IO.Reader, just using the struct to store the stack of
filters, and functions on the struct to add filters in a comfortable
left-right pattern.

Howard

*/

package main

import (
	"fmt"
	"strings"
)

func PassThrough(s string) string {
	return s
}

type MagicGeFilteFish struct {
	FuncStack func(string) string
}

// Make a new magic filter and start adding bits
func NewFilter() *MagicGeFilteFish {
	return &MagicGeFilteFish{FuncStack: PassThrough}
}

// Call this on the returned object to actually process a string
func (m *MagicGeFilteFish) Apply(s string) string {
	return m.FuncStack(s)
}

// Use this to apply arbitrary filters
func (m *MagicGeFilteFish) AddFilter(f func(string) string) *MagicGeFilteFish {
	m.FuncStack = func(a func(string) string, b func(string) string) func(string) string {
		return func(s string) string {
			return a(b(s))
		}
	}(f, m.FuncStack)
	return m
}

func (m *MagicGeFilteFish) ToLower() *MagicGeFilteFish {
	m.AddFilter(strings.ToLower)
	return m
}

func (m *MagicGeFilteFish) ToUpper() *MagicGeFilteFish {
	m.AddFilter(strings.ToUpper)
	return m
}

func (m *MagicGeFilteFish) Replace(old, new string, times int) *MagicGeFilteFish {
	m.AddFilter(func(s string) string {
		return strings.Replace(s, old, new, times)
	})
	return m
}

func main() {
	// Construct pipelines
	UpCase := NewFilter().ToUpper()
	LCase := NewFilter().ToLower()
	Replace := NewFilter().Replace("test", "biscuit", -1)

	// Test pipelines
	fmt.Printf("%s\n", UpCase.Apply("This is a test."))
	fmt.Printf("%s\n", LCase.Apply("This is a test."))
	fmt.Printf("%s\n", Replace.Apply("This is a test."))
	// Note that we can reuse these stacks as many times as we like
	fmt.Printf("%s\n", Replace.Apply("This is also a test. Testificate."))
	// We can also add stages later on - though we cannot remove stages using this style
	Replace.ToUpper()
	fmt.Printf("%s\n", Replace.Apply("This is also a test. Testificate."))
	LCase.Replace("test", "biscuit", -1)
	fmt.Printf("%s\n", LCase.Apply("This is also a test. Testificate."))

	fmt.Printf("Finished.\n")
}

/*

Output :

THIS IS A TEST.
this is a test.
This is a biscuit.
This is also a biscuit. Testificate.
THIS IS ALSO A BISCUIT. TESTIFICATE.
this is also a biscuit. biscuitificate.
Finished.

*/
