// https://leetcode.com/problems/permutations/
// https://medium.com/@CalvinChankf/a-general-approach-for-subsets-combinations-and-permutations-5c8fe3aff0ae

package main

import (
	"fmt"
)

func main() {
	fmt.Print(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	res := [][]int{}
	dfs(nums, []int{}, &res)
	return res
}

func dfs(nums, path []int, res *[][]int) {
	b := make([]int, len(path))
	copy(b, path)
	if len(nums) == 0 {
		*res = append(*res, b)
	}
	for i := 0; i < len(nums); i++ {
		available := []int{}
		available = append(available, nums[:i]...)
		// fmt.Println(i, available)
		available = append(available, nums[i+1:]...)
		// fmt.Println(i, available)
		dfs(available, append(b, nums[i]), res)
	}
}
