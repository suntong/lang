// From: http://www.dotnetperls.com/json-go

package main

import (
	"encoding/json"
	"fmt"
)

type boxItems struct {
	SeqNo int
}

type Box struct {
	Width  int
	Height int
	Color  string
	Open   bool
	Items  []boxItems
}

func main() {
	// Create an instance of the Box struct.
	box := Box{
		Width:  10,
		Height: 20,
		Color:  "blue",
		Open:   false,
		Items:  []boxItems{boxItems{1}, boxItems{2}, boxItems{3}},
	}
	// Create JSON from the instance data.
	// ... Ignore errors.
	b, _ := json.Marshal(box)
	// Convert bytes to string.
	s := string(b)
	fmt.Println(s)
}

/*

$ go run json.Marshal.go
{"Width":10,"Height":20,"Color":"blue","Open":false,"Items":[{"SeqNo":1},{"SeqNo":2},{"SeqNo":3}]}

*/
