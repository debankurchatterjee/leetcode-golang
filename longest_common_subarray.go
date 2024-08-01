package main

func main() {
	findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7})
}
func findLength(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	dp := make([][]int, m+1)
	maxSize := 0
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if nums1[i-1] == nums2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = 0
			}
			maxSize = max(dp[i][j], maxSize)
		}
	}
	return maxSize
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
