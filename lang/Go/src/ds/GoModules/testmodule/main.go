package main

import (
	"fmt"

	"testmodule/mypackage"
)

func main() {
	fmt.Println("Hello, Modules!")

	mypackage.PrintHello()

	fmt.Println(mypackage.Double(3))
}
