package main

func jump(nums []int) bool {
	right := len(nums) - 1
	for left := len(nums) - 2; left >= 0; left-- {
		currentDistanceToFinalStep := nums[left] + left
		if currentDistanceToFinalStep >= right {
			right = left
		}

	}
	return right == 0
}
