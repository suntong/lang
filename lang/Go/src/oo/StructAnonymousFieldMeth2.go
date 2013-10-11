////////////////////////////////////////////////////////////////////////////
// Porgram: StructAnonymousFieldMeth2
// Purpose: Go struct anonymous fields methods overriding demo
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

// By Big Yuuta, http://go-book.appspot.com/more-methods.html

package main

import "fmt"

type Human struct {
  name  string
  age   int
  phone string
}

type Student struct {
  Human  //an anonymous field of type Human
  school string
}

type Employee struct {
  *Human  //an anonymous field of type Human
  company string
}

type Manager struct {
  *Human  //an anonymous field of type *Human
  company string
}

//A human method to say hi
func (h *Human) SayHi() {
  fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Employee's method overrides Human's one
func (e *Employee) SayHi() {
  fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
    e.company, e.phone) //Yes you can split into 2 lines here.
}

func main() {
  mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
  sam := Employee{&Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
  tom := Manager{&Human{"Tom", 55, "222-999-XXXX"}, "Golang Inc"}

  mark.SayHi()
  tom.SayHi()
  sam.SayHi()
  sam.Human.SayHi()

  /*

  Hi, I am Mark you can call me on 222-222-YYYY
  Hi, I am Tom you can call me on 222-999-XXXX
  Hi, I am Sam, I work at Golang Inc. Call me on 111-888-XXXX
  Hi, I am Sam you can call me on 111-888-XXXX

  The way anonymous embedding works with functions is that when you attempt to
  call a method on the struct, the compiler will first look for methods
  directly on that struct that match the function name. If it finds one,
  that's the one it'll assume you mean. If it does not find one, it'll look at
  embedded structs for matching function names.  Note that only the name
  matters, not the signature, so a function with the name Foo will hide any
  function named Foo in an embedded struct, even if they take different
  arguments and return different values.

    -- Nate Finch

  */
}
