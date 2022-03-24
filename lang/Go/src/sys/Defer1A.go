////////////////////////////////////////////////////////////////////////////
// Porgram: Defer2.go
// Purpose: Demo defer that returns a function to the caller
// Authors: Tong Sun (c) 2022, All rights reserved
// Credits: gopl.io/ch5/trace
////////////////////////////////////////////////////////////////////////////

package main

import (
	"log"
	"time"
)

func main() {
	log.SetPrefix("Defer1A: ")
	//log.SetFlags(0)
	log.SetFlags(log.Ltime)

	bigSlowOperation()
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work…
	time.Sleep(5 * time.Second) // simulate slow operation by sleeping
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
func trace(msg string) func() {
	start := time.Now()
	log.Printf("%s:", msg)
	return func() {
		log.Printf("%s. (%s)", msg, time.Since(start))
	}
}

/*

$ go run Defer1A.go
Defer1A: 16:17:19 bigSlowOperation:
Defer1A: 16:17:24 bigSlowOperation. (5.004292356s)

*/
