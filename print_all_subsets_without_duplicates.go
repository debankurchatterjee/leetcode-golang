package main

import (
	"fmt"
	"sort"
)

// generateSubsets generates all subsets, making sure to avoid duplicates.
func generateSubsets(nums []int, index int, currentSubset []int, result *[][]int) {
	if index == len(nums) {
		// Make a copy of currentSubset and add it to result
		subsetCopy := make([]int, len(currentSubset))
		copy(subsetCopy, currentSubset)
		*result = append(*result, subsetCopy)
		return
	}

	// Include the current element
	generateSubsets(nums, index+1, append(currentSubset, nums[index]), result)

	// Skip duplicates
	for index+1 < len(nums) && nums[index] == nums[index+1] {
		index++
	}

	// Exclude the current element and move to the next unique element
	generateSubsets(nums, index+1, currentSubset, result)
}

func main() {
	nums := []int{1, 2, 2, 3}
	sort.Ints(nums) // Sort the array to bring duplicates together

	var result [][]int
	generateSubsets(nums, 0, []int{}, &result)

	// Print the subsets
	for _, subset := range result {
		fmt.Println(subset)
	}
}
