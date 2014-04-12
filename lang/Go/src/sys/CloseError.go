////////////////////////////////////////////////////////////////////////////
// Porgram: CloseError
// Purpose: To check an error on a Close() or similar
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2
// https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/sfo0tvPIDuk

package main

import (
  "errors"
  "fmt"
)

type myT int

func (x myT) Close() error {
  return errors.New("error occured")
}

// http://golang.org/ref/spec#Defer_statements
// For instance, if the deferred function is a function literal and the
// surrounding function has named result parameters that are in scope within
// the literal, the deferred function may access and modify the result
// parameters before they are returned.

func test() (err, errC error) {
  var x myT
  defer func() { 
    fmt.Println("before close")
    errC = x.Close() 
    fmt.Println("after close")
  }()
  fmt.Println("test was here")
  return
}

func main() {
  err, errC := test()
  fmt.Println(err, errC)
}
