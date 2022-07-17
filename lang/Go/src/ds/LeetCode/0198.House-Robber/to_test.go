package leetcode

import (
	"fmt"
	"testing"
)

type question198 struct {
	para198
	ans198
}

// para 是参数
// one 代表第一个参数
type para198 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans198 struct {
	one int
}

func Test_Problem198(t *testing.T) {

	qs := []question198{

		{
			para198{[]int{1, 2}},
			ans198{2},
		},

		{
			para198{[]int{1, 2, 3, 1}},
			ans198{4},
		},
		{
			para198{[]int{2, 7, 9, 3, 1}},
			ans198{12},
		},
		{
			para198{[]int{2, 1, 7, 9, 3, 1}},
			ans198{12},
		},
		{
			para198{[]int{1, 2, 1, 7, 9, 3, 1}},
			ans198{12},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 198------------------------\n")

	for _, q := range qs {
		a, p := q.ans198.one, q.para198
		res := rob198(p.one)
		if a != res {
			t.Errorf(`For "%v", expected "%v" but got "%v"`, p, a, res)
		}
	}
	fmt.Printf("\n\n\n")
}
