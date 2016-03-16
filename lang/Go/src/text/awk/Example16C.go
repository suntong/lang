////////////////////////////////////////////////////////////////////////////
// Purpose: Go awk demo, with NewValueArray
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
)

import (
	"github.com/spakin/awk"
)

func main() {
	s := awk.NewScript()

	sa := s.NewValueArray()
	sa.Set("prev", "")
	s.AppendStmt(func(s *awk.Script) bool {
		return !s.F(1).StrEqual(sa.Get("prev").String())
	},
		func(s *awk.Script) {
			//s.Println()
			fmt.Println(s.F(0))
			sa.Set("prev", s.F(1))
			fmt.Println(sa.Get("prev").String())
		})

	s.End = func(s *awk.Script) {
		fmt.Println(sa.Get("no-exist").String())
		fmt.Println(sa.Get("no-exist").Int())
		sa.Set("no-exist", 3)
		fmt.Println(sa.Get("no-exist").String())
		fmt.Println(sa.Get("prev").String())
	}

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

/*

Example (16)
https://godoc.org/github.com/spakin/awk

Write all lines whose first field is different from the previous line's
(AWK: $1 != prev {print; prev = $1}).

$ seq 5 | go run Example16C.go
1
1
2
2
3
3
4
4
5
5

*/
