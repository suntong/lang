////////////////////////////////////////////////////////////////////////////
// Porgram: FileRW
// Purpose: Go 1-compatible list of all the ways to read and write files in Go
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2 -w

/*

 To run,
 go run FileRW-Fprintf.go
 cat output.txt

*/

package main

import (
  "fmt"
  "os"
)

func main() {

  // open output file
  file, err := os.Create("output.txt")
  if err != nil {
    panic(err)
  }
  // close file on exit and check for its returned error
  defer func() {
    if err := file.Close(); err != nil {
      panic(err)
    }
  }()

  file.WriteString("Test starts.\n\n")

  for i := 0; i <= 20; i++ {
    fmt.Fprintf(file, "%02d\n", i)
  }
  
  file.WriteString("\nTest done.\n")
}
