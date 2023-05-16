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

type Truck struct {
	Vehicle //anonymous field Vehicle
	Payload string
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
		c.wheelCount = 5
		fmt.Printf("Wheels: %d\n", c.wheelCount)
	}
	{
		t1 := Truck{Vehicle{4}, "1.5T"}
		fmt.Printf("Hello, %#v\n", t1)

		v := Vehicle{6}
		t2 := Truck{v, "2.5T"}
		fmt.Printf("Hello, %#v\n", t2)
	}
}
