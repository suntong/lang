////////////////////////////////////////////////////////////////////////////
// Porgram: jsonfile.go
// Purpose: GO jsonfile package demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: http://www.dotnetperls.com/json-go
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/go-jsonfile/jsonfile"
)

type Box struct {
	Width  int
	Height int
	Color  string
	Open   bool
}

var filename = "jsonfile.json"

func main() {
	// Create an instance of the Box struct.
	box1 := Box{
		Width:  10,
		Height: 20,
		Color:  "blue",
		Open:   false,
	}
	jsonfile.WriteJSONToFile(filename, &box1)

	box2 := Box{Width: 20}
	jsonfile.ReadJSONFromFile(filename, &box2)

	if box1 != box2 {
		fmt.Println("Boo! Badly wrong.")
	}

	fmt.Printf("%+v\n", box2)
}

/*

$ go run jsonfile.go
{Width:10 Height:20 Color:blue Open:false}

$ cat jsonfile.json
{"Width":10,"Height":20,"Color":"blue","Open":false}

$ rm jsonfile.json

*/
