////////////////////////////////////////////////////////////////////////////
// Porgram: StructInit.go
// Purpose: Struct initilaization with private variables
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Listed inline
////////////////////////////////////////////////////////////////////////////

package main

import (
	"log"
	"os"
)

// Constructors example
// https://github.com/luciotato/golang-notes/blob/master/OOP.md

type Rectangle struct {
	Name          string
	Width, Height float64
}

func constructors() {

	var a Rectangle
	var b = Rectangle{"I'm b.", 10, 20}
	var c = Rectangle{Height: 12, Width: 14}

	Println(a)
	Println(b)
	Println(c)
}

// Methods on structs
// http://golangtutorials.blogspot.ca/2011/06/methods-on-structs.html

type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func methods_demo() {
	r1 := Rectangle{4, 3}
	fmt.Println("Rectangle is: ", r1)
	fmt.Println("Rectangle area is: ", r1.Area())
}

/*

The output is:

    Rectangle is: {4 3}
    Rectangle area is: 12

Many object oriented languages have a concept of this or self that implicitly refers to the current instance. Go has no such keyword. When defining a function or method associated with a type, it is given as a named variable - in this case (r Rectangle) and then within the function the variable r is used.

In the above call to Area, the instance of Rectangle is passed as a value. You could also pass it by reference. In calling the function, there would be no difference whether the instance that you call it with is a pointer or a value because Go will automatically do the conversion for you.

*/

func (r Rectangle) Area_by_value() int {
	return r.length * r.width
}

func (r *Rectangle) Area_by_reference() int {
	return r.length * r.width
}

func aok() {
	r1 := Rectangle{4, 3}
	fmt.Println("Rectangle is: ", r1)
	fmt.Println("Rectangle area is: ", r1.Area_by_value())
	fmt.Println("Rectangle area is: ", r1.Area_by_reference())
	fmt.Println("Rectangle area is: ", (&r1).Area_by_value())
	fmt.Println("Rectangle area is: ", (&r1).Area_by_reference())
}

/*
Adapted from embedding section of `Effective Go`
https://groups.google.com/d/msg/golang-nuts/40vBKgflGjI/mtSyaOvvmHsJ
*/

type Job struct {
	*log.Logger
	size int
	b    []byte
}

/*
question is, it would be easier to just initialize these embeds by composite
literal, and default remaining (unspecified) fields to zero values so as to
set them appropriate later.
*/
func question() {
	x := &Job{
		log.New(os.Stderr, nil, "Job: ", log.Ldate),
		0, nil, // is it really necessary ?
	}
	x.size = 2
	x.b = make([]byte, x.size)
	x.Logf("Job size=%d", x.size)
}

/*
answer is, yes, you can omit fields, but if you do
Go asks that you name the fields you're intending
to list, so that it's clear which ones you intended
to supply and that it isn't just a matter of forgetting one.

Russ Cox
*/
func answer() {
	x := &Job{
		Logger: log.New(os.Stderr, nil, "Job: ", log.Ldate),
	}
}

type T struct {
	Variable string
	private  int
}

/*
idiom shows common idiom for initialization.

Private variables are package private. If you were to try initializing a struct from another package this way, you would get an error. If you want to assign selectively to public variables, you can do:

    t := &T{Variable: "a"}

A common idiom is for a package to provide a factory function for its types that need initialization:

    func NewT() *T { return &T{"a", 3} }
    ...
    t := NewT()

Or an initializer function (less common, but still used):

    func (t *T) Init() *T { *t := T{"a", 3}; return t }

You could also include parameters in the function/method signatures to allow the user to provide values required for initialization.

Steven Blenkinsop
*/
func idiom() {
	t := new(T).Init()
	// OR
	t := new(T)
	t.Init()
}

/*

To initialize members in go struct
http://stackoverflow.com/questions/4498998/how-to-initialize-members-in-go-struct

A common used pattern is to use a constructor
*/

type SyncMap struct {
	lock *sync.RWMutex
	hm   map[string]string
}

func NewSyncMap() *SyncMap {
	return &SyncMap{hm: make(map[string]string)}
}

func main() {
	question()
	answer()
	idiom()
}
