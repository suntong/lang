////////////////////////////////////////////////////////////////////////////
// Program: MapToArray.go
// Purpose: Go map to array
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

/*

A data structure in Go that maps into a slice.

The following code will assign different people into different age groups.
E.g., for age group of 65, there are people "ABC", and "XYZ".

*/

package main

import "fmt"

type Person struct {
	first string
	last  string
}

type People []Person

// Age to different peoples
type PeopleGroup map[int]People

func (pg PeopleGroup) Set(key int, value Person) {
	pg[key] = append(pg[key], value)
}

func (pg PeopleGroup) Get(key int) (People, bool) {
	val, ok := pg[key]
	return val, ok
}

func main() {
	//m := make(map[int]People)
	m := make(PeopleGroup)
	fmt.Println("Values in map (after creating): ", m)
	m[65] = People{Person{"ABC", "L1"}, Person{"XYZ", "L2"}}
	m[60] = People{}

	fmt.Println("Length of map: ", len(m))
	fmt.Println("Values in map(after adding values): ", m)
	fmt.Println(m.Get(60))
	fmt.Println()

	m[60] = append(m[60], Person{"XYZ", "LL2"})
	m.Set(60, Person{"ABC", "LL1"})
	m.Set(50, Person{"ABC", "L50"})
	fmt.Println("Values in map(after adding values): ", m)
	fmt.Println("Length of value:", len(m[20]))
	fmt.Println(m.Get(60))
	fmt.Println(m.Get(30))
	fmt.Println()

	for k, v := range m {
		fmt.Println("Key :", k, " Value :", v)
	}
}

/*

Values in map (after creating):  map[]
Length of map:  2
Values in map(after adding values):  map[65:[{ABC L1} {XYZ L2}] 60:[]]
[] true

Values in map(after adding values):  map[65:[{ABC L1} {XYZ L2}] 60:[{XYZ LL2} {ABC LL1}] 50:[{ABC L50}]]
Length of value: 0
[{XYZ LL2} {ABC LL1}] true
[] false

Key : 65  Value : [{ABC L1} {XYZ L2}]
Key : 60  Value : [{XYZ LL2} {ABC LL1}]

*/
