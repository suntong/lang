////////////////////////////////////////////////////////////////////////////
// Porgram: iota
// Purpose: Go iota data demo
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// Snail-moving Directions
const (
	dRight = iota
	dDown
	dLeft
	dUp
	dNumbers // total number of directions
)

// Option for output format
type tOptFmt int

// A list of valid formats
const (
	fNormal tOptFmt = iota
	fFlip           // Flip the snail direction and flip the count from up to down
	fMirror         // Mirror the starting porint and count direction
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	a1 int = iota
	b1
	c1
	d1 = 9
)

const (
	a int = iota
	b
	c
	d
)

const (
	_ int = iota
	e
	f
	g = 9
)

////////////////////////////////////////////////////////////
// Global variables definitions

// Function main
func main() {

	fmt.Printf("%d, %d, %d, %d\n", dRight, dDown, fNormal, fFlip)
	fmt.Printf("%v, %v\n", dRight == fNormal, dDown == fFlip)
	// X: fmt.Printf("%v\n", fNormal == Sunday)
	// invalid operation: fNormal == Sunday (mismatched types tOptFmt and Weekday)
	fmt.Println(a1, b1, c1, d1)
	fmt.Println(a, b, c, d, e, f, g)
	// X: fmt.Printf("%v, %v\n", a == fNormal, b == fFlip)
	// invalid operation: a == fNormal (mismatched types int and tOptFmt)

	iota_with_gaps()
}

/*

> > How to use `iota` to define consts that have gap(s) in them?
> > E.g., If my consts are,
> > 1, 2, 3, 7, 8, 9
> > How to use iota to define them?

> D = iota + 3 // for skip 3

I think this is slightly misleading - this idiom doesn't
skip three - it adds 3 to the current value of iota

roger peppe
https://groups.google.com/d/msg/golang-nuts/R813n17CS58/NIDM-qEOAQAJ

*/

const (
	A = iota + 1
	B
	C
	D = iota + 20
	E
	F = iota + 2
	G
)

func iota_with_gaps() {
	fmt.Println(C, D, E, F, G)
	// 3 23 24 7 8
}
