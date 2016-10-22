////////////////////////////////////////////////////////////////////////////
// Porgram: Set-Go_Diff.go
// Purpose: Demo the golang-set usage (diff)
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: http://stackoverflow.com/questions/23870102/
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	set "github.com/deckarep/golang-set"
)

func main() {
	//note that the set accepts []interface{}
	X := []interface{}{10, 12, 12, 12, 13}
	Y := []interface{}{12, 14, 15}

	Sx := set.NewSetFromSlice(X)
	Sy := set.NewSetFromSlice(Y)
	result1 := Sx.Difference(Sy)
	result2 := Sy.Difference(Sx)

	fmt.Println(result1)
	fmt.Println(result2)
}

/*

$ go run Set-Go.go
Set{10, 13}
Set{14, 15}

*/
