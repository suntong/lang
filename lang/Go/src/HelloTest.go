////////////////////////////////////////////////////////////////////////////
// Porgram: HelloTest.go
// Purpose: to run go testing code with "go run"
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: https://gobyexample.com/hello-world
//          https://semaphoreci.com/community/tutorials/how-to-test-in-go
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/suntong/testing"
)

func main() {
	fmt.Println("Hello World")
	TestThem()
}

/*

Hello World
true
--- PASS:  (0.00s)
Say hi
--- PASS:  (0.00s)
        HelloTest.go:42: Say bye
--- PASS:  (0.00s)

*/

// https://golang.org/src/testing/testing.go
func TestThem() {
	var t *testing.T = testing.NewT()
	fmt.Println(testing.Verbose())
	TestSomething(t)
	t.Report()
	TestPrintSomething(t)
	t.Report()
	TestSum(t)
	t.Report()
}

// https://smartystreets.com/blog/2015/02/go-testing-part-1-vanillla
func TestSomething(t *testing.T) {
	//t.Fail()
}

// http://stackoverflow.com/questions/23205419
func TestPrintSomething(t *testing.T) {
	fmt.Println("Say hi")
	t.Log("Say bye")
	//t.Error("Error is equivalent to Log followed by Fail")
}

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := 15
	actual := Sum(numbers)

	if actual != expected {
		t.Errorf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual)
	}
}

func Sum(numbers []int) int {
	sum := 0
	// This bug is intentional
	for _, n := range numbers {
		sum += n
	}
	return sum
}
