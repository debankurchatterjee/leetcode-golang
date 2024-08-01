package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	closestSum := nums[0] + nums[1] + nums[len(nums)-1]
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		l := i + 1
		r := n - 1
		for l < r {
			currentSum := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(target-currentSum)) < math.Abs(float64(target-closestSum)) {
				closestSum = currentSum
			} else if currentSum < target {
				l++
			} else {
				r--
			}
		}
	}
	return closestSum
}
