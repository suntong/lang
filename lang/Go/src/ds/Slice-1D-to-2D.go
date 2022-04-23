// https://leetcode.com/problems/convert-1d-array-into-2d-array/

package main

import "fmt"

func main() {
	fmt.Println(construct2DArray([]int{1, 2, 3, 4}, 2, 2))
	fmt.Println(construct2DArray([]int{1, 2, 3}, 1, 3))
	fmt.Println(construct2DArray([]int{1, 2}, 1, 1))
}

func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return [][]int{}
	}
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = original[n*i : n*(i+1)]
	}
	return res
}
