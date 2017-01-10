////////////////////////////////////////////////////////////////////////////
// Purpose: selfLog will keep a local log (instead of using syslog)
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits:
//   http://stackoverflow.com/questions/32619318/logging-to-a-file-in-golang
//   http://stackoverflow.com/questions/27629380/#27630092
////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Open log file in init()

var logF os.File

func init() {
	fmt.Println("Logging to " + os.Args[0] + ".log")
	logF, err := os.OpenFile(os.Args[0]+".log",
		os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	log.SetOutput(logF)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// don't forget to close it, but can't do it here
	// defer logF.Close()

}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Close log file in final()

/*

To exit a go program while honoring deferred calls

The problem is that os.Exit skips any deferred instruction.
So to exit a go program honoring declared defer calls, we have to panic,
because by nature, it will honor defer calls.
But will also always exit with non 0 status code and dump a stack trace.

To exit, use

  panic(Exit{3}) // 3 is the exit code

*/

type Exit struct{ Code int }

// safe exit handler
func final() {
	fmt.Println("exiting...")
	defer logF.Close()

	if e := recover(); e != nil {
		if exit, ok := e.(Exit); ok == true {
			os.Exit(exit.Code)
		}
		log.Panic(e) // not an Exit, bubble up
	}
}

func main() {
	defer final() // plug the exit handler
	defer fmt.Println("cleaning...")
	log.Print("Hello, log file!")
	panic(errors.New("Simlating a panic"))
	panic(Exit{3}) // 3 is the exit code
}
