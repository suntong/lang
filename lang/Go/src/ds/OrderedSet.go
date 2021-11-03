package main

import (
	"fmt"

	"github.com/goombaio/orderedset"
)

func main() {
	s := orderedset.NewOrderedSet()
	s.Add("First element")
	s.Add("A")
	s.Add("B")
	s.Add("A")
	s.Add("A")
	s.Add("Second element")
	s.Add("B")
	s.Add("Last element")

	for _, entry := range s.Values() {
		fmt.Println(entry)
	}
	// Output:
	// First element
	// Second element
	// Last element
}
