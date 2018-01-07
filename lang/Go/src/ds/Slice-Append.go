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
	fmt.Println()
	stringArrayTest()
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

	// Slices can also be created with the make function.
	b := make([]int, 5) // len(b)=5, cap(b)=5
	printIntSlice("b", b)

	// To specify a capacity, pass a third argument to make
	c := make([]int, 0, 5) // len(c)=0, cap(c)=5
	printIntSlice("c", c)

	c = c[:cap(c)] // len(c)=5, cap(c)=5
	printIntSlice("c", c)
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

	fmt.Println()

	fmt.Println("a[1:3] ==", a[1:3])

	// missing low index implies 0
	fmt.Println("a[:3] ==", a[:3])

	// missing high index implies len(a)
	fmt.Println("a[2:] ==", a[2:])

	fmt.Println()

	// init slice with predetermined values
	var b []string = []string{"333", "222"}
	printStringSlice("b", b)

	b = append(b, "111")
	printStringSlice("b", b)

	// How to append slice to slice
	// X: a = append(a, b) -- cannot use b (type []string) as type string in append
	a = append(a, b...)
	printStringSlice("a", a)

	a = appendString(a, "some", "more", "string")
	printStringSlice("a", a)

}

/*

a len=0 cap=0 []
a len=1 cap=1 [0]
a len=2 cap=2 [0 1]
a len=5 cap=6 [0 1 2 3 4]
b len=5 cap=5 [0 0 0 0 0]
c len=0 cap=5 []
c len=5 cap=5 [0 0 0 0 0]

a len=0 cap=0 []
a len=1 cap=1 [000]
a len=2 cap=2 [000 111]
a len=5 cap=5 [000 111 222 333 444]

a[1:3] == [111 222]
a[:3] == [000 111 222]
a[2:] == [222 333 444]

b len=2 cap=2 [333 222]
b len=3 cap=4 [333 222 111]
a len=8 cap=10 [000 111 222 333 444 333 222 111]
a len=11 cap=11 [000 111 222 333 444 333 222 111 some more string]

*/

func appendString(x []string, y ...string) []string {
	var z []string
	zlen := len(x) + len(y)
	// var z [zlen]string X: non-constant array bound zlen
	// expand z to at least zlen
	z = make([]string, zlen)
	copy(z, x)
	copy(z[len(x):], y)
	return z
}

func printStringSlice(s string, x []string) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func stringArrayTest() {
	var a [2]string
	//a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//a = append(a, "Howdy")
	//first argument to append must be slice
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
