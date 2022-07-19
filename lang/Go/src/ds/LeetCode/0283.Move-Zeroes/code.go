package leetcode

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	j := len(nums) - 1
	for i := 0; i < j; i++ {
		for ; j > i; j-- {
			if nums[j] != 0 {
				break
			}
		}
		if nums[i] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}
}
