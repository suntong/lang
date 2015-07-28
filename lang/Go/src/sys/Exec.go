package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	app := "echo"
	//app := "buah"

	// the first command line arguments
	if len(os.Args) > 1 {
		app = os.Args[1]
	}

	arg0 := "-e"
	arg1 := "Hello world"
	arg2 := "\n\tfrom"
	arg3 := "golang"

	// From http://stackoverflow.com/questions/6182369/exec-a-shell-command-in-go
	{
		cmd := exec.Command(app, arg0, arg1, arg2, arg3)
		out, err := cmd.Output()

		if err != nil {
			println(err.Error())
			return
		}

		print(string(out))
	}

	// Based on http://golang.org/pkg/os/exec/#example_Command
	{
		args2 := []string{
			"-c:a", "libopus",
			"-c:v", "libx265", "-x265-params",
			"me=star:subme=7:bframes=16:b-adapt=2:ref=16",
		}

		args := append([]string{},
			"ffmpeg", "-i",
			"Harry Potter and the Philosopher's Stone",
		)
		for _, a := range args2 {
			args = append(args, a)
		}

		cmd := exec.Command(app, args...)
		cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n== Out:\n%s\n", out.String())
	}

}
