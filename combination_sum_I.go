package main

import "fmt"

func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	f(len(candidates), 0, target, candidates, make([]int, 0), &ans)
	return ans
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func f(n, index, target int, candidates []int, ds []int, ans *[][]int) {
	if index == n {
		if target == 0 {
			*ans = append(*ans, append([]int{}, ds...))
		}
		return
	}
	// take case
	if candidates[index] <= target {
		ds = append(ds, candidates[index])
		f(n, index+1, target-candidates[index], candidates, ds, ans)
		// remove the last element for not take
		ds = ds[:len(ds)-1]
	}
	f(n, index+1, target, candidates, ds, ans)

}
