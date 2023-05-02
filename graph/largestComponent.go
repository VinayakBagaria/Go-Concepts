package graph

import "fmt"

func dfsLargest(node string, adjList map[string][]string, visited map[string]bool) int {
	if _, ok := visited[node]; ok {
		return 0
	}

	visited[node] = true
	nodeCount := 1

	for _, neighbour := range adjList[node] {
		nodeCount += dfsLargest(neighbour, adjList, visited)
	}

	return nodeCount
}

func DoLargestComponentCount() {
	var adjList = map[string][]string{
		"0": {"8", "1", "5"},
		"1": {"0"},
		"5": {"0", "8"},
		"8": {"0", "5"},
		"2": {"3", "4"},
		"3": {"2", "4"},
		"4": {"3", "2"},
	}

	largestCount := 0
	var largestHead string
	visited := map[string]bool{}

	for node := range adjList {
		nodeCount := dfsLargest(node, adjList, visited)
		if nodeCount > largestCount {
			largestCount = nodeCount
			largestHead = node
		}
	}

	fmt.Printf("Larget Graph/Component Count: %d\nHead: %s\n", largestCount, largestHead)
}
