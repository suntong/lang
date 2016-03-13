////////////////////////////////////////////////////////////////////////////
// Porgram: MethExtend.go
// Purpose: Extend Methord to Go Struct
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Howard C. Shaw III
//          https://groups.google.com/d/msg/golang-nuts/snoIyANd-8c/V_IC57y4AwAJ
//          https://groups.google.com/d/msg/golang-nuts/snoIyANd-8c/hhOnu-lFAgAJ
//          https://groups.google.com/d/msg/golang-nuts/snoIyANd-8c/ZiC1wxhZAgAJ
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

import (
	"github.com/suntong/shaper"
)

func main() {
	test1()
	fmt.Print("\n")
	test2()
}

func test1() {
	Replace := shaper.NewFilter().ApplyReplace("test", "biscuit", -1)
	// Add stages later on
	Replace.ApplyToUpper()

	// Demonstrating copy
	Extra := Replace.Copy()
	// The copy has everything Replace had (note the change to Testificate.)
	Extra.ApplyReplace("THIS", "THAT", -1)
	fmt.Printf("%s\n", Extra.Process("This is also a test. Testificate."))
	fmt.Print("^^^^\n")
	// but Replace still has the old original effect.
	fmt.Printf("%s\n", Replace.Process("This is also a test. Testificate."))
	fmt.Print("^^^^\n")
}

// Extending shaper.Shaper
type Shaper struct {
	*shaper.Shaper
}

func (sp *Shaper) Dummy() *Shaper {
	sp.AddFilter(func(s string) string {
		return "I'm in your function replacing your strings with hardcoded text."
	})
	return sp

}

// Make a new Shaper filter and start adding bits
func NewFilter() *Shaper {
	return &Shaper{Shaper: shaper.NewFilter()}
}

func test2() {
	Replace := NewFilter()
	Replace.ApplyReplace("test", "biscuit", -1)
	fmt.Printf("%s\n", Replace.Process("This is also a test. Testificate."))
	Replace.Dummy()
	fmt.Printf("%s\n", Replace.Process("This is also a test. Testificate."))

	// Note that this does NOT work,
	// because all the child functions return *shaper.Shaper, which does not have Dummy()
	//	Replace := NewFilter().ApplyReplace("test", "biscuit", -1)
	//	fmt.Printf("%s\n", Replace.Process("This is also a test. Testificate."))
	//	Replace.Dummy()
	//	fmt.Printf("%s\n", Replace.Process("This is also a test. Testificate."))
	// would give
	// ./StructMethExtend.go:52: Replace.Dummy undefined (type *shaper.Shaper has no field or method Dummy)
	// because the ApplyReplace returned a shaper.Shaper.
}
