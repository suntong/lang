////////////////////////////////////////////////////////////////////////////
// Porgram: Set
// Purpose: Demo the SET data structure in GO
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2

package main

import "fmt"

////// Debugging //////
import "github.com/davecgh/go-spew/spew"
var _ = spew.Config
//////////////////////

/*

Go 101, the SET data structure
// https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/NvwxAJgD3KI

What the Go's way for the SET data structure? Like the list in Perl that
offers FIFO/Stack operation and delete from random places.
http://en.wikipedia.org/wiki/Set_(computer_science)#Dynamic_sets

By Peter Bourgon:

>> Thanks for http://golang.org/pkg/container/.
>> I'll settle with the simple container/heap for the being.
>
> I would usually use a map[Thing]bool or a map[Thing]struct{} for the job.

+1 to using a map. "container" is usually not what you want.

http://play.golang.org/p/rWdn0vts1c

*/

type Set interface {
  Insert(string)
  Has(string) bool
  Remove(string)
}

type set map[string]struct{}

func (s set) Insert(a string)   { s[a] = struct{}{} }
func (s set) Remove(a string)   { delete(s, a) }
func (s set) Has(a string) bool { _, ok := s[a]; return ok }

func main1() {
  s := &set{}
  fmt.Printf("Has(abc) = %v\n", s.Has("abc"))
  s.Insert("abc")
  fmt.Printf("Has(abc) = %v\n", s.Has("abc"))
  s.Remove("abc")
  fmt.Printf("Has(abc) = %v\n", s.Has("abc"))
}

/*

On Mon, Mar 4, 2013 at 5:37 AM, Rémy Oudompheng wrote:

>> Now a problem that still bewilders me -- how to iterate through the
>> dynamically changing set? Think of the web crawler. It visit the first page
>> and put all the urls from within the page into the toVisit set, then iterate
>> through this toVisit set while putting more urls into the set (with
>> predefined conditions of course), until all items have been removed from the
>> toVisit set to the hasSeen set.
>
>
> It's generally a bad idea to modify a collection while you're iterating
> through it. In most implementations this will either fail with a runtime
> exception, or just have "undefined behavior" (including crashing.)

Go correctly defines the behaviour of range loops when you mutate the
collection. In particular, the very common pattern of deleting the key
you are currently sitting on, is well-defined, deterministic and
usually what you want.

Notably, the following pattern will work for a recursive visitor:

for len(set) > 0 {
  for item := range set {
     do(item)
     add some elements to set
     delete(item)
  }
}

*/

const N = 6

var v []int = make([]int, N)

type SetNp interface {
  Insert(*int)
  Has(*int) bool
  Remove(*int)
}

type setNp map[*int]struct{}

func (s setNp) Insert(a *int)   { s[a] = struct{}{} }
func (s setNp) Remove(a *int)   { delete(s, a) }
func (s setNp) Has(a *int) bool { _, ok := s[a]; return ok }


func main2() {
  for i := range v {
    v[i] = i * 11
  }

  s := &setNp{}

  s.Insert(&v[1])
  s.Insert(&v[2])

  // http://golang.org/ref/spec#Comparison_operators
  // Pointer values are comparable. Two pointer values are equal if they point to the same variable or if both have value nil. 
  // So even trying to add the &v[1] & &v[2] twice to the set, there should be only one

  i := 1
  for len(*s) > 0 {
    for item := range *s {
      spew.Dump(item)
      if i<5 { s.Insert(&v[i]); i++; s.Insert(&v[i]); }
      s.Remove(item)
    }
    fmt.Println("---")
  }

}


func main() {
  main1()
  main2()
}
