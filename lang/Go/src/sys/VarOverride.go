package main

import "fmt"

type personT struct {
	Name string
	Age  int
}

var (
	thePerson        = personT{"Bob", 20}
	theTime          = "2016-05-26"
	name      string = "Sean"
	age       int    = 30
)

func main() {

	fmt.Println(name, thePerson)
	fmt.Println(theTime)
}

/*

https://blog.cloudflare.com/setting-go-variables-at-compile-time/

You can use go run (or other build commands like go build or go install)
with the -ldflags option to modify the value of the global variables:

However,

    $ go run VarOverride.go
    Sean {Bob 20}
    2016-05-26

    $ go run  -ldflags="-X main.theTime=`date -I`" VarOverride.go
    Sean {Bob 20}
    2016-05-27

    $ go run  -ldflags="-X main.theTime=`date -I` -X main.thePerson.Age=35" VarOverride.go
    Sean {Bob 20}
    2016-05-27

    $ go run  -ldflags="-X main.thePerson.Name=Fred" VarOverride.go
    Sean {Bob 20}
    2016-05-26

    $ go run  -ldflags="-X main.thePerson.XX=Fred" VarOverride.go
    Sean {Bob 20}
    2016-05-26

    $ go run -ldflags="-X main.name=Fred" VarOverride.go
    Fred {Bob 20}
    2016-05-26

    $ go run -ldflags="-X main.name=Fred -X main.theTime=`date -I`" VarOverride.go
    Fred {Bob 20}
    2016-05-27

    $ go run -ldflags="-X main.name=Fred -X main.theTime=`date -I` -X main.age=40" VarOverride.go
    # command-line-arguments
    cannot use -X with non-string symbol main.age

The format is importpath.name string, so it's possible to set the value of
any string anywhere in the Go program, not just in main.

NB, haven't found a way to specify values with space yet.
I.e., while `go run -ldflags="-X main.name=Fred" VarOverride.go` works,
the following don't:

    go run -ldflags="-X main.name='John Doe'" VarOverride.go
    go run -ldflags="-X main.name='John\ Doe'" VarOverride.go
    go run -ldflags="-X main.name=John\ Doe" VarOverride.go
    go run -ldflags="-X main.name=\"John Doe\"" VarOverride.go

*/
