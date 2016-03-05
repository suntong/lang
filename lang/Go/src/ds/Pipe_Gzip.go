////////////////////////////////////////////////////////////////////////////
// Porgram: Pipe_Gzip.go
// Purpose: io.Pipe demo with base64 & gzip
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://www.socketloop.com/references/golang-io-pipe-function-example
////////////////////////////////////////////////////////////////////////////

package main

import (
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	// read and write with pipe
	pReader, pWriter := io.Pipe()

	// use base64 encoder
	b64Writer := base64.NewEncoder(base64.StdEncoding, pWriter)

	gWriter := gzip.NewWriter(b64Writer)

	// write text to be gzipped and push to pipe
	go func() {
		fmt.Println("Start writing")
		n, err := gWriter.Write([]byte("These words will be compressed and push into pipe"))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		gWriter.Close()
		b64Writer.Close()
		pWriter.Close()
		fmt.Printf("Written %d bytes \n", n)
	}()

	// start reading from the pipe(reverse of the above process)

	// use base64 decoder to wrap pipe Reader
	b64Reader := base64.NewDecoder(base64.StdEncoding, pReader)

	// read gzipped text and decompressed the text
	gReader, err := gzip.NewReader(b64Reader)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// print out the text
	text, err := ioutil.ReadAll(gReader)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", text)
}

/*

Output :

Start writing
Written 49 bytes
These words will be compressed and push into pipe

References :

http://golang.org/pkg/io/#Pipe

*/
