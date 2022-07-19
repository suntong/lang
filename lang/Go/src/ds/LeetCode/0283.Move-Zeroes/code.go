package leetcode

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	j := len(nums) - 1
	for ; j > 0; j-- {
		if nums[j] != 0 {
			break
		}
	}

	for i := 0; i < j; i++ {
		if nums[i] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}
}
