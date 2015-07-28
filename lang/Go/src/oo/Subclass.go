package main

import "fmt"

type Vehicle struct {
	wheelCount int
}

// Vehicler is an interface and has a single function numberOfWheels that returns an int.
// In Go, the convention is to "er" a type to indicate that it is an interface.
type Vehicler interface {
	numberOfWheels() int
}

// define a behavior for Vehicle
func (vehicle Vehicle) numberOfWheels() int {
	return vehicle.wheelCount
}

type Car struct {
	Vehicle //anonymous field Vehicle
	Maker   string
}

func main() {
	{
		// From http://golangtutorials.blogspot.ca/2011/06/inheritance-and-subclassing-in-go-or.html
		c := Car{Vehicle{4}, "Ford"}
		fmt.Println("A Car has this many wheels: ", c.numberOfWheels()) //no method defined for Car, but we have the same behavior as Vehicle.

		/*

		   we have only defined a method or behavior for Vehicle. Since we then
		   defined Vehicle as an anonymous field in Car, the latter class
		   automatically can call on all the visible behaviors/methods of the
		   anonymous field type. So here, we have not subclassed a parent class,
		   but composed it. But the effect is the very same

		*/

	}
	{
		/*

			From http://play.golang.org/p/iUuuyxY9Yy
			By Roberto Zanotto @gmail

			from the language specification, https://golang.org/ref/spec

			For struct literals the following rules apply:

			- A key must be a field name declared in the LiteralType.
			- An element list that does not contain any keys must list an element for each struct field in the order in which the fields are declared.
			- If any element has a key, every element must have a key.
			- An element list that contains keys does not need to have an element for each struct field. Omitted fields get the zero value for that field.
			- A literal may omit the element list; such a literal evaluates to the zero value for its type.
			- It is an error to specify an element for a non-exported field of a struct belonging to a different package


		*/
		c := Car{Vehicle{4}, "Ferrari"}
		v := Vehicle{3}
		c.Vehicle = v
		// or
		c = Car{Vehicle: v}
		fmt.Println("A Car has this many wheels: ", c.wheelCount) //not directly defined in Car, but use as the same.

		// From http://golangtutorials.blogspot.ca/2011/06/interfaces-in-go.html
		v2 := Vehicler(c)
		fmt.Println("A Vehicle has this many wheels: ", v2.numberOfWheels())
	}
}

/*

A Car has this many wheels:  4
A Car has this many wheels:  3
A Vehicle has this many wheels:  3

*/
