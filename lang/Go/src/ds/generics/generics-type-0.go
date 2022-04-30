// https://www.freecodecamp.org/news/generics-in-golang/

package main

import (
	"fmt"
)

type Number interface {
	int | float64
}

func MultiplyTen[T Number](a T) T {
	return a * 10
}

func main() {
	fmt.Println(MultiplyTen(10))
	fmt.Println(MultiplyTen(5.55))
}
