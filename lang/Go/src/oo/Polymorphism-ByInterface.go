// Interfaces for Polymorphism
// https://medium.com/@gianbiondi/interfaces-in-go-59c3dc9c2d98

/*

Gian Biondi, Mar 11, 2016

Interfaces in Go provide a way to specify behavior for a type. An interface
is basically a set of methods.

- For a type to satisfy an interface, it must implement at least all the
  methods specified in the interface.

- All interfaces are satisfied implicitly (no implements declaration), so
  you can write interfaces for code that you don’t own.

- As a result all types in Go automatically satisfy the empty interface. We
  can take advantage of this quality by using an empty interface as an input
  parameter for a function which we want to accept any type, and then using
  Reflection to determine the type and act on the input accordingly.

Go facilitates polymorphism through interfaces. Any type that satisfies an
interface can be used interchangeably as that interface type. Here’s a quick
example.

This example is a bit contrived, but it gets the point across. This pattern
is used in many places throughout the language. A more practical example is
in Go’s Sort functionality. This built-in package sorts collections. The
developer can define how a type should be sorted by implementing the
“Interface” interface in the Sort package. The “Interface” contains just
three methods: “Len”, “Less”, and “Swap”. The sort package calls these three
methods in a Quicksort algorithm.

*/

package main

import "fmt"

type Animal interface {
	Type() string
	Swim() string
}
type Dog struct {
	Name  string
	Breed string
}
type Frog struct {
	Name  string
	Color string
}

func main() {
	f := new(Frog)
	d := new(Dog)
	zoo := [...]Animal{f, d}
	for _, a := range zoo {
		fmt.Println(a.Type(), " can ", a.Swim())
	}
}

func (f *Frog) Type() string {
	return "Frog"
}
func (f *Frog) Swim() string {
	return "Kick"
}

func (d *Dog) Swim() string {
	return "Paddle"
}
func (d *Dog) Type() string {
	return "Doggie"
}
