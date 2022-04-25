// https://leetcode.com/problems/subsets/
// https://medium.com/@CalvinChankf/a-general-approach-for-subsets-combinations-and-permutations-5c8fe3aff0ae

package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute("123", 3))  // 213
	fmt.Println(permute("1234", 9)) // 2314
}

func permute(nums string, _k int) string {
	res, k := "", _k
	dfs(nums, "", &res, &k)
	return res
}

func dfs(nums, path string, res *string, k *int) {
	if len(nums) == 0 {
		*k--
		if *k == 0 {
			*res = path
		}
	}
	for i := 0; i < len(nums); i++ {
		dfs(nums[:i]+nums[i+1:], path+string(nums[i]), res, k)
	}
}
