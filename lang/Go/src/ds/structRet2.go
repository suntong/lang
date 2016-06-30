// By Shawn Milochik
package main

import (
	"fmt"
)

/*

Return different structs
https://groups.google.com/d/msg/golang-nuts/OnRivRrXE40/XKM3idGVAwAJ

I'm having a hard time figuring out if it's possible to return a function but having different structs as the return type.

I'm talking to a REST API that does something like this:

GET /endpoint?option=1


{
 "fieldA": 1,
 "fieldB": 2
}


GET /endpoint?option=2
{
 "fieldA": 1,
 "fieldC": 3
}


So I have a golang code like this.


type FirstOption struct {
  FieldA string `json:"fieldA,omitempty"`
  FieldB string `json:"fieldB,omitempty"`
}



type SecondOption struct {
  FieldA string `json:"fieldA,omitempty"`
  FieldC string `json:"fieldC,omitempty"`
}


I would normally do something like this:

func (option string) FirstOption { //I know this line is bad because I might want to return SecondOption


// Do request, unmarshall json to FirstOption or SecondOption depending on the option string variable


}


But I need to know if it's possible to return either FirstOption or SecondOption.

*/

// from your example
type FirstOption struct {
	FieldA string `json:"fieldA,omitempty"`
	FieldB string `json:"fieldB,omitempty"`
}

// satisfies fmt.Stringer
func (first FirstOption) String() string {
	return fmt.Sprintf("%s: %s\n", first.FieldA, first.FieldB)
}

// from your example
type SecondOption struct {
	FieldA string `json:"fieldA,omitempty"`
	FieldC string `json:"fieldC,omitempty"`
}

// satisfies fmt.Stringer
func (second SecondOption) String() string {
	return fmt.Sprintf("%s: %s\n", second.FieldA, second.FieldC)
}

func main() {
	fmt.Println(getThing("apple"))
	fmt.Println(getThing("orange"))
}

func getThing(msg string) fmt.Stringer {
	if msg == "apple" {
		return FirstOption{FieldA: "Johnny Appleseed", FieldB: "Keeps the doctor away!"}
	}

	return SecondOption{FieldA: msg, FieldC: "of unknown usefulness"}
}

/*

This is covered by in Donovan & Kernighan's new book The Go Programming Language.  On pages 197 and 198 of the first edition, they describe an interface Expr, representing an expression, then five concrete types that satisfy that interface and contain different sorts of expressions and sub expressions.  Since they all implement the Expr type, you can write a method that can receive and return any or all of them.

However, there are some subtleties which can trip you up.  Once you start using interfaces rather than concrete types, you are pretty quickly forced into using pointers rather than objects.  Go does a pretty good job of hiding some of the details of this, but not a complete job.

Interfaces are also crucial to a lot of testing strategies, because using them allows you to easily mock up objects with the values that you need for your tests.

Unfortunately, this means that somebody new to Go has to get their head around quite a few tricky ideas before they can do anything useful.  They are all explained in the book, but you need to read the relevant sections carefully, do a few experiments and then read them again and again until it becomes clear.

The great thing about Kernighan's writing style is that he says everything he needs to say exactly once and using the minimum of words.  The thing that often causes his readers problems is that he says everything he needs to say exactly once and using the minimum of words.

Simon Ritchie

*/
