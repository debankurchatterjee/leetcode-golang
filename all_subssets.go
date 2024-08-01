package main

import "fmt"

// generateSubsets generates all subsets of nums and prints them.
func generateSubsets1(nums []int, index int, currentSubset []int) {
	if index == len(nums) {
		fmt.Println(currentSubset)
		return
	}

	// Include the current element
	currentSubset = append(currentSubset, nums[index])
	generateSubsets1(nums, index+1, currentSubset)

	// Backtrack: Exclude the current element (pop it)
	currentSubset = currentSubset[:len(currentSubset)-1]
	generateSubsets1(nums, index+1, currentSubset)
}

func main() {
	nums := []int{1, 2, 3}
	generateSubsets1(nums, 0, []int{})
}
