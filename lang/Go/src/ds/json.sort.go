////////////////////////////////////////////////////////////////////////////
// Porgram: json.sort.go
// Purpose: Go Sorting json fields recursive
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: Jakob Borg https://groups.google.com/d/msg/golang-nuts/iCZG3bixy-0/-udkZE8zBgAJ
////////////////////////////////////////////////////////////////////////////

/*

RE: Sorting json fields recursive

Q: Is there any existing code/package out there that can sort an arbitrary json string in all its nesting levels, so that the attributes at any level are in sorted order?

A:

Unmarshalling and (re)marshalling the message will do that for you.

https://play.golang.org/p/UrHedCa6sO

You get error checking and reformatting for free.

Jakob Borg

*/

package main

import (
	"encoding/json"
	"fmt"
)

var in = `{
  "foo": {
    "b": 42,
    "c": 12,
    "a": 33
  },
  "bar": {
    "z": false,
    "w": true
  }
}`

func main() {
	var res interface{}
	json.Unmarshal([]byte(in), &res)
	bs, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(bs))
}

/*

Output:

{
  "bar": {
    "w": true,
    "z": false
  },
  "foo": {
    "a": 33,
    "b": 42,
    "c": 12
  }
}

*/
