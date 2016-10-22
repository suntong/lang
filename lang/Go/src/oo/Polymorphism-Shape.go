////////////////////////////////////////////////////////////////////////////
// Porgram: Polymorphism-Shape.go
// Purpose: Demo the Go "polymorphism" feature with Shapes
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://play.golang.org/p/Zn7TjiFQik
////////////////////////////////////////////////////////////////////////////

/*

Problem (From Polymorphism-Subtype.go):

https://groups.google.com/d/msg/golang-nuts/N4MBApd09M8/tOO5ZXtwbhYJ

LRN: Subtype polymorphism: Not applicable (Go doesn't have subtyping).

Goal:

This is to demo that "polymorphism" is still doable in Go.

*/

package main

import (
	"fmt"
)

type Shape interface {
	Area() float32
}

type Point struct {
	x float32
	y float32
}

// Make sure the structs are different sizes so we're sure it'll work with
// all sorts of types
type Circle struct {
	center Point
	radius float32
}

func (c Circle) Area() float32 {
	return 3.1415 * c.radius * c.radius
}

type Rectangle struct {
	ul Point
	lr Point
}

func (r Rectangle) Area() float32 {
	xDiff := r.lr.x - r.ul.x
	yDiff := r.ul.y - r.lr.y
	return xDiff * yDiff
}

func main() {
	mtDict := make(map[string]Shape)
	// No problem storing different custom types in the multitype dict
	mtDict["circ"] = Circle{Point{3.0, 3.0}, 2.0}
	mtDict["rect"] = Rectangle{Point{2.0, 4.0}, Point{4.0, 2.0}}

	for k, v := range mtDict {
		fmt.Printf("[%v] [%0.2f]\n", k, v.Area())
	}
}

/*

$ go run Polymorphism-Shape.go
[circ] [12.57]
[rect] [4.00]

*/
