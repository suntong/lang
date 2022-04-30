// https://www.freecodecamp.org/news/generics-in-golang/

package main

import (
	"fmt"
)

type worker string

func (w worker) Work() {
	fmt.Printf("%s is working\n", w)
}

func DoWork[T any](things []T) {
	for _, v := range things {
		v.Work()
	}
}

func main() {
	var a, b, c worker
	a = "A"
	b = "B"
	c = "C"
	DoWork([]worker{a, b, c})
}
