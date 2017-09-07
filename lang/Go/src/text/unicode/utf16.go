////////////////////////////////////////////////////////////////////////////
// Porgram: utf16
// Purpose: Go utf16 string/file reading demo, convert utf16 to string
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open("utf16.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)

	// `ReadString` returns the next string from the
	// input up to the given separator byte.
	// `ReadLine` will read input line-by-line
	for {
		switch line, _, err := r.ReadLine(); err {

		// If the read succeeded (the read `err` is nil),
		case nil:
			// Output utf16 bytes array
			//fmt.Println(line)
			// Convert it to string, i.e., treat the utf16 bytes as ascii
			fmt.Println(string(line))

		// The `EOF` error is expected when we reach the
		// end of input, so exit gracefully in that case.
		case io.EOF:
			os.Exit(0)

		// Otherwise there's a problem; print the
		// error and exit with non-zero status.
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}

}
