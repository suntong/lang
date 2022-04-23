// https://leetcode.com/problems/combinations/
// https://medium.com/@CalvinChankf/a-general-approach-for-subsets-combinations-and-permutations-5c8fe3aff0ae

package main

import (
	"fmt"
)

func main() {
	fmt.Print(combine(4, 2))
}

func combine(n, k int) [][]int {
	res := [][]int{}
	if k < 1 || k > n {
		return res
	}
	// construct the arrary from 1 to n
	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, i+1)
	}
	dfs(nums, []int{}, k, &res)
	return res
}

func dfs(nums, path []int, k int, res *[][]int) {
	b := make([]int, len(path))
	copy(b, path)
	if len(path) == k {
		*res = append(*res, b)
		return
	}
	for i := 0; i < len(nums); i++ {
		// last n elements of slice
		dfs(nums[i+1:], append(b, nums[i]), k, res)
	}
}

// [[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]
