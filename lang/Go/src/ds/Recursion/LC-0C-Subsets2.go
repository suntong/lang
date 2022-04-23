// https://leetcode.com/problems/subsets-ii/
// https://medium.com/@CalvinChankf/a-general-approach-for-subsets-combinations-and-permutations-5c8fe3aff0ae

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(subsets([]int{1, 2, 2}))
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums) // 这里是去重的关键逻辑, first sort the numbers
	dfs(nums, []int{}, &res)
	return res
}

func dfs(nums, path []int, res *[][]int) {
	b := make([]int, len(path))
	copy(b, path)
	*res = append(*res, b)
	for i := 0; i < len(nums); i++ {
		// check if the current number is the same as the previous number.
		if i == 0 || (i > 0 && nums[i-1] != nums[i]) {
			dfs(nums[i+1:], append(b, nums[i]), res)
		}
	}
}
