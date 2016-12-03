////////////////////////////////////////////////////////////////////////////
// Porgram: Polymorphism-AnimalVF3.go
// Purpose: Demo the Go polymorphism & virtual function feature with Animals
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Xingtao Zhao
// https://groups.google.com/d/msg/golang-nuts/f_62HEOIBV4/3wfAj4pDCAAJ
////////////////////////////////////////////////////////////////////////////

/*

The best way to implement *virtual function* so far, that satisfies the
following challenges:

- Consider the "func Output()" as a *very complicated* function that I only
  want to define once at the base level, not to duplicate into each sub
  classes, yet it need to access member functions from sub classes.

- Meanwhile I have a *huge list* of common variables that I defined in my
  "base class", and using simple pure function will cause almost every
  single variable to be undefined: https://play.golang.org/p/QjCtD9rGpa

Note from Xingtao Zhao:

You could merge the interface Animal and ExtraFactsor together if this extra
behavior is not needed. It is listed there just for demonstration.

*/

package main

import "fmt"

////////////////////////////////////////////////////////////////////////////
// Data type/structure definitions

// AnimalBase contains all the base fields for animals.
type AnimalBase struct {
	Name     string
	Age      int
	IsMale   bool
	IsMammal bool
}

func (a *AnimalBase) GetBase() *AnimalBase { return a }

// Animal provide a common interface for all concrete types
// to follow if they want to be a part of this group. This
// is a contract for these concrete types to follow.
type Animal interface {
	GetBase() *AnimalBase
	IsA() string
	Speak() string
}

// ExtraFactsor is for demonstration only
type ExtraFactsor interface {
	ExtraFacts() string
}

func (a AnimalBase) IsA() string {
	return "Animal"
}

func (a AnimalBase) Speak() string {
	return "Something"
}

func Output(d Animal) {
	base := d.GetBase()
	fmt.Println(d.Speak(),
		"\n  I am a", d.IsA(),
		"\n  My name is", base.Name,
		", aged", base.Age,
		"\n  It is", base.IsMammal, "I am a mammal.",
		"\n  It is", base.IsMale, "I am male.")
	if s, ok := d.(ExtraFactsor); ok {
		fmt.Println(" ", s.ExtraFacts())
	}
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Dog contains everything an Animal is but specific
// attributes that only a Dog has.
type Dog struct {
	AnimalBase
	PackFactor int
}

func (d Dog) IsA() string {
	return "Dog"
}

// Speak knows how to speak like a dog.
// This makes a Dog now part of a group of concrete
// types that know how to speak.
func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) ExtraFacts() string {
	return fmt.Sprintf("I have a a pack factor of %v", d.PackFactor)
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Cat contains everything an Animal is but specific
// attributes that only a Cat has.
type Cat struct {
	AnimalBase
	ClimbFactor int
}

func (c Cat) IsA() string {
	return "Cat"
}

// Speak knows how to speak like a cat.
// This makes a Cat now part of a group of concrete
// types that know how to speak.
func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) ExtraFacts() string {
	return fmt.Sprintf("I have a climb factor of %v", c.ClimbFactor)
}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
type Sheep struct {
	AnimalBase
	PackFactor int
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

func main() {

	// Create a list of Animals that know how to speak.
	animals := []Animal{

		&Sheep{}, // OK to use Sheep as type Animal now

		// Create a Dog by initializing its AnimalBase parts
		// and then its specific Dog attributes.
		&Dog{
			AnimalBase: AnimalBase{
				Name:     "Fido",
				Age:      8,
				IsMale:   true,
				IsMammal: true,
			},
			PackFactor: 5,
		},

		// Create a Cat by initializing its AnimalBase parts
		// and then its specific Cat attributes.
		&Cat{
			AnimalBase: AnimalBase{
				Name:     "Milo",
				Age:      5,
				IsMale:   false,
				IsMammal: true,
			},
			ClimbFactor: 4,
		},
	}

	// Have the Animals speak.
	for _, animal := range animals {
		Output(animal)
	}
}

/*

Woof!
  I am a Dog
  My name is Fido , aged 8
  It is true I am a mammal.
  It is true I am male.
  I have a a pack factor of 5
Meow!
  I am a Cat
  My name is Milo , aged 5
  It is true I am a mammal.
  It is false I am male.
  I have a climb factor of 4

*/
