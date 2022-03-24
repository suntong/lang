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

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Printf("%s [%s:%s] %s", time.Now().Format(logFmt),
		prog, writer.curr, string(bytes))
}

func main() {
	l := logWriter{"main"}
	log.SetFlags(0)
	log.SetOutput(l)
	log.Println("This is something being logged!")
	time.Sleep(3 * time.Second) // simulate slow operation by sleeping
	log.Println("This is something to be logged next.")
}

/*

$ go run LogWriter.go
16:31:02 [LogWriter] This is something being logged!
16:31:05 [LogWriter] This is something to be logged next.

*/
