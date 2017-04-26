////////////////////////////////////////////////////////////////////////////
// Porgram: Interface0.go
// Purpose: Go interface intro & demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: As listed below
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Name of function: " + funcName())
	x()

	type_of_variable()
	gotour_empty_interface()

	type_from_fmt()
	type_of_object()

	// Type assertions in Go
	PrimaryExpression()
	DynamicType()
	InterfaceType()
	NotInterfaceType()
}

////////////////////////////////////////////////////////////////////////////
// Type of a Variable
// https://gistpages.com/posts/go-lang-get-type-of-a-variable

func type_of_variable() {
	fmt.Println("\n## " + funcName())

	var now time.Time = time.Now().UTC()
	fmt.Println("now is a type of: ", reflect.TypeOf(now))
	var name string = "Carl Johannes"
	fmt.Println("name is a type of: ", reflect.TypeOf(name))
	var age int = 5
	fmt.Println("age is a type of: ", reflect.TypeOf(age))
}

/*

now is a type of:  time.Time
name is a type of:  string
age is a type of:  int

*/

////////////////////////////////////////////////////////////////////////////
// https://tour.golang.org/methods/14

func gotour_empty_interface() {
	fmt.Println("\n## " + funcName())

	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*

(<nil>, <nil>)
(42, int)
(hello, string)

*/

////////////////////////////////////////////////////////////////////////////
// Name of function
// http://stackoverflow.com/questions/10742749/

func funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

// func funcName() string {
// 	pc, _, _, _ := runtime.Caller(1)
// 	nameFull := runtime.FuncForPC(pc).Name() // main.foo
// 	nameEnd := filepath.Ext(nameFull)        // .foo
// 	name := strings.TrimPrefix(nameEnd, ".") // foo
// 	return name
// }

func x() {
	fmt.Println("Name of function: " + funcName())
	y()
}

func y() {
	fmt.Println("Name of function: " + funcName())
	z()
}
func z() {
	fmt.Println("Name of function: " + funcName())
}

////////////////////////////////////////////////////////////////////////////
// http://stackoverflow.com/questions/20170275/

func type_from_fmt() {
	fmt.Println("\n## " + funcName())

	types := []interface{}{"s", 6, 6.0, true, []string{}}
	for _, v := range types {
		fmt.Printf("%T\n", v)
	}
}

/*

string
int
float64
bool
[]string

*/

func type_of_object() {
	fmt.Println("\n## " + funcName())

	b := true
	s := "str"
	n := 1
	f := 1.0
	a := []string{"foo", "bar", "baz"}

	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(n))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(a))
	fmt.Println()

	fmt.Println(reflect.ValueOf(b).Kind())
	fmt.Println(reflect.ValueOf(s).Kind())
	fmt.Println(reflect.ValueOf(n).Kind())
	fmt.Println(reflect.ValueOf(f).Kind())
	fmt.Println(reflect.ValueOf(a).Index(0).Kind()) // For slices and strings
	fmt.Println("\n### type_from_assertion")

	type_from_assertion(b)
	type_from_assertion(s)
	type_from_assertion(n)
	type_from_assertion(f)
	type_from_assertion(a)

}

/*

bool
string
int
float64
[]string

bool
string
int
float64
string

*/

func type_from_assertion(v interface{}) {
	// fmt.Println(v.(type)) // X use of .(type) outside type switch
	// switch v := anything.(type) {
	// case string:
	// 	fmt.Println(v)
	// case int32, int64:
	// 	fmt.Println(v)
	// case SomeCustomType:
	// 	fmt.Println(v)
	// default:
	// 	fmt.Println("unknown")
	// }
	switch t := v.(type) {
	case int32, int64:
		fmt.Println(t)
	default:
		fmt.Println(t)
	}
}

////////////////////////////////////////////////////////////////////////////
// Type assertions in Go
// https://medium.com/golangspec/type-assertions-in-go-e609759c42e1

type IB interface {
	walk()
	quack()
}
type S struct{}

func (s S) walk()  {}
func (s S) quack() {}

func PrimaryExpression() {
	fmt.Println("\n## " + funcName())

	var i IB
	i = S{}

	fmt.Println(i.(interface {
		walk()
	}))

	// PrimaryExpression must always evaluate to interface type otherwise it's a
	// compile-time error:
	// S{}.(IB) // invalid type assertion: S literal.(IB) (non-interface type S on left)

}

/*

Dynamic type

Besides static type that all variables have (it's a type from variable's
declaration), variables of interface type also have a dynamic type. It's a
type of value currently set in interface type variable. Over the course of
program execution variable of interface type has the same static type but
its dynamic type can change as different values implementing desired
interface will be assigned:

*/

type I interface {
	walk()
}

type A struct{}

func (a A) walk() {}

type B struct{}

func (b B) walk() {}

func DynamicType() {
	fmt.Println("\n## " + funcName())

	var i I
	i = A{} // dynamic type of i is A
	fmt.Printf("%T\n", i.(A))
	i = B{} // dynamic type of i is B
	fmt.Printf("%T\n", i.(B))
}

/*

Interface type

If T from v.(T) is an interface type then such assertion checks if dynamic type of v implements interface T:

*/

type J interface {
	quack()
}
type K interface {
	bark()
}

// type S struct{}
// func (s S) walk()  {}
// func (s S) quack() {}

func InterfaceType() {
	fmt.Println("\n## " + funcName())

	var i I
	i = S{}
	fmt.Printf("%T\n", i.(J))
	// Xfmt.Printf(?%T\n?, i.(K))
	// panic: interface conversion: main.S is not main.K: missing method bark
}

/*

Not interface type

If T from v.(T) is not an interface type then such assertion checks if dynamic type of v is identical to T:

*/

// type A struct{}

// func (a A) walk() {}

// type B struct{}

// func (b B) walk() {}

func NotInterfaceType() {
	fmt.Println("\n## " + funcName())

	var i I
	i = A{}
	fmt.Printf("%T\n", i.(A))
	// X: fmt.Printf("%T\n", i.(B))
	// panic: interface conversion: main.I is main.A, not main.B
}

/*

Type passed in non-interface type case must implement interface I as not fulfilling this requirement will be caught while compilation:

	type C struct{}
	fmt.Printf("%T\n", i.(C))

outputs:

impossible type assertion: C does not implement I (missing walk method)

No panic!

In all above cases when assertion doesn't hold run-time panic will be triggered. To handle failures gracefully there are special forms of assignments or initializations:


package nopanic

type I interface {
    walk()
}

type A struct {
    name string
}
func (a A) walk() {}
type B struct {
    name string
}
  func (b B) walk() {}
}

func NoPanic() {
    var i I
    i = A{name: "foo"}
    valA, okA := i.(A)
    fmt.Printf("%#v %#v\n", valA, okA)
    valB, okB := i.(B)
    fmt.Printf("%#v %#v\n", valB, okB)
}

*/
