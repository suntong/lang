package main

import "fmt"

func main() {
	gobyexampleMain()
	gotourMain()
}

////////////////////////////////////////////////////////////////////////////
// https://gobyexample.com/closures

// Go supports [_anonymous functions_](http://en.wikipedia.org/wiki/Anonymous_function),
// which can form <a href="http://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closures</em></a>.
// Anonymous functions are useful when you want to define
// a function inline without having to name it.

// This function `intSeq` returns another function, which
// we define anonymously in the body of `intSeq`. The
// returned function _closes over_ the variable `i` to
// form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func gobyexampleMain() {

	// We call `intSeq`, assigning the result (a function)
	// to `nextInt`. This function value captures its
	// own `i` value, which will be updated each time
	// we call `nextInt`.
	nextInt := intSeq()

	// See the effect of the closure by calling `nextInt`
	// a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that
	// particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println()
}

/*

1
2
3
1

*/

////////////////////////////////////////////////////////////////////////////
// https://tour.golang.org/moretypes/25

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func gotourMain() {
	pos, neg := adder(), adder()
	for i := 0; i < 8; i++ {
		fmt.Println(
			i,
			pos(i),
			neg(-2*i),
		)
	}
}

/*

0 0 0
1 1 -2
2 3 -6
3 6 -12
4 10 -20
5 15 -30
6 21 -42
7 28 -56

*/
