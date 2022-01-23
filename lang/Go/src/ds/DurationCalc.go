// https://go.dev/play/p/AvzauOB_pzY

package main

import (
	"fmt"
	"time"
)

func main() {
	// https://pkg.go.dev/time#pkg-constants

	// To count the number of units in a Duration, divide:
	{
		second := time.Second
		fmt.Println(int64(second / time.Millisecond)) // prints 1000
	}
	{
		second := 3 * time.Second
		fmt.Println(int64(second / time.Second)) // prints 3
	}
	// To convert an integer number of units to a Duration, multiply:
	{
		seconds := 10
		fmt.Println(time.Duration(seconds) * time.Second) // prints 10s
	}

	d := 100 * time.Microsecond
	fmt.Println(d) // Output: 100µs

	value := 100 // value is of type int

	d2 := time.Duration(value) * time.Millisecond
	fmt.Println(d2) // Output: 100ms

	ms := int64(d2 / time.Millisecond)
	fmt.Println("ms:", ms) // Output: ms: 100

	fmt.Println("ns:", int64(d2/time.Nanosecond))  // ns: 100000000
	fmt.Println("µs:", int64(d2/time.Microsecond)) // µs: 100000
	fmt.Println("ms:", int64(d2/time.Millisecond)) // ms: 100
}
