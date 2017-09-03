////////////////////////////////////////////////////////////////////////////
// Program: MapSort.go
// Purpose: Go map sort
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"fmt"
	"sort"
)

type Concordance map[string]int

func main() {
	ages := Concordance{
		"a": 1,
		"c": 3,
		"d": 4,
		"b": 2,
	}

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names) //sort by key
	for _, k := range names {
		fmt.Println("Key:", k, "Value:", ages[k])
	}
	fmt.Println("=====")
	fmt.Println(ages.String())
	fmt.Println("=====")
}

/*

Key: a Value: 1
Key: b Value: 2
Key: c Value: 3
Key: d Value: 4
=====
[ a:1 b:2 c:3 d:4 ]
=====

*/

func (con Concordance) String() string {
	buf := bytes.NewBufferString("")

	var keys []string
	for k := range con {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	buf.WriteString("[ ")
	for _, k := range keys {
		fmt.Fprintf(buf, "%s:%d ", k, int(con[k]))
	}
	buf.WriteByte(']')
	return buf.String()
}
