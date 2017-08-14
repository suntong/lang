## TOC
- [Usage](#usage)
  - [> Introduction.go](#-introductiongo)
  - [> cc-dict.go](#-cc-dictgo)
  - [> Example02.go](#-example02go)
  - [> Example05.go](#-example05go)
  - [> Example11.go](#-example11go)
  - [> Example11B.go](#-example11bgo)
  - [> Example13.go](#-example13go)
  - [> Example15b.go](#-example15bgo)
  - [> Example15bB.go](#-example15bbgo)
  - [> Example16A.go](#-example16ago)
  - [> Example16B.go](#-example16bgo)
  - [> Example16C.go](#-example16cgo)
  - [> GetLine.go](#-getlinego)
  - [> SkipLine1.go](#-skipline1go)
  - [> SkipLine2.go](#-skipline2go)
  - [> Wp2Hugo.go](#-wp2hugogo)
  - [> Comments.go](#-commentsgo)

## Usage

#### > Introduction.go
```go
package main

import (
	"github.com/spakin/awk"
	"os"
)

func main() {
	s := awk.NewScript()
	s.AppendStmt(func(s *awk.Script) bool { return s.F(1).Int()%2 == 1 }, nil)
	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

/*

Introduction
https://godoc.org/github.com/spakin/awk#hdr-Introduction

For first column is an odd number:

$5 % 2 == 1

$ seq 10 | go run Introduction.go
1
3
5
7
9

*/
```

#### > cc-dict.go
```go
////////////////////////////////////////////////////////////////////////////
// Purpose: Turn the OpenCC STCharacters.txt file to dictionary
// Authors: Tong Sun (c) 2017
// Sources: https://github.com/go-cc/opencc-dict/blob/master/data/dictionary/STCharacters.txt
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spakin/awk"
)

func main() {
	// https://godoc.org/github.com/spakin/awk
	s := awk.NewScript()

	// == BEGIN
	sa := s.NewValueArray()
	sa.Set("cS", "")
	sa.Set("cT", "")

	// == Match & Process
	s.AppendStmt(nil, func(s *awk.Script) {
		print(s.NR)
		c1 := s.F(1).String()
		for ii := 2; ii <= s.NF; ii++ {
			c2 := s.F(ii).String()
			if c1 == c2 {
				continue
			}
			c2 = s.F(ii).String()
			fmt.Printf("%v:%v\n", c1, c2)
			// fmt.Printf("%+q:%+q\n", c1, c2)
			sa.Set("cS", fmt.Sprintf("%s%+q", sa.Get("cS"), c1))
			sa.Set("cT", fmt.Sprintf("%s%+q", sa.Get("cT"), c2))
		}
		print(" ")
	})

	// == END
	s.End = func(s *awk.Script) {
		dqRegex := regexp.MustCompile(`"`)
		// s.Println(sa.Get("cS"), "\n", sa.Get("cT"))
		cS := dqRegex.ReplaceAllString(sa.Get("cS").String(), "")
		cT := dqRegex.ReplaceAllString(sa.Get("cT").String(), "")
		fmt.Printf("%s\n%s\n", cS, cT)
	}

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

/*
 */
```

#### > Example02.go
```go
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
```

#### > Example05.go
```go
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
```

#### > Example11.go
```go
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
```

#### > Example11B.go
```go
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
```

#### > Example13.go
```go
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
```

#### > Example15b.go
```go
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
```

#### > Example15bB.go
```go
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
```

#### > Example16A.go
```go
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

	s.State = s.NewValue("")
	s.AppendStmt(func(s *awk.Script) bool { return !s.F(1).StrEqual(s.State) },
		func(s *awk.Script) {
			s.Println()
			s.State = s.F(1)
		})

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

/*

Example (16)
https://godoc.org/github.com/spakin/awk

Write all lines whose first field is different from the previous line's
(AWK: $1 != prev {print; prev = $1}).

$ { seq 5; seq 5; } | sort | tee /dev/tty | go run Example16A.go
1
1
2
2
3
3
4
4
5
5
1
2
3
4
5

*/
```

#### > Example16B.go
```go
////////////////////////////////////////////////////////////////////////////
// Purpose: Go awk demo
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
)

import (
	"github.com/spakin/awk"
)

func main() {
	s := awk.NewScript()

	s.State = s.NewValue("")
	s.AppendStmt(func(s *awk.Script) bool { return !s.F(1).StrEqual(s.State) },
		func(s *awk.Script) {
			//s.Println()
			fmt.Println(s.F(0))
			s.State = s.F(1)
			fmt.Println(s.State)
		})

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

/*

Example (16)
https://godoc.org/github.com/spakin/awk

Write all lines whose first field is different from the previous line's
(AWK: $1 != prev {print; prev = $1}).

$ seq 5 | go run Example16B.go
1
1
2
2
3
3
4
4
5
5

*/
```

#### > Example16C.go
```go
////////////////////////////////////////////////////////////////////////////
// Purpose: Go awk demo, with NewValueArray
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
)

import (
	"github.com/spakin/awk"
)

func main() {
	s := awk.NewScript()

	sa := s.NewValueArray()
	sa.Set("prev", "")
	s.AppendStmt(func(s *awk.Script) bool {
		return !s.F(1).StrEqual(sa.Get("prev").String())
	},
		func(s *awk.Script) {
			//s.Println()
			fmt.Println(s.F(0))
			sa.Set("prev", s.F(1))
			fmt.Println(sa.Get("prev").String())
		})

	s.End = func(s *awk.Script) {
		fmt.Println(sa.Get("no-exist").String())
		fmt.Println(sa.Get("no-exist").Int())
		sa.Set("no-exist", 3)
		fmt.Println(sa.Get("no-exist").String())
		fmt.Println(sa.Get("prev").String())
	}

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

/*

Example (16)
https://godoc.org/github.com/spakin/awk

Write all lines whose first field is different from the previous line's
(AWK: $1 != prev {print; prev = $1}).

$ seq 5 | go run Example16C.go
1
1
2
2
3
3
4
4
5
5

*/
```

#### > GetLine.go
```go
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
```

#### > SkipLine1.go
```go
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
```

#### > SkipLine2.go
```go
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
		for !s.F(0).Match("^7") {
			v, err := s.GetLine(nil)
			if err != nil {
				return // panic(err)
			}
			s.SetF(0, v)
			println(">", s.NR, v.String())
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

$ seq 12 | xargs -n 3  | go run SkipLine2.go
. 1 1 2 3
> 2 4 5 6
> 3 7 8 9
. 4 10 11 12
Passed.

*/
```

#### > Wp2Hugo.go
```go
////////////////////////////////////////////////////////////////////////////
// Porgram: Wp2Hugo
// Purpose: From wordpress meta to Hugo's
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

import (
	"github.com/spakin/awk"
)

/*

Wp2Hugo converts wordpress's markdown meta data format into Hugo's.

E.g., for an input of wordpress's markdown like this,

    # Dbab From Start To Finish

    [category Tech][tags Debian,Ubuntu,Linux,DHCP,DNS,WPAD,dnsmasq,dbab]

The output Hugo's meta data is:

    ---
    title: "Dbab From Start To Finish"
    date: "2015-12-06T09:57:45-05:00"
    categories: ["Tech"]
    tags: ["Debian","Ubuntu","Linux","DHCP","DNS","WPAD","dnsmasq","dbab"]
    ---

Usage:

    Wp2Hugo < wordpress.md > path/to/hugo.md

*/

func main() {
	s := awk.NewScript()

	s.Begin = func(s *awk.Script) {
		s.SetFS("[][ ]")
	}
	s.State = s.NewValue("")

	// NR == 1 & $0 ~ /^# +/ { State = $0 }
	s.AppendStmt(func(s *awk.Script) bool { return s.NR == 1 && s.F(0).Match("^# +") },
		func(s *awk.Script) {
			s.State = s.F(0)
			s.Next()
		})

	// /^\[category/ { convert meta data format }
	s.AppendStmt(func(s *awk.Script) bool { return s.F(0).Match("^\\[category") },
		func(s *awk.Script) {
			re := regexp.MustCompile("^# +(.*)$")
			fmt.Println(re.ReplaceAllString(s.State.(*awk.Value).String(),
				"---\ntitle: \"$1\""))
			fmt.Printf("date: \"%s\"\n", time.Now().Format(time.RFC3339))
			fmt.Printf("categories: [\"%s\"]\n", s.F(3))
			fmt.Printf("%s: [%s]\n", s.F(5), quoteEach(s.F(6).String()))
			fmt.Println("---\n")
			//s.Exit()
			s.Next()
		})

	// 1; # i.e., print all
	s.AppendStmt(nil, nil)

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

func quoteEach(tags string) string {
	t := strings.Split(tags, ",")
	for i, tag := range t {
		t[i] = "\"" + tag + "\""
	}
	return strings.Join(t, ",")
}
```

#### > Comments.go
```go
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
	s.MaxRecordSize = 24 * 1024 * 1024
	s.MaxFieldSize = 24 * 1024 * 1024

	s.AppendStmt(func(s *awk.Script) bool {
		return s.F(1).Match("<Comment")
	}, func(s *awk.Script) {})
	// 1; # i.e., print all
	s.AppendStmt(nil, nil)

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}
```
