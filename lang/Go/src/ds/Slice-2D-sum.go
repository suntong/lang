// https://leetcode.com/problems/shift-2d-grid/

package main

import "fmt"

var matrix [][]int = [][]int{
	{3, 0, 1, 4, 2},
	{5, 6, 3, 2, 1},
	{1, 2, 0, 1, 5},
	{4, 1, 0, 1, 7},
	{1, 0, 3, 0, 5},
}

func main() {
	m := Constructor(matrix)
	fmt.Println(m.SumRegion(2, 1, 4, 3)) // 8
	fmt.Println(m.SumRegion(1, 1, 2, 2)) // 11
	fmt.Println(m.SumRegion(1, 2, 2, 4)) // 12
}

type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{nil}
	}
	dp := make([][]int, len(matrix)+1)
	dp[0] = make([]int, len(matrix[0])+1)
	for i := range matrix {
		dp[i+1] = make([]int, len(matrix[i])+1)
		for j := range matrix[i] {
			dp[i+1][j+1] = matrix[i][j] + dp[i][j+1] + dp[i+1][j] - dp[i][j]
		}
	}
	return NumMatrix{dp}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	dp := this.dp
	return dp[row2+1][col2+1] - dp[row1][col2+1] - dp[row2+1][col1] + dp[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
