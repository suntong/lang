////////////////////////////////////////////////////////////////////////////
// Porgram: Regexp
// Purpose: Go regexp manipulating demo
// Authors: Tong Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	Summary()
	RegexNotes()
	FromDoc()
	reDemystified()

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

var p = fmt.Println

func Summary() {
	r := regexp.MustCompile(`(?i)^http`)
	p(r.MatchString("http://site/"))  // => true
	p(r.MatchString("https://site/")) // => true
	p(r.MatchString("HTTPS://site/")) // => true

	// -- "(?i)" Case insensitive string match/replace
	fmt.Println(regexp.MustCompile(`(?i)html|uml`).
		ReplaceAllLiteralString("html HTML Html aa uml bb Uml", "XML"))
	// XML XML XML aa XML bb XML

	// These methods return a []string which is indexed by the match group position.
	// - The 0th item of the slice corresponds to the entire match.
	fmt.Println(regexp.MustCompile(`(\d+)(\D+(\d+))`).FindStringSubmatch("1000abcd123"))
	// [1000abcd123 1000 abcd123 123]

	re := regexp.MustCompile(`(?i)t(his|h[eo]se)`)
	sourceStr := "This and these are for THOSE people"
	fmt.Println(re.FindStringSubmatch("Nothing match"))
	// []
	fmt.Println(re.FindStringSubmatch(sourceStr))
	// [This his]
	// Use FindAllStringSubmatch to get $0, $1, $2 etc sub matches!
	fmt.Println(regexp.MustCompile(`(?i)(th)(is|[eo]se)`).
		FindAllStringSubmatch(sourceStr, -1))
	// [[This Th is] [these th ese] [THOSE TH OSE]]
	fmt.Println(re.ReplaceAllString(sourceStr, "<b>${0}</b>"))
	// <b>This</b> and <b>these</b> are for <b>THOSE</b> people
}

func RegexNotes() {
	// https://ruk.si/notes/go/regex
	// Basic regular expression match does not require compiling.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	p(match) // => true

	// Normally you need to compile the regular expression.
	pchRegex, err := regexp.Compile("p([a-z]+)ch")
	if err != nil {
		panic(err)
	}

	p(pchRegex.MatchString("peach"))   // => true
	p(pchRegex.Match([]byte("peach"))) // => true

	p(pchRegex.FindString("peach punch"))        // => peach
	p(pchRegex.FindAllString("peach punch", -1)) // => []string{peach, punch}
	p(pchRegex.FindAllString("peach punch", 1))  // => []string{peach}

	p(pchRegex.FindStringIndex("peach punch")) // => []int{0, 5}
	p(pchRegex.FindStringIndex("plum fump"))   // => []int{}

	p(pchRegex.FindStringSubmatch("peach punch"))      // => []string{peach, ea}
	p(pchRegex.FindStringSubmatchIndex("peach punch")) // => []int{0, 5, 1, 3}

	p(pchRegex.FindAllStringSubmatchIndex("peach punch pinch", -1))
	// => [][]int{[0 5 1 3], [6 11 7 9], [12 17 13 15]}

	p(pchRegex.ReplaceAllString("a peach", "<fruit>")) // => a <fruit>

	in := []byte("a peach")
	out := pchRegex.ReplaceAllFunc(in, bytes.ToUpper)
	p(string(out)) // => a PEACH

	// MustCompile automatically panics if it fails.
	mustPchRegex := regexp.MustCompile("p([a-z]+)ch")
	p(mustPchRegex) // => p([a-z]+)ch
}

func FromDoc() {
	// -- ReplaceAllString
	// https://golang.org/pkg/regexp/#Regexp.ReplaceAll
	{
		re := regexp.MustCompile("a(x*)b")
		fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))
		fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
		fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))   // ! NOT working !!!
		fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W")) // This works!
	}
	// -T-T-
	// --xx-
	// ---
	// -W-xxW-

	// -- ReplaceAllLiteralString
	// https://golang.org/pkg/regexp/#Regexp.ReplaceAllLiteralString
	// The replacement is substituted directly, without using $ Expand
	{
		re := regexp.MustCompile("a(x*)b")
		fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "T"))
		fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "$1"))
		fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "${1}"))
	}
	// -T-T-
	// -$1-$1-
	// -${1}-${1}-

	// -- "(?i)" Case insensitive string match/replace
	fmt.Println(regexp.MustCompile(`(?i)html|uml`).
		ReplaceAllLiteralString("html HTML Html aa uml bb Uml", "XML"))
	// XML XML XML aa XML bb XML
	fmt.Println(regexp.MustCompile(`(?i)(this|th[eo]se)`).
		ReplaceAllString("This and these are for THOSE people", "<b>${1}</b>"))
	// <b>This</b> and <b>these</b> are for <b>THOSE</b> people

	// -- Just to check if a RE can be found
	{
		re := regexp.MustCompile("ab?")
		fmt.Println(re.FindStringIndex("foo") == nil)
		// true
	}
	fmt.Println()
}

////////////////////////////////////////////////////////////////////////////
// red: Regular Expressions demystified
// https://appliedgo.net/regexp/

var (
	exps = []string{"b.*tter", "b(i|u)tter", `batter (\w+)`}

	text = `Betty Botter bought some butter 
But she said the butter?s bitter 
If I put it in my batter, it will make my batter bitter 
But a bit of better butter will make my batter better 
So ?twas better Betty Botter bought a bit of better butter`
)

func reDemystified() {
	for _, e := range exps {
		re := regexp.MustCompile(e)
		fmt.Println(e + ":")
		fmt.Println("1. FindString: ", re.FindString(text))
		fmt.Println("2. FindStringIndex: ", re.FindStringIndex(text))
		fmt.Println("3. FindStringSubmatch: ", re.FindStringSubmatch(text))
		fmt.Printf("4. FindAllString: %v\n", prettyMatches(re.FindAllString(text, -1)))
		fmt.Printf("5. FindAllStringIndex: %v\n", re.FindAllStringIndex(text, -1))
		fmt.Printf("6. FindAllStringSubmatch: %v\n\n", prettySubmatches(re.FindAllStringSubmatch(text, -1)))
	}
}

/*

b.*tter:
1. FindString:  bought some butter
2. FindStringIndex:  [13 31]
3. FindStringSubmatch:  [bought some butter]
4. FindAllString: [bought some butter|butter?s bitter|batter, it will make my batter bitter|bit of better butter will make my batter better|better Betty Botter bought a bit of better butter]
5. FindAllStringIndex: [[13 31] [50 65] [85 122] [130 177] [188 237]]
6. FindAllStringSubmatch: [
    [bought some butter]
    [butter?s bitter]
    [batter, it will make my batter bitter]
    [bit of better butter will make my batter better]
    [better Betty Botter bought a bit of better butter]
]

b(i|u)tter:
1. FindString:  butter
2. FindStringIndex:  [25 31]
3. FindStringSubmatch:  [butter u]
4. FindAllString: [butter|butter|bitter|bitter|butter|butter]
5. FindAllStringIndex: [[25 31] [50 56] [59 65] [116 122] [144 150] [231 237]]
6. FindAllStringSubmatch: [
    [butter|u]
    [butter|u]
    [bitter|i]
    [bitter|i]
    [butter|u]
    [butter|u]
]

batter (\w+):
1. FindString:  batter bitter
2. FindStringIndex:  [109 122]
3. FindStringSubmatch:  [batter bitter bitter]
4. FindAllString: [batter bitter|batter better]
5. FindAllStringIndex: [[109 122] [164 177]]
6. FindAllStringSubmatch: [
    [batter bitter|bitter]
    [batter better|better]
]

*/

func prettySubmatches(m [][]string) string {
	s := "[\n"
	for _, e := range m {
		s += "    " + prettyMatches(e) + "\n"
	}
	s += "]"
	return s
}

func prettyMatches(m []string) string {
	s := "["
	for i, e := range m {
		s += e
		if i < len(m)-1 {
			s += "|"
		}
	}
	s += "]"
	return s
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
