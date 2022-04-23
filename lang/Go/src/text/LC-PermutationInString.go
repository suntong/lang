// https://leetcode.com/problems/permutation-in-string/
// https://books.halfrost.com/leetcode/ChapterFour/0500~0599/0567.Permutation-in-String/

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
	s, t   string
	result bool
}

func mustEqual(t *testing.T, res, result bool) {
	if res != result {
		t.Errorf(`expected "%v" but got "%v"`, result, res)
	} else {
		fmt.Println("matched")
	}
}

func TestIt(t *testing.T) {
	testData := []testCase{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"uvw", "eidboaoo", false},
		{"daibo", "eidbaooo", true},
		{"xyz", "pfijyxzskm", true},
		{"XYZ", "PFYIJXZSKM", false},
		{"", "", false},
	}

	for _, tc := range testData {
		mustEqual(t, checkInclusion(tc.s, tc.t), tc.result)
	}
}

func checkInclusion(s1 string, s2 string) bool {
	var freq [256]int
	if len(s2) == 0 || len(s2) < len(s1) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		freq[s1[i]]++
	}
	left, right, count := 0, 0, len(s1)

	for right < len(s2) {
		if freq[s2[right]] >= 1 {
			count--
			fmt.Printf("%c", s2[right])
		}
		if count == 0 {
			fmt.Println()
			return true
		}
		freq[s2[right]]--
		right++
		if right-left == len(s1) {
			if freq[s2[left]] >= 0 {
				count++
			}
			freq[s2[left]]++
			left++
			fmt.Println()
		}
	}
	return false
}
