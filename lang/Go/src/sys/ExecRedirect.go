////////////////////////////////////////////////////////////////////////////
// Porgram: ExecRedirect.go
// Purpose: Go exec redirect examples
// Authors: Tong Sun (c) 2017, All rights reserved
// Credict: as listed below
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/suntong/testing"
)

func main() {
	redirect1()

	var t *testing.T = testing.NewT()
	redirect2(t)
}

/*

https://stackoverflow.com/questions/18986943/

run a shell command, capture stdout and write that output to a file

*/

func redirect1() {

	cmd := exec.Command("echo", "'WHAT THE HECK IS UP'", "Huh?")

	// open the out file for writing
	outfile, err := os.Create("./out.txt")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	cmd.Wait()
}

/*

https://github.com/golang/go/issues/11866

The > is a feature of your shell language. To use it, you need to run your
command inside a shell. Perhaps something like

cmd := exec.Command("bash", "-c", "echo hello > /tmp/aaa")

Be careful with this though. If you construct a string and pass it to a
shell, you need to think carefully about escaping and security. For the
particular example you give, you may want to consider using os.Open,
os.Create, and io.Copy.


*/

func redirect2(t *testing.T) {

	out := &bytes.Buffer{}

	cmd := exec.Command("bash", "-c", "echo Hello~~~ > /tmp/aaa")
	err := cmd.Start()
	if err != nil {
		t.Errorf("start error %s [%s]", err, out.String())
	}
	err = cmd.Wait()
	if err == nil {
		t.Errorf("exit error %s [%s]", err, out.String())
	}

	cmd = exec.Command("bash", "-c", "cat ./out.txt >> /tmp/aaa")
	err = cmd.Start()
	if err != nil {
		t.Errorf("start error %s [%s]", err, out.String())
	}
	err = cmd.Wait()
	if err != nil {
		t.Errorf("exit error %s [%s]", err, out.String())
	}
}
