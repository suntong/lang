////////////////////////////////////////////////////////////////////////////
// Porgram: RegexpFunc.go
// Purpose: demo Go manipulating regexp with func
// Authors: Tong Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	ExReplaceAllFunc()

	// Using the Go Regexp Package
	ugrp_NamedCapturingGroups2()

	replaceSome()

	// cg: captured group
	cgFixed()
	cgGeneric()

	// rs: replace submatch
	rsMethod1()
	rsMethod2()

	ExExpand()
}

////////////////////////////////////////////////////////////////////////////
// Go Regexp
// https://ruk.si/notes/go/regex

func ExReplaceAllFunc() {

	pchRegex, err := regexp.Compile("p([a-z]+)ch")
	if err != nil {
		panic(err)
	}

	in := []byte("a peach")
	out := pchRegex.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out)) // => a PEACH
}

////////////////////////////////////////////////////////////////////////////
// Using the Go Regexp Package
// http://blog.kamilkisiel.net/blog/2012/07/05/using-the-go-regexp-package/

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
	// AEIo AEIo eo eo

	i = -1
	src2 := search.ReplaceAllFunc(src, func(s []byte) []byte {
		if i != 0 {
			i -= 1
			return repl
		}
		return s
	})

	fmt.Println(string(src2))
	// AEIo AEIo AEIo AEIo
}

////////////////////////////////////////////////////////////////////////////
// cg: Access a capturing group from regexp.ReplaceAllFunc
// http://stackoverflow.com/questions/28000832/

/*

Q: How can I access a capture group from inside ReplaceAllFunc()?

package main

    import (
        "fmt"
        "regexp"
    )

    func main() {
        body := []byte("Visit this page: [PageName]")
        search := regexp.MustCompile("\\[([a-zA-Z]+)\\]")

        body = search.ReplaceAllFunc(body, func(s []byte) []byte {
            // How can I access the capture group here?
        })

        fmt.Println(string(body))
    }

The goal is to replace [PageName] with <a href="/view/PageName">PageName</a>.

A: I agree that having access to capture group while inside of your function would be ideal, but impossible.


*/

func cgFixed() {
	body := []byte("Visit this page: [PageName] [OtherPageName]")
	search := regexp.MustCompile("\\[[a-zA-Z]+\\]")
	body = search.ReplaceAllFunc(body, func(s []byte) []byte {
		m := string(s[1 : len(s)-1])
		return []byte("\n\t<a href=\"/view/" + m + "\">" + m + "</a>")
	})
	fmt.Println(string(body))
	/*
	   Visit this page:
	           <a href="/view/PageName">PageName</a>
	           <a href="/view/OtherPageName">OtherPageName</a>
	*/

}

func cgGeneric() {
	body := []byte("Visit this page: [PageName] [OtherPageName][XYZ]     [XY]")
	search := regexp.MustCompile("(?:\\[)([a-zA-Z]+)(?:\\])")

	body = ReplaceAllSubmatchFunc(search, body, func(s []byte) []byte {
		m := string(s)
		return []byte("\n\t<a href=\"/view/" + m + "\">" + m + "</a>")
	})

	fmt.Println(string(body))
	/*
	   Visit this page:
	           <a href="/view/PageName">PageName</a>
	           <a href="/view/OtherPageName">OtherPageName</a>
	           <a href="/view/XYZ">XYZ</a>
	           <a href="/view/XY">XY</a>
	*/
}

func ReplaceAllSubmatchFunc(re *regexp.Regexp, b []byte, f func(s []byte) []byte) []byte {
	idxs := re.FindAllSubmatchIndex(b, -1)
	if len(idxs) == 0 {
		return b
	}
	l := len(idxs)
	ret := append([]byte{}, b[:idxs[0][0]]...)
	for i, pair := range idxs {
		// replace internal submatch with result of user supplied function
		ret = append(ret, f(b[pair[2]:pair[3]])...)
		if i+1 < l {
			ret = append(ret, b[pair[1]:idxs[i+1][0]]...)
		}
	}
	ret = append(ret, b[idxs[len(idxs)-1][1]:]...)
	return ret
}

////////////////////////////////////////////////////////////////////////////
// rs: replace submatch
// Replace a regular expression submatch using a function
// http://stackoverflow.com/questions/17065465/

/*

For

    input := `bla bla b:foo="hop" blablabla b:bar="hu?"`

I want to replace the parts between quotes in b:foo="hop" or b:bar="hu?" using a function.

The problem is that the callback receives the whole match and not the submatch.

How can I replace the submatch ?

*/

func rsMethod1() {
	input := `bla bla b:foo="hop" blablabla b:bar="hu?"`
	r := regexp.MustCompile(`(\bb:\w+=")([^"]+)`)
	fmt.Println(r.ReplaceAllStringFunc(input, func(m string) string {
		parts := r.FindStringSubmatch(m)
		return parts[1] + complexFunc(parts[2])
	}))
	// bla bla b:foo="dbvalue(hop)" blablabla b:bar="dbvalue(hu?)"
}

func rsMethod2() {
	input := `bla bla b:foo="hop" blablabla b:bar="hu?"`
	r := regexp.MustCompile(`\bb:\w+="([^"]+)"`)
	r2 := regexp.MustCompile(`"([^"]+)"`)
	fmt.Println(r.ReplaceAllStringFunc(input, func(m string) string {
		match := string(r2.Find([]byte(m)))
		return r2.ReplaceAllString(m, complexFunc(match))
	}))
	// bla bla b:foo="dbvalue(hop)" blablabla b:bar="dbvalue(hu?)"
}

func complexFunc(s string) string {
	return "dbvalue(" + s + ")"
}

////////////////////////////////////////////////////////////////////////////
// Expand Example
// http://www.cnblogs.com/golove/p/3270918.html

func ExExpand() {
	pat := `(((abc.)def.)ghi)`
	reg := regexp.MustCompile(pat)

	src := []byte(`abc-def-ghi abc+def+ghi`)
	template := []byte(`$0   $1   $2   $3`)

	// Replace 1st match only
	match := reg.FindSubmatchIndex(src)
	fmt.Printf("%v\n", match)
	// [0 11 0 11 0 8 0 4]
	dst := reg.Expand(nil, template, src, match)
	fmt.Printf("%s\n\n", dst)
	// abc-def-ghi   abc-def-ghi   abc-def-   abc-

	// Replace all
	for _, match := range reg.FindAllSubmatchIndex(src, -1) {
		fmt.Printf("%v\n", match)
		dst := reg.Expand(nil, template, src, match)
		fmt.Printf("%s\n", dst)
	}
	// [0 11 0 11 0 8 0 4]
	// abc-def-ghi   abc-def-ghi   abc-def-   abc-
	// [12 23 12 23 12 20 12 16]
	// abc+def+ghi   abc+def+ghi   abc+def+   abc+
}
