package graph

import "fmt"

type queueNode struct {
	val      string
	distance int
}

func calculateShortest(source, destination string, adjList map[string][]string) int {
	sourceNode := queueNode{val: source, distance: 0}
	queue := []queueNode{sourceNode}
	visited := map[string]bool{}
	visited[source] = true

	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]

		if popped.val == destination {
			return popped.distance
		}

		for _, neighbour := range adjList[popped.val] {
			if _, ok := visited[neighbour]; !ok {
				neighbourNode := queueNode{val: neighbour, distance: popped.distance + 1}
				queue = append(queue, neighbourNode)
				visited[neighbour] = true
			}
		}
	}

	return -1
}

func DoShortestPathDistance() {
	edges := [][]string{
		{"w", "x"},
		{"x", "y"},
		{"z", "y"},
		{"z", "v"},
		{"w", "v"},
	}
	adjList := MakeAdjacencyList(edges)
	fmt.Printf("Shortest path from w => z is %d\n", calculateShortest("w", "z", adjList))
}
