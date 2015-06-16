////////////////////////////////////////////////////////////////////////////
// Porgram: ExecExamples
// Purpose: Go exec examples
// Authors: Tong Sun (c) 2015, All rights reserved
// based on https://golang.org/src/os/exec/example_test.go
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Testing started")
	ExampleLookPath()
	ExampleCommand()
	ExampleCmd_Output()
	ExampleCmd_Start()
	ExampleCmd_StdoutPipe()
	fmt.Println("Testing done")
}

func ExampleLookPath() {
	path, err := exec.LookPath("fortune")
	if err != nil {
		log.Printf("installing fortune is in your future")
		out, err := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The echo result is %s\n", out)
		return
	}
	fmt.Printf("fortune is available at %s\n", path)
}

func ExampleCommand() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}

func ExampleCmd_Output() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}

func ExampleCmd_Start() {
	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

func ExampleCmd_StdoutPipe() {
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}
