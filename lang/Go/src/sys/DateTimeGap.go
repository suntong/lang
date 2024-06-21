////////////////////////////////////////////////////////////////////////////
// Program: DateTimeGap.go
// Purpose: DateTime Gap Calculation
// Authors: Tong Sun (c) 2024, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const (
	progname = "DateTimeGap"
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	timeStrFmt = "2006-01-02 15:04:05.000000"
	timeStrLen = 26
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {

	// the 1st & 2nd command line argument is to overide timeStrFmt & Len
	if len(os.Args) > 1 {
		//println("The first command line argument: ", os.Args[1])
		timeStrFmt = os.Args[1]
	}
	var err error
	if len(os.Args) > 2 {
		timeStrLen, err = strconv.Atoi(os.Args[2])
		abortOn("parsing 2nd command line argument", err)
	}

	firstLine := true
	scanner := bufio.NewScanner(os.Stdin)
	var prevTime time.Time

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < timeStrLen {
			// Skip lines that are too short to contain a timestamp
			continue
		}

		// Parse the timestamp
		timestampStr := line[:timeStrLen]
		timestamp, err := time.Parse(timeStrFmt, timestampStr)
		abortOn("parsing timestamp: "+timestampStr, err)

		if !firstLine {
			// Calculate and print the time gap
			timeGap := timestamp.Sub(prevTime)
			fmt.Printf("%10v ", timeGap)
		} else {
			firstLine = false
			fmt.Printf("%10v ", " ")
		}

		// Print the full log line
		fmt.Println(line)
		prevTime = timestamp
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

//==========================================================================

// abortOn will quit on anticipated errors gracefully without stack trace
func abortOn(errCase string, e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "[%s] Error, %s: %v\n",
			progname, errCase, e)
		os.Exit(1)
	}
}

// go build -o DateTimeGap DateTimeGap.go
