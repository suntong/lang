////////////////////////////////////////////////////////////////////////////
// Porgram: FileRW
// Purpose: Go 1-compatible list of all the ways to read and write files in Go
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

/*
 Originally From: http://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file
 Author: Mostafa

 To run,
 seq 20 > input.txt
 go run FileRW-ioutil.go

 */

package main

import (
    "io/ioutil"
)

func main() {
    // read whole the file
    b, err := ioutil.ReadFile("input.txt")
    if err != nil { panic(err) }

    // write whole the body
    err = ioutil.WriteFile("output.txt", b, 0644)
    if err != nil { panic(err) }
}