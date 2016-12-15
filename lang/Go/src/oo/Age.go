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

Although none methods explicitly defined on a pointer type *T are methods of
T, receiver values of type T can also call methods of *T if the receivers
are addressable. For callings of this case, Go compilers will automatically
take the addresses of the non-pointer receivers. This is totally for sake of
convenience. It doesn't mean the pointer methods are methods of T.

(Methods defined for non-pointer types are called non-pointer
methods. Methods defined for pointer types are called pointer methods.)

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
