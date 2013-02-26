////////////////////////////////////////////////////////////////////////////
// Porgram: StructAnonymousFieldDef
// Purpose: Go struct anonymous fields definition demo
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

package main

import "fmt"

type Kitchen struct {
  numOfPlates int
}

type Room struct {
  numOfLamps int
}

type Bedroom struct {
  Room
}

type House struct {
  numOfRooms int
  Kitchen    //anonymous field
  Bedroom    //anonymous field
}

func main() {

  // use composed type name to initialize 
  h := House{3, Kitchen{10}, Bedroom{Room{2}}}
  // numOfRooms is a field of House
  fmt.Println("House h has this many rooms:", h.numOfRooms)
  //numOfPlates is a field of anonymous field Kitchen, so it can be referred
  //to like a field of House
  fmt.Println("House h has this many plates:", h.numOfPlates)
  //we can refer to the embedded struct in its entirety by referring to the
  //name of the struct type
  fmt.Println("The Kitchen contents of this house are:", h.Kitchen)
  fmt.Println("House h has this many lamps:", h.numOfLamps, h.Bedroom)

  // if Kitchen also has a field numOfLamps, i.e., both the Kitchen and the
  // Bedroom have a field numOfLamps, and both are available as an anonymous
  // field within House. We cannot refer to h.numOfLamps any more. We have to 
  // refer to fields via type name: h.Kitchen.numOfLamps | h.Bedroom.numOfLamps
  // Also we can still reach the number of lamps in the Bedroom by using the
  // type name h.Bedroom
 
  // == Assignment between types

  //var k Kitchen = h
  // Error: cannot use h (type House) as type Kitchen in assignment
  var k Kitchen = h.Kitchen // value copy
  fmt.Println("The Kitchen:", k)

  r := Room{5}
  var br Bedroom = Bedroom{r}
  fmt.Println("Bedroom br has this many lamps:", br.Room.numOfLamps)

  /*
  House h has this many rooms: 3
  House h has this many plates: 10
  The Kitchen contents of this house are: {10}
  House h has this many lamps: 2 {{2}}
  The Kitchen: {10}
  Bedroom br has this many lamps: 5
  */

}
