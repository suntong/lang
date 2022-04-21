// https://leetcode.com/problems/find-all-anagrams-in-a-string/
// https://www.geeksforgeeks.org/anagram-substring-search-search-permutations/

package main

import (
	"fmt"
	"reflect"

	"github.com/suntong/testing"
)

type testCase struct {
	s, p   string
	result []int
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
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
		{"BACDGABCDA", "ABCD", []int{0, 5, 6}},
		{"AAABABAA", "AABA", []int{0, 1, 4}},
	}

	for _, tc := range testData {
		mustEqual(t, findAnagrams(tc.s, tc.p), tc.result)
	}
}

func mustEqual(t *testing.T, res, result []int) {
	//fmt.Printf(`expected "%v", got "%v"\n`, result, res)
	// https://appdividend.com/2020/04/15/how-to-compare-two-slices-in-golang-with-example/
	if reflect.DeepEqual(res, result) {
		fmt.Println("matched")
	} else {
		t.Errorf(`expected "%v" but got "%v" matches`, result, res)
	}
}

func findAnagrams(s string, p string) []int {
	var freq [256]int
	result := []int{}
	if len(s) == 0 || len(s) < len(p) {
		return result
	}
	for i := 0; i < len(p); i++ {
		freq[p[i]]++
	}
	left, right, count := 0, 0, len(p)

	for right < len(s) {
		if freq[s[right]] >= 1 {
			count--
		}
		freq[s[right]]--
		right++
		if count == 0 {
			result = append(result, left)
		}
		if right-left == len(p) {
			if freq[s[left]] >= 0 {
				count++
			}
			freq[s[left]]++
			left++
		}

	}
	return result
}
