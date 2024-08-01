package main

import "fmt"

var cycleArr = make([]int, 0)

func distanceToCycle(n int, edges [][]int) []int {
	cycleArr = make([]int, 0)
	adjacencyList := make(map[int][]int, 0)
	for _, edge := range edges {
		adjacencyList[edge[0]] = append(adjacencyList[edge[0]], edge[1])
		adjacencyList[edge[1]] = append(adjacencyList[edge[1]], edge[0])
	}
	fmt.Println(cycleDetectionUsingDegree(adjacencyList))
	visitedDfs := make([]bool, n)
	x := dfsCycleDetection(visitedDfs, 0, -1, adjacencyList)
	if !x {
		return []int{}
	}
	for i := len(cycleArr) - 2; i > 0; i-- {
		if cycleArr[i-1] == cycleArr[len(cycleArr)-1] {
			cycleArr = cycleArr[i:]
			break
		}
	}
	//	cycleArr = cycleDetectionUsingDegree(adjacencyList)
	res := make([]int, n)
	visited := make([]bool, n)
	var queue []int
	for _, node := range cycleArr {
		queue = append(queue, node)
		visited[node] = true
	}
	var dist int
	for len(queue) > 0 {
		dist = dist + 1
		for _, _ = range queue {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adjacencyList[u] {
				if !visited[v] {
					queue = append(queue, v)
					visited[v] = true
					res[v] = dist
				}
			}
		}
	}
	return res
}

func dfsCycleDetection(visited []bool, currentNode int, parent int, al map[int][]int) bool {
	visited[currentNode] = true
	cycleArr = append(cycleArr, currentNode)
	for _, v := range al[currentNode] {
		if !visited[v] {
			if dfsCycleDetection(visited, v, currentNode, al) {
				return true
			}
		} else if parent != v {
			cycleArr = append(cycleArr, v)
			return true
		}
	}
	cycleArr = cycleArr[:len(cycleArr)-1]
	return false
}

func cycleDetectionUsingDegree(al map[int][]int) []int {
	degree := make(map[int]int, 0)
	visited := make(map[int]bool, 0)
	results := make([]int, 0)
	q := make([]int, 0)
	// Generating the degree map form AL
	for k, v := range al {
		degree[k] = len(v)
	}
	// Adding those nodes whose degree is 1 to the queue
	for {
		for k, v := range degree {
			if v == 1 && !visited[k] {
				q = append(q, k)
			}
		}
		if len(q) == 0 {
			break
		}
		for len(q) != 0 {
			currentVertex := q[0]
			q = q[1:]
			visited[currentVertex] = true
			for _, nbrs := range al[currentVertex] {
				degree[nbrs] = degree[nbrs] - 1
			}
		}
	}
	for k, _ := range al {
		if visited[k] == false {
			results = append(results, k)
		}
	}
	return results
}

func cycleDetectionPlainDfsStack(al map[int]int, node int) {
	visited := make(map[int]bool, 0)
	s := make([]int, 0)
	s = append(s, node)
	for len(s) != 0 {
		item := s[len(s)-1]
		if !visited[item] {
			visited[item] = true
			s = s[:len(s)-1]
		}
		for _, nbr := range al {
			if !visited[nbr] {
				s = append(s, nbr)
			}
		}
	}
}

func main() {
	fmt.Println(distanceToCycle(7, [][]int{{1, 2}, {2, 4}, {4, 3}, {3, 1}, {0, 1}, {5, 2}, {6, 5}}))
	fmt.Println(distanceToCycle(9, [][]int{{0, 1}, {1, 2}, {0, 2}, {2, 6}, {6, 7}, {6, 8}, {0, 3}, {3, 4}, {3, 5}}))
	fmt.Println(distanceToCycle(8, [][]int{{4, 0}, {7, 5}, {1, 2}, {3, 2}, {5, 2}, {7, 0}, {3, 6}, {6, 1}}))
}
