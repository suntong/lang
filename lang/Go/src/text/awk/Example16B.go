////////////////////////////////////////////////////////////////////////////
// Purpose: Go awk demo
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

	s.State = s.NewValue("")
	s.AppendStmt(func(s *awk.Script) bool { return !s.F(1).StrEqual(s.State) },
		func(s *awk.Script) {
			//s.Println()
			fmt.Println(s.F(0))
			s.State = s.F(1)
			fmt.Println(s.State)
		})

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

/*

Example (16)
https://godoc.org/github.com/spakin/awk

Write all lines whose first field is different from the previous line's
(AWK: $1 != prev {print; prev = $1}).

$ seq 5 | go run Example16B.go
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
