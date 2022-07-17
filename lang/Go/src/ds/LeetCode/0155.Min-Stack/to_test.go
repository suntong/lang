package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem155(t *testing.T) {
	obj1 := Constructor155()
	obj1.Push(1)
	fmt.Printf("obj1 = %v\n", obj1)
	obj1.Push(0)
	fmt.Printf("obj1 = %v\n", obj1)
	obj1.Push(10)
	fmt.Printf("obj1 = %v\n", obj1)
	obj1.Pop()
	fmt.Printf("obj1 = %v\n", obj1)
	param3 := obj1.Top()
	fmt.Printf("param_3 = %v\n", param3)
	param4 := obj1.GetMin()
	fmt.Printf("param_4 = %v\n", param4)
}

func Test_MinStack(t *testing.T) {
	minStack := Constructor155()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	r := minStack.GetMin() // --> Returns -3.
	if r != -3 {
		t.Errorf("Returns != -3")
	}
	minStack.Pop()
	minStack.Top()
	r = minStack.GetMin() // --> Returns -2.
	if r != -2 {
		t.Errorf("Returns != -2")
	}
}
