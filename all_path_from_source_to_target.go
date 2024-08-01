package main

//func allPathsSourceToTarget(graph [][]int) [][]int {
//	result := make([][]int, 0)
//	var dfs func(int, []int, int)
//	dfs = func(currentNode int, currentPath []int, dest int) {
//		for _, v := range graph[currentNode] {
//			if v == dest {
//				var tempPath []int
//				for _, path := range currentPath {
//					tempPath = append(tempPath, path)
//				}
//				result = append(results, tempPath)
//			} else {
//				currentPath = append(currentPath, v)
//				dfs(v, currentPath, dest)
//			}
//		}
//	}
//	dfs()
//	return results
//}
