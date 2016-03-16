////////////////////////////////////////////////////////////////////////////
// Purpose: Go Rich pipeline demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: http://labix.org/pipe
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

import (
	"gopkg.in/pipe.v2"
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Function main

func main() {
	TestRich()
	fmt.Printf("\n--\n")
	TestRicher()
	fmt.Printf("\n--\n")
	TestExtending()
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Rich pipeline

/*

TestRich grabs the free space information for the /dev* mount points, and
writes it both to a file named "dev.txt" in the local directory, and to an
in-memory buffer. It would be more easily implemented via pipe.TeeFile, but
this shows more clearly the flexibility of the system.

*/
func TestRich() {
	b := &bytes.Buffer{}
	p := pipe.Script(
		pipe.Line(
			pipe.Exec("df"),
			pipe.Filter(func(line []byte) bool {
				return bytes.Contains(line, []byte(" /dev"))
			}),
			pipe.Tee(b),
			pipe.WriteFile("dev.tmp", 0644),
		),
		pipe.RenameFile("dev.tmp", "dev.txt"),
	)
	err := pipe.Run(p)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Print(b.String())
}

/*

TestRicher demonstrates that concept being used in a richer pipe. It outputs
a passwd line for the root user, and then streams all the content from the
/etc/passwd file containing string ":x:3". The result is then streamed to
os.Stdout.

*/
func TestRicher() {
	prefix := []byte(":x:3")
	script := pipe.Script(
		pipe.Println("root:x:0:0:root:/root:/bin/sh"),
		pipe.Line(
			pipe.ReadFile("/etc/passwd"),
			pipe.Filter(func(line []byte) bool {
				return bytes.Contains(line, prefix)
			}),
		),
	)
	p := pipe.Line(
		script,
		pipe.Write(os.Stdout),
	)
	err := pipe.Run(p)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Extending with custom logic

// Echo implements a trivial echo-like function
func Echo(str string) pipe.Pipe {
	return pipe.TaskFunc(func(s *pipe.State) error {
		_, err := s.Stdout.Write([]byte(str))
		return err
	})
}

// MyReplace filters lines read from the pipe's stdin and writes
// the returned values to stdout.
func MyReplace(f func(line []byte) []byte) pipe.Pipe {
	return pipe.TaskFunc(func(s *pipe.State) error {
		r := bufio.NewReader(s.Stdin)
		for {
			line, err := r.ReadBytes('\n')
			if len(line) > 0 {
				line := f(line)
				if len(line) > 0 {
					_, err := s.Stdout.Write(line)
					if err != nil {
						return err
					}
				}
			}
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
		}
		panic("unreachable")
	})
}

func TestExtending() {
	p := pipe.Line(
		Echo("out1\nout2\nout3"),
		MyReplace(func(line []byte) []byte {
			if bytes.HasPrefix(line, []byte("out")) {
				// skip line 2
				if line[3] == '2' {
					return nil
				}
				// replace other lines
				return []byte{'l', line[3], ','}
			}
			return line
		}),
	)
	output, err := pipe.Output(p)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%s", output)
}
