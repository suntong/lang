////////////////////////////////////////////////////////////////////////////
// Porgram: StringPointers.go
// Purpose: Go string pointer demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits:
//          https://dhdersch.github.io/golang/2016/01/23/golang-when-to-use-string-pointers.html
////////////////////////////////////////////////////////////////////////////

package main

func main() {
	/*

	  A string in Go is a value. Thus, a string cannot be nil.

	  x := "I am a string!"
	  x = nil // Won't compile, strings can't be nil in Go

	  However, a pointer to a string (or *string) can be nil.

	*/

	var x *string

	x = nil // Compiles! String pointers in GoLang can be nil
	UseString(x)

	testQuote("test it", false)
	testQuote("test it", true)

}

func UseString(s *string) string {
  // Pointers require you to write more code because
  // you need to check that a *string has a value before dereferencing
	if s == nil {
		temp := "" // *string cannot be initialized
		s = &temp  // in one statement
	}
	return *s // safe to dereference the *string
}

func testQuote(str string, quote bool) {
	sPreE, sSurE := "", ""
	sPreQ, sSurQ := "<b>", "</b>"

	var sPre, sSur *string
	if quote {
		sPre, sSur = &sPreQ, &sSurQ
	} else {
		sPre, sSur = &sPreE, &sSurE
	}

	print(*sPre + str + *sSur + "\n")
}
