////////////////////////////////////////////////////////////////////////////
// Porgram: Logger.go
// Purpose: Logging to a file with logger
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://godoc.org/log
//  http://stackoverflow.com/questions/27629380/how-to-exit-a-go-program-honoring-deferred-calls
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

////////////////////////////////////////////////////////////////////////////

type Exit struct{ Code int }

// exit code handler
func handleExit() {
	// open & write to log file
	f, err := os.OpenFile("LogToFile.log",
		os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	// don't forget to close it
	defer f.Close()

	fmt.Fprint(f, &buf)

	if e := recover(); e != nil {
		if exit, ok := e.(Exit); ok == true {
			os.Exit(exit.Code)
		}
		panic(e) // not an Exit, bubble up
	}
}

////////////////////////////////////////////////////////////////////////////

var buf bytes.Buffer
var logger *log.Logger

func init() {
	logger = log.New(&buf, "logger: ", log.LstdFlags|log.Lshortfile)
}

func main() {
	defer handleExit() // plug the exit handler

	log.Print("Hello, log file!")
	log.Output(1, "this is an event")

	logger.Print("Hello, log file!")
	logger.Output(1, "this is an event")

	//panic(Exit{3}) // 3 is the exit code
}
