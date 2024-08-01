/*
Find all valid combinations of k numbers that sum up to n such that the following conditions are true:

Only numbers 1 through 9 are used.
Each number is used at most once.
Return a list of all possible valid combinations. The list must not contain the same combination twice, and the combinations may be returned in any order.



Example 1:

Input: k = 3, n = 7
Output: [[1,2,4]]
Explanation:
1 + 2 + 4 = 7
There are no other valid combinations.

*/

package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	backtrackSum(k, n, []int{}, &ans, 1)
	return ans
}

func backtrackSum(k int, n int, ds []int, ans *[][]int, start int) {
	if len(ds) == k && n == 0 {
		*ans = append(*ans, append([]int{}, ds...))
		return
	}
	for i := start; i <= 9; i++ {
		ds = append(ds, i)
		backtrackSum(k, n-i, ds, ans, i+1)
		ds = ds[:len(ds)-1]
	}
	return
}

func main() {
	fmt.Println(combinationSum3(3, 7))
}
