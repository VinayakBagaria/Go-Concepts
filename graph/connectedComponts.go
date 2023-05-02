package graph

import "fmt"

func dfsConnected(node string, adjList map[string][]string, visited map[string]bool) bool {
	// If this node is already visited, its not a new graph
	if _, ok := visited[node]; ok {
		return false
	}

	visited[node] = true

	for _, neighbour := range adjList[node] {
		dfsConnected(neighbour, adjList, visited)
	}

	return true
}

func DoConnectedComponents() {
	var adjList = map[string][]string{
		"0": {"8", "1", "5"},
		"1": {"0"},
		"5": {"0", "8"},
		"8": {"0", "5"},
		"2": {"3", "4"},
		"3": {"2", "4"},
		"4": {"3", "2"},
	}

	visited := map[string]bool{}
	counter := 0
	headOfGraphs := []string{}

	for node := range adjList {
		// If this node is not already visited, its part of a new graph
		if dfsConnected(node, adjList, visited) {
			counter += 1
			headOfGraphs = append(headOfGraphs, node)
		}
	}

	fmt.Printf("Connected Graphs/Components Count: %d\nHeads: %v\n", counter, headOfGraphs)
}
