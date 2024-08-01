package main

import "sort"

func threeSumSmaller(nums []int, target int) int {
	// sort the array to bring all the repeated elements at one place
	sort.Ints(nums)
	n := len(nums)
	count := 0
	for i := 0; i < n-2; i++ {
		l := i + 1
		r := n - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < target {
				count = count + r - l
				l++
			} else {
				r--
			}
		}
	}
	return count
}
