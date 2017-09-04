////////////////////////////////////////////////////////////////////////////
// Porgram: CamelCaseSplit.go
// Purpose: Go splitting CamelCaseWords demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Seth Bunce https://play.golang.org/p/-FM2wC22A0
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"unicode"
)

func main() {
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
