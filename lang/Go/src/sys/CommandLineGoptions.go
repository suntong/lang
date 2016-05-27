////////////////////////////////////////////////////////////////////////////
// Porgram: CommandLineGoptions
// Purpose: Go goptions command line options handling template
// authors: Tong Sun (c) 2015, All rights reserved
// Credits: https://github.com/voxelbrain/goptions/tree/master/examples
//          https://github.com/jwilder/github-release/blob/master/github-release.go
//          https://github.com/daaku/summon/blob/master/summon.go
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"time"
)

import (
	"github.com/voxelbrain/goptions"
)

type Options struct {
	Server    string        `goptions:"-s, --server, description='Server to connect to'"`
	Password  string        `goptions:"-p, --password, description='Don\\'t prompt for password'"`
	Timeout   time.Duration `goptions:"-t, --timeout, description='Connection timeout in seconds'"`
	Verbosity []bool        `goptions:"-v, --verbose, description='Be verbose'"`
	Quiet     bool          `goptions:"-q, --quiet, description='Do not print anything, even errors (except if --verbose is specified)'"`
	Help      goptions.Help `goptions:"-h, --help, description='Show this help\n\nSub-commands (Verbs):\n\n\texecute\t\tExecute it\n\t\t\tExecute the given command\n\n\tdelete\t\tDelete it'"`

	goptions.Verbs
	Execute struct {
		Command string   `goptions:"-c, --command, mutexgroup='input', description='Command to exectute', obligatory"`
		Script  *os.File `goptions:"--script, mutexgroup='input', description='Script to exectute', rdonly"`
	} `goptions:"execute"`
	Delete struct {
		Path  string `goptions:"-n, --name, obligatory, description='Name of the entity to be deleted'"`
		Force bool   `goptions:"-f, --force, description='Force removal'"`
	} `goptions:"delete"`
}

var options = Options{ // Default values goes here
	Timeout: 10 * time.Second,
}

type Command func(Options) error

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
		err := cmd(options)
		if err != nil {
			if !options.Quiet {
				fmt.Println("error:", err)
			}
			os.Exit(1)
		}
	}
}

func executecmd(options Options) error {
	fmt.Printf("Selected verb: %s\n", options.Verbs)
	fmt.Printf("Execute.Command: %s\n", options.Execute.Command)
	fmt.Printf(" with verbosity: %d\n", VERBOSITY)
	return nil
}

func deletecmd(opt Options) error {
	return nil
}

/*

$ go run CommandLineGoptions.go
Usage: CommandLineGoptions [global options] <verb> [verb options]

Global options:
        -s, --server   Server to connect to
        -p, --password Don't prompt for password
        -t, --timeout  Connection timeout in seconds (default: 10s)
        -v, --verbose  Be verbose
        -q, --quiet    Do not print anything, even errors (except if --verbose is specified)
        -h, --help     Show this help

Verbs:
    delete:
        -n, --name     Name of the entity to be deleted (*)
        -f, --force    Force removal
    execute:
        -c, --command  Command to exectute (*)
            --script   Script to exectute

$ go run CommandLineGoptions.go execute -c 'test'
Selected verb: execute
Execute.Command: test
 with verbosity: 0

$ go run CommandLineGoptions.go -v execute -c 'test'
Selected verb: execute
Execute.Command: test
 with verbosity: 1

$ go run CommandLineGoptions.go -v -v execute -c 'test'
command-line-arguments
Selected verb: execute
Execute.Command: test
 with verbosity: 2

$ go run CommandLineGoptions.go execute -c 'test' --script CommandLineGoptions.go
Error: Exactly one of --command, --script must be specified
Usage: CommandLineGoptions ...

$ go run CommandLineGoptions.go cmd_verb
Error: Invalid trailing arguments: [cmd_verb]
Usage: CommandLineGoptions ...

*/
