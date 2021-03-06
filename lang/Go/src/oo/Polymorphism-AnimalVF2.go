////////////////////////////////////////////////////////////////////////////
// Porgram: Polymorphism-AnimalVF2.go
// Purpose: Demo the Go polymorphism & virtual function feature with Animals
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Nick Patavalis https://play.golang.org/p/FsorWRaLKk
//    Tahir Hashmi https://tech.t9i.in/2014/01/22/inheritance-semantics-in-go/
////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

/*

Another way to implement virtual function:

This is the easiest way to make virtual function works, that satisfies the
following restrains:

- Consider the "func Output()" as a very complicated function that I only
  want to define once at the base level, not to duplicate into each sub classes.
- Changing it from a member function to a pure function is problematic as I
  have a huge list of common variables that I defined in my "base class",
  such change will cause almost every single variable now undefined.

Note that the signature (API) of Output() of based & sub class can be different.

*/

// Animal contains all the base fields for animals.
type Animal struct {
	Name     string
	Age      int
	IsMale   bool
	IsMammal bool
}

// Speaker provide a common behavior for all concrete types
// to follow if they want to be a part of this group. This
// is a contract for these concrete types to follow.
type Speaker interface {
	IsA() string
	Speak()
	Output()
}

func (a Animal) Output(s Speaker) {
	// Complicated stuff that must not be re-implemented
	fmt.Print("I am a ", s.IsA(),
		". My name is ", a.Name,
		", aged ", a.Age,
		", it is ", a.IsMale,
		" I am male.\n  ")
	s.Speak()
}

// Dog contains everything an Animal is but specific
// attributes that only a Dog has.
type Dog struct {
	Animal
	PackFactor int
}

func (d Dog) IsA() string {
	return "Dog"
}

// Speak knows how to speak like a dog.
// This makes a Dog now part of a group of concrete
// types that know how to speak.
func (d Dog) Speak() {
	fmt.Println("Woof!",
		"I have a a pack factor of", d.PackFactor)
}

func (d Dog) Output() {
	// Presumably complicated stuff, not reimplemented
	d.Animal.Output(d)
}

// Cat contains everything an Animal is but specific
// attributes that only a Cat has.
type Cat struct {
	Animal
	ClimbFactor int
}

func (c Cat) IsA() string {
	return "Cat"
}

// Speak knows how to speak like a cat.
// This makes a Cat now part of a group of concrete
// types that know how to speak.
func (c Cat) Speak() {
	fmt.Println("Meow!",
		"I have a climb factor of", c.ClimbFactor)
}

func (c Cat) Output() {
	// Presumably complicated stuff, not reimplemented
	c.Animal.Output(c)
}

func main() {

	// Create a list of Animals that know how to speak.
	speakers := []Speaker{

		// Create a Dog by initializing its Animal parts
		// and then its specific Dog attributes.
		Dog{
			Animal: Animal{
				Name:     "Fido",
				Age:      8,
				IsMale:   true,
				IsMammal: true,
			},
			PackFactor: 5,
		},

		// Create a Cat by initializing its Animal parts
		// and then its specific Cat attributes.
		Cat{
			Animal: Animal{
				Name:     "Milo",
				Age:      5,
				IsMale:   false,
				IsMammal: true,
			},
			ClimbFactor: 4,
		},
	}

	// Have the Animals speak.
	for _, spkr := range speakers {
		// spkr.Speak()
		spkr.Output()
	}
}

/*

I am a Dog. My name is Fido, aged 8, it is true I am male.
  Woof! I have a a pack factor of 5
I am a Cat. My name is Milo, aged 5, it is false I am male.
  Meow! I have a climb factor of 4

*/
