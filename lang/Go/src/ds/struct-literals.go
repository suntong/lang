package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	// https://tour.golang.org/moretypes/5
	fmt.Println(v1, p, v2, v3)
	// {1 2} &{1 2} {1 0} {0 0}

	// == struct Slice
	var a []Vertex
	// append works on nil slices.
	a = append(a, v1)
	// the slice grows as needed.
	a = append(a, v2)
	fmt.Println(a)
	// [{1 2} {1 0}]
}
