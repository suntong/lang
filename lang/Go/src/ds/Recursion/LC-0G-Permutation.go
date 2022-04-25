// https://leetcode.com/problems/subsets/
// https://medium.com/@CalvinChankf/a-general-approach-for-subsets-combinations-and-permutations-5c8fe3aff0ae

package main

import (
	"fmt"
)

func main() {
	fmt.Print(subsets("123"))
}

func subsets(nums string) []string {
	res := []string{}
	dfs(nums, "", &res)
	return res
}

func dfs(nums, path string, res *[]string) {
	*res = append(*res, path)
	for i := 0; i < len(nums); i++ {
		dfs(nums[i+1:], path+string(nums[i]), res)
	}
}
