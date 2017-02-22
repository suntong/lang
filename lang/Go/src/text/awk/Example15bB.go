////////////////////////////////////////////////////////////////////////////
// Purpose: Go awk read in a block demo
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/spakin/awk"
)

const (
	startTag = "3:"
	stopTag  = "6:"
)

func main() {
	s := awk.NewScript()
	// Read in a block
	// https://github.com/spakin/awk/issues/8#issuecomment-281255053

	// AWK: /3:/,/6:/ { sub(/^/, "+ "); print; next; }
	s.State = make([]byte, 0)
	s.AppendStmt(awk.Auto(startTag, stopTag), func(s *awk.Script) {
		// if s.F(0).Match(startTag) {
		//   s.State = make([]byte, 0)
		// }
		str := []byte("+ " + s.F(0).String() + "\n")
		s.State = append(s.State.([]byte), str...)
		if s.F(0).Match(stopTag) {
			s.Println("=== I HAVE A TABLE ===")
			fmt.Print(string(s.State.([]byte)))
			s.Println("=== ALL DONE ===")
			s.State = make([]byte, 0)
		}
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

How to read in a block
https://github.com/spakin/awk/issues/8

How to define the awk.Script func so that it will process the lines between
"start" and "stop" as a block? -- My "start" and "stop" case is "<table>"
and "</table>", and I got the error of invalid XML format, and I believe I'm
reading them one line at a time, instead of all the lines between "<table>"
and "</table>" as a block.


$ seq 18 | tac | cat -n | sed 's/\t/: /' | go run Example15bB.go
     1: 18
     2: 17
=== I HAVE A TABLE ===
+      3: 16
+      4: 15
+      5: 14
+      6: 13
=== ALL DONE ===
     7: 12
     8: 11
     9: 10
    10: 9
    11: 8
    12: 7
=== I HAVE A TABLE ===
+     13: 6
+     14: 5
+     15: 4
+     16: 3
=== ALL DONE ===
    17: 2
    18: 1

*/
