// https://leetcode.com/problems/permutations-ii/
// https://medium.com/@CalvinChankf/a-general-approach-for-subsets-combinations-and-permutations-5c8fe3aff0ae

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print(permute([]int{1, 1, 2}))
}

func permute(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums) // 这里是去重的关键逻辑, first sort the numbers
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
		// check if the current number is the same as the previous number.
		if i == 0 || (i > 0 && nums[i-1] != nums[i]) {
			available := []int{}
			available = append(available, nums[:i]...)
			available = append(available, nums[i+1:]...)
			dfs(available, append(b, nums[i]), res)
		}
	}
}
