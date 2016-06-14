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
		for line, err := s.GetLine(nil); !line.Match("^7"); {
			if err != nil {
				panic(err)
			}
			println(">", s.NR, line.String())
			s.Next()
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

$ seq 12 | xargs -n 3  | go run SkipLine.go
.
.
Passed.

*/
