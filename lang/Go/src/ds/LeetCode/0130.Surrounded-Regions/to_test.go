package leetcode

import (
	"fmt"
	"testing"
)

type question130 struct {
	para130
	ans130
}

// para 是参数
// one 代表第一个参数
type para130 struct {
	one [][]byte
}

// ans 是答案
// one 代表第一个答案
type ans130 struct {
	one [][]byte
}

func Test_Problem130(t *testing.T) {

	qs := []question130{

		{
			para130{[][]byte{}},
			ans130{[][]byte{}},
		},

		{
			para130{[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'}}},
			ans130{[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}},
		},

		{
			para130{[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'O', 'X', 'X'}}},
			ans130{[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 130------------------------\n")

	for _, q := range qs {
		_, p := q.ans130, q.para130
		fmt.Printf("【input】:%v      ", p)
		solve1(p.one)
		fmt.Printf("【output】:%v      \n", p)
	}
	fmt.Printf("\n\n\n")
}
