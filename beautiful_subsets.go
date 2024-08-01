package main

import (
	"fmt"
)

func beautifulSubsets(nums []int, k int) int {
	count := 0
	return backtrack(nums, k, 0, []int{}, &count)
}

func backtrack(nums []int, k int, index int, currentSubset []int, count *int) int {
	// Base case: If we've considered all elements
	if index == len(nums) {
		if len(currentSubset) > 0 {
			*count++
		}
		return *count
	}
	// Decision to include nums[index]
	if isBeautiful(currentSubset, nums[index], k) {
		currentSubset = append(currentSubset, nums[index])
		backtrack(nums, k, index+1, currentSubset, count)
		currentSubset = currentSubset[:len(currentSubset)-1]
	}
	// Decision to exclude nums[index]
	backtrack(nums, k, index+1, currentSubset, count)
	return *count
}

// Helper function to check if adding a new number keeps the subset beautiful
func isBeautiful(subset []int, num, k int) bool {
	for _, s := range subset {
		if abs(s-num) == k {
			return false
		}
	}
	return true
}

// Helper function to calculate the absolute value
func abs(a int) int {
	if a < 0 {
		return a
	}
	return a
}
func main() {
	nums := []int{1}
	k := 1
	fmt.Println("Number of beautiful subsets:", beautifulSubsets(nums, k))
}
