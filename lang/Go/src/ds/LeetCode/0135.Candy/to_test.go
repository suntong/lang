package leetcode

import (
	"fmt"
	"testing"
)

type question135 struct {
	para135
	ans135
}

// para 是参数
// one 代表第一个参数
type para135 struct {
	ratings []int
}

// ans 是答案
// one 代表第一个答案
type ans135 struct {
	one int
}

func Test_Problem135(t *testing.T) {

	qs := []question135{

		{
			para135{[]int{}},
			ans135{0},
		},

		{
			para135{[]int{0}},
			ans135{1},
		},

		{
			para135{[]int{1, 0, 2}},
			ans135{5},
		},

		{
			para135{[]int{1, 2, 2}},
			ans135{4},
		},

		{
			para135{[]int{1, 2, 3, 3}},
			ans135{7},
		},

		{
			para135{[]int{3, 2, 1, 3}},
			ans135{8},
		},

		{
			para135{[]int{1, 3, 2, 1, 3}},
			ans135{9},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 135------------------------\n")

	for _, q := range qs {
		a, p := q.ans135.one, q.para135
		res := candy(p.ratings)
		if a != res {
			t.Errorf(`For "%v", expected "%v" but got "%v"`, p, a, res)
		}
	}
	fmt.Printf("\n\n\n")
}
