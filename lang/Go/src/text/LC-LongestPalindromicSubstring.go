// https://leetcode.com/problems/longest-palindromic-substring/
// https://www.geeksforgeeks.org/longest-palindrome-substring-set-1/

package main

import (
	"fmt"

	"github.com/suntong/testing"
)

type testCase struct {
	s      string
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
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"forgeeksskeegfor", "geeksskeeg"},
		{"Geeks", "ee"},
		{"", ""},
	}

	for _, tc := range testData {
		mustEqual(t, longestPalindrome(tc.s), tc.result)
	}
}

func mustEqual(t *testing.T, res, result string) {
	if res != result {
		t.Errorf(`expected "%s" but got "%s" matches`, result, res)
	} else {
		fmt.Println("matched")
	}
}

// 滑动窗口解法，时间复杂度 O(n^2)，空间复杂度 O(1)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		// 移动到相同字母的最右边（如果有相同字母）
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		// 找到回文的边界
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > pr-pl {
			pl, pr = left, right
		}
		// 重置到下一次寻找回文的中心
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}
