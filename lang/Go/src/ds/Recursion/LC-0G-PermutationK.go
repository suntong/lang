// https://leetcode.com/problems/subsets/
// https://medium.com/@CalvinChankf/a-general-approach-for-subsets-combinations-and-permutations-5c8fe3aff0ae

package main

import (
	"fmt"
)

func main() {
	fmt.Print(permute("123"))
}

func permute(nums string) []string {
	res := []string{}
	dfs(nums, "", &res)
	return res
}

func dfs(nums, path string, res *[]string) {
	if len(nums) == 0 {
		*res = append(*res, path)
	}
	for i := 0; i < len(nums); i++ {
		dfs(nums[:i]+nums[i+1:], path+string(nums[i]), res)
	}
}
