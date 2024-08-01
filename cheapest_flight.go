package main

import "math"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	distance := make([]int, n)
	for i := 0; i < len(distance); i++ {
		distance[i] = math.MaxInt
	}
	distance[src] = 0
	for i := 0; i <= k; i++ {
		tmpDistance := make([]int, n)
		copy(tmpDistance, distance)
		for _, flight := range flights {
			u := flight[0]
			v := flight[1]
			c := flight[2]
			// if the distance is reachable
			if distance[u] != math.MaxInt {
				tmpDistance[v] = min(tmpDistance[v], distance[u]+c)
			}
		}
		copy(distance, tmpDistance)
	}
	if distance[dst] == math.MaxInt {
		return -1
	}
	return distance[dst]
}
