////////////////////////////////////////////////////////////////////////////
// Porgram: StructAnonymousFieldMeth1
// Purpose: Go struct anonymous fields methods demo
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

// By Big Yuuta, http://go-book.appspot.com/more-methods.html

package main

import "fmt"

type Human struct {
  name  string
  age   int
  phone string //his own mobile number
}

type Student struct {
  Human  //an anonymous field of type Human
  school string
}

type Employee struct {
  Human   //an anonymous field of type Human
  company string
}

//A human method to say hi
func (h *Human) SayHi() {
  fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
  mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
  sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

  mark.SayHi()
  sam.SayHi()
}
