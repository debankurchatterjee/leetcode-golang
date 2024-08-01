package main

func findJudge(n int, trust [][]int) int {
	outdegree := make([]int, n+1)
	indegree := make([]int, n+1)

	for _, eachTrust := range trust {
		outdegree[eachTrust[0]] = outdegree[eachTrust[0]] + 1
		indegree[eachTrust[1]] = indegree[eachTrust[1]] + 1
	}
	for i := 1; i <= n; i++ {
		if indegree[i] == n-1 && outdegree[i] == 0 {
			return i
		}
	}
	return -1
}

func main() {
	findJudge(3, [][]int{{1, 3}, {2, 3}})
}
