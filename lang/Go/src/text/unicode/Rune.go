////////////////////////////////////////////////////////////////////////////
// Porgram: Rune.go
// Purpose: Go Rune manipulation demo
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"regexp/syntax"
	"strconv"
)

func main() {
	runeFromCode()
	runeAndString()
}

////////////////////////////////////////////////////////////////////////////

/*

Q: Given a unicode code string, be it "4e16", or "0x4e16", or "u4e16",
how to turn it into a single char rune?

A:

1. Convert it to a number
2. Use `rune()`

Ivan Kurnosov

*/

func runeFromCode() {
	fmt.Printf("Hello \u4e16\u754c (%c%c): %+q\n", 0x4e16, 0x754c, "世界")
	code10 := "4e16"
	//code11 := "0x4e16"
	//code12 := "u4e16"
	// hex to int
	c1, _ := strconv.ParseInt(code10, 16, 16)
	fmt.Printf("Hello %c\u754c %v %v\n", c1, c1 == '世', "\u4e16" == "世")
}

////////////////////////////////////////////////////////////////////////////
// Convert(cast) string to rune and back to string example
// https://www.socketloop.com/tutorials/golang-convert-cast-string-to-rune-and-back-to-string-example

/*

Tutorial on how how to convert a string to rune type and back to
string. Dealing with rune can be confusing sometimes because single quotes
and double quotes can have different meaning in Golang. This example code
below is pretty straightforward and ... is self explanatory.

*/

func runeAndString() {
	str := []rune("beta")   // use rune slice
	acharacter := rune('a') // use single quote, instead of double quote
	onerune := rune('吃')

	fmt.Printf("%v \n", string(str)) // convert back to string

	fmt.Printf("%v \n", string(acharacter)) // convert back to string

	fmt.Printf("%v \n", string(onerune)) // convert back to string

	// there are times when accessing str is not acceptable because
	// it is a slice. Therefore, you just have to reference the first
	// element

	// for example :

	ok := syntax.IsWordChar(str[0]) // won't work without [0]

	fmt.Printf("%v is a word ? : %v \n", string(str), ok)
}
