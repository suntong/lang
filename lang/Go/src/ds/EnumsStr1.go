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

func Ciota(s string) Enum {
	enums = append(enums, s)
	return Enum(len(enums) - 1)
}

func Get(s string) (Enum, bool) {
	for ii, vv := range enums {
		if vv == s {
			return Enum(ii), true
		}
	}
	return -1, false
}

var (
	Alpha = Ciota("Alpha")
	Beta  = Ciota("Beta")
)

type Example struct {
	Enum
}

func main() {
	fmt.Printf("%+v\n", Example{Alpha})
	fmt.Printf("%+v\n", Example{Beta})
	fmt.Println("=======")
	fmt.Printf("%d\t%d\n", Alpha, Alpha+1)
	fmt.Printf("%+v\t%+v\n", Example{Beta - 1}, Example{Alpha + 1})
	fmt.Println("=======")
	fmt.Printf("%+v\n", enums)
	fmt.Println("=======")
	if a, ok := Get("Alpha"); ok {
		fmt.Printf("%d\n", a)
	}
	if b, ok := Get("Beta"); ok {
		fmt.Printf("%d: %+v\n", b, Example{b})
	}
}

/*

Alpha
Beta
=======
0	1
Alpha	Beta
=======
[Alpha Beta]
=======
0
1: Beta

*/
