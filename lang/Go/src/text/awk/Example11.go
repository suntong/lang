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
	s.AppendStmt(nil, func(s *awk.Script) { s.Println(s.F(2), s.F(1)) })
	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

/*

Example (11)
https://godoc.org/github.com/spakin/awk

Write the first two fields in opposite order (AWK: {print $2, $1}).

$ seq 8 | cat -n | sed 's/$/:/' | go run Example11.go
1: 1
2: 2
3: 3
4: 4
5: 5
6: 6
7: 7
8: 8

*/
