/*

There are n cities. Some of them are connected, while some are not. If city a is connected directly with city b, and city b is connected directly with city c, then city a is connected indirectly with city c.
A province is a group of directly or indirectly connected cities and no other cities outside of the group.
You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.
Return the total number of provinces.

Example 1:
Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
Output: 2

Example 2:
Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
Output: 3

*/

package main

var visited = make([]bool, 0)

func findCircleNum(isConnected [][]int) int {
	numProvinces := 0
	n := len(isConnected)
	for i := 0; i < n; i++ {
		if !visited[i] {
			numProvinces++
			cirDfs(i, visited, isConnected)
		}
	}
	return numProvinces
}

func cirDfs(node int, visited []bool, isConnected [][]int) {
	s := make([]int, 0)
	s = append(s, node)
	visited[node] = true
	for len(s) > 0 {
		item := s[len(s)-1]
		s = s[:len(s)-1]
		for i := 0; i < len(isConnected); i++ {
			if isConnected[item][i] == 1 && !visited[item] {
				visited[item] = true
				s = append(s, item)
			}
		}
	}
}
