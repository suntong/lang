////////////////////////////////////////////////////////////////////////////
// Porgram: Defer2.go
// Purpose: Demo defer and error recovering
// Authors: Tong Sun (c) 2018, All rights reserved
// Credits: https://play.golang.org/p/evabhcjvNs
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var ok = "OK"

func main() {
	log.SetPrefix("Defer2: ")
	log.SetFlags(0)

	defer final() // plug the exit handler
	defer trace("main")()
	if len(os.Args) > 1 {
		ok = os.Args[1]
	}

	if ok == "OK" {
		test1()
	}
	if ok == "1" {
		test1()
	}
	if ok == "2" {
		test2()
	}
	if ok == "3" {
		test3()
	}
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// log.Fatalf will abandon any deferred instructions.
func test1() {
	defer trace("test")()
	r, err := Open("a")
	if err != nil {
		log.Fatalf("error opening 'a'\n")
	}
	defer r.Close()

	r.Use()
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// panic will honor defer calls, but will print strack trace
func test2() {
	defer trace("test")()
	r, err := Open("a")
	if err != nil {
		panic("error opening 'a'\n")
	}
	defer r.Close()

	r.Use()
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// to honor defer calls, but also exit gracefully without strack trace
func test3() {
	defer trace("test")()
	r, err := Open("a")
	if err != nil {
		log.Printf("%v\n", err)
		panic(Exit{3})
	}
	defer r.Close()

	r.Use()
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	time.Sleep(1 * time.Second) // simulate slow
	return func() {
		log.Printf("exit %s in %s", msg, time.Since(start))
	}
}

type Exit struct{ Code int }

// safe exit handler
func final() {
	if e := recover(); e != nil {
		if exit, ok := e.(Exit); ok == true {
			os.Exit(exit.Code)
		}
		log.Panic(e) // not an Exit, bubble up
	}
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
type Resource struct {
	name string
}

func Open(name string) (*Resource, error) {
	log.Printf("opening %s\n", name)
	if ok == "OK" {
		return &Resource{name}, nil
	}
	return nil, fmt.Errorf("opening '%s' failed", name)
}

func (r *Resource) Use() error {
	log.Printf("using %s\n", r.name)
	return nil
}

func (r *Resource) Close() error {
	log.Printf("closing %s\n", r.name)
	return nil
}

/*

$ go run Defer2.go
Defer2: enter main
Defer2: enter test
Defer2: opening a
Defer2: using a
Defer2: closing a
Defer2: exit test in 1.000289555s
Defer2: exit main in 2.001040054s

// log.Fatalf will abandon any deferred cleanup instructions.
$ go run Defer2.go 1
Defer2: enter main
Defer2: enter test
Defer2: opening a
Defer2: error opening 'a'
exit status 1

// panic will honor defer calls, but will print strack trace
$ go run Defer2.go 2
Defer2: enter main
Defer2: enter test
Defer2: opening a
Defer2: exit test in 1.000222516s
Defer2: exit main in 2.000471549s
Defer2: error opening 'a'
panic: error opening 'a'
...
goroutine 1 [running]:
log.Panic(0xc42004be20, 0x1, 0x1)
...
exit status 2

// to honor defer calls, but also exit gracefully without strack trace
$ go run Defer2.go 3
Defer2: enter main
Defer2: enter test
Defer2: opening a
Defer2: opening 'a' failed
Defer2: exit test in 1.000314955s
Defer2: exit main in 2.000555199s
exit status 3

*/
