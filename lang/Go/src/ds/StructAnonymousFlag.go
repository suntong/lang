package main

import (
	"flag"
	"fmt"
)

type employee struct {
	name   string
	age    int
	salary int
	address
}

type address struct {
	city    string
	country string
}

func main() {
	address := address{city: "London", country: "UK"}

	emp := employee{name: "Sam", age: 31, salary: 2000, address: address}

	fmt.Printf("City: %s\n", emp.address.city)
	fmt.Printf("Country: %s\n", emp.address.country)

	// nested struct’s fields are directly accessed
	emp.country = "Canada"
	flag.StringVar(&emp.city, "c", "foo", "a city var")
	flag.Parse()
	fmt.Printf("City: %s\n", emp.city)
	fmt.Printf("Country: %s\n", emp.country)
	fmt.Printf("%+v\n", emp)
}

/*

$ go run StructAnonymousFlag.go
City: London
Country: UK
City: foo
Country: Canada

$ go run StructAnonymousFlag.go -c Toronto
City: London
Country: UK
City: Toronto
Country: Canada
{name:Sam age:31 salary:2000 address:{city:Toronto country:Canada}}

*/
