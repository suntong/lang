package main

import (
	"fmt"
	"os"
	"time"
)

import (
	"github.com/voxelbrain/goptions"
)

type Execute struct {
	Command string   `goptions:"-c, --command, mutexgroup='input', description='Command to exectute', obligatory"`
	Script  *os.File `goptions:"--script, mutexgroup='input', description='Script to exectute', rdonly"`
	Fo      *os.File `goptions:"-o, --output, description='The output', wronly"`
	Force   bool     `goptions:"-f, --force, description='Force removal'"`
	Check   string   `goptions:"--check, description='Check str'"`
}

type Options struct {
	Server    string        `goptions:"-s, --server, description='Server to connect to'"`
	Password  string        `goptions:"-p, --password, description='Don\\'t prompt for password'"`
	Timeout   time.Duration `goptions:"-t, --timeout, description='Connection timeout in seconds'"`
	Verbosity []bool        `goptions:"-v, --verbose, description='Be verbose'"`
	Quiet     bool          `goptions:"-q, --quiet, description='Do not print anything, even errors (except if --verbose is specified)'"`
	Help      goptions.Help `goptions:"-h, --help, description='Show this help'"`

	goptions.Verbs
	Execute `goptions:"execute"` // Embedding!
	Delete  struct {
		Path  string `goptions:"-n, --name, obligatory, description='Name of the entity to be deleted'"`
		Force bool   `goptions:"-f, --force, description='Force removal'"`
	} `goptions:"delete"`
}

var options = Options{ // Default values goes here
	Timeout: 10 * time.Second,
	Execute: Execute{
		Check: "something",
	},
}

type Command func() error

var commands = map[goptions.Verbs]Command{
	"execute": executecmd,
	"delete":  deletecmd,
}

var (
	VERBOSITY = 0
)

func main() {
	goptions.ParseAndFail(&options)

	if len(options.Verbs) == 0 {
		goptions.PrintHelp()
		os.Exit(2)
	}

	VERBOSITY = len(options.Verbosity)

	if cmd, found := commands[options.Verbs]; found {
		err := cmd()
		if err != nil {
			if !options.Quiet {
				fmt.Println("error:", err)
			}
			os.Exit(1)
		}
	}
}

func executecmd() error {
	fmt.Printf("Selected verb: %s\n", options.Verbs)
	fmt.Printf("Execute.Command: %s\n", options.Execute.Command)
	fmt.Printf(" with verbosity: %d\n", VERBOSITY)
	//options.Execute.Check = "something else"
	if options.Execute.Fo != nil {
		fmt.Fprintf(options.Execute.Fo, "To output, Check str: '%s'\n",
			options.Execute.Check)
		fmt.Fprintf(os.Stdout, "To os.Stdout, Check str: '%s'\n",
			options.Execute.Check)
	}
	options.Execute.Force = true
	fmt.Printf("Force: %v\n", options.Execute.Force)
	testMore()
	return nil
}

func deletecmd() error {
	return nil
}

func testMore() {
	fmt.Printf("Force: %v\n", options.Execute.Force)
}
