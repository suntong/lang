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

func test1() {
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
		fmt.Printf("%d-wheelTruck, %#v\n", t1.Vehicle.wheelCount, t1)

		v := Vehicle{6}
		t2 := Truck{v, "2.5T"}
		fmt.Printf("%d-wheelTruck, %+v\n", t2.wheelCount, t2)
	}
	{
		// unknown field 'wheelCount' in struct literal
		// mixture of field:value and value elements in struct literal
		// t1 := Truck{wheelCount: 3, "0.5T"}
		// t1 := Car{wheelCount: 3, "0.5T"}
		// fmt.Printf("Three-wheel car, %#v\n", t1)
	}
}

type T1 struct {
	T1_Text string
}

type T2 struct {
	T2_Text string
	T1
}

func test2() {
	// unknown field 'T1_Text' in struct literal of type T2
	// t2 := T2{T2_Text: "Test2", T1_Text: "Test1"}

	// mixture of field:value and value elements in struct literal
	// X: t := T2{T2_Text: "Test", T1{T1_Text: "Test"}}
	// https://stackoverflow.com/a/47379416/2125837
	t2 := T2{T2_Text: "Test2", T1: T1{T1_Text: "Test1"}}
	fmt.Println(t2)
}

func main() {
	test1()
	test2()
}
