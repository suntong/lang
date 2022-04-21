// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

import (
	"fmt"

	"github.com/suntong/testing"
)

type testCase struct {
	tCase  string
	result int
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
		{"ABCABCBB", 3},
		{"ABDEFGABEF", 6},
		{"BBBB", 1},
		{"GEEKSFORGEEKS", 7},
	}

	for _, tc := range testData {
		tcs := tc.tCase
		mustEqual(t, lengthOfLongestSubstringHashSet(tcs), tc.result)
	}
}

func mustEqual(t *testing.T, res, result int) {
	if res != result {
		t.Errorf(`expected "%d" but got "%d" matches`, result, res)
	} else {
		fmt.Println("matched")
	}
}

// HashSet -- Optimised Sliding Window
// https://www.geeksforgeeks.org/length-of-the-longest-substring-without-repeating-characters/
// https://www.interviewbit.com/blog/longest-substring-without-repeating-characters/
/*

   This solution uses extra space to store the last indexes of already
   visited characters. The idea is to scan the string from left to right,
   keep track of the maximum length Non-Repeating Character Substring seen
   so far in res. When we traverse the string, to know the length of current
   window we need two indexes.

   1) Ending index ( j ) : We consider current index as ending index.

   2) Starting index ( i ) : It is same as previous window if current
   character was not present in the previous window. To check if the current
   character was present in the previous window or not, we store last index
   of every character in an array lasIndex[]. If lastIndex[str[j]] + 1 is
   more than previous start, then we updated the start index i. Else we keep
   same i.

*/

func lengthOfLongestSubstringHashSet(s string) int {
	fmt.Printf("------ %s\n", s)
	i, j, res := 0, 0, 0
	lastIndex := make(map[byte]int, len(s))
	for j < len(s) {
		// Find the last index of str[j]
		// Update i (starting index of current window)
		// as maximum of current value of i and last index plus 1
		// if idx, ok := lastIndex[s[j]]; ok && idx >= i {
		// 	i = idx + 1
		// }
		i = max(i, lastIndex[s[j]]+1)
		fmt.Println(i, j, lastIndex[s[j]], j-i, res)
		// Update last index of j.
		lastIndex[s[j]] = j
		j++
		// Update result if we get a larger window
		res = max(res, j-i)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
