////////////////////////////////////////////////////////////////////////////
// Porgram: Set by map
// Purpose: Demo the SET data structure in GO
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Valentin Deleplace, https://play.golang.org/p/_n56yMhlRt
////////////////////////////////////////////////////////////////////////////

/*

Unlike Map-Set0.go, that illustrates how to make things works technically,
this apporacho is more idiom in go -- the "constructor-like" idiom in go is
a function NewXxx, not a method.

*/

package main

import (
	"fmt"
)

type set map[string]struct{}

var validSub set

func init() {
	validSub = NewSet([]string{"aa", "bb", "cc"})
	fmt.Printf("%+v\n", validSub)
}

func NewSet(slice []string) set {
	s := make(set, len(slice))
	for _, s1 := range slice {
		s[s1] = struct{}{}
	}
	return s
}

func (s set) Has(a string) bool { _, ok := s[a]; return ok }

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

> > map is a reference type
>
> Does that means that, this function
>
>     func (s set) Has(a string) bool { _, ok := s[a]; return ok }
>
> is exactly the same as the following?
>
>     func (s *set) Has(a string) bool { _, ok := (*s)[a]; return ok }
>
> I.e., even the set is super super big, there is no pass-by-value penalty in the first version?

No penalty, passing a map around never copies the values.

But a "pointer to a pointer" receiver is not strictly the same as a pointer receiver.

Valentin Deleplace

*/

/*

> Since a map is a pointer, for any function that take map as a parameter, or function like,
>
>  func (s set) Something(xxx)
>
> Does it means the map values can be changed inside the function?
>
> If yes, what's the proper way to declare it so as to make sure the values cannot be changed inside the function?

Good question, "by reference" vs "by values" is a fundamental question in most languages, and the source of many headaches when used incorrectly so it's worth taking some time to ask and get it right.

Yes you may change the values of an existing map, inside any function or method.

Go doesn't provide idiomatic ways to have "immutable-like" or "view-like" builtin containers, except string which are always immutable.  When you have an array, or a slice, or a map, you can always alter its content.

Having said this, if you don't want to caller to modify the data, you may :

- clone the data (make a copy of the map, you have to explicitly loop for that) and give the clone to the caller. This is called "defensive copying".
- or create custom struct type, containing unexported map, with exported read-only accessors e.g.   func (c *MyCustomContainer) Get(key string) Value

Valentin Deleplace

*/
