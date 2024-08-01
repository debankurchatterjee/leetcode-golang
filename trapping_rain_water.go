package main

import "fmt"

func trap(height []int) int {
	size := len(height)
	left := make([]int, size)
	right := make([]int, size)
	left[0] = height[0]
	right[size-1] = height[size-1]
	ans := 0

	// Populate the left array
	for i := 1; i < size; i++ {
		left[i] = max(left[i-1], height[i])
	}
	// populate the right array
	for i := size - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	for i := 0; i < size; i++ {
		ans = ans + min(left[i], right[i]) - height[i]
	}
	fmt.Println(left, right)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	fmt.Println(trap([]int{3, 1, 2, 4, 0, 1, 3, 2}))
}
