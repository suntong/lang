package main

import (
	"fmt"
	"os"
	"time"
)

import (
	"github.com/voxelbrain/goptions"
)

var (
	options = struct {
		Server   string        `goptions:"-s, --server, obligatory, description='Server to connect to'"`
		Password string        `goptions:"-p, --password, description='Don\\'t prompt for password'"`
		Timeout  time.Duration `goptions:"-t, --timeout, description='Connection timeout in seconds'"`
		Help     goptions.Help `goptions:"-h, --help, description='Show this help'"`

		Verb    goptions.Verbs
		Execute struct {
			Command string   `goptions:"-c, --command, mutexgroup='input', description='Command to exectute', obligatory"`
			Script  *os.File `goptions:"--script, mutexgroup='input', description='Script to exectute', rdonly"`
		} `goptions:"execute"`
		Delete struct {
			Path  string `goptions:"-n, --name, obligatory, description='Name of the entity to be deleted'"`
			Force bool   `goptions:"-f, --force, description='Force removal'"`
		} `goptions:"delete"`
	}{ // Default values goes here
		Timeout: 10 * time.Second,
	}
)

func main() {
	goptions.ParseAndFail(&options)
	fmt.Printf("Selected verb: %s\n", options.Verb)
	fmt.Printf("Execute.Command: %s\n", options.Execute.Command)
}

/*

$ go run CommandLineGoptions.go
Error: --server must be specified
Usage: CommandLineGoptions [global options] <verb> [verb options]

Global options:
        -s, --server   Server to connect to (*)
        -p, --password Don't prompt for password
        -t, --timeout  Connection timeout in seconds (default: 10s)
        -h, --help     Show this help

Verbs:
    delete:
        -n, --name     Name of the entity to be deleted (*)
        -f, --force    Force removal
    execute:
        -c, --command  Command to exectute (*)
            --script   Script to exectute
exit status 1

$ go run CommandLineGoptions.go -s sss execute -c 'test'
Selected verb: execute
Execute.Command: test

$ go run CommandLineGoptions.go -s sss execute -c 'test' --script CommandLineGoptions.go
Error: Exactly one of --command, --script must be specified
Usage: CommandLineGoptions ...

*/
