////////////////////////////////////////////////////////////////////////////
// Purpose: Go awk GetLine/SkipLine demo
// Authors: Tong Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"

	"github.com/spakin/awk"
)

func main() {
	s := awk.NewScript()

	s.AppendStmt(nil, func(s *awk.Script) {
		// Validate the current line.
		for i := 1; i <= 3; i++ {
			if s.F(i).Int() != (s.NR-1)*3+i {
				check((s.NR-1)*3+i, s.F(i).Int())
			}
		}
		println(".", s.NR, s.F(0).String())

		// Skip till ^7.
		for !s.F(0).Match("^7") {
			v, err := s.GetLine(nil)
			if err != nil {
				return // panic(err)
			}
			s.SetF(0, v)
			println(">", s.NR, v.String())
		}
	})

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}

	println("Passed.")
}

func check(a, b int) {
	if a == b {
		return
	}
	panic("Received != expected")
}

/*

$ seq 12 | xargs -n 3  | go run SkipLine2.go
. 1 1 2 3
> 2 4 5 6
> 3 7 8 9
. 4 10 11 12
Passed.

*/
