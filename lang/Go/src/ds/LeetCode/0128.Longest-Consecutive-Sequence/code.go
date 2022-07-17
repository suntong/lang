package leetcode

// 解法一 map，时间复杂度 O(n)
func longestConsecutive(nums []int) int {
	res, numMap := 0, map[int]int{}
	for _, num := range nums {
		if numMap[num] == 0 {
			left, right, sum := 0, 0, 0
			if numMap[num-1] > 0 {
				left = numMap[num-1]
			} else {
				left = 0
			}
			if numMap[num+1] > 0 {
				right = numMap[num+1]
			} else {
				right = 0
			}
			// sum: length of the sequence n is in
			sum = left + right + 1
			// keep track of the max length
			res = max(res, sum)

			numMap[num] = sum
			// extend the length to the boundary(s) of the sequence
			// will do nothing if n has no neighbors
			numMap[num-left] = sum
			numMap[num+right] = sum
		} else {
			continue
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
