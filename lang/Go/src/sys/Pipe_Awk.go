////////////////////////////////////////////////////////////////////////////
// Purpose: Go pipeline + awk demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: http://labix.org/pipe, https://github.com/spakin/awk/
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

import (
	"github.com/spakin/awk"
	"gopkg.in/pipe.v2"
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Function main

func main() {
	TestExtending()
	fmt.Printf("\n--\n")
	TestWrite()
}

// Awk1 writes the first two fields in opposite order if the second field does
// not match the regular expression "4:" and the first field does not match "7"
// (AWK: $2 !~ /4:/ && $1 !~ /7/ {print $2, $1}).
func Awk1() pipe.Pipe {
	return pipe.TaskFunc(func(st *pipe.State) error {
		// == Setup
		s := awk.NewScript()
		s.Output = st.Stdout

		// == Transform
		s.AppendStmt(func(s *awk.Script) bool {
			return !s.F(2).Match("4:") && !s.F(1).Match("7")
		}, func(s *awk.Script) { s.Println(s.F(2), s.F(1)) })

		// == Run it
		return s.Run(st.Stdin)
	})
}

func TestExtending() {
	p := pipe.Line(
		pipe.System("seq 8 | cat -n | sed 's/$/:/'"),
		Awk1(),
	)
	output, err := pipe.Output(p)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%s", output)
}

func TestWrite() {
	p := pipe.Line(
		pipe.System("seq 8 | cat -n | sed 's/$/:/'"),
		Awk1(),
		pipe.WriteFile("awk.tmp", 0644),
	)
	err := pipe.Run(p)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

/*

Output:

1: 1
2: 2
3: 3
5: 5
6: 6
8: 8

*/
