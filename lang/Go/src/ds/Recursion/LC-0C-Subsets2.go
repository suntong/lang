// https://leetcode.com/problems/subsets-ii/
// https://medium.com/@CalvinChankf/a-general-approach-for-subsets-combinations-and-permutations-5c8fe3aff0ae

package main

import (
	"fmt"
)

func main() {
	fmt.Print(subsets([]int{1, 2, 2}))
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, []int{}, &res)
	return res
}

func dfs(nums, path []int, res *[][]int) {
	b := make([]int, len(path))
	copy(b, path)
	*res = append(*res, b)
	for i := 0; i < len(nums); i++ {
		// last n elements of slice
		dfs(nums[i+1:], append(b, nums[i]), res)
	}
}
