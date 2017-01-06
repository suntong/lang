////////////////////////////////////////////////////////////////////////////
// Porgram: LogToFile.go
// Purpose: Logging to a file in golang
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: http://stackoverflow.com/questions/32619318/
////////////////////////////////////////////////////////////////////////////

// https://godoc.org/log

package main

import (
	"log"
	"os"
)

func main() {
	// open a file
	f, err := os.OpenFile("LogToFile.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}

	// don't forget to close it
	defer f.Close()

	// assign it to the standard logger
	log.SetOutput(f)

	log.Output(1, "this is an event")

}
