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

	s.Begin = func(s *awk.Script) {
		s.SetFS(",[ \t]*|[ \t]+")
		s.State = 0.0
	}

	s.AppendStmt(nil, func(s *awk.Script) { s.State = s.State.(float64) + s.F(1).Float64() })

	s.End = func(s *awk.Script) {
		sum := s.State.(float64)
		s.Println("sum is", sum, "average is", sum/float64(s.NR))
	}

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

/*

Example (12,13)
https://godoc.org/github.com/spakin/awk

Add up the first column and print the sum and average
with input fields separated by a comma, space and tab characters
(AWK:

BEGIN { FS = ",[ \t]*|[ \t]+" }
    {s += $1 }
END {print "sum is", s, "average is", s/NR}

)

$ seq 10 | go run Example13.go
sum is 55 average is 5.5

*/
