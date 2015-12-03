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

	s.AppendStmt(nil, func(s *awk.Script) {
		// Validate the current line.
		for i := 1; i <= 3; i++ {
			if s.F(i).Int() != (s.NR-1)*3+i {
				check((s.NR-1)*3+i, s.F(i).Int())
			}
		}

		// Read and validate the next line.
		line, err := s.GetLine(nil)
		if err != nil {
			panic(err)
		}
		s.SetF(0, line)
		for i := 1; i <= 3; i++ {
			if s.F(i).Int() != (s.NR-1)*3+i {
				check((s.NR-1)*3+i, s.F(i).Int())
			}
		}
		println(".")
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

TestGetLineSetF
https://github.com/spakin/awk/commit/81db0f480a1e2200650c2663ecfa531c725de619
tests that GetLine + SetF can replace the current input line

$ seq 12 | xargs -n 3  | go run GetLine.go
.
.
Passed.

*/
