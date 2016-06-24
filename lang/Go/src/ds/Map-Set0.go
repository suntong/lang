////////////////////////////////////////////////////////////////////////////
// Porgram: Set by map
// Purpose: Demo the SET data structure in GO
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Val @golang-nuts, https://play.golang.org/p/DTtGs3FiwZ
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

type set map[string]struct{}

var validSub set

func init() {
	validSub.Init([]string{"aa", "bb", "cc"})
	fmt.Printf("%+v\n", validSub)
}

func (s *set) Init(slice []string) {
	*s = make(map[string]struct{}, len(slice))
	for _, s1 := range slice {
		(*s)[s1] = struct{}{}
	}
}

func (s *set) Has(a string) bool { _, ok := (*s)[a]; return ok }

func main() {
	fmt.Println(validSub.Has("aa"))
	fmt.Println(validSub.Has("dd"))
}

/*

Output:

map[aa:{} bb:{} cc:{}]
true
false

*/

/*

> Is it possible to fix and make the
>
> validSub.Init([]string{"aa", "bb", "cc"})
>
> works instead?
>
> Coming from the OO world, this is a form that I feel more comfort with.

Even in OO style (e.g. java), you would not be able to write

 Set s = null;
 s.init( list(1,2,3) );

This is (more or less) the same problem with value receiver.

However, map is a reference type, so once the map has been created you really can use value receiver : https://play.golang.org/p/FePU2I-u2-

```
var validSub set

func init() {
	validSub = make(set)
	validSub.Init([]string{"aa", "bb", "cc"})
	fmt.Printf("%+v\n", validSub)
}

func (s set) Init(slice []string) {
	for _, s1 := range slice {
		s[s1] = struct{}{}
	}
}

func (s set) Has(a string) bool { _, ok := s[a]; return ok }
```

If your concern is about "allocating a map with make, exactly the size of the input slice", then consider that the "constructor-like" idiom in go is a function NewSet, not a method :
https://play.golang.org/p/_n56yMhlRt

See Map-Set1.go

*/
