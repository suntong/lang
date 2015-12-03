////////////////////////////////////////////////////////////////////////////
// Purpose: Go awk demo
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/spakin/awk"
	"os"
)

func main() {
	s := awk.NewScript()

	s.State = s.NewValue("")
	s.AppendStmt(func(s *awk.Script) bool { return !s.F(1).StrEqual(s.State) },
		func(s *awk.Script) {
			s.Println()
			s.State = s.F(1)
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

$ seq 5 | go run Example16A.go
1
2
3
4
5

*/
