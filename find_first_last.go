package main

import "fmt"

/*
Input: nums = [5,7,7,7,8,10], target = 8
Output: [3,4]
*/
func searchRange(nums []int, target int) []int {
	return []int{findPosition(nums, target, true), findPosition(nums, target, false)}
}

func findPosition(nums []int, target int, isFirst bool) int {
	lo := 0
	hi := len(nums) - 1
	result := -1
	for lo <= hi {
		mid := (lo + hi) / 2
		if target == nums[mid] {
			result = mid
			if isFirst {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if target > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}

func main() {
	fmt.Println(searchRange([]int{7, 7, 7, 7, 7, 10}, 7))
}
