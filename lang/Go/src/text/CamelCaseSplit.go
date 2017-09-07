////////////////////////////////////////////////////////////////////////////
// Porgram: CamelCaseSplit.go
// Purpose: Go splitting CamelCaseWords demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: See blow
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	algorithm1()
	algorithm2()
}

//==========================================================================
// Credits: Seth Bunce https://play.golang.org/p/-FM2wC22A0
func algorithm1() {
	s := "fooBarBaz GNU PYTHON Standard"
	var parts []string
	start, lastCap := 0, 0
	for end, r := range s {
		if end != 0 && unicode.IsUpper(r) {
			//fmt.Printf("%d, %d\n",lastCap, end)
			if end != lastCap+1 {
				parts = append(parts, s[start:end])
				start = end
			}
			lastCap = end
		}
	}
	if start != len(s) {
		parts = append(parts, s[start:])
	}
	fmt.Printf("%q\n", parts)
}

/*

["foo" "Bar" "Baz " "GNU " "PYTHON " "Standard"]

*/

//==========================================================================
// Michael Jones https://play.golang.org/p/AZ7Ptg4HYP
const testIn = "FooBar3Baz GNU PYTHON Standard"
const testOut = "Foo Bar3 Baz GNU PYTHON Standard"

func algorithm2() {
	out := Decamel(testIn)
	if out != testOut {
		fmt.Println("error:", out)
	}
	fmt.Printf("%q ==> %q\n", testIn, out)
}

func Decamel(s string) string {
	var b bytes.Buffer
	splittable := false
	for _, v := range s {
		if splittable && unicode.IsUpper(v) {
			b.WriteByte(' ')
		}
		b.WriteRune(v)
		splittable = unicode.IsLower(v) || unicode.IsNumber(v)
	}
	return b.String()
}

/*

"FooBar3Baz GNU PYTHON Standard" ==> "Foo Bar3 Baz GNU PYTHON Standard"

*/
