// https://leetcode.com/problems/minimum-window-substring/
// https://www.geeksforgeeks.org/find-the-smallest-window-in-a-string-containing-all-characters-of-another-string/

package main

import (
	"fmt"

	"github.com/suntong/testing"
)

type testCase struct {
	s, t   string
	result string
}

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

func TestIt(t *testing.T) {
	testData := []testCase{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"this is a test string", "tist", "t stri"},
		{"geeksforgeeks", "ork", "ksfor"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"", "", ""},
	}

	for _, tc := range testData {
		mustEqual(t, minWindow(tc.s, tc.t), tc.result)
	}
}

func mustEqual(t *testing.T, res, result string) {
	if res != result {
		t.Errorf(`expected "%s" but got "%s" matches`, result, res)
	} else {
		fmt.Println("matched")
	}
}

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	var tFreq, sFreq [256]int
	result, left, right, finalLeft, finalRight, minW, count := "", 0, -1, -1, -1, len(s)+1, 0

	// creating map
	for i := 0; i < len(t); i++ {
		tFreq[t[i]-'a']++
	}

	// Traversing the window
	for left < len(s) {
		if right+1 < len(s) && count < len(t) {
			sFreq[s[right+1]-'a']++
			if sFreq[s[right+1]-'a'] <= tFreq[s[right+1]-'a'] {
				count++
			}
			right++
		} else {
			if right-left+1 < minW && count == len(t) {
				minW = right - left + 1
				finalLeft = left
				finalRight = right
			}
			if sFreq[s[left]-'a'] == tFreq[s[left]-'a'] {
				count--
			}
			sFreq[s[left]-'a']--
			left++
		}
	}
	if finalLeft != -1 {
		result = string(s[finalLeft : finalRight+1])
	}
	return result
}
