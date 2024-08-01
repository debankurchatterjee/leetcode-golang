package main

func longestCommonSubsequenceIterative(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	maxLen := 0
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = maxU(dp[i-1][j], dp[i][j-1])
			}
			maxLen = max(maxLen, dp[i][j])
		}
	}
	return maxLen
}

func maxU(a, b int) int {
	if a > b {
		return a
	}
	return b
}
