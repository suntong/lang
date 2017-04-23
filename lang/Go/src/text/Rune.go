////////////////////////////////////////////////////////////////////////////
// Porgram: Rune.go
// Purpose: Go Rune manipulation demo
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

/*

Q: Given a unicode code string, be it "4e16", or "0x4e16", or "u4e16",
how to turn it into a single char rune?

A:

1. Convert it to a number
2. Use `rune()`

Ivan Kurnosov

*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Hello \u4e16\u754c (%c%c): %+q\n", 0x4e16, 0x754c, "世界")
	code10 := "4e16"
	//code11 := "0x4e16"
	//code12 := "u4e16"
	// hex to int
	c1, _ := strconv.ParseInt(code10, 16, 16)
	fmt.Printf("Hello %c\u754c %v\n", c1, c1 == '世')
}
