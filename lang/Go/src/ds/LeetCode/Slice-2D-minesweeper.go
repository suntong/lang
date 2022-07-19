// https://github.com/halfrost/LeetCode-Go/blob/master/leetcode/0529.Minesweeper/

package main

import (
	"fmt"

	"github.com/suntong/testing"
)

func main() {
	var t *testing.T = testing.NewT()
	Test_Problem529(t)
	t.Report()
}

//==========================================================================
// tests

type question529 struct {
	para529
	ans529
}

// para 是参数
// one 代表第一个参数
type para529 struct {
	b     [][]byte
	click []int
}

// ans 是答案
// one 代表第一个答案
type ans529 struct {
	one [][]byte
}

func Test_Problem529(t *testing.T) {

	qs := []question529{

		{
			para529{[][]byte{
				{'M', 'M', 'E', 'E', 'E'},
				{'M', 'E', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'M', 'E', 'E', 'M', 'E'},
			}, []int{2, 1}},
			ans529{[][]byte{
				{'M', 'M', 'E', 'E', 'E'},
				{'M', 'E', 'M', 'E', 'E'},
				{'E', 3, 'E', 'E', 'E'},
				{'M', 'E', 'E', 'M', 'E'},
			}},
		},

		{
			para529{[][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			}, []int{3, 0}},
			ans529{[][]byte{
				{0, 1, 'E', 1, 0},
				{0, 1, 'M', 1, 0},
				{0, 1, 1, 1, 0},
				{0, 0, 0, 0, 0},
			}},
		},

		{
			para529{[][]byte{
				{'E', 'E', 'E', 'E', 'M'},
				{'E', 'E', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			}, []int{3, 0}},
			ans529{[][]byte{
				{0, 1, 'E', 'E', 'M'},
				{0, 1, 'M', 2, 1},
				{0, 1, 1, 1, 0},
				{0, 0, 0, 0, 0},
			}},
		},

		{
			para529{[][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			}, []int{1, 2}},
			ans529{[][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'X', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 529------------------------\n")

	for _, q := range qs {
		a, p := q.ans529, q.para529
		fmt.Printf("\t%v\n\t%v\n", a.one, updateBoard(p.b, p.click))
	}
}

var dir8 = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
}

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	mineMap := make([][]int, len(board))
	for i := range board {
		mineMap[i] = make([]int, len(board[i]))
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'M' {
				mineMap[i][j] = -1
				for _, d := range dir8 {
					nx, ny := i+d[0], j+d[1]
					if isInBoard(board, nx, ny) && mineMap[nx][ny] >= 0 {
						mineMap[nx][ny]++
					}
				}
			}
		}
	}
	mineSweeper(click[0], click[1], board, mineMap, dir8)
	fmt.Println(mineMap)
	return board
}

func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func mineSweeper(x, y int, board [][]byte, mineMap [][]int, dir8 [][]int) {
	if board[x][y] != 'M' && board[x][y] != 'E' {
		return
	}
	if mineMap[x][y] == -1 {
		board[x][y] = 'X'
	} else if mineMap[x][y] > 0 {
		board[x][y] = byte(mineMap[x][y])
	} else {
		board[x][y] = 0
		for _, d := range dir8 {
			nx, ny := x+d[0], y+d[1]
			if isInBoard(board, nx, ny) && mineMap[nx][ny] >= 0 {
				mineSweeper(nx, ny, board, mineMap, dir8)
			}
		}
	}
}
