// https://www.freecodecamp.org/news/generics-in-golang/

package main

import (
	"fmt"
)

type GenericSlice[T any] []T

func (g GenericSlice[T]) Print() {
	for _, v := range g {
		fmt.Println(v)
	}
}

func Print[T any](g GenericSlice[T]) {
	for _, v := range g {
		fmt.Println(v)
	}
}

func main() {
	g := GenericSlice[int]{1, 2, 3}

	g.Print() //1 2 3
	fmt.Println()
	Print(g) //1 2 3

	s := GenericSlice[string]{"A", "foo", "B", "bar"}

	fmt.Println()
	s.Print() // A foo B bar
}
