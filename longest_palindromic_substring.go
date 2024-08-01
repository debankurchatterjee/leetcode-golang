package main

import "fmt"

func longestPalindromicSubString(s string) string {
	if len(s) <= 1 {
		return s
	}
	n := len(s)
	result := ""
	for i := 0; i < n; i++ {
		// odd length
		a := palindromeLength(i, i, n, s)
		// even length
		b := palindromeLength(i, i+1, n, s)
		if len(a) > len(result) {
			result = a
		}
		if len(b) > len(result) {
			result = b
		}
	}
	return result
}

func palindromeLength(l, r, n int, s string) string {
	currentSubString := ""
	maxLength := 0
	for l >= 0 && r < n && s[l] == s[r] {
		if r-l+1 > maxLength {
			maxLength = r - l + 1
			currentSubString = s[l : r+1]
		}
		l--
		r++
	}
	return currentSubString
}

func main() {
	fmt.Println(longestPalindromicSubString("babad"))
}
