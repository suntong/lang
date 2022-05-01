// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

import (
	"fmt"

	"github.com/suntong/testing"
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Function main

func main() {
	var t *testing.T = testing.NewT()
	TestIt(t)
	t.Report()
}

//==========================================================================
// test functions

type testCase struct {
	tCase  string
	result bool
}

func TestIt(t *testing.T) {
	testData := []testCase{
		{"121", true},
		{"1221", true},
		{"BBB", true},
		{"BBBB", true},
		{"ABDBA", true},
		{"ABDEA", false},
	}

	for _, tc := range testData {
		fmt.Print("Checking  palindrome  with  loop: ")
		mustEqual(t, isPalindromeLoop(tc.tCase), tc.result)
		fmt.Print("Checking palindrome without loop: ")
		mustEqual(t, isPalindromeRecursion(tc.tCase), tc.result)
	}
}

func mustEqual(t *testing.T, res, result bool) {
	if res != result {
		t.Errorf(`expected "%v" but got "%v"`, result, res)
	} else {
		fmt.Println("matched")
	}
}

func isPalindromeLoop(s string) bool {
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func isPalindromeRecursion(s string) bool {
	length := len(s)
	if length <= 1 {
		return true
	}
	if s[0] != s[length-1] {
		return false
	}
	return isPalindromeRecursion(s[1 : length-1])
}
