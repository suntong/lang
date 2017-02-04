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
	"testing"
)

func main() {
	fmt.Println("Hello World")
	TestThem()
}

func TestThem() {
	var t *testing.T
	TestSum(t)
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
