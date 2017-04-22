////////////////////////////////////////////////////////////////////////////
// Porgram: EnumsStr1.go
// Purpose: Go Enum and its string demo1
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: Egon https://groups.google.com/d/msg/golang-nuts/fCdBSRNNUY8/P45qC_03LoAJ
////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

var enums []string

type Enum int

func (e Enum) String() string {
	return enums[int(e)]
}

func ciota(s string) Enum {
	enums = append(enums, s)
	return Enum(len(enums) - 1)
}

var (
	Alpha = ciota("A")
	Beta  = ciota("B")
)

type Example struct {
	X Enum
}

func main() {
	fmt.Printf("%+v\n", Example{Alpha})
	fmt.Printf("%+v\n", Example{Beta})
}

/*

{X:A}
{X:B}

*/
