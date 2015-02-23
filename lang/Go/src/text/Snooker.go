////////////////////////////////////////////////////////////////////////////
// Porgram: Snooker
// Purpose: play the "snooker" game
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2 -w

package main

import (
  "fmt"
  "math/rand"
  "time"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// snooker table size
const (
  YMax = 80
  XMax = 60
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// the snooker table
var sTable [XMax][YMax]byte

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Abs returns the absolute value of x.
func Abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

/* 

newLineFunc returns the linear function y=f(x) given x1, y1, x2, & y2.

The function can be as simple as

   y = (y1-y2)/(x1-x2)*(x-x1)+y1

however, the concern is that then those x1, y1, x2, & y2 will be treated
exactly as x. I.e., the x1, y1, x2, & y2, are fixed value, as long as the
linear function is concerned. However, if programming this way, the
`(y1-y2)/(x1-x2)` will be calculated each time when the function `f()` is
called.

The soultion is to use a closure as follows:

*/

func newLineFunc(x1, y1, x2, y2 int) func(int) int {
  m := float64(y1-y2) / float64(x1-x2)
  return func(x int) int {
    return int(m*float64(x-x1)) + y1
  }
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Function main
func main() {

  //========================================================================
  // snooker table init
  i := 0
  j := 0

  // empty the snooker table
  for i = 0; i < XMax; i++ {
    for j = 0; j < YMax; j++ {
      sTable[i][j] = ' '
    }
  }

  // draw the snooker table borders
  for i = 0; i < XMax; i++ {
    sTable[i][0] = '#'
    sTable[i][YMax-1] = '#'
  }
  for i = 0; i < YMax; i++ {
    sTable[0][i] = '#'
    sTable[XMax-1][i] = '#'
  }

  //========================================================================
  // two Random points
  rand.Seed(time.Now().UTC().UnixNano())
  x1 := rand.Intn(XMax)
  y1 := rand.Intn(YMax)
  x2 := rand.Intn(XMax)
  y2 := rand.Intn(YMax)
  x1 = rand.Intn(XMax/2) + XMax/2
  sTable[x1][y1] = 'X'
  sTable[x2][y2] = 'X'

  //========================================================================
  // hit

  // == direct hit
  // f := newLineFunc(x1, y1, x2, y2)
  // for i = x1-1; i>x2; i-- {
  //   j= f(i)
  //   sTable[i][j]='O'
  // }
  
  // == bounce two times
  f := newLineFunc(x1, y1, -x2, -y2)
  for i = x1 - 1; i > -x2; i-- {
    j = Abs(f(i))
    sTable[Abs(i)][j] = 'O'
  }

  //========================================================================
  // print out the snooker table
  for i = 0; i < XMax; i++ {
    for j = 0; j < YMax; j++ {
      fmt.Printf("%c", sTable[i][j])
    }
    fmt.Println()
  }

}
