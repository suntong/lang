////////////////////////////////////////////////////////////////////////////
// Purpose: Go awk demo
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"

	"github.com/spakin/awk"
)

func main() {
	s := awk.NewScript()

	// AWK: /3:/,/6:/ { sub(/^/, "+ "); print; next; }
	s.AppendStmt(awk.Auto("3:", "6:"),
		func(s *awk.Script) {
			s.Println("+ " + s.F(0).String())
			s.Next()
		})
	// 1; # i.e., print all
	s.AppendStmt(nil, nil)

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

/*

Example (15b)
https://godoc.org/github.com/spakin/awk#ex-package--15b

Write all lines between occurrences of the strings "start" and "stop" (AWK: /start/, /stop/). This version of the Go code uses awk.Auto to define the begin and end conditions as simple regular-expression matches.


$ seq 8 | tac | cat -n | sed 's/\t/: /' | go run Example15b.go
     1: 8
     2: 7
+      3: 6
+      4: 5
+      5: 4
+      6: 3
     7: 2
     8: 1

*/
