////////////////////////////////////////////////////////////////////////////
// Porgram: PersistentData-GOB.go
// Purpose: Go Persistent Data w/ GOB Demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits:
//          https://blog.golang.org/gobs-of-data
//          https://play.golang.org/p/wT8_H44crC by Michael Jones
//          Matt Silverlock in go-nuts on "Serialization internal data to disk"
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

import (
	//"encoding/gob"
	//"fmt"
	//"log"
	"os"
	"time"
)

func main() {
  test0()
  test1()
  test2()
}

/*

$ go run PersistentData-GOB.go
"Pythagoras": {3,4}
&main.Data{ID:"113131", Payload:[]uint8{0x73, 0x74, 0x75, 0x66, 0x66}, Created:1476543890}
{123 1.6777216e+07}
{123 1.6777216e+07}

$ rm -v *.gob
removed 'data.gob'
removed 'persist.gob'

*/

////////////////////////////////////////////////////////////////////////////
// From https://blog.golang.org/gobs-of-data

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func test0() {
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	// Encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) the value.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)
}

////////////////////////////////////////////////////////////////////////////
// From Matt Silverlock in go-nuts

/*

> I only see the encoding/gob serializing into internal buffer or network.
The challenge for me is to store the internal data to disk then read them back next time.

You can just encode into an os.File (and decode from an os.File) without
having to use a buffer - os.File implements io.Writer/io.Reader -
here's a naive example:

*/



type Data struct {
	ID      string
	Payload []byte
	Created int64
}

func test1() {
	data := &Data{
		ID:      "113131",
		Payload: []byte("stuff"),
		Created: time.Now().Unix(),
	}

	f, err := os.Create("data.gob")
	if err != nil {
		log.Fatal(err)
	}

	// os.File impleements io.Writer
	enc := gob.NewEncoder(f)
	err = enc.Encode(&data)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	existing, err := os.Open("data.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer existing.Close()

	retrieved := &Data{}
	dec := gob.NewDecoder(existing)
	err = dec.Decode(&retrieved)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", retrieved)
}

////////////////////////////////////////////////////////////////////////////
// From https://play.golang.org/p/wT8_H44crC by Michael Jones

type example struct {
	Count  int32
	Amount float64
}

func test2() {
	state := example{123, 16777216.0}

	// print the state
	fmt.Printf("%v\n", state)

	if err := SaveState(state); err != nil {
		log.Fatal("SaveState failed:", err)
	}

	restored := example{} // empty state
	err := RestoreState(&restored)
	if err != nil {
		log.Fatal("RestoreState failed:", err)
	}

	// print the restored state
	fmt.Printf("%v\n", restored)
}

const persistName = "persist.gob"

func SaveState(state interface{}) error {
	// create persistence file
	f, err := os.Create(persistName)
	if err != nil {
		return err
	}
	defer f.Close()

	// write persistemce file
	e := gob.NewEncoder(f)
	if err = e.Encode(state); err != nil {
		return err
	}
	return nil
}

func RestoreState(state interface{}) error {
	// open persistence file
	f, err := os.Open(persistName)
	if err != nil {
		return err
	}
	defer f.Close()

	// read persistemce file
	e := gob.NewDecoder(f)
	if err = e.Decode(state); err != nil {
		return err
	}
	return nil
}
