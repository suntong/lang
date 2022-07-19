// https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0063.Unique-Paths-II/

package main

import (
	"fmt"

	"github.com/suntong/testing"
)

func main() {
	var t *testing.T = testing.NewT()
	Test_Problem63(t)
	t.Report()
}

//==========================================================================
// tests

type question63 struct {
	para63
	ans63
}

// para 是参数
// one 代表第一个参数
type para63 struct {
	og [][]int
}

// ans 是答案
// one 代表第一个答案
type ans63 struct {
	one int
}

func Test_Problem63(t *testing.T) {

	qs := []question63{

		{
			para63{[][]int{
				{0, 1, 0},
				{0, 0, 0},
				{0, 0, 0},
			}},
			ans63{3},
		},

		{
			para63{[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			}},
			ans63{2},
		},

		{
			para63{[][]int{
				{0, 0},
				{1, 1},
				{0, 0},
			}},
			ans63{0},
		},

		{
			para63{[][]int{
				{0, 1, 0, 0},
				{1, 0, 0, 0},
				{0, 0, 0, 0},
			}},
			ans63{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 63------------------------\n")

	for _, q := range qs {
		_, p := q.ans63, q.para63
		fmt.Printf("%v\t:%v\n", p, uniquePathsWithObstacles(p.og))
	}
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}
