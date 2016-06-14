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
	scr := awk.NewScript()
	scr.AppendStmt(func(s *awk.Script) bool {
		return s.F(0).Match("^B")
	}, func(s *awk.Script) {
		for !s.F(0).Match("^E") {
			v, err := s.GetLine(nil)
			if err != nil {
				return
			}
			s.SetF(0, v)
		}
		s.Next()
	})
	scr.AppendStmt(nil, nil)

	if err := scr.Run(os.Stdin); err != nil {
		panic(err)
	}
}

/*

Skip lines with given regex for start & end (from /^B/ to /^E/):

$ jot -c 6 'A'
A
B
C
D
E
F

$ jot -c 6 'A' | awk '/^B/{ while ($0 !~/^E/) { getline;} next; }; 1'
A
F

$ jot -c 6 'A' | go run SkipLine1.go
A
F

*/
