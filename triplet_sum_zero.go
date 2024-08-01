package main

import "sort"

func threeSum1(s []int) {
	sort.Ints(s)
	n := len(s)
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && s[i] == s[i-1] {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			sum := s[i] + s[l] + s[r]
			if sum == 0 {
				res = append(res, sum)
				for l < r && s[l] == s[l+1] {
					l++
				}
				for l < r && s[r] == s[r-1] {
					r--
				}
				l++
				r--
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
}
