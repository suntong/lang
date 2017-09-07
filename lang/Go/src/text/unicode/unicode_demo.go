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

func main() {
	unicodeIsControl()
	unicodeIs()
	unicodeIn()
}

/*

A is a control rune code ? :  false
\f is a control rune code ? :  true
Ө(theta) is in the range of unicode.Greek range table ? :  false
世 is in the range of unicode.Han range table ? :  true
A is a member of tab ? :  false
界 is a member of unicode.Han ? :  true

*/

// https://www.socketloop.com/references/golang-unicode-iscontrol-function-example
// unicode.IsControl() checks if the input rune is a ASCII control character.
// See https://en.wikipedia.org/wiki/Control_character
// Reference: http://golang.org/pkg/unicode/#IsControl

func unicodeIsControl() {
	control := unicode.IsControl('A')
	fmt.Println("A is a control rune code ? : ", control)

	control = unicode.IsControl('\f')
	fmt.Println("\\f is a control rune code ? : ", control)

	// see https://en.wikipedia.org/wiki/Control_character
	// for the list of control characters
}

// https://www.socketloop.com/references/golang-unicode-is-function-example
func unicodeIs() {
	memberShip := unicode.Is(unicode.Greek, 'Ө')
	fmt.Println("Ө(theta) is in the range of unicode.Greek range table ? : ", memberShip)

	memberShip = unicode.Is(unicode.Han, '世')
	fmt.Println("世 is in the range of unicode.Han range table ? : ", memberShip)
}

// https://www.socketloop.com/references/golang-unicode-in-function-example
// unicode.In() checks if a given rune is from within ranges of RangeTable.
// http://golang.org/pkg/unicode/#RangeTable
// http://golang.org/pkg/unicode/#In

func unicodeIn() {
	var tab unicode.RangeTable
	memberShip := unicode.In('A', &tab)
	fmt.Println("A is a member of tab ? : ", memberShip)

	memberShip = unicode.In('界', unicode.Han)
	fmt.Println("界 is a member of unicode.Han ? : ", memberShip)
}

/*

func () {
}


*/
