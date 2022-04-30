// https://gosamples.dev/generics-map-function/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func mapSlice[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

func main() {
	numbers := []float64{4, 9, 16, 25}

	newNumbers := mapSlice(numbers, math.Sqrt)
	fmt.Println(newNumbers)

	words := []string{"a", "b", "c", "d"}
	quoted := mapSlice(words, func(s string) string {
		return "\"" + s + "\""
	})
	fmt.Println(quoted)

	stringPowNumbers := mapSlice(numbers, func(n float64) string {
		return strconv.FormatFloat(math.Pow(n, 2), 'f', -1, 64)
	})
	fmt.Println(stringPowNumbers)
}

/*

[2 3 4 5]
["a" "b" "c" "d"]
[16 81 256 625]

*/
