////////////////////////////////////////////////////////////////////////////
// Porgram: unicode_demo.go
// Purpose: Go unicode demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: See blow
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"unicode"
)

// https://github.com/go-cc/cc-table/tree/master/text/info/UnicodeCJK#unicode-character-ranges-for-cjk
var runes = []rune{
	'\u2E80',
	'\u2EFF',
	'\u3000',
	'\u303F',
	'\u3200',
	'\u32FF',
	'\u3300',
	'\u33FF',
	'\u3400',
	'\u4DBF',
	'\u4E00',
	'\u9FFF',
	'\uF900',
	'\uFAFF',
	'\uFE30',
	'\uFE4F',
	// https://groups.google.com/d/msg/golang-nuts/u0KCxYsTp3A/6Z-yyo_2AQAJ
	'\U00020000',
	'\U0002A6DF',
	'\U0002F800',
	'\U0002FA1F',
}

func main() {
	for ii, r := range runes {
		unicodeRangeCheck(r)
		if ii%2 != 0 {
			fmt.Println()
		}
	}
}

func unicodeRangeCheck(r rune) {
	unicodeIs(r - 1)
	unicodeIs(r)
	unicodeIs(r + 1)
	fmt.Println()
}

func unicodeIs(r rune) {
	// https://stackoverflow.com/a/16760489/2125837
	fmt.Printf("\\u%X: %v\t", r,
		// https://www.socketloop.com/references/golang-unicode-is-function-example
		unicode.Is(unicode.Han, r))
}
