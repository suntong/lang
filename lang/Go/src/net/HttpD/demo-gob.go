// https://go.dev/play/p/QIWKQD4EL9

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

var m = map[int]string{1: "one", 2: "two", 3: "three"}

func main() {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	err := encoder.Encode(m)
	if err != nil {
		panic(err)
	}

	// your encoded stuff
	fmt.Println(buf.Bytes())

	var decodedMap map[int]string
	decoder := gob.NewDecoder(buf)

	err = decoder.Decode(&decodedMap)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", decodedMap)
}
