// https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0074.Search-a-2D-Matrix

package main

import (
	"fmt"

	"github.com/suntong/testing"
)

func main() {
	var t *testing.T = testing.NewT()
	Test_Problem74(t)
	t.Report()
}

//==========================================================================
// tests

type question74 struct {
	para74
	ans74
}

// para 是参数
type para74 struct {
	matrix [][]int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans74 struct {
	one bool
}

func Test_Problem74(t *testing.T) {

	qs := []question74{

		{
			para74{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 3},
			ans74{true},
		},

		{
			para74{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 13},
			ans74{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 74------------------------\n")

	for _, q := range qs {
		a, p := q.ans74, q.para74
		fmt.Printf("%v\t%v:%v\n", p, a.one, searchMatrix(p.matrix, p.target))
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, low, high := len(matrix[0]), 0, len(matrix[0])*len(matrix)-1
	for low <= high {
		mid := low + (high-low)>>1
		if matrix[mid/m][mid%m] == target {
			return true
		} else if matrix[mid/m][mid%m] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
