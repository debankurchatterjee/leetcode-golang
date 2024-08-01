package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	ans := math.MinInt
	prefix, suffix := 1, 1
	n := len(nums)
	for i := 0; i < n; i++ {
		// Resetting the Prefix and suffix
		if prefix == 0 {
			prefix = 1
		}
		if suffix == 0 {
			suffix = 1
		}
		prefix = prefix * nums[i]
		suffix = suffix * nums[n-i-1]
		ans = max(ans, max(prefix, suffix))
	}
	return ans
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, 0, 4, 8}))
}
