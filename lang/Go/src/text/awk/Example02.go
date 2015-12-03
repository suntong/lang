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
	s.AppendStmt(func(s *awk.Script) bool { return s.NR%10 == 0 }, nil)
	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

/*

Example (02)
https://godoc.org/github.com/spakin/awk

Write every tenth line (AWK: (NR % 10) == 0).

$ seq 30 | go run Example02.go
10
20
30

*/
