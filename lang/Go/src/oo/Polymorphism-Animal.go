////////////////////////////////////////////////////////////////////////////
// Porgram: Polymorphism-Animal.go
// Purpose: Demo the Go "polymorphism" feature with Animals
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://www.goinggo.net/2016/10/reducing-type-hierarchies.html
////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

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
	Speak()
}

// Dog contains everything an Animal is but specific
// attributes that only a Dog has.
type Dog struct {
	Animal
	PackFactor int
}

// Speak knows how to speak like a dog.
// This makes a Dog now part of a group of concrete
// types that know how to speak.
func (d Dog) Speak() {
	fmt.Println("Woof!",
		"My name is", d.Name,
		", aged", d.Age,
		", it is", d.IsMammal,
		"I am a mammal.\n  it is", d.IsMale,
		"I am male. I have a a pack factor of", d.PackFactor)
}

// Cat contains everything an Animal is but specific
// attributes that only a Cat has.
type Cat struct {
	Animal
	ClimbFactor int
}

// Speak knows how to speak like a cat.
// This makes a Cat now part of a group of concrete
// types that know how to speak.
func (c Cat) Speak() {
	fmt.Println("Meow!",
		"My name is", c.Name,
		", aged", c.Age,
		", it is", c.IsMammal,
		"I am a mammal.\n  it is", c.IsMale,
		"I am male. I have a climb factor of", c.ClimbFactor)
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
		spkr.Speak()
	}
}
