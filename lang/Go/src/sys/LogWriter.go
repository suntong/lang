// Custom log date/time format
// https://stackoverflow.com/questions/26120698/

package main

import (
	"fmt"
	"log"
	"time"
)

type logWriter struct {
	curr string
}

const prog = "LogWriter"

//const logFmt = "2006-01-02T15:04:05.999Z"
const logFmt = "15:04:05"

var l = logWriter{"main"}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Printf("%s [%s:%s] %s", time.Now().Format(logFmt),
		prog, writer.curr, string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(l)
	log.Println("This is something being logged!")
	sub1()
	log.Println("This is something to be logged in the end.")
}

func sub1() {
	defer track("sub1")() // don't forget the extra parentheses
	log.Println(" working in sub1...")
	time.Sleep(3 * time.Second) // simulate slow operation by sleeping
}

func track(curr string) func() {
	prev := l.curr
	l.curr = curr
	log.SetOutput(l) // must!
	log.Printf("%s:", curr)
	return func() {
		log.Printf("%s.", curr)
		l.curr = prev
		log.SetOutput(l)
	}
}

/*

$ go run LogWriter.go
16:50:38 [LogWriter:main] This is something being logged!
16:50:38 [LogWriter:sub1] sub1:
16:50:38 [LogWriter:sub1]  working in sub1...
16:50:41 [LogWriter:sub1] sub1.
16:50:41 [LogWriter:main] This is something to be logged in the end.

*/
