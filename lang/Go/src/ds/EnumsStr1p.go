////////////////////////////////////////////////////////////////////////////
// Porgram: EnumsStr1p.go
// Purpose: Go Enum and its string representation lib demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: Egon https://groups.google.com/d/msg/golang-nuts/fCdBSRNNUY8/P45qC_03LoAJ
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/suntong/enum"
)

var (
	Alpha = enum.Ciota("Alpha")
	Beta  = enum.Ciota("Beta")
)

type Example struct {
	enum.Enum
}

func main() {
	fmt.Printf("%+v\n", Example{Alpha})
	fmt.Printf("%+v\n", Example{Beta})
	fmt.Println("=======")
	fmt.Printf("%d\t%d\n", Alpha, Alpha+1)
	fmt.Printf("%+v\t%+v\n", Example{Beta - 1}, Example{Alpha + 1})
	fmt.Println("=======")
	if a, ok := enum.Get("Alpha"); ok {
		fmt.Printf("%d\n", a)
	}
	if b, ok := enum.Get("Beta"); ok {
		fmt.Printf("%d: %+v\n", b, Example{b})
	}
}
