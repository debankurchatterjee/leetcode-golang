package main

import (
	"fmt"
)

func minimumFuelCost(roads [][]int, seats int) int64 {
	numberOfCities := len(roads) + 1
	adjList := make(map[int][]int, 0)
	// in-degree and out-degree
	numberOfIncomingOutgoingRoads := make([]int, numberOfCities)
	for _, road := range roads {
		adjList[road[0]] = append(adjList[road[0]], road[1])
		adjList[road[1]] = append(adjList[road[1]], road[0])
		numberOfIncomingOutgoingRoads[road[0]]++
		numberOfIncomingOutgoingRoads[road[1]]++
	}
	// Bottom-up BFS
	var totalFuelCost int64 = 0
	// at-least every city will have one representative
	totalRepresentatives := make([]int, numberOfCities)
	for i := 1; i < numberOfCities; i++ {
		totalRepresentatives[i] = 1
	}
	q := make([]int, 0)
	// inserting all the leaf nodes to queue
	for i := 1; i < numberOfCities; i++ {
		if numberOfIncomingOutgoingRoads[i] == 1 {
			q = append(q, i)
		}
	}
	for len(q) != 0 {
		city := q[0]
		q = q[1:]
		// fuel required for the representative to reach parent
		testVal := 0
		if totalRepresentatives[city]%seats != 0 {
			testVal = 1
		}
		totalFuelCost += int64((totalRepresentatives[city] / seats) + testVal)
		for _, nbr := range adjList[city] {
			numberOfIncomingOutgoingRoads[nbr]--
			totalRepresentatives[nbr] += totalRepresentatives[city]
			if numberOfIncomingOutgoingRoads[nbr] == 1 && nbr != 0 {
				q = append(q, nbr)
			}
		}
	}
	return totalFuelCost
}

func main() {
	fmt.Println(minimumFuelCost([][]int{{0, 1}, {1, 2}}, 3))
}
