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
 go run FileRW-bufio.go

 */

package main

import (
    "bufio"
    "io"
    "os"
)

func main() {
    // open input file
    fi, err := os.Open("input.txt")
    if err != nil { panic(err) }
    // close fi on exit and check for its returned error
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()
    // make a read buffer
    r := bufio.NewReader(fi)

    // open output file
    fo, err := os.Create("output.txt")
    if err != nil { panic(err) }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
    // make a write buffer
    w := bufio.NewWriter(fo)

    // make a buffer to keep chunks that are read
    buf := make([]byte, 1024)
    for {
        // read a chunk
        n, err := r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }

        // write a chunk
        if _, err := w.Write(buf[:n]); err != nil {
            panic(err)
        }
    }

    if err = w.Flush(); err != nil { panic(err) }
}