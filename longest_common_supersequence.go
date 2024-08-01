package main

func lcs(s, t string) {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			}
			if s[i-1] == t[j-1] {
				dp[i][i] = 1 + dp[i-1][j-1]
			} else {
				dp[i][i] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
}
