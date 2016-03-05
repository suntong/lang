////////////////////////////////////////////////////////////////////////////
// Porgram: Pipe_JsonPost.go
// Purpose: io.Pipe demo with Json Post
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://garbagecollected.org/2015/05/30/io-with-go-io-pipe/
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type msg struct {
	Text string
}

func handleErr(err error) {
	if err != nil {
		log.Fatalf("%s\n", err)
	}
}

// use a io.Pipe to connect a JSON encoder to an HTTP POST: this way you do
// not need a temporary buffer to store the JSON bytes
func main() {
	r, w := io.Pipe()

	// writing without a reader will deadlock so write in a goroutine
	go func() {
		// it is important to close the writer or reading from the other end of the
		// pipe will never finish
		defer w.Close()

		m := msg{Text: "brought to you by io.Pipe()"}
		err := json.NewEncoder(w).Encode(&m)
		handleErr(err)
	}()

	resp, err := http.Post("https://httpbin.org/post", "application/json", r)
	handleErr(err)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	handleErr(err)

	log.Printf("%s\n", b)
}
