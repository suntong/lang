package main

import (
	"fmt"
)

type Vehicle struct {
	wheelCount int
}

type Car struct {
	Vehicle //anonymous field Vehicle
}

func main() {
	{
		v := Vehicle{4}
		fmt.Printf("Hello, %#v\n", v)

		var c Car
		// c = v // cannot use v (type Vehicle) as type Car in assignment
		// c = Car(v) // cannot convert v (type Vehicle) to type Car
		c = Car{v}
		fmt.Printf("Hello, %#v\n", c)
	}
	{
		v := &Vehicle{4}
		fmt.Printf("Hello, %#v\n", v)

		var c *Car
		c = &Car{*v}
		fmt.Printf("Hello, %#v\n", c)
		fmt.Printf("Wheels: %d\n", c.wheelCount)
	}
}
