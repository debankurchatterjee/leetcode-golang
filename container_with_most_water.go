package main

func maxArea(heights []int) int {
	maxArea := 0
	for left := 0; left < len(heights); left++ {
		for right := left + 1; right < len(heights); right++ {
			currentArea := min(heights[left], heights[right]) * (right - left)
			maxArea = max(maxArea, currentArea)
		}
	}
	return maxArea
}

func maxAreaTwoPointers(height []int) int {
	left := 0
	right := len(height) - 1
	maxA := min(height[left], height[right]) * (right - left)
	for left < right {
		currentArea := min(height[left], height[right]) * (right - left)
		maxA = max(currentArea, maxA)
		if height[right] >= height[left] {
			left++
		} else {
			right--
		}
	}
	return maxA
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
