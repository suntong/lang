////////////////////////////////////////////////////////////////////////////
// Program: GoMap.go
// Purpose: go map demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: as credited below
////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {
	gobyexample()

	dotnetperls1()
	dotnetperls2()
	dotnetperls3()
	dotnetperls4()
	dotnetperls5()
	dotnetperls6()
	dotnetperls7()
	dotnetperls8()
	dotnetperls9()
	dotnetperls10()

	gotour()
}

////////////////////////////////////////////////////////////////////////////
// https://gobyexample.com/maps

// _Maps_ are Go's built-in [associative data type](http://en.wikipedia.org/wiki/Associative_array)
// (sometimes called _hashes_ or _dicts_ in other languages).
func gobyexample() {

	// To create an empty map, use the builtin `make`:
	// `make(map[key-type]val-type)`.
	m := make(map[string]int)

	// Set key/value pairs using typical `name[key] = val`
	// syntax.
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. `Println` will show all of
	// its key/value pairs.
	fmt.Println("map:", m)

	// Get a value for a key with `name[key]`.
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// The builtin `len` returns the number of key/value
	// pairs when called on a map.
	fmt.Println("len:", len(m))

	// The builtin `delete` removes key/value pairs from
	// a map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// The optional second return value when getting a
	// value from a map indicates if the key was present
	// in the map. This can be used to disambiguate
	// between missing keys and keys with zero values
	// like `0` or `""`.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// You can also declare and initialize a new map in
	// the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

/*

map: map[k1:7 k2:13]
v1:  7
len: 2
map: map[k1:7]
prs: false
map: map[foo:1 bar:2]

*/

////////////////////////////////////////////////////////////////////////////
// https://www.dotnetperls.com/map-go

// == First example, use map
func dotnetperls1() {
	// Map animal names to color strings.
	// ... Create a map with composite literal syntax.
	colors := map[string]string{
		"bird":  "blue",
		"snake": "green",
		"cat":   "black",
	}

	// Get color of snake.
	c := colors["snake"]

	// Display string.
	fmt.Println(c) // green
}

// == Add with assignment, adds keys with assignments
// The len operator, also used on arrays and slices, can be used on a
// map. It counts keys (which is the same as the count of values).
func dotnetperls2() {
	// Create an empty map.
	names := map[int]string{}
	// Add three pairs to the map in separate statements.
	names[990] = "file.txt"
	names[1009] = "data.xls"
	names[1209] = "image.jpg"

	// There are three pairs in the map.
	fmt.Println(len(names)) // 3

	// missing key?
	fmt.Printf("has '%s', '%s'not", names[990], names[991])
}

// == Delete, uses delete built-in
// We can add keys to a map by assigning them. But to delete a key and its
// value, we must use the delete built-in: delete, the pair is entirely erased.
func dotnetperls3() {
	// Create an empty map and add three pairs to it.
	ids := map[string]int{}
	ids["steve"] = 10
	ids["mark"] = 20
	ids["adnan"] = 30
	fmt.Println(len(ids)) // 3

	// Delete one key from it.
	delete(ids, "steve")
	fmt.Println(len(ids)) // 2
}

// == Range, loop. uses for-loop on map
func dotnetperls4() {
	// Create a string to string map.
	animals := map[string]string{}
	animals["cat"] = "Mittens"
	animals["dog"] = "Spot"

	// Loop over the map.
	for key, value := range animals {
		fmt.Println(key, "=", value)
	}
}

/*

cat = Mittens
dog = Spot

*/

// == Ok syntax. uses map, ok syntax
// A key lookup in a map returns two values: the value (if found) and a
// value that indicates success or failure. With the "comma ok" idiom, we
// can test for existence and store the lookup result in one statement.
func dotnetperls5() {
	meds := map[string]int{
		"Ativan":   15,
		"Xanax":    20,
		"Klonopin": 30,
	}

	// The ok variable is set to true.
	if dosage, ok := meds["Xanax"]; ok {
		fmt.Println("Xanax", dosage)
	}

	// The ok variable is set to false.
	// ... The string does not exist in the map.
	if dosage, ok := meds["Atenolol"]; ok {
		fmt.Println("Atenolol", dosage)
	}
}

/*

Xanax 20

*/

// == Get keys from map. gets slice of keys from map
func dotnetperls6() {
	// Create map with three string keys.
	sizes := map[string]int{
		"XL": 20,
		"L":  10,
		"M":  5,
	}

	// Loop over map and append keys to empty slice.
	keys := []string{}
	for key, _ := range sizes {
		keys = append(keys, key)
	}

	// This is a slice of the keys.
	fmt.Println(keys) // [XL L M]
}

// == Get values.
func dotnetperls7() {
	// A simple map.
	birds := map[string]string{
		"finch":    "yellow",
		"parakeet": "blue",
	}

	// Place values in a string slice.
	values := []string{}
	for _, value := range birds {
		values = append(values, value)
	}

	// The values.
	fmt.Println(values) // [yellow blue]
}

// == Make, capacity.
// Make() can create a map with a capacity. This is the number of elements
// that can be placed in the map without resizing the map. A capacity
// optimizes the map.
func dotnetperls8() {
	// Create a map with a capacity of 200 pairs.
	// ... This makes adding the first 200 pairs faster.
	lookup := make(map[string]int, 200)

	// Use the new map.
	lookup["cat"] = 10
	result := lookup["cat"]
	fmt.Println(result) // 10
}

// == Randomized range loop
// The Go runtime randomizes the loop order over a map when the range
// keyword is used. This means programs that rely on a certain ordering of
// elements will fail sooner.
func dotnetperls9() {
	// Create a map with three key-value pairs.
	lookup := map[int]int{
		1: 10,
		2: 20,
		3: 30,
	}

	// Loop ten times.
	for i := 0; i < 10; i++ {
		// Print all keys in range loop over map.
		// ... Ordering is randomized.
		for key := range lookup {
			fmt.Print(key)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

/*

2 3 1
1 2 3
3 1 2
1 2 3
1 2 3
...

*/

// == Func argument. passes map as func argument
// In Go we place important units of logic in funcs. We can pass a map
// object to a func like any other type. We must specify the key and value
// type in the argument of the func.
func PrintGreen(colors map[string]int) {
	// Handle map argument.
	fmt.Println(colors["green"])
}

func dotnetperls10() {
	// This map has two string keys.
	colors := map[string]int{
		"blue":  10,
		"green": 20,
	}
	// Pass map to func.
	PrintGreen(colors) // 20
}

/*

Other examples,

Sort. A map stores its keys in an order that makes them fast to look up. We cannot sort a map. But we can extract the keys (or values) from a map and sort them.

Sort: keys in map
https://www.dotnetperls.com/sort-go

String slices. One common kind of map has string slices as the values. In another example, we append strings to the values in a string slice map.

Map, String Slices
https://www.dotnetperls.com/map-string-slice-go

Struct keys. A struct can be used as the key to a map. With structs, we can create compound keys. A key can contain an int and a string together.

Struct: map keys
https://www.dotnetperls.com/struct-go

*/

////////////////////////////////////////////////////////////////////////////
// https://tour.golang.org/moretypes/19

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func gotour() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

/*

{40.68433 -74.39967}

*/
