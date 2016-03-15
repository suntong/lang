////////////////////////////////////////////////////////////////////////////
// Purpose: Go pipe demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: http://labix.org/pipe
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

import (
	"gopkg.in/pipe.v2"
)

func main() {
	TestLine()
	fmt.Printf("\n--\n")
	TestLineNesting()
	fmt.Printf("\n--\n")
	TestFilter()
	fmt.Printf("\n--\n")
	TestReplace()
	fmt.Printf("\n--\n")
	TestSetEnvVar()
}

func TestLine() {
	p := pipe.Line(
		//pipe.ReadFile("article.txt"),
		//pipe.Read(bytes.NewBufferString("hello")),
		pipe.Exec("/bin/sh", "-c", "echo out1; echo err1 1>&2; echo out2; echo err2 1>&2"),
		pipe.Exec("sed", `s/\(...\)\([12]\)/\1-\2/`),
	)
	output, err := pipe.CombinedOutput(p)
	checkError(err)
	fmt.Printf("%s", output)
}

func TestLineNesting() {
	b := &bytes.Buffer{}
	p := pipe.Line(
		pipe.Print("hello"),
		pipe.Line(
			pipe.Filter(func(line []byte) bool { return true }),
			pipe.Exec("sed", "s/l/k/g"),
		),
		pipe.Write(b),
	)
	err := pipe.Run(p)
	checkError(err)
	fmt.Printf("%v", b)
}

func TestFilter() {
	p := pipe.Line(
		pipe.System("echo out1; echo err1 1>&2; echo out2; echo err2 1>&2; echo out3"),
		pipe.Filter(func(line []byte) bool { return string(line) != "out2" }),
	)
	output, err := pipe.Output(p)
	checkError(err)
	fmt.Printf("%s", output)
}

func TestReplace() {
	p := pipe.Line(
		pipe.System("echo out1; echo err1 1>&2; echo out2; echo err2 1>&2; echo out3"),
		pipe.Replace(func(line []byte) []byte {
			if bytes.HasPrefix(line, []byte("out")) {
				if line[3] == '3' {
					return nil
				}
				return []byte{'l', line[3], ','}
			}
			return line
		}),
	)
	output, err := pipe.Output(p)
	checkError(err)
	fmt.Printf("%s", output)
}

func TestSetEnvVar() {
	os.Setenv("PIPE_NEW_VAR", "")
	os.Setenv("PIPE_OLD_VAR", "old")
	defer os.Setenv("PIPE_OLD_VAR", "")
	p := pipe.Script(
		pipe.SetEnvVar("PIPE_NEW_VAR", "new"),
		pipe.System("echo $PIPE_OLD_VAR $PIPE_NEW_VAR"),
		pipe.SetEnvVar("PIPE_NEW_VAR", "after"),
		func(s *pipe.State) error {
			count := 0
			prefix := "PIPE_NEW_VAR="
			for _, kv := range s.Env {
				if strings.HasPrefix(kv, prefix) {
					count++
				}
			}
			if count != 1 {
				return fmt.Errorf("found %d environment variables", count)
			}
			return nil
		},
	)
	output, err := pipe.Output(p)
	checkError(err)
	fmt.Printf("%s", output)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

/*

Output:

err1
err2
out-1
out-2

--
hekko
--
out1
out3

--
l1,l2,
--
old new


*/
