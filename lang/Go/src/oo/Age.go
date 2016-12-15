// http://www.tapirgames.com/blog/golang-oo-programming

package main

import "fmt"

type Age struct {
	age int
}

func (age *Age) GrowUp() {
	age.age++
}

func (age Age) CanDrink() bool {
	age.age++
	return age.age >= 18
}

// Defining the above method makes
// the following method is also defined implicitly.

/*
func (age *Age) CanDrink() bool {
	return (*age).CanDrink()
}

When a method is defined for a non-interface type and non-pointer named type
T, another method with the same name is also defined for type *T,
implicitly.

If a method is defined for type *T, there is not an implicit method with
same name defined for type T.

*/

func main() {
	age := Age{16}

	fmt.Println(age.CanDrink()) // false
	fmt.Println(age.CanDrink()) // false

	fmt.Println((&age).CanDrink()) // false
	fmt.Println((&age).CanDrink()) // false

	// age is a variable. Variables are always addressable.
	// Compilers will automatically convert age into (&age).
	age.GrowUp()                // <=> (&age).GrowUp()
	fmt.Println(age)            // 17
	fmt.Println(age.CanDrink()) // true

	age2 := &Age{16}
	fmt.Println(age2.CanDrink()) // false
	fmt.Println(age2.CanDrink()) // false

}
