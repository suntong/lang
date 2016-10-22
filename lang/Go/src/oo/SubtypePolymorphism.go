////////////////////////////////////////////////////////////////////////////
// Porgram: SubtypePolymorphism
// Purpose: Subtype polymorphism showcase
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: hutch
//          https://groups.google.com/d/msg/golang-nuts/N4MBApd09M8/0ij9yGHK_8EJ
////////////////////////////////////////////////////////////////////////////

/*

https://groups.google.com/d/msg/golang-nuts/N4MBApd09M8/tOO5ZXtwbhYJ

LRN:

Subtype polymorphism: Not applicable (Go doesn't have subtyping).
Although if you embed a struct A implementing interface X into a struct B,
struct B will implement interface X, and can be used instead of struct A in
places where struct A is expected. So, kind of yes.

Robert Johnstone:

interfaces behave similarly to virtual functions, but they are not identical.  See the (following) example program by hutch.

*/

package main

import "fmt"

type A struct {
	astring string
}

type B struct {
	A
	bstring string
}

type Funny interface {
	strange()
	str() string
}

func (this *A) strange() {
	fmt.Printf("my string is %q\n", this.str())
}

func (this *A) str() string {
	return this.astring
}

func (this *B) str() string {
	return this.bstring
}

func main() {
	b := new(B)
	b.A.astring = "this is an A string"
	b.bstring = "this is a B string"

	b.strange()
	// Output: my string is "this is an A string"

	// Many people familiar with OO (and unfamiliar with Go) will be quite
	// surprised at the output of that program.
}
