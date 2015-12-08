////////////////////////////////////////////////////////////////////////////
// Porgram: AddToList
// Purpose: Adding elements to a slice
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: A Tour of Go
//          https://tour.golang.org/moretypes/14
////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {
	intSliceTest()
	fmt.Println()
	stringSliceTest()
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

func stringSliceTest() {
	var a []string
	printStringSlice("a", a)

	// append works on nil slices.
	a = append(a, "000")
	printStringSlice("a", a)

	// the slice grows as needed.
	a = append(a, "111")
	printStringSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, "222", "333", "444")
	printStringSlice("a", a)
}

func printStringSlice(s string, x []string) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
