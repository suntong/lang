// https://go.dev/play/p/GdEUnzrIuYp

// https://github.com/wk8/go-ordered-map
// go get -v -u github.com/wk8/go-ordered-map

package main

import (
	"fmt"

	orderedmap "github.com/wk8/go-ordered-map"
)

func main() {
	test1()
	test2()
}

func test1() {
	om := orderedmap.New()

	om.Set("foo", "bar")
	om.Set("bar", "baz")
	om.Set("coucou", "toi")

	fmt.Println(om.Get("foo"))          // => bar, true
	fmt.Println(om.Get("i dont exist")) // => <nil>, false

	// iterating pairs from oldest to newest:
	for pair := om.Oldest(); pair != nil; pair = pair.Next() {
		fmt.Printf("%s => %s\n", pair.Key, pair.Value)
	} // prints:
	// foo => bar
	// bar => baz
	// coucou => toi
	fmt.Println()

	// iterating over the 2 newest pairs:
	i := 0
	for pair := om.Newest(); pair != nil; pair = pair.Prev() {
		fmt.Printf("%s => %s\n", pair.Key, pair.Value)
		i++
		if i >= 2 {
			break
		}
	} // prints:
	// coucou => toi
	// bar => baz
	fmt.Println()
}

type myStruct struct {
	payload string
}

func test2() {
	om := orderedmap.New()

	om.Set(12, &myStruct{"foo"})
	om.Set(1, &myStruct{"bar"})

	value, present := om.Get(12)
	if !present {
		panic("should be there!")
	}
	fmt.Println(value.(*myStruct).payload) // => foo

	for pair := om.Oldest(); pair != nil; pair = pair.Next() {
		fmt.Printf("%d => %s\n", pair.Key, pair.Value.(*myStruct).payload)
	} // prints:
	// 12 => foo
	// 1 => bar
}
