package main

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	prefixSum := make([]int, n)
	candlesLeft := make([]int, n)
	candlesRight := make([]int, n)
	results := make([]int, len(queries))
	if s[0] == '*' {
		prefixSum[0] = 1
	} else {
		prefixSum[0] = 0
	}
	if s[0] == '|' {
		candlesLeft[0] = 0
	} else {
		candlesLeft[0] = -1
	}
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1]
		if s[i] == '*' {
			prefixSum[i] = prefixSum[i-1] + 1
		}
		if s[i] == '|' {
			candlesLeft[i] = i
		} else {
			candlesLeft[i] = candlesLeft[i-1]
		}
	}
	if s[n-1] == '|' {
		candlesRight[n-1] = n - 1
	} else {
		candlesRight[n-1] = n
	}
	for i := n - 2; i >= 0; i-- {
		if s[i] == '|' {
			candlesRight[i] = i
		} else {
			candlesRight[i] = candlesRight[i+1]
		}
	}
	for i := 0; i < len(queries); i++ {
		start := candlesRight[queries[i][0]]
		end := candlesLeft[queries[i][1]]
		if start > end {
			results[i] = 0
		} else {
			results[i] = prefixSum[end] - prefixSum[start]
		}
	}
	return results
}
