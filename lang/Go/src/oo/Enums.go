////////////////////////////////////////////////////////////////////////////
// Porgram: Enums
// Purpose: enums demo for golang with constants and iota
// Authors: Tong Sun (c) 2014, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2 -w

/*

In C, enums and variables are associated. I.e., when associate a variable with
enum values, we know for sure that such variable will take no values other
than the ones from defined enums. What's the best way to establish such
association in go?

The following example demonstrates the logical association betwee a groups of
constants and variables, expressing the idea that the variable are associated
with the constants.

*/

package main

import "fmt"

////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type weekday int

const (
	Sunday weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	var x weekday = Friday
	y := weekday(Friday)
	xx := Tuesday
	xx++

	fmt.Println(Friday, x, xx, y)

	// cannot use Wednesday (type weekday) as type int in assignment
	//var yy int = Wednesday

	// invalid operation: (mismatched types int and weekday)
	//var yy int = 3
	//fmt.Println(yy == Wednesday)
	
}
