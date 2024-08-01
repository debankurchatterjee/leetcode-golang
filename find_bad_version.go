package main

func firstBadVersion(n int) int {
	left := 1
	right := n

	for left < right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isBadVersion(m int) bool {
	return true
}
