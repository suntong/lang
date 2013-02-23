////////////////////////////////////////////////////////////////////////////
// Porgram: DebugOutput
// Purpose: Go Debugging Output Demo
// Authors: Tong Sun, Gerard (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

// Goal: to have a simple C like macros for debug output

/*

Problem definition in details:

On Saturday, February 23, 2013 9:53:11 AM UTC-5, Gerard wrote:

> > Expressed in C macros, I'd like to define something like this:
> > 
> >  #define Debugln(X) if debugging { println(X) } 
> > 
> > What's the closest way to do it in go? 
> 
> const debug = true 
> 
> ... 
> >         if debug { // if body is not linked in for debug == false 
> >                 log.Printf("%v", foo) 
> >         } 
> > ... 
> > 
> > 
> > Thanks Jan, but the spirit of C macros is that all your above 3
> > line would condensed to "Debugln(foo)", so that I don't need to
> > do duplicate such 3-line code everywhere.
> 
> I see, but that's what I suggest to do in Go. Of course one can have: 
> 
> func DebugLn(s string, args ...interface{}) { 
>         if debug { 
>                 log.Printf(s, args...) 
>         } 
> } 
> 
> But that: 
> 
> - Gets included in the production binary 
> - Arguments are evaluated even for debug == false. 
> 
> So I cannot recommend it. 

Another approch

http://play.golang.org/p/mOSbdHwSYR

However, if you have a multiple package piece application the previous
example may not work as expected.

With this http://play.golang.org/p/gjszhNzoOk package, a "debug" package
with fmt Print like functions (Rob Pike had this idea first), you have one
global variable for your app: dbg.Debug.

And maybe it's better to use the log package instead of fmt (or print to
os.Stderr).

Gerard

*/

package main

import "fmt"

type Debug bool

func (d Debug) Printf(s string, a ...interface{}) {
  if d {
    fmt.Printf(s, a...)
  }
}

func main() {
  var dbg Debug
  dbg.Printf("Debugging off")
  dbg = true
  dbg.Printf("Debugging on")
}
