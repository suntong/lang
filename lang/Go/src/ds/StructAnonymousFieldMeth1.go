////////////////////////////////////////////////////////////////////////////
// Porgram: StructAnonymousFieldMeth1
// Purpose: Go struct anonymous fields methods demo
// Authors: Tong Sun (c) 2013-2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

// By Big Yuuta, http://go-book.appspot.com/more-methods.html

package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string //his own mobile number
}

type Student struct {
	Human  //an anonymous field of type Human
	school string
}

type Employee struct {
	Human   //an anonymous field of type Human
	company string
}

//A human method to say hi
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()

	h1 := Human{"Tom", 25, "222-333-YYYY"}
	var s1 Student
	// X: s1 = Student{h1} too few values in struct initializer
	s1 = Student{h1, "MIT"}
	// X: s1 = Student{h1, school: "MIT"} mixture of field:value and value
	s1 = Student{Human: h1, school: "MIT"}
	s1.SayHi()
}
