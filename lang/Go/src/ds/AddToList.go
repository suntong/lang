////////////////////////////////////////////////////////////////////////////
// Porgram: Composition.go
// Purpose: Go composition & field promotion demo
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: A Tour of Go
//          https://tour.golang.org/moretypes/14
////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {
	intSliceTest()
}

func intSliceTest() {
	var a []int
	printIntSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printIntSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printIntSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printIntSlice("a", a)
}

func printIntSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
