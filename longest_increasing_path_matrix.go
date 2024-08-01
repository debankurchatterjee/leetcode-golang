package main

func longestIncreasingPath(matrix [][]int) int {
	row, col := len(matrix), len(matrix[0])
	// store result in the dp map for every row and column
	dp := make(map[[2]int]int)
	// Path in-creasing order we will use the previous value
	var dfs func(r, c, previousVal int) int
	dfs = func(r, c, previousVal int) int {
		// check for out of bounce and current position is less than or equal to previous value
		if r < 0 || r == row || c < 0 || c == col || matrix[r][c] <= previousVal {
			return 0
		}
		// return from cache
		if _, ok := dp[[2]int{r, c}]; ok {
			return dp[[2]int{r, c}]
		}
		// at least one result will be there
		res := 1
		res = max(res, 1+dfs(r+1, c, matrix[r][c]))
		res = max(res, 1+dfs(r-1, c, matrix[r][c]))
		res = max(res, 1+dfs(r, c+1, matrix[r][c]))
		res = max(res, 1+dfs(r, c-1, matrix[r][c]))
		dp[[2]int{r, c}] = res
		return res
	}
	// run DFS for each and every cell
	result := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			result = max(dfs(i, j, -1), result)
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
