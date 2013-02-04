////////////////////////////////////////////////////////////////////////////
// Porgram: BailOut
// Purpose: A demo of mechanism that makes sure all clean up and done properly before closing down the program
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

package main

import (
  "log"
  "runtime/debug"
  "time"
)

// defer/trace demo
// http://tip.golang.org/doc/effective_go.html#defer

// Arguments to deferred functions are evaluated when the defer executes. 
// The tracing routine can set up the argument to the untracing routine.

func trace(s string) string {
  log.Println("entering:", s)
  return s
}

func un(s string) {
  log.Println("leaving:", s)
}

func a() {
  defer un(trace("a"))
  log.Println("in a")
  // Deferred functions are executed in LIFO order, will print 4 3 2 1 0
  for i := 0; i < 5; i++ {
    defer log.Printf("%d ", i)
  }
}

func b() {
  defer un(trace("b"))
  log.Println("in b")
  a()
}

// panic & recover demo
// By Jan Mercl, http://play.golang.org/p/m96skGjRjo

func deep2(n int) {
  defer func() {
    log.Printf("Resource %d closed", n)
    if err := recover(); err != nil {
      panic(err)
    }
  }()

  for i := 1; i <= 5; i++ {
    log.Println("Working...")
    time.Sleep(time.Second)

  }
  panic("I'm afraid I can't do that")
}

func deep1(n int) {
  defer func() {
    log.Printf("Resource %d closed", n)
    if err := recover(); err != nil {
      panic(err)
    }
  }()
  deep2(n + 1)
}

func main() {
  defer func() {
    log.Print("Main exiting")
    if err := recover(); err != nil {
      log.Fatalf("Stack trace:\n%s----\n%s", debug.Stack(), err)
    }
  }()

  log.Println("Main started")
  b()

  deep1(1)
}
