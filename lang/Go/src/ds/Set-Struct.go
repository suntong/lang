////////////////////////////////////////////////////////////////////////////
// Porgram: Set-Struct
// Purpose: Demo the SET struct data structure in GO
// Authors: Tong Sun (c) 2013, All rights reserved
// Credits: Sets in Golang http://davidkaya.sk/sets-in-golang/
////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

var exists = struct{}{}

type set struct {
	m map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(value string) {
	s.m[value] = exists
}

func (s *set) Remove(value string) {
	delete(s.m, value)
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

func main() {
	s := NewSet()

	s.Add("Peter")
	s.Add("David")

	fmt.Println(s.Contains("Peter"))  // True
	fmt.Println(s.Contains("George")) // False

	s.Remove("David")
	fmt.Println(s.Contains("David")) // False
}
