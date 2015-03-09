////////////////////////////////////////////////////////////////////////////
// Porgram: DataStructure-Nested
// Purpose: Go nested data structure demo
// Authors: Tong Sun (c) 2014, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2 -w

/*

Array of structs
https://groups.google.com/d/msg/golang-nuts/XSgUJCTnhU8/aW58iq5o20MJ

I am trying to figure out the right syntax to initialize an array of structs.

It would be a bit wiser to remove the inner anonymous struct. Then you could simply: http://play.golang.org/p/pNU33mFTaf

Péter Szilágyi

*/

package main

import "fmt"

type Color struct {
  Name  string
  Title string
}

type ColorGroup struct {
  ID     int
  Name   string
  Colors []Color
}

func main() {
  group := ColorGroup{
    ID:   1,
    Name: "notify",
    Colors: []Color{
      {"red", "bad"},
      {"green", "good"},
    },
  }
  fmt.Printf("%v\n", group)
  // {1 notify [{red bad} {green good}]}
}

