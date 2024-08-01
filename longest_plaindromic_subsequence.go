package main

import "fmt"

/*
LCS(x,y) = if i,j > 0 and x==y ->1 + LCS(x-1,y-1)
*/
// Recursion solution
func longestPalindromicSubsequence(s string, st, end int) int {
	if st > end {
		return 0
	}
	if st == end {
		return 1
	}
	if s[st] == s[end] {
		return 2 + longestPalindromicSubsequence(s, st+1, end-1)
	}
	return max(longestPalindromicSubsequence(s, st+1, end), longestPalindromicSubsequence(s, st, end-1))
}

func longestPalindromicSubsequenceMemo(s string, st, end int, dp map[[2]int]int) int {
	if st > end {
		return 0
	}
	if st == end {
		return 1
	}
	if _, ok := dp[[2]int{st, end}]; ok {
		return dp[[2]int{st, end}]
	}
	if s[st] == s[end] {
		dp[[2]int{st, end}] = 2 + longestPalindromicSubsequence(s, st+1, end-1)
		return dp[[2]int{st, end}]
	}
	dp[[2]int{st, end}] = max(longestPalindromicSubsequence(s, st+1, end), longestPalindromicSubsequence(s, st, end-1))
	return dp[[2]int{st, end}]
}

func longestPalindromeBottomUp(s string) {

}

func main() {
	s := "cddpd"
	fmt.Println(longestPalindromicSubsequence(s, 0, len(s)-1))
	dp := make(map[[2]int]int)
	fmt.Println(longestPalindromicSubsequenceMemo(s, 0, len(s)-1, dp))
}
