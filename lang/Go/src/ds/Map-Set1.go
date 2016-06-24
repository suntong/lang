////////////////////////////////////////////////////////////////////////////
// Porgram: Set by map
// Purpose: Demo the SET data structure in GO
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Val @golang-nuts, https://play.golang.org/p/_n56yMhlRt
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
