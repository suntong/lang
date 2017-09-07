////////////////////////////////////////////////////////////////////////////
// Porgram: unicode_category.go
// Purpose: Go unicode category demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: https://codexample.org/questions/966388/get-unicode-category-from-rune.c
////////////////////////////////////////////////////////////////////////////

/*

How to get the unicode category (RangeTable) from a rune in Go?

Q: For example, the character a maps to the Ll category. The unicode package specifies all of the categories (http://golang.org/pkg/unicode/#pkg-variables), but I don't see any way to lookup the category from a given rune.

A: The docs for the "unicode" package does not have a method that returns ranges for the rune but it is not very tricky to build one:

*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	catShow('\f')
	catShow('1')
	catShow('A')
	catShow('Ө')
	catShow('世')
	// Differences between IsDigit and IsNumber in unicode in Go
	// https://stackoverflow.com/questions/25540951/
	catShow('Ⅷ')
	catShow('½')
}

/*
: [C Cc]
1: [Nd N]
A: [L Lu]
Ө: [L Lu]
世: [L Lo]
Ⅷ: [N Nl]
½: [No N]

*/

func catShow(r rune) {
	fmt.Printf("%c: %v\n", r, cat(r))
}

func cat(r rune) (names []string) {
	names = make([]string, 0)
	for name, table := range unicode.Categories {
		if unicode.Is(table, r) {
			names = append(names, name)
		}
	}
	return
}
