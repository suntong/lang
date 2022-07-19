package main

import "fmt"

/*
Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""

Example 3:
Input: strs = ["flower","flow","flowing"]
Output: "flow"

Example 4:
Input: strs = ["flower","glow","flowing"]
Output: ""

*/

func main() {
	fmt.Println(findPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println("---")
	fmt.Println(findPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println("---")
	fmt.Println(findPrefix([]string{"flower", "flow", "flowing"}))
}

func findPrefix(input []string) string {
	prefix := ""

Loop:
	for j := 0; j < len(input[0]); j++ {
		prefix += string(input[0][j])
		for i := 1; i < len(input); i++ {
			if len(input[i]) <= j || input[i][j] != prefix[j] {
				prefix = prefix[:j]
				break Loop
			}
		}
	}
	return prefix
}

func findPrefix0(input []string) string {
	prefix := input[0]

	for i := 1; i < len(input); i++ {
		for j := 0; j < len(prefix); j++ {
			// compare prefix with input string
			if len(input[i]) <= j || input[i][j] != prefix[j] {
				prefix = prefix[:j]
			}
		}
	}
	return prefix
}
