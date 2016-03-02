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
	s.AppendStmt(func(s *awk.Script) bool {
		return s.F(2).Match("6:") && !s.F(1).Match("6:")
	}, nil)
	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

/*

Example (05)
https://godoc.org/github.com/spakin/awk

Write any line in which the second field matches the regular expression "6:" and the first field does not (AWK: $2 ~ /6:/ && $1 !~ /6:/).

$ seq 8 | cat -n | sed 's/$/:/' | go run Example05.go
     6  6:

*/
