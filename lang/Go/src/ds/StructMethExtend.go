////////////////////////////////////////////////////////////////////////////
// Porgram: MethExtend.go
// Purpose: Extend Methord to Go Struct
// Authors: Tong Sun (c) 2013, All rights reserved
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
	return &Shaper{ShaperStack: PassThrough}
}

func test2() {
	Replace := Shaper.NewFilter().ApplyReplace("test", "biscuit", -1)
	fmt.Printf("%s\n", Replace.Process("This is also a test. Testificate."))
}
