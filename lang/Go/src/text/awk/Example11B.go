////////////////////////////////////////////////////////////////////////////
// Purpose: Go awk demo
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"

	"github.com/spakin/awk"
)

func main() {
	s := awk.NewScript()

	s.Begin = func(s *awk.Script) {
		s.SetFS(",")
	}

	s.AppendStmt(nil, func(s *awk.Script) {
		s.Println(s.F(1), s.F(2), s.F(3), s.F(4))
	})

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

/*


Write all four fields.

$ seq 12 | xargs -n 4 | sed 's/ /,/g'
1,2,3,4
5,6,7,8
9,10,11,12

!! | go run Example11B.go
1 2 3 4
5 6 7 8
9 10 11 12

*/
