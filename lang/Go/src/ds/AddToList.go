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

/*

Note, polymorphism in Go:

OK to:

	a = append(a, 1)
	a = append(a, 2, 3, 4)

	a = append(a, "111")
	a = append(a, "222", "333", "444")

But it is NOK to define:

func printSlice(s string, x []int) {
}

func printSlice(s string, x []string) {
}

printSlice redeclared in this block
        previous declaration at ...

To summarize (by LRN),
https://groups.google.com/d/msg/golang-nuts/N4MBApd09M8/tOO5ZXtwbhYJ

- Ad-hoc polymorphism: Supported with interfaces. You call a method with the
same name on any number of objects that implement a particular interface,
and they will behave differently, depending on the implementation.

- Parametric polymorphism: No generics. But if you write a function that
works with interfaces (well, object implementing interfaces), not types,
then you can have parametric polymorphism as well (a single function
implementation that works correctly, and with the same algorithm, on any
type that implements the right interface correctly).

- Subtype polymorphism: Not applicable (Go doesn't have subtyping).
Although if you embed a struct A implementing interface X into a struct B,
struct B will implement interface X, and can be used instead of struct A in
places where struct A is expected. So, kind of yes.

*/
