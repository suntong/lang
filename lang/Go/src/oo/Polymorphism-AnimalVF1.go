////////////////////////////////////////////////////////////////////////////
// Porgram: Polymorphism-AnimalVF.go
// Purpose: Demo the Go polymorphism & virtual function feature with Animals
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: paraiso.marc https://play.golang.org/p/o6Ot4IdJZ1
////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

/*

How to architect the OO's virtual function in Go?
https://groups.google.com/d/topic/golang-nuts/f_62HEOIBV4/

Q:

>> Please take a look at this (not working) Go program
>> https://play.golang.org/p/qrBX6ScABp

>> Please think of the "func Output()" as a very complicated function that I only want to define once at the base level, not to duplicate into each sub classes.

> You define a function:
>
> func Output(s Shape) string {
>    return s.Name() + "'s area size is " // + s.Area()
> }
>
> Go uses interfaces for polymorphism.
> Other OOP languages can use inheritance for polymorphism too, but Go
> doesn't have inheritance.

Thanks Jesse. That works.

However I just realized that it is only part of my problem.

I have a huge list of common variables that I defined in my "base class", changing it from a member function to a pure function causes almost every single variable now undefined.

I can demonstrate the problem with this code
https://play.golang.org/p/QjCtD9rGpa

So, once again, thinking in OO, I'll define all of the common variables in base class, and common functionalities in virtual functions. How to make that idea work in Go?

For the above specific code, how to easily make "func Output" works?

A:

interfaces only work on methods.

https://play.golang.org/p/o6Ot4IdJZ1

Name, Age , ... are not methods they are fields. You need to make them part of Speaker interface by using methods instead of fields.

paraiso.marc @gmail.com

*/

// Animal contains all the base fields for animals.
type Animal struct {
	Name     string
	Age      int
	IsMale   bool
	IsMammal bool
}

func (a Animal) GetAge() int       { return a.Age }
func (a Animal) GetIsMale() bool   { return a.IsMale }
func (a Animal) GetIsMammal() bool { return a.IsMammal }
func (a Animal) GetName() string   { return a.Name }

// Speaker provide a common behavior for all concrete types
// to follow if they want to be a part of this group. This
// is a contract for these concrete types to follow.
type Speaker interface {
	IsA() string
	Speak()
	GetName() string
	GetAge() int
	GetIsMammal() bool
	GetIsMale() bool
	GetPackFactor() int
}

func Output(d Speaker) {
	fmt.Println("I am a", d.IsA(),
		"My name is", d.GetName(),
		", aged", d.GetAge(),
		", it is", d.GetIsMammal(),
		"I am a mammal.\n  it is", d.GetIsMale(),
		"I am male. I have a a pack factor of", d.GetPackFactor())
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

func (d Dog) GetPackFactor() int { return d.PackFactor }

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

func (c Cat) IsA() string {
	return "Cat"
}

func (c Cat) GetPackFactor() int { return 0 }

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
		fmt.Println("===")
		spkr.Speak()
		Output(spkr)
	}
}

/*

===
Woof! My name is Fido , aged 8 , it is true I am a mammal.
  it is true I am male. I have a a pack factor of 5
I am a Dog My name is Fido , aged 8 , it is true I am a mammal.
  it is true I am male. I have a a pack factor of 5
===
Meow! My name is Milo , aged 5 , it is true I am a mammal.
  it is false I am male. I have a climb factor of 4
I am a Cat My name is Milo , aged 5 , it is true I am a mammal.
  it is false I am male. I have a a pack factor of 0

*/
