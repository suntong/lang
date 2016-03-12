////////////////////////////////////////////////////////////////////////////
// Porgram: Regexp
// Purpose: Go regexp manipulating demo
// Authors: Tong Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Using the Go Regexp Package
	ugrp_Basic_Matching()
	ugrp_Submatches()
	ugrp_NamedCapturingGroups1()
	ugrp_NamedCapturingGroups2()

	// Golang-Regex-Tutorial
	grt_SimpleMatching()
	grt_CompilePOSIX()
	grt_CharacterClasses()
	grt_Rest()

	//
	replaceSome()
}

////////////////////////////////////////////////////////////////////////////
// Using the Go Regexp Package
// http://blog.kamilkisiel.net/blog/2012/07/05/using-the-go-regexp-package/

func ugrp_Basic_Matching() {
	digitsRegexp := regexp.MustCompile(`\d+`)

	someString := "1000abcd123"

	// Find just the leftmost
	fmt.Println(digitsRegexp.FindString(someString))

	// Find all (-1) the matches
	fmt.Println(digitsRegexp.FindAllString(someString, -1))

	/*

		1000
		[1000 123]

		Notes:

		- use the backticks (`...`) instead of quotes ("...") to avoid having to escape backslashes.
		- using the FindString() method because input is string, not []bytes

	*/
}

func ugrp_Submatches() {
	var digitsRegexp = regexp.MustCompile(`(\d+)(\D+(\d+))`)

	someString := "1000abcd123"
	fmt.Println(digitsRegexp.FindStringSubmatch(someString))

	/*

		[1000abcd123 1000 abcd123 123]

		Notes,

		- These methods return a []string which is indexed by the match group position.
		- The 0th item of the slice corresponds to the entire match.

	*/

}

func ugrp_NamedCapturingGroups1() {
	var myExp = regexp.MustCompile(`(?P<first>\d+)\.(\d+).(?P<second>\d+)`)
	fmt.Printf("%+v\n", myExp.FindStringSubmatch("1234.5678.9"))

	/*

		[1234.5678.9 1234 5678 9]

		Notes,

		- it’s useful to be able to document the purpose of the matching groups using the named capturing groups
		- A named capturing group is created with the (?P<name>re) syntax

	*/
}

/*

	The names of the capturing groups can be retrieved via the SubExpNames()
	method and their index within the slice will match the corresponding index
	of the slice returned by FindStringSubmatch(). Capturing groups without a
	name such as the middle one in the example expression will simply have an
	empty string.

  The following example ignores capturing groups without names but they
  could possibly be returned as a second return value or via special names
  in the map.

*/

// embed regexp.Regexp in a new type so we can extend it
type myRegexp struct {
	*regexp.Regexp
}

// add a new method to our new regular expression type
func (r *myRegexp) FindStringSubmatchMap(s string) map[string]string {
	captures := make(map[string]string)

	match := r.FindStringSubmatch(s)
	if match == nil {
		return captures
	}

	for i, name := range r.SubexpNames() {
		// Ignore the whole regexp match and unnamed groups
		if i == 0 || name == "" {
			continue
		}

		captures[name] = match[i]

	}
	return captures
}

// an example regular expression
var myExp = myRegexp{regexp.MustCompile(`(?P<first>\d+)\.(\d+).(?P<second>\d+)`)}

func ugrp_NamedCapturingGroups2() {
	fmt.Printf("%+v\n\n", myExp.FindStringSubmatchMap("1234.5678.9"))
	// map[first:1234 second:9]
}

////////////////////////////////////////////////////////////////////////////
// Golang-Regex-Tutorial
// https://github.com/StefanSchroeder/Golang-Regex-Tutorial

func grt_SimpleMatching() {
	if regexp.MustCompile(`Hello`).MatchString("Hello Regular Expression.") == true {
		fmt.Printf("Match\n") // Will print 'Match' again
	} else {
		fmt.Printf("No match\n")
	}
}

func grt_CompilePOSIX() {

	// the POSIX engine will prefer the leftmost-longest match. It will not
	// return after finding the first match, but will check that the found
	// match is indeed the longest one.

	s := "ABCDEEEEE"
	rr := regexp.MustCompile(`ABCDE{2}|ABCDE{4}`)
	rp := regexp.MustCompilePOSIX(`ABCDE{2}|ABCDE{4}`)
	fmt.Println(rr.FindAllString(s, 2)) // <- first acceptable match
	fmt.Println(rp.FindAllString(s, 2)) // POSIX wants the longer match

}

func grt_CharacterClasses() {
	{
		r, _ := regexp.Compile(`H\wllo`)
		// Will print 'true'.
		fmt.Printf("%v\n", r.MatchString("Hello Regular Expression."))
	}

	{
		r, _ := regexp.Compile(`\d`)
		// Will print 'true':
		fmt.Printf("%v\n", r.MatchString("Seven times seven is 49."))
		// Will print 'false':
		fmt.Printf("%v\n", r.MatchString("Seven times seven is forty-nine."))
	}

	{
		r, _ := regexp.Compile(`\s`)
		// Will print 'true':
		fmt.Printf("%v\n", r.MatchString("/home/bill/My Documents"))
	}

}

func grt_Rest() {
	fmt.Println("Rest of the simple examples ignored.\nhttps://regex-golang.appspot.com/assets/html/index.html explains them all\n")
}

////////////////////////////////////////////////////////////////////////////
// Replace a number of matches of a Regexp
// https://groups.google.com/forum/#!topic/golang-nuts/imllm0To8_E

/*

Q: The package regexp only have functions to replace all matches of the
Regexp. But, how do it if you only want replace a number of them?

Given the number of replacements, when it is -1, then it will return all replacements.

A: do an actual replacement rather than append/slice magic,
using one variable is just fine http://play.golang.org/p/M5zegNp226

DisposaBoy

*/

func replaceSome() {
	src := []byte("eo eo eo eo")
	search := regexp.MustCompile("e")
	repl := []byte("AEI")

	i := 0
	src1 := search.ReplaceAllFunc(src, func(s []byte) []byte {
		if i < 2 {
			i += 1
			return repl
		}
		return s
	})

	fmt.Println(string(src1))

	i = -1
	src2 := search.ReplaceAllFunc(src, func(s []byte) []byte {
		if i != 0 {
			i -= 1
			return repl
		}
		return s
	})

	fmt.Println(string(src2))
}
